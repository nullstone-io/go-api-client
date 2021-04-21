package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type AppEnvs struct {
	Client *Client
}

func (e AppEnvs) Get(appName, envName string) (*types.ApplicationEnvironment, error) {
	res, err := e.Client.Do(http.MethodGet, path.Join("apps", appName, "envs", envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var appEnv types.ApplicationEnvironment
	if err := e.Client.ReadJsonResponse(res, &appEnv); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &appEnv, nil
}

func (e AppEnvs) Update(appEnv types.ApplicationEnvironment) (*types.ApplicationEnvironment, error) {
	res, err := e.Client.Do(http.MethodGet, path.Join("apps", appEnv.AppName, "envs", appEnv.EnvName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var updated types.ApplicationEnvironment
	if err := e.Client.ReadJsonResponse(res, &updated); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updated, nil
}
