package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type AppEnvs struct {
	Client *Client
}

func (e AppEnvs) Get(appId int, envName string) (*types.AppEnv, error) {
	res, err := e.Client.Do(http.MethodGet, path.Join("apps", strconv.Itoa(appId), "envs", envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var appEnv types.AppEnv
	if err := e.Client.ReadJsonResponse(res, &appEnv); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &appEnv, nil
}

func (e AppEnvs) Update(appId int, envName string, version string) (*types.AppEnv, error) {
	rawPayload, _ := json.Marshal(map[string]interface{}{
		"version": version,
	})
	res, err := e.Client.Do(http.MethodPut, path.Join("apps", strconv.Itoa(appId), "envs", envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updated types.AppEnv
	if err := e.Client.ReadJsonResponse(res, &updated); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updated, nil
}
