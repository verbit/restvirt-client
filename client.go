package restvirt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

type Client struct {
	Host       string
	Username   string
	Password   string
	HTTPClient *http.Client
}

type ClientConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	HostCA   string `yaml:"ca"`
}

type RestvirtRoundTripper struct {
	username string
	password string
	wrapped  http.RoundTripper
}

type ErrorMessage struct {
	Error string
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "not found"
}

func NewRoundTripper(wrapped http.RoundTripper, username string, password string) *RestvirtRoundTripper {
	return &RestvirtRoundTripper{
		username: username,
		password: password,
		wrapped:  wrapped,
	}
}

func (r *RestvirtRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(r.username, r.password)
	return r.wrapped.RoundTrip(req)
}

func NewClient(host string, username string, password string) (*Client, error) {
	c := Client{
		Host:       host,
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}

	return &c, nil
}

func NewClientFromConfig(config ClientConfig) (*Client, error) {
	var tlsConfig *tls.Config = nil
	if config.HostCA != "" {
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}

		hostCAPEM, err := base64.StdEncoding.DecodeString(config.HostCA)
		if err != nil {
			return nil, err
		}

		if ok := rootCAs.AppendCertsFromPEM(hostCAPEM); !ok {
			return nil, fmt.Errorf("couldn't append hosts's CA to the bundle")
		}

		tlsConfig = &tls.Config{
			RootCAs: rootCAs,
		}
	}

	var httpTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
	httpTransport = NewRoundTripper(httpTransport, config.Username, config.Password)
	client := &http.Client{Transport: httpTransport}

	c := Client{
		Host:       config.Host,
		Username:   config.Username,
		Password:   config.Password,
		HTTPClient: client,
	}

	return &c, nil
}

func NewClientFromEnvironment() (*Client, error) {
	configPath, exists := os.LookupEnv("RESTVIRT_PROFILE")
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

func (c *Client) doRequest(method string, body io.Reader, paths ...string) (*http.Response, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, err
	}
	ps := append([]string{u.Path}, paths...)
	u.Path = path.Join(ps...)
	request, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return c.HTTPClient.Do(request)
}

func hasValidStatusCode(resp *http.Response) bool {
	if resp == nil {
		return false
	}
	code := resp.StatusCode
	req := resp.Request
	if req == nil {
		return code == http.StatusOK
	}
	switch req.Method {
	case "POST":
		return code == http.StatusOK ||
			code == http.StatusCreated
	case "PUT":
		return code == http.StatusOK ||
			code == http.StatusCreated ||
			code == http.StatusNoContent
	default:
		return code == http.StatusOK
	}
}

func checkForErrors(resp *http.Response) error {
	if hasValidStatusCode(resp) {
		return nil
	}

	if resp.StatusCode == http.StatusNotFound {
		return &NotFoundError{}
	}

	var errorMessage ErrorMessage
	err := json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return err
	}
	return fmt.Errorf("restvirt client error: %s", errorMessage.Error)
}
