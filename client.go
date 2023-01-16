package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Config Config
}

// WithApiKey returns a copy configured with an API key
func (c *Client) WithApiKey(apiKey string) *Client {
	clone := *c
	clone.Config.ApiKey = apiKey
	return &clone
}

func (c *Client) Org(orgName string) *Client {
	cfg := c.Config
	cfg.OrgName = orgName
	return &Client{Config: cfg}
}

func (c *Client) Organizations() Organizations {
	return Organizations{Client: c}
}

func (c *Client) Stacks() Stacks {
	return Stacks{Client: c}
}

func (c *Client) StacksByName() StacksByName {
	return StacksByName{Client: c}
}

func (c *Client) Environments() Environments {
	return Environments{Client: c}
}

func (c *Client) PreviewEnvs() PreviewEnvs {
	return PreviewEnvs{Client: c}
}

func (c *Client) EnvironmentsByName() EnvironmentsByName {
	return EnvironmentsByName{Client: c}
}

func (c *Client) EnvRuns() EnvRuns {
	return EnvRuns{Client: c}
}

func (c *Client) Providers() Providers {
	return Providers{Client: c}
}

func (c *Client) ProviderCredentials() ProviderCredentials {
	return ProviderCredentials{Client: c}
}

func (c *Client) Blocks() Blocks {
	return Blocks{Client: c}
}

func (c *Client) Apps() Apps {
	return Apps{Client: c}
}

func (c *Client) AppCapabilities() AppCapabilities {
	return AppCapabilities{Client: c}
}

func (c *Client) AppEnvs() AppEnvs {
	return AppEnvs{Client: c}
}

func (c *Client) EnvVariables() EnvVariables {
	return EnvVariables{Client: c}
}

func (c *Client) Deploys() Deploys {
	return Deploys{Client: c}
}

func (c *Client) DeployLogs() DeployLogs {
	return DeployLogs{Client: c}
}

func (c *Client) Workspaces() Workspaces {
	return Workspaces{Client: c}
}

func (c *Client) WorkspaceChanges() WorkspaceChanges {
	return WorkspaceChanges{Client: c}
}

func (c *Client) WorkspaceOutputs() WorkspaceOutputs {
	return WorkspaceOutputs{Client: c}
}

func (c *Client) WorkspaceVariables() WorkspaceVariables {
	return WorkspaceVariables{Client: c}
}

func (c *Client) Runs() Runs {
	return Runs{Client: c}
}

func (c *Client) RunConfigs() RunConfigs {
	return RunConfigs{Client: c}
}

func (c *Client) PromotionConfigs() PromotionConfigs {
	return PromotionConfigs{Client: c}
}

func (c *Client) RunLogs() RunLogs {
	return RunLogs{Client: c}
}

func (c *Client) AutogenSubdomain() AutogenSubdomain {
	return AutogenSubdomain{Client: c}
}

func (c *Client) AutogenSubdomainDelegation() AutogenSubdomainDelegation {
	return AutogenSubdomainDelegation{Client: c}
}

func (c *Client) Domains() Domains {
	return Domains{Client: c}
}

func (c *Client) Subdomains() Subdomains {
	return Subdomains{Client: c}
}

func (c *Client) Modules() Modules {
	return Modules{Client: c}
}

func (c *Client) ModuleVersions() ModuleVersions {
	return ModuleVersions{Client: c}
}

func (c *Client) Do(method string, relativePath string, query url.Values, headers map[string]string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if jrm, ok := body.(json.RawMessage); ok {
		bodyReader = bytes.NewReader(jrm)
		if headers == nil {
			headers = map[string]string{}
		}
		if headers["Content-Type"] == "" {
			headers["Content-Type"] = "application/json"
		}
	} else {
		bodyReader, _ = body.(io.Reader)
	}

	req, err := c.CreateRequest(method, relativePath, query, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	httpClient := &http.Client{
		Transport: c.Config.CreateTransport(http.DefaultTransport),
	}
	return httpClient.Do(req)
}

func (c *Client) CreateRequest(method string, relativePath string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := c.Config.ConstructUrl(relativePath, query)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(method, u.String(), body)
}
