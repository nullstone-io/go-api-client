package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/ws"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type WorkspaceWorkflows struct {
	Client *Client
}

func (ww WorkspaceWorkflows) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflows", ww.Client.Config.OrgName, stackId, blockId, envId)
}

func (ww WorkspaceWorkflows) path(stackId, blockId, envId, workspaceWorkflowId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflows/%d", ww.Client.Config.OrgName, stackId, blockId, envId, workspaceWorkflowId)
}

func (ww WorkspaceWorkflows) activitiesPath(stackId, blockId, envId, workspaceWorkflowId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflow_activities/%d", ww.Client.Config.OrgName, stackId, blockId, envId, workspaceWorkflowId)
}

func (ww WorkspaceWorkflows) List(ctx context.Context, stackId, blockId, envId int64, page, perPage int) ([]types.WorkspaceWorkflow, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	res, err := ww.Client.Do(ctx, http.MethodGet, ww.basePath(stackId, blockId, envId), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceWorkflow](res)
}

func (ww WorkspaceWorkflows) Get(ctx context.Context, stackId, blockId, envId, workspaceWorkflowId int64) (*types.WorkspaceWorkflow, error) {
	res, err := ww.Client.Do(ctx, http.MethodGet, ww.path(stackId, blockId, envId, workspaceWorkflowId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceWorkflow](res)
}

func (ww WorkspaceWorkflows) GetActivities(ctx context.Context, stackId, blockId, envId, workspaceWorkflowId int64) (*types.WorkspaceWorkflowActivities, error) {
	res, err := ww.Client.Do(ctx, http.MethodGet, ww.activitiesPath(stackId, blockId, envId, workspaceWorkflowId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceWorkflowActivities](res)
}

func (ww WorkspaceWorkflows) WatchGet(ctx context.Context, stackId, blockId, envId, workspaceWorkflowId int64, retryFn ws.StreamerRetryFunc) (*types.WorkspaceWorkflow, <-chan types.StreamObject[types.WorkspaceWorkflowUpdate], error) {
	endpoint, headers, err := ww.Client.Config.ConstructWsEndpoint(ctx, ww.path(stackId, blockId, envId, workspaceWorkflowId))
	if err != nil {
		return nil, nil, err
	}
	return ws.StreamObject[types.WorkspaceWorkflow, types.WorkspaceWorkflowUpdate](ctx, endpoint, headers, retryFn)
}

type CreateWorkspaceWorkflowInput struct {
	Actions       []string  `json:"actions"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	Status        string    `json:"status"`
	StatusMessage string    `json:"statusMessage"`
	StatusAt      time.Time `json:"statusAt"`
}

// Deprecated
// Used for migrations; remove once on v3 engine
func (ww WorkspaceWorkflows) Create(ctx context.Context, stackId, blockId, envId int64, input CreateWorkspaceWorkflowInput) (*types.WorkspaceWorkflow, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := ww.Client.Do(ctx, http.MethodPost, ww.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceWorkflow](res)
}
