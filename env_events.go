package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type EnvEvents struct {
	Client *Client
}

func (s EnvEvents) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/events", s.Client.Config.OrgName, stackId, envId)
}

func (s EnvEvents) eventPath(stackId, envId int64, eventUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/events/%s", s.Client.Config.OrgName, stackId, envId, eventUid)
}

// List - GET /org/:orgName/stacks/:stackId/envs/:envId/events
func (s EnvEvents) List(ctx context.Context, stackId, envId int64) ([]types.EnvEvent, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.EnvEvent](res)
}

// Create - POST /org/:orgName/stacks/:stackId/envs/:envId/events
func (s EnvEvents) Create(ctx context.Context, stackId, envId int64, event types.EnvEvent) (*types.EnvEvent, error) {
	rawPayload, _ := json.Marshal(event)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.EnvEvent](res)
}

// Get - GET /org/:orgName/stacks/:stackId/envs/:envId/events/:eventUid
func (s EnvEvents) Get(ctx context.Context, stackId, envId int64, eventUid uuid.UUID) (*types.EnvEvent, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.eventPath(stackId, envId, eventUid), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.EnvEvent](res)
}

// Update - PUT /org/:orgName/stacks/:stackId/envs/:envId/events/:eventUid
func (s EnvEvents) Update(ctx context.Context, stackId, envId int64, eventUid uuid.UUID, event types.EnvEvent) (*types.EnvEvent, error) {
	rawPayload, _ := json.Marshal(event)
	res, err := s.Client.Do(ctx, http.MethodPut, s.eventPath(stackId, envId, eventUid), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.EnvEvent](res)
}

// Delete - DELETE /org/:orgName/stacks/:stackId/envs/:envId/events/:eventUid
func (s EnvEvents) Delete(ctx context.Context, stackId, envId int64, eventUid uuid.UUID) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.eventPath(stackId, envId, eventUid), nil, nil, nil)
	if err != nil {
		return false, err
	}
	err = response.Verify(res)
	if err != nil {
		if response.IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
