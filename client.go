package duros

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

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

	defaultUserAgent = "duros-go"
)

// client for the duros grpc endpoint
type client struct {
	eps  EPs
	conn *grpc.ClientConn
	log  *zap.SugaredLogger
}

// DialConfig is the configuration to create a duros-api connection
type DialConfig struct {
	Endpoints EPs
	Scheme    GRPCScheme
	Token     string
	Log       *zap.SugaredLogger
	// UserAgent to use, if empty duros-go is used
	UserAgent string
	TLSConfig *tls.Config
}

// Credentials specify the TLS Certificate based authentication for the grpc connection
type Credentials struct {
	Certificates []tls.Certificate
	RootCAs      *x509.CertPool
	ServerName   string
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
	if !config.Endpoints.isValid() {
		return nil, status.Errorf(codes.InvalidArgument,
			"invalid target endpoints specified: [%s]", config.Endpoints)
	}
	id := fmt.Sprintf("%07s", strconv.FormatUint(uint64(prng.Uint32()), 36))
	if config.Log == nil {
		return nil, fmt.Errorf("logger is nil")
	}
	log := config.Log

	ua := defaultUserAgent
	if config.UserAgent != "" {
		ua = config.UserAgent
	}

	log.Infow("connecting...",
		"client", ua,
		"targets", config.Endpoints,
		"client-id", id,
	)

	res := &client{
		eps: config.Endpoints.clone(),
		log: log,
	}

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithLevels(grpcToZapLevel),
	}
	interceptors := []grpc.UnaryClientInterceptor{
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

	opts := []grpc.DialOption{
		//grpc.WithBlock(),
		grpc.WithDisableRetry(),
		grpc.WithUserAgent(ua),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(false)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(interceptors...)),
		grpc.WithKeepaliveParams(kal),
		grpc.WithConnectParams(cp),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, roundrobin.Name)),
		grpc.WithPerRPCCredentials(tokenAuth{
			token: config.Token,
		}),
	}
	// Configure tls ca certificate based auth if credentials are given
	switch config.Scheme {
	case GRPC:
		log.Infof("connecting insecurely")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case GRPCS:
		log.Infof("connecting securely")
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config.TLSConfig)))
	default:
		return nil, fmt.Errorf("unsupported scheme:%v", config.Scheme)
	}

	var err error
	res.conn, err = grpc.DialContext(
		ctx,
		"ipv4:"+config.Endpoints.String(),
		opts...,
	)
	if err != nil {
		log.Errorw("failed to connect", "endpoints", config.Endpoints.String(), "error", err.Error())
		return nil, err
	}

	log.Infof("connected")

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
