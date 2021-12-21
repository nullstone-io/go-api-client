package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AppEnvs struct {
	Client *Client
}

func (e AppEnvs) basePath(appId int64, envName string) string {
	return fmt.Sprintf("apps/%d/envs/%s", appId, envName)
}

func (e AppEnvs) Get(appId int64, envName string) (*types.AppEnv, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(appId, envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.AppEnv](res)
}

func (e AppEnvs) Update(appId int64, envName string, version string) (*types.AppEnv, error) {
	rawPayload, _ := json.Marshal(map[string]interface{}{
		"version": version,
	})
	res, err := e.Client.Do(http.MethodPut, e.basePath(appId, envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.AppEnv](res)
}
