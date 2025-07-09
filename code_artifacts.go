package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type CodeArtifacts struct {
	Client *Client
}

func (s CodeArtifacts) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/apps/%d/envs/%d/artifacts", s.Client.Config.OrgName, stackId, appId, envId)
}

func (s CodeArtifacts) List(ctx context.Context, stackId, appId, envId int64) ([]types.CodeArtifact, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId, appId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.CodeArtifact](res)
}

type UpsertCodeArtifactInput struct {
	Version    string           `json:"version"`
	CommitInfo types.CommitInfo `json:"commitInfo"`
}

func (s CodeArtifacts) Upsert(ctx context.Context, stackId, appId, envId int64, version string, commitInfo types.CommitInfo) (*types.CodeArtifact, error) {
	input := UpsertCodeArtifactInput{
		Version:    version,
		CommitInfo: commitInfo,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPut, s.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.CodeArtifact](res)
}
