package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"log"
	"net/http"
)

type CreatePreviewEnvInput struct {
	Name string `json:"name"`
}

type PreviewEnvs struct {
	Client *Client
}

func (pe PreviewEnvs) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/preview_envs", pe.Client.Config.OrgName, stackId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/envs
func (pe PreviewEnvs) Create(stackId int64, env *CreatePreviewEnvInput) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := pe.Client.Do(http.MethodPost, pe.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	log.Printf("res: %v", res)
	log.Printf("err: %v", err)
	if err != nil {
		return nil, err
	}

	var updatedEnv types.Environment
	if err := response.ReadJson(res, &updatedEnv); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedEnv, nil
}
