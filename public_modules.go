package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type PublicModules struct {
	Client *Client
}

func (m PublicModules) List() ([]types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, path.Join("public-modules"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var modules []types.Module
	if err := response.ReadJson(res, &modules); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return modules, nil
}

func (m PublicModules) Get(moduleName string) (*types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, path.Join("public-modules", moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var module types.Module
	if err := response.ReadJson(res, &module); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &module, nil
}
