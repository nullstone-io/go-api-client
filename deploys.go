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

// DeployCreatePayload is the payload for creating a deploy
//
//		If `FromSource` is specified, we deploy from the latest commit on the configured branch for this environment
//		Otherwise, version is required
//		commitSha and reference are optional and get populated on the deploy no matter whether we are deploying
//	   fromSource or by version
type DeployCreatePayload struct {
	FromSource bool   `json:"fromSource"`
	CommitSha  string `json:"commitSha"`
	Version    string `json:"version"`
	Reference  string `json:"reference"`
}

func (d Deploys) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/deploys", d.Client.Config.OrgName, stackId, appId, envId)
}

func (d Deploys) path(stackId, appId, envId, deployId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/deploys/%d", d.Client.Config.OrgName, stackId, appId, envId, deployId)
}

func (d Deploys) Create(stackId, appId, envId int64, payload DeployCreatePayload) (*types.Deploy, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := d.Client.Do(http.MethodPost, d.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Deploy](res)
}

func (d Deploys) Get(stackId, appId, envId, deployId int64) (*types.Deploy, error) {
	res, err := d.Client.Do(http.MethodGet, d.path(stackId, appId, envId, deployId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Deploy](res)
}
