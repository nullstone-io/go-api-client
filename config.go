package api

import (
	"fmt"
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
	cfg := Config{
		BaseAddress: DefaultAddress,
		ApiKey:      os.Getenv(ApiKeyEnvVar),
	}
	if val := os.Getenv(AddressEnvVar); val != "" {
		cfg.BaseAddress = val
	}
	cfg.IsTraceEnabled = trace.IsEnabled()
	return cfg
}

type Config struct {
	BaseAddress    string
	ApiKey         string
	IsTraceEnabled bool
	OrgName        string
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
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	return endpoint.String(), headers, nil
}

func (c *Config) CreateTransport(baseTransport http.RoundTripper) http.RoundTripper {
	bt := baseTransport
	if c.IsTraceEnabled {
		bt = &trace.HttpTransport{BaseTransport: bt}
	}
	return &apiKeyTransport{BaseTransport: bt, ApiKey: c.ApiKey}
}

var _ http.RoundTripper = &apiKeyTransport{}

type apiKeyTransport struct {
	BaseTransport http.RoundTripper
	ApiKey        string
}

func (t *apiKeyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer "+t.ApiKey)
	return t.BaseTransport.RoundTrip(r)
}
