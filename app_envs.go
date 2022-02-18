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
	return fmt.Sprintf("orgs/%s/apps/%d/envs/%s", e.Client.Config.OrgName, appId, envName)
}

func (e AppEnvs) Get(appId int64, envName string) (*types.AppEnv, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(appId, envName), nil, nil, nil)
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

func (e AppEnvs) Update(appId int64, envName string, version string) (*types.AppEnv, error) {
	rawPayload, _ := json.Marshal(map[string]interface{}{
		"version": version,
	})
	res, err := e.Client.Do(http.MethodPut, e.basePath(appId, envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updated types.AppEnv
	if err := response.ReadJson(res, &updated); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updated, nil
}
