package restvirt

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/verbit/restvirt-client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
)

type Client struct {
	Username                    string
	Password                    string
	DNSClient                   pb.DNSClient
	DomainServiceClient         pb.DomainServiceClient
	PortForwardingServiceClient pb.PortForwardingServiceClient
	RouteServiceClient          pb.RouteServiceClient
	VolumeServiceClient         pb.VolumeServiceClient
	ControllerServiceClient     pb.ControllerServiceClient
	HostServiceClient           pb.HostServiceClient
}

type ClientConfig struct {
	Host       string `yaml:"host"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	HostCA     string `yaml:"ca"`
	ClientCert string `yaml:"cert"`
	ClientKey  string `yaml:"key"`
}

type ErrorMessage struct {
	Error string
}

type NotFoundError struct {
	grpcStatus *status.Status
}

func (e *NotFoundError) Error() string {
	return "not found"
}

func (e *NotFoundError) GRPCStatus() *status.Status {
	return e.grpcStatus
}

func errorWrapper(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	sErr := status.Convert(err)
	if sErr.Code() == codes.NotFound {
		return &NotFoundError{
			grpcStatus: sErr,
		}
	}
	return err
}

func NewClientFromConfig(config ClientConfig) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(errorWrapper),
	}
	var tlsConfig *tls.Config = nil
	if config.HostCA != "" {
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}

		if ok := rootCAs.AppendCertsFromPEM([]byte(config.HostCA)); !ok {
			return nil, fmt.Errorf("couldn't append hosts's CA to the bundle")
		}

		var certs = []tls.Certificate{}
		if config.ClientCert != "" || config.ClientKey != "" {
			clientKeyPair, err := tls.X509KeyPair([]byte(config.ClientCert), []byte(config.ClientKey))
			if err != nil {
				return nil, err
			}
			certs = append(certs, clientKeyPair)
		}

		tlsConfig = &tls.Config{
			RootCAs:      rootCAs,
			Certificates: certs,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(config.Host, opts...)
	if err != nil {
		return nil, err
	}
	c := Client{
		Username:                    config.Username,
		Password:                    config.Password,
		DNSClient:                   pb.NewDNSClient(conn),
		DomainServiceClient:         pb.NewDomainServiceClient(conn),
		PortForwardingServiceClient: pb.NewPortForwardingServiceClient(conn),
		RouteServiceClient:          pb.NewRouteServiceClient(conn),
		VolumeServiceClient:         pb.NewVolumeServiceClient(conn),
		ControllerServiceClient:     pb.NewControllerServiceClient(conn),
		HostServiceClient:           pb.NewHostServiceClient(conn),
	}

	return &c, nil
}

func NewClientFromEnvironment() (*Client, error) {
	configPath, exists := os.LookupEnv("RESTVIRT_CONFIG")
	if !exists {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		configPath = path.Join(homeDir, ".restvirt", "config.yaml")
	}
	configRaw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var configMap map[string]ClientConfig
	err = yaml.Unmarshal(configRaw, &configMap)
	if err != nil {
		return nil, err
	}

	profile, exists := os.LookupEnv("RESTVIRT_PROFILE")
	if !exists {
		profile = "default"
	}

	config, exists := configMap[profile]
	if !exists {
		return nil, fmt.Errorf(`couldn't find restvirt profile "%s"`, profile)
	}

	return NewClientFromConfig(config)
}
