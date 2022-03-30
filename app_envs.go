package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AppEnvs struct {
	Client *Client
}

func (e AppEnvs) basePath(stackId, appId int64, envName string) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%s", e.Client.Config.OrgName, stackId, appId, envName)
}

func (e AppEnvs) Get(stackId, appId int64, envName string) (*types.AppEnv, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(stackId, appId, envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var appEnv types.AppEnv
	if err := response.ReadJson(res, &appEnv); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &appEnv, nil
}
