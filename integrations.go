package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Integrations struct {
	Client *Client
}

func (s Integrations) basePath() string {
	return fmt.Sprintf("/orgs/%s/integrations", s.Client.Config.OrgName)
}
func (s Integrations) integrationPath(id int64) string {
	return fmt.Sprintf("/orgs/%s/integrations/%d", s.Client.Config.OrgName, id)
}
func (s Integrations) oauthPath(id int64) string {
	return fmt.Sprintf("/orgs/%s/integrations/%d/oauth", s.Client.Config.OrgName, id)
}
func (s Integrations) statusPath(id int64) string {
	return fmt.Sprintf("/orgs/%s/integrations/%d/status", s.Client.Config.OrgName, id)
}

// List - GET /orgs/:orgName/integrations
func (s Integrations) List(ctx context.Context) ([]types.Integration, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Integration](res)
}

// Create - POST /orgs/:orgName/integrations
func (s Integrations) Create(ctx context.Context, integration types.Integration) (*types.Integration, error) {
	rawPayload, _ := json.Marshal(integration)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Integration](res)
}

// Get - GET /orgs/:orgName/integrations/:integrationId
func (s Integrations) Get(ctx context.Context, id int64) (*types.Integration, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.integrationPath(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Integration](res)
}

// Delete - DELETE /orgs/:orgName/integrations/:integrationId
func (s Integrations) Delete(ctx context.Context, id int64) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.integrationPath(id), nil, nil, nil)
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

type GetOauthResponse struct {
	Url string `json:"url"`
}

// GetOauth - GET /orgs/:orgName/integrations/:integrationId/oauth
func (s Integrations) GetOauth(ctx context.Context, id int64) (*GetOauthResponse, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.oauthPath(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[GetOauthResponse](res)
}

// DeleteOauth - DELETE /orgs/:orgName/integrations/:integrationId/oauth
func (s Integrations) DeleteOauth(ctx context.Context, id int64) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.oauthPath(id), nil, nil, nil)
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

// GetStatus - GET /orgs/:orgName/integrations/:integrationId/status
func (s Integrations) GetStatus(ctx context.Context, id int64) (*types.IntegrationStatus, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.statusPath(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.IntegrationStatus](res)
}
