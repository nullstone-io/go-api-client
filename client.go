package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Config Config
}

func (c *Client) Org(orgName string) *Client {
	cfg := c.Config
	cfg.OrgName = orgName
	return &Client{Config: cfg}
}

func (c *Client) Stacks() Stacks {
	return Stacks{Client: c}
}

func (c *Client) Environments() Environments {
	return Environments{Client: c}
}

func (c *Client) EnvironmentsByName() EnvironmentsByName {
	return EnvironmentsByName{Client: c}
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

func (c *Client) AppEnvs() AppEnvs {
	return AppEnvs{Client: c}
}

func (c *Client) Workspaces() Workspaces {
	return Workspaces{Client: c}
}

func (c *Client) RunConfigs() RunConfigs {
	return RunConfigs{Client: c}
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

func (c *Client) ReadJsonResponse(res *http.Response, obj interface{}) error {
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		raw, _ := ioutil.ReadAll(res.Body)
		return &HttpError{
			StatusCode: res.StatusCode,
			Status:     res.Status,
			Body:       string(raw),
		}
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(obj); err != nil {
		return fmt.Errorf("error decoding json body: %w", err)
	}
	return nil
}

func (c *Client) ReadFileResponse(res *http.Response, file io.Writer) error {
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		raw, _ := ioutil.ReadAll(res.Body)
		return &HttpError{
			StatusCode: res.StatusCode,
			Status:     res.Status,
			Body:       string(raw),
		}
	}

	_, err := io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("error reading file body: %w", err)
	}
	return nil
}

func (c *Client) VerifyResponse(res *http.Response) error {
	if res.StatusCode >= 400 {
		raw, _ := ioutil.ReadAll(res.Body)
		return &HttpError{
			StatusCode: res.StatusCode,
			Status:     res.Status,
			Body:       string(raw),
		}
	}
	return nil
}

func (c *Client) CreateRequest(method string, relativePath string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := c.Config.ConstructUrl(relativePath, query)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(method, u.String(), body)
}
