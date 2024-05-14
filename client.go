package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/auth"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Config Config
}

// Org
// Deprecated
func (c *Client) Org(orgName string) *Client {
	clone := &Client{Config: c.Config}
	clone.Config.OrgName = orgName
	return clone
}

// WithApiKey
// Deprecated
func (c *Client) WithApiKey(apiKey string) *Client {
	clone := &Client{Config: c.Config}
	clone.Config.AccessTokenSource = auth.RawAccessTokenSource{AccessToken: apiKey}
	return clone
}

func (c *Client) CurrentUser() CurrentUser {
	return CurrentUser{Client: c}
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

func (c *Client) EnvInfraConfigurations() EnvInfraConfigurations {
	return EnvInfraConfigurations{Client: c}
}

func (c *Client) PipelineInfraConfigurations() PipelineInfraConfigurations {
	return PipelineInfraConfigurations{Client: c}
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

func (c *Client) BlockSyncs() BlockSyncs {
	return BlockSyncs{Client: c}
}

func (c *Client) CapabilityCopies() CapabilityCopies {
	return CapabilityCopies{Client: c}
}

func (c *Client) PipelineBlockSyncs() PipelineBlockSyncs {
	return PipelineBlockSyncs{Client: c}
}

func (c *Client) Apps() Apps {
	return Apps{Client: c}
}

func (c *Client) AppCapabilities() AppCapabilities {
	return AppCapabilities{Client: c}
}

func (c *Client) AppPipelineCapabilities() AppPipelineCapabilities {
	return AppPipelineCapabilities{Client: c}
}

func (c *Client) AppEnvs() AppEnvs {
	return AppEnvs{Client: c}
}

func (c *Client) EnvVariables() EnvVariables {
	return EnvVariables{Client: c}
}

func (c *Client) Builds() Builds { return Builds{Client: c} }

func (c *Client) BuildLogs() BuildLogs { return BuildLogs{Client: c} }

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

func (c *Client) WorkspaceModule() WorkspaceModule {
	return WorkspaceModule{Client: c}
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

func (c *Client) Events() Events {
	return Events{Client: c}
}

func (c *Client) Do(ctx context.Context, method string, relativePath string, query url.Values, headers map[string]string, body interface{}) (*http.Response, error) {
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

	u, err := c.Config.ConstructUrl(relativePath, query)
	if err != nil {
		return nil, fmt.Errorf("invalid request url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if err := c.Config.AddAuthorizationHeader(req.Context(), req.Header); err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: c.Config.CreateTransport(http.DefaultTransport),
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error issuing request: %w", err)
	}
	return res, nil
}

func (c *Client) CreateRequest(ctx context.Context, method string, relativePath string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := c.Config.ConstructUrl(relativePath, query)
	if err != nil {
		return nil, err
	}
	return http.NewRequestWithContext(ctx, method, u.String(), body)
}
