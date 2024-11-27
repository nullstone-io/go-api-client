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

type WorkspaceOutputCredentials struct {
	Client *Client
}

func (c WorkspaceOutputCredentials) path(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/output-credentials", c.Client.Config.OrgName, stackId, workspaceUid)
}

type GenerateCredentialsInput struct {
	Provider       string   `json:"provider"`
	OutputNames    []string `json:"outputNames"`
	GcpOauthScopes []string `json:"gcpOauthScopes,omitempty"`
}

// Create - POST /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/output-credentials
func (c WorkspaceOutputCredentials) Create(ctx context.Context, stackId int64, workspaceUid uuid.UUID, input GenerateCredentialsInput) (*types.OutputCredentials, error) {
	rawInput, _ := json.Marshal(input)
	res, err := c.Client.Do(ctx, http.MethodPost, c.path(stackId, workspaceUid), nil, nil, json.RawMessage(rawInput))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.OutputCredentials](res)
}
