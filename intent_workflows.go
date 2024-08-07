package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/ws"
	"net/http"
)

type IntentWorkflows struct {
	Client *Client
}

func (s IntentWorkflows) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/intent_workflows", s.Client.Config.OrgName, stackId, envId)
}

func (s IntentWorkflows) path(stackId, intentWorkflowId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/intent_workflows/%d", s.Client.Config.OrgName, stackId, intentWorkflowId)
}

func (s IntentWorkflows) List(ctx context.Context, stackId, envId int64) ([]types.IntentWorkflow, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.IntentWorkflow](res)
}

func (s IntentWorkflows) Get(ctx context.Context, stackId, intentWorkflowId int64) (*types.IntentWorkflow, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.path(stackId, intentWorkflowId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.IntentWorkflow](res)
}

func (s IntentWorkflows) WatchGet(ctx context.Context, stackId, intentWorkflowId int64, retryFn ws.StreamerRetryFunc) (*types.IntentWorkflow, <-chan types.StreamObject[types.IntentWorkflowUpdate], error) {
	endpoint, headers, err := s.Client.Config.ConstructWsEndpoint(ctx, s.path(stackId, intentWorkflowId))
	if err != nil {
		return nil, nil, err
	}
	return ws.StreamObject[types.IntentWorkflow, types.IntentWorkflowUpdate](ctx, endpoint, headers, retryFn)
}
