package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type Modules struct {
	Client *Client
}

func (m Modules) List() ([]types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, path.Join("modules"), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.JsonArray[types.Module](res)
}

func (m Modules) Get(moduleName string) (*types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, path.Join("modules", moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Module](res)
}

func (m Modules) Create(module *types.Module) error {
	rawPayload, _ := json.Marshal(module)
	res, err := m.Client.Do(http.MethodPost, path.Join("modules"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}

	return response.Verify(res)
}
