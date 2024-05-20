package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceWorkflows struct {
	Client *Client
}

func (ww WorkspaceWorkflows) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflows", ww.Client.Config.OrgName, stackId, blockId, envId)
}

func (ww WorkspaceWorkflows) Create(ctx context.Context, stackId, blockId, envId int64, input types.CreateWorkspaceWorkflowPayload) (*types.WorkspaceWorkflow, error) {
	raw, _ := json.Marshal(input)
	res, err := ww.Client.Do(ctx, http.MethodPost, ww.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceWorkflow](res)
}
