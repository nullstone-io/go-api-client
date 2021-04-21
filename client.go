package api

import (
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

func (c *Client) AutogenSubdomains() AutogenSubdomains {
	return AutogenSubdomains{Client: c}
}

func (c *Client) AutogenSubdomainsDelegation() AutogenSubdomainsDelegation {
	return AutogenSubdomainsDelegation{Client: c}
}

func (c *Client) Domains() Domains {
	return Domains{Client: c}
}

func (c *Client) Subdomains() Subdomains {
	return Subdomains{Client: c}
}

func (c *Client) Do(method string, relativePath string, query url.Values, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := c.CreateRequest(method, relativePath, query, body)
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
