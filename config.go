package api

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gopkg.in/nullstone-io/go-api-client.v0/auth"
	"gopkg.in/nullstone-io/go-api-client.v0/trace"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

var (
	ApiKeyEnvVar   = "NULLSTONE_API_KEY"
	AddressEnvVar  = "NULLSTONE_ADDR"
	DefaultAddress = "https://api.nullstone.io"
)

func DefaultConfig() Config {
	cfg := Config{BaseAddress: DefaultAddress}
	if val := os.Getenv(AddressEnvVar); val != "" {
		cfg.BaseAddress = val
	}
	if apiKey := os.Getenv(ApiKeyEnvVar); apiKey != "" {
		cfg.AccessTokenSource = auth.RawAccessTokenSource{AccessToken: apiKey}
	}
	cfg.IsTraceEnabled = trace.IsEnabled()
	return cfg
}

type Config struct {
	BaseAddress    string
	IsTraceEnabled bool
	OrgName        string

	InsecureSkipVerify bool

	// AccessTokenSource provides a hook for authenticating requests
	// GetAccessToken() is performed on every request using http.RoundTripper
	AccessTokenSource auth.AccessTokenSource
}

func (c *Config) ConstructUrl(relativePath string, query url.Values) (*url.URL, error) {
	u, err := url.Parse(c.BaseAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid nullstone API base address (%s): %w", c.BaseAddress, err)
	}
	u.Path = path.Join(u.Path, relativePath)
	if query != nil {
		u.RawQuery = query.Encode()
	}
	return u, nil
}

func (c *Config) ConstructWsEndpoint(ctx context.Context, relativePath string) (string, http.Header, error) {
	endpoint, err := url.Parse(c.BaseAddress)
	if err != nil {
		return "", http.Header{}, fmt.Errorf("invalid url: %w", err)
	}
	endpoint.Scheme = strings.Replace(endpoint.Scheme, "http", "ws", 1)
	endpoint.Path = path.Join(endpoint.Path, relativePath)

	headers := http.Header{}
	if err := c.AddAuthorizationHeader(ctx, headers); err != nil {
		return "", nil, err
	}
	return endpoint.String(), headers, nil
}

func (c *Config) AddAuthorizationHeader(ctx context.Context, headers http.Header) error {
	if c.AccessTokenSource != nil {
		accessToken, err := c.AccessTokenSource.GetAccessToken(ctx, c.OrgName)
		if err != nil {
			return fmt.Errorf("unable to retrieve access token to authenticate request: %w", err)
		}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	return nil
}

func (c *Config) CreateTransport() http.RoundTripper {
	baseTransport := http.DefaultTransport.(*http.Transport)
	if c.InsecureSkipVerify {
		baseTransport = baseTransport.Clone()
		baseTransport.TLSClientConfig.InsecureSkipVerify = true
	}
	transport := otelhttp.NewTransport(baseTransport)
	if c.IsTraceEnabled {
		return &trace.HttpTransport{BaseTransport: transport}
	}
	return transport
}
