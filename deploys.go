package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Deploys struct {
	Client *Client
}


func (d Deploys) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/deploys", d.Client.Config.OrgName, stackId, appId, envId)
}

func (d Deploys) List(stackId, appId, envId int64) ([]types.Deploy, error) {
	res, err := d.Client.Do(http.MethodGet, d.basePath(stackId, appId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var deploys []types.Deploy
	if err := response.ReadJson(res, &deploys); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return deploys, nil
}

func (d Deploys) Create(stackId, appId, envId int64, version string) (*types.Deploy, error) {
	rawPayload, _ := json.Marshal(map[string]interface{}{
		"version": version,
	})
	res, err := d.Client.Do(http.MethodPut, d.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updated types.Deploy
	if err := response.ReadJson(res, &updated); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updated, nil
}
