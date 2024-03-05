package duros

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

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
	hostRegex *regexp.Regexp
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

// DialConfig is the configuration to create a duros-api connection
type DialConfig struct {
	Endpoint        string
	Scheme          GRPCScheme
	Token           string
	Credentials     *Credentials
	ByteCredentials *ByteCredentials
	Log             *slog.Logger
	// UserAgent to use, if empty duros-go is used
	UserAgent string
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

func init() {
	hostRegex = regexp.MustCompile(`^([a-zA-Z0-9.\[\]:%-]+)$`)
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
	if err := isValid(config.Endpoint); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"invalid target endpoints specified: %v", err)
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

	log.Info("connecting...",
		"client", ua,
		"target", config.Endpoint,
		"client-id", id,
	)

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
		grpc.WithDisableRetry(),
		grpc.WithUserAgent(ua),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(interceptorLogger(log))),
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
		log.Info("connecting insecurely")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case GRPCS:
		log.Info("connecting securely")
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
		return nil, fmt.Errorf("unsupported scheme:%v", config.Scheme)
	}

	conn, err := grpc.DialContext(ctx, config.Endpoint, opts...)
	if err != nil {
		log.Error("failed to connect", "endpoints", config.Endpoint, "error", err.Error())
		return nil, err
	}

	log.Info("connected")

	return durosv2.NewDurosAPIClient(conn), nil
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

func (c Credentials) getTransportCredentials() (credentials.TransportCredentials, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system credentials: %w", err)
	}
	if c.CAFile == "" || c.Certfile == "" || c.Keyfile == "" || c.ServerName == "" {
		return nil, fmt.Errorf("all credentials properties must be configured")
	}
	ca, err := os.ReadFile(c.CAFile)
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

func isValid(endpoint string) error {
	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		return err
	}
	if strings.TrimSpace(host) == "" {
		return fmt.Errorf("invalid empty host")
	}
	if !hostRegex.MatchString(host) {
		return fmt.Errorf("invalid host %q", host)
	}
	if _, err = strconv.ParseUint(port, 10, 16); err != nil {
		return fmt.Errorf("invalid port number %q", port)
	}
	return nil
}

// interceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg, fields...)
		case logging.LevelInfo:
			l.Info(msg, fields...)
		case logging.LevelWarn:
			l.Warn(msg, fields...)
		case logging.LevelError:
			l.Error(msg, fields...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
