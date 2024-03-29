package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Modules struct {
	Client *Client
}

func (m Modules) basePath(orgName string) string {
	return fmt.Sprintf("orgs/%s/modules", orgName)
}

func (m Modules) path(orgName, moduleName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s", orgName, moduleName)
}

func (m Modules) List(ctx context.Context, orgName string) ([]types.Module, error) {
	res, err := m.Client.Do(ctx, http.MethodGet, m.basePath(orgName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Module](res)
}

func (m Modules) Get(ctx context.Context, orgName, moduleName string) (*types.Module, error) {
	res, err := m.Client.Do(ctx, http.MethodGet, m.path(orgName, moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Module](res)
}

func (m Modules) Create(ctx context.Context, orgName string, module *types.Module) error {
	rawPayload, _ := json.Marshal(module)
	res, err := m.Client.Do(ctx, http.MethodPost, m.basePath(orgName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}
	return response.Verify(res)
}
