package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type AutomationCredentials struct {
	Client *Client
}

func (s AutomationCredentials) path(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/automation_credentials", s.Client.Config.OrgName, stackId, blockId, envId)
}

type AcquireAutomationCredentialsInput struct {
	ProviderType   string   `json:"providerType"`
	Purpose        string   `json:"purpose"`
	OutputNames    []string `json:"outputNames"`
	GcpOauthScopes []string `json:"gcpOauthScopes,omitempty"`
}

// Acquire - POST /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/automation_credentials
func (s AutomationCredentials) Acquire(ctx context.Context, stackId, blockId, envId int64, input AcquireAutomationCredentialsInput) (*types.OutputCredentials, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPost, s.path(stackId, blockId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.OutputCredentials](res)
}
