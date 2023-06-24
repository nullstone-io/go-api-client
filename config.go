package api

import (
	"fmt"
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
	cfg.UseApiKey(os.Getenv(ApiKeyEnvVar))
	cfg.IsTraceEnabled = trace.IsEnabled()
	return cfg
}

type Config struct {
	BaseAddress    string
	IsTraceEnabled bool
	OrgName        string

	// AccessTokenSource provides a hook for authenticating requests
	// GetAccessToken() is performed on every request using http.RoundTripper
	AccessTokenSource auth.AccessTokenSource
}

func (c *Config) UseApiKey(apiKey string) {
	if apiKey == "" {
		c.AccessTokenSource = nil
	} else {
		c.AccessTokenSource = auth.RawAccessTokenSource{AccessToken: apiKey}
	}
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

func (c *Config) ConstructWsEndpoint(relativePath string) (string, http.Header, error) {
	endpoint, err := url.Parse(c.BaseAddress)
	if err != nil {
		return "", http.Header{}, fmt.Errorf("invalid url: %w", err)
	}
	endpoint.Scheme = strings.Replace(endpoint.Scheme, "http", "ws", 1)
	endpoint.Path = path.Join(endpoint.Path, relativePath)

	headers := http.Header{}
	if c.AccessTokenSource != nil {
		accessToken, err := c.AccessTokenSource.GetAccessToken()
		if err != nil {
			return "", nil, fmt.Errorf("unable to retrieve access token to authenticate websocket request: %w", err)
		}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	return endpoint.String(), headers, nil
}

func (c *Config) CreateTransport(baseTransport http.RoundTripper) http.RoundTripper {
	bt := baseTransport
	if c.IsTraceEnabled {
		bt = &trace.HttpTransport{BaseTransport: bt}
	}
	return &auth.AccessTokenSourceTransport{BaseTransport: bt, AccessTokenSource: c.AccessTokenSource}
}
