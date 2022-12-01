package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Modules struct {
	Client *Client
}

func (m Modules) basePath() string {
	return fmt.Sprintf("orgs/%s/modules", m.Client.Config.OrgName)
}

func (m Modules) path(moduleName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s", m.Client.Config.OrgName, moduleName)
}

func (m Modules) List() ([]types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, m.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Module](res)
}

func (m Modules) Get(moduleName string) (*types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, m.path(moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Module](res)
}

func (m Modules) Create(module *types.Module) error {
	rawPayload, _ := json.Marshal(module)
	res, err := m.Client.Do(http.MethodPost, m.basePath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}
	return response.Verify(res)
}
