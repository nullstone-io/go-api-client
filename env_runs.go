package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type EnvRuns struct {
	Client *Client
}

func (er EnvRuns) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/runs", er.Client.Config.OrgName, stackId, envId)
}

func (er EnvRuns) Create(stackId, envId int64, input types.CreateEnvRunInput) ([]types.Run, error) {
	raw, _ := json.Marshal(input)
	res, err := er.Client.Do(http.MethodPost, er.basePath(stackId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	var runs []types.Run
	if err := response.ReadJson(res, &runs); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return runs, nil
}
