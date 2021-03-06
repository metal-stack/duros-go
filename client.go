package duros

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/gogo/status"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"

	durosv2 "github.com/metal-stack/duros-go/api/duros/v2"
)

var (
	//nolint
	prng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// GRPCScheme the scheme to talk to the duros api endpoint, can be plaintext or https
type GRPCScheme string

const (
	// GRPC defines a plaitext communication
	GRPC GRPCScheme = "grpc"
	// GRPCS defines https protocol for the communication
	GRPCS GRPCScheme = "grpcs"
)

// client for the duros grpc endpoint
type client struct {
	eps         EPs
	conn        *grpc.ClientConn
	DurosClient durosv2.DurosAPIClient

	id   string
	tgts string // cached string repr of `eps`
	log  *zap.SugaredLogger

	// peerMu protects all peer-related fields:
	peerMu   sync.Mutex
	lastPeer peer.Peer
	switched bool // a matter of aesthetics: 1st conn shouldn't warn
}

// DialConfig is the configuration to create a duros-api connection
type DialConfig struct {
	Endpoints       EPs
	Scheme          GRPCScheme
	Token           string
	Credentials     *Credentials
	ByteCredentials *ByteCredentials
	Log             *zap.SugaredLogger
}

// Credentials specify the TLS Certificate based authentication for the grpc connection
// If you provide credentials, provide either these or byte credentials but not both.
type Credentials struct {
	ServerName string
	Certfile   string
	Keyfile    string
	CAFile     string
}

// Credentials specify the TLS Certificate based authentication for the grpc connection
// without having to use certificate files.
// If you provide credentials, provide either these or file path credentials but not both.
type ByteCredentials struct {
	ServerName string
	Cert       []byte
	Key        []byte
	CA         []byte
}

// Dial creates a LightOS cluster client. it is a blocking call and will only
// return once the connection to [at least one of the] `targets` has been
// actually established - subject to `ctx` limitations. if `ctx` specified
// timeout or duration - dialling (and only dialling!) timeout will be set
// accordingly. `ctx` can also be used to cancel the dialling process, as per
// usual.
//
// the cluster client will make an effort to transparently reconnect to one of
// the `targets` in case of connection loss. if the process of finding a live
// and responsive target amongst `targets` and establishing the connection takes
// longer than the actual operation context timeout (as opposed to the `ctx`
// passed here) - `DeadlineExceeded` will be returned as usual, and the caller
// can retry the operation.
func Dial(ctx context.Context, config DialConfig) (durosv2.DurosAPIClient, error) {
	if config.Credentials != nil && config.ByteCredentials != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"if you provide credentials, provide either file or byte credentials but not both")
	}
	if !config.Endpoints.isValid() {
		return nil, status.Errorf(codes.InvalidArgument,
			"invalid target endpoints specified: [%s]", config.Endpoints)
	}
	id := fmt.Sprintf("%07s", strconv.FormatUint(uint64(prng.Uint32()), 36))
	if config.Log == nil {
		return nil, fmt.Errorf("logger is nil")
	}
	log := config.Log

	log.Infow("connecting...",
		"client", "duros-go",
		"targets", config.Endpoints,
		"client-id", id,
	)

	res := &client{
		eps:  config.Endpoints.clone(),
		id:   id,
		tgts: config.Endpoints.String(),
		log:  log,
	}

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithLevels(grpcToZapLevel),
	}
	interceptors := []grpc.UnaryClientInterceptor{
		mkUnaryClientInterceptor(res),
		grpc_zap.UnaryClientInterceptor(log.Desugar(), zapOpts...),
		grpc_zap.PayloadUnaryClientInterceptor(log.Desugar(),
			func(context.Context, string) bool { return true },
		),
	}

	// these are broadly in line with the expected server SLOs:
	kal := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}

	// these control only dialling, both initial and subsequent (on server
	// switch-over). it's relatively tight, but then the LightOS clusters
	// controlled by a LB CSI plugin will typically be on the same DC
	// network. given that some COs are fairly aggressive about their call
	// deadlines (for K8s - often 10sec), this should, hopefully, give the
	// client a decent chance to try out at least one more server before
	// the CO call will time out, saving a top-level retry cycle.
	dialBackoffConfig := backoff.Config{
		BaseDelay:  1.0 * time.Second,
		Multiplier: 1.2,
		Jitter:     0.1,
		MaxDelay:   7 * time.Second,
	}
	cp := grpc.ConnectParams{
		Backoff:           dialBackoffConfig,
		MinConnectTimeout: 6 * time.Second,
	}

	scheme := "lightos-" + id
	lbr := newLbResolver(log, scheme, res.eps)

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithDisableRetry(),
		grpc.WithUserAgent("duros-go"), // TODO enable setting this by client
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(interceptors...)),
		grpc.WithKeepaliveParams(kal),
		grpc.WithConnectParams(cp),
		grpc.WithResolvers(lbr),
		grpc.WithPerRPCCredentials(tokenAuth{
			token: config.Token,
		}),
	}
	// Configure tls ca certificate based auth if credentials are given
	switch config.Scheme {
	case GRPC:
		log.Infof("connecting insecurely")
		opts = append(opts, grpc.WithInsecure())
	case GRPCS:
		log.Infof("connecting securely")
		if config.Credentials != nil {
			creds, err := config.Credentials.getTransportCredentials()
			if err != nil {
				return nil, err
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else if config.ByteCredentials != nil {
			creds, err := config.ByteCredentials.getTransportCredentials()
			if err != nil {
				return nil, err
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			//nolint
			opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
		}
	default:
		return nil, fmt.Errorf("unsupported scheme:%v", scheme)
	}

	var err error
	res.conn, err = grpc.DialContext(
		ctx,
		scheme+":///lb-resolver", // use our resolver instead of explicit target
		opts...,
	)
	if err != nil {
		log.Errorw("failed to connect", "error", err.Error())
		return nil, err
	}

	c := durosv2.NewDurosAPIClient(res.conn)
	return c, nil
}

type tokenAuth struct {
	token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}

// TODO: add stream interceptor *IF* LB API adds streaming entrypoints...
func mkUnaryClientInterceptor(c *client) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, rep interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		return c.peerReviewUnaryInterceptor(ctx, method, req, rep, cc, invoker, opts...)
	}
}

