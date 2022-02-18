package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type PublicModules struct {
	Client *Client
}

func (m PublicModules) basePath() string {
	return fmt.Sprintf("orgs/%s/public-modules", m.Client.Config.OrgName)
}

func (m PublicModules) modulePath(moduleName string) string {
	return fmt.Sprintf("orgs/%s/public-modules/%s", m.Client.Config.OrgName, moduleName)
}

func (m PublicModules) List() ([]types.Module, error) {
	res, err := m.Client.Do(http.MethodGet, m.basePath(), nil, nil, nil)
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
	res, err := m.Client.Do(http.MethodGet, m.modulePath(moduleName), nil, nil, nil)
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
