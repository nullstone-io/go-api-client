package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type CostProviders struct {
	Client *Client
}

func (s CostProviders) basePath() string {
	return fmt.Sprintf("orgs/%s/cost_providers", s.Client.Config.OrgName)
}

func (s CostProviders) providerPath(id int64) string {
	return fmt.Sprintf("orgs/%s/cost_providers/%d", s.Client.Config.OrgName, id)
}

// List - GET /orgs/:orgName/cost_providers
func (s CostProviders) List(ctx context.Context) ([]*types.CostProvider, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]*types.CostProvider](res)
}

// Get - GET /orgs/:orgName/cost_providers/:cost_provider_id
func (s CostProviders) Get(ctx context.Context, id int64) (*types.CostProvider, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.providerPath(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.CostProvider](res)
}

// Create - POST /orgs/:orgName/cost_providers
func (s CostProviders) Create(ctx context.Context, providerId int64) (*types.CostProvider, error) {
	rawPayload, _ := json.Marshal(map[string]any{"providerId": providerId})
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.CostProvider](res)
}

// Destroy - DELETE /orgs/:orgName/providers/:name
func (s CostProviders) Destroy(ctx context.Context, id int64) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.providerPath(id), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := response.Verify(res); response.IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