func (c *client) peerReviewUnaryInterceptor( // sic!
	ctx context.Context, method string, req, rep interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {
	var currPeer peer.Peer
	opts = append(opts, grpc.Peer(&currPeer))
	err := invoker(ctx, method, req, rep, cc, opts...)
	c.peerMu.Lock()
	if currPeer.Addr != c.lastPeer.Addr {
		// TODO: introduce rate-limiter to spare logs and perf!
		lastPeer := c.lastPeer
		c.lastPeer = currPeer
		c.peerMu.Unlock()
		curr := "<NONE>"
		if currPeer.Addr != nil {
			curr = currPeer.Addr.String()
		}
		last := "<NONE>"
		if lastPeer.Addr != nil {
			last = lastPeer.Addr.String()
		}
		// don't want to warn on healthy flow...
		if c.switched {
			c.log.Warnf("switched target: %s -> %s", last, curr)
		} else {
			c.switched = true
			c.log.Infof("switched target: %s -> %s", last, curr)
		}
	} else {
		c.peerMu.Unlock()
	}
	return err
}

func grpcToZapLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.OK,
		codes.Canceled,
		codes.DeadlineExceeded,
		codes.NotFound,
		codes.Unavailable:
		return zapcore.InfoLevel
	case codes.Aborted,
		codes.AlreadyExists,
		codes.FailedPrecondition,
		codes.InvalidArgument,
		codes.OutOfRange,
		codes.PermissionDenied,
		codes.ResourceExhausted,
		codes.Unauthenticated:
		return zapcore.WarnLevel
	case codes.DataLoss,
		codes.Internal,
		codes.Unimplemented,
		codes.Unknown:
		return zapcore.ErrorLevel
	default:
		return zapcore.ErrorLevel
	}
}

func (c Credentials) getTransportCredentials() (credentials.TransportCredentials, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system credentials: %w", err)
	}
	if c.CAFile == "" || c.Certfile == "" || c.Keyfile == "" || c.ServerName == "" {
		return nil, fmt.Errorf("all credentials properties must be configured")
	}
	ca, err := ioutil.ReadFile(c.CAFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %w", err)
	}

	ok := certPool.AppendCertsFromPEM(ca)
	if !ok {
		return nil, fmt.Errorf("failed to append ca certs: %s", c.CAFile)
	}

	clientCertificate, err := tls.LoadX509KeyPair(c.Certfile, c.Keyfile)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %w", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   c.ServerName,
		Certificates: []tls.Certificate{clientCertificate},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
	})
	return creds, nil
}

func (c ByteCredentials) getTransportCredentials() (credentials.TransportCredentials, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system credentials: %w", err)
	}
	if string(c.CA) == "" || string(c.Cert) == "" || string(c.Key) == "" || c.ServerName == "" {
		return nil, fmt.Errorf("all credentials properties must be configured")
	}

	ok := certPool.AppendCertsFromPEM(c.CA)
	if !ok {
		return nil, fmt.Errorf("failed to append ca certs: %s", c.CA)
	}

	clientCertificate, err := tls.X509KeyPair(c.Cert, c.Key)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %w", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   c.ServerName,
		Certificates: []tls.Certificate{clientCertificate},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
	})
	return creds, nil
}
