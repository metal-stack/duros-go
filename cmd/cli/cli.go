package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/uuid"
	duros "github.com/metal-stack/duros-go"
	v2 "github.com/metal-stack/duros-go/api/duros/v2"
	"go.uber.org/zap"
)

const (
	adminPub = `
-----BEGIN CERTIFICATE-----
MA0GA1UECBMGSXNyYWVsMQswCQYDVQQHEwJLUzESMBAGA1UEChMJTGlnaHRiaXRz
MQswCQYDVQQLEwJCWTEOMAwGA1UEAxMFYWRtaW4wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQC1BiAg7ojVhGfaGGaAiz01KZhVKJPK2XqGUuaoecZ+psxe
3L9phD82fJKmD2CbuKBCU2wPQieMzLYXqdSWmbtffFQr0hMkgdVJATwoSZHtIMnr
zFm9HqSemPqC0EnpkpHC3nJHSsZPvs088I5NpBI8PZNEqO4BDzy58ajG7Q1ZE697
tYoNBfrSNB07M8SDQWxBhEv0jvxrLdFEeWkMStXegCdpMdKNgEOH0rcPOhu3mSfX
vnQFChodkqXXrs+t0jXblqPyc97vmWqWIXLlGWsQcAF0fdXJRfAbu2kLgiZBqlW2
ozwntfhxRAEAJlnNPy+QuMPl18MemsZGYDZ5ARDTAgMBAAGjfzB9MA4GA1UdDwEB
/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/
BAIwADAdBgNVHQ4EFgQUbe3QvVd3Bz/vmT+Z/PvigRJ0EeMwHwYDVR0jBBgwFoAU
quTzHAJ8iP9siRG09QGi14FmxBowDQYJKoZIhvcNAQELBQADggEBAKCQrVwnsDo6
hjL4cYxdaf4V/RLAnSHV3whuiIiXi9gbek9hbD1IY1UWUHJBMAHqSDBoBr475DY/
nMO5Ow8w5dJ6zvNzSZP2rCQSpZZy03hSVvlZqp2Ag5iaxPQjPk0Ns5KLhERPrbh+
NpeTXSb+Kxj3R2j7Ri9qqLl1uux/eg4YBbFy2VocxVVQRkWaxabUaydzyaLYFKmt
HUfJjCsupU5ZztNFTaBAacZrITPDL3JM9sagS3iCZq7JrA/M+flRs1BoRtgc2jYP
lI+Z9+7r0XUW2k1qej3ktWgsn68bOmu2uIQFrNskFPT4tHma9igc1/bJdpv9qjJ4
htcrucZWyL4=
-----END CERTIFICATE-----
`
)

func main() {
	var (
		endpoints  string
		token      string
		scheme     string
		caFile     string
		certFile   string
		keyFile    string
		serverName string
	)
	flag.StringVar(&token, "token", "", "The token to authenticate against the lightbits api.")
	flag.StringVar(&scheme, "scheme", "grpcs", "The scheme to connect to the lightbits api, can be grpc|grpcs")
	flag.StringVar(&endpoints, "endpoints", "localhost:443", "The endpoints, in the form host:port,host:port of the lightbits api.")
	flag.StringVar(&caFile, "ca-file", "", "the filename of the ca for certificate based authentication")
	flag.StringVar(&certFile, "cert-file", "", "the filename of the ca certificate for certificate based authentication")
	flag.StringVar(&keyFile, "key-file", "", "the filename of the key  for certificate based authentication")
	flag.StringVar(&serverName, "server-name", "", "the servername to validate against for certificate based authentication")

	flag.Parse()

	zlog, err := zap.NewProduction()
	if err != nil {
		panic((err))
	}
	var grpcScheme duros.GRPCScheme
	switch scheme {
	case "grpc":
		grpcScheme = duros.GRPC
	case "grpcs":
		grpcScheme = duros.GRPCS
	default:
		panic(fmt.Errorf("unsupported scheme:%s", scheme))
	}

	ctx := context.Background()
	dialConfig := duros.DialConfig{
		Endpoints: duros.MustParseCSV(endpoints),
		Scheme:    grpcScheme,
		Token:     token,
		Log:       zlog.Sugar(),
		UserAgent: "duros-go-cli",
	}
	if caFile != "" && certFile != "" && keyFile != "" && serverName != "" {
		creds := &duros.Credentials{
			CAFile:     caFile,
			Certfile:   certFile,
			Keyfile:    keyFile,
			ServerName: serverName,
		}
		dialConfig.Credentials = creds
	}

	durosClient, err := duros.Dial(ctx, dialConfig)
	if err != nil {
		panic(err)
	}

	info, err := durosClient.GetClusterInfo(ctx, &v2.GetClusterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cluster:%v\n", info)

	err = createProjectWithCredentials(ctx, durosClient)
	if err != nil {
		panic(err)
	}
}

func createProjectWithCredentials(ctx context.Context, client v2.DurosAPIClient) error {
	projectID := "project-a"
	lpr, err := client.ListProjects(ctx, &v2.ListProjectsRequest{})
	if err != nil {
		return err
	}
	projectExists := false
	for _, p := range lpr.Projects {
		fmt.Printf("Project:%v\n", p.Name)
		if p.Name == projectID {
			projectExists = true
		}
	}
	if !projectExists {
		project, err := client.CreateProject(ctx, &v2.CreateProjectRequest{Name: projectID})
		if err != nil {
			return err
		}
		fmt.Printf("Project created:%q\n", project)
	}

	creds, err := client.ListCredentials(ctx, &v2.ListCredentialsRequest{ProjectName: projectID})
	if err != nil {
		return err
	}
	for _, cred := range creds.Credentials {
		fmt.Printf("Credential: ID:%s Project:%s Type:%s\n", cred.ID, cred.ProjectName, cred.Type)
		_, err := client.DeleteCredential(ctx, &v2.DeleteCredentialRequest{ID: cred.ID, ProjectName: cred.ProjectName})
		if err != nil {
			return err
		}
	}
	id := uuid.New()
	credential, err := client.CreateCredential(ctx, &v2.CreateCredentialRequest{
		ProjectName: projectID,
		ID:          id.String(),
		Type:        v2.CredsType_RS256PubKey,
		Payload:     []byte(adminPub),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Credentials created:%s type:%s\n", credential.ID, credential.Type)
	return nil
}
