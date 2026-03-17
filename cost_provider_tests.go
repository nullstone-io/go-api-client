package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
)

type CostProviderTests struct {
	Client *Client
}

func (s CostProviderTests) testExistingPath(costProviderId int64) string {
	return fmt.Sprintf("orgs/%s/cost_providers/%d/test", s.Client.Config.OrgName, costProviderId)
}

func (s CostProviderTests) testNewPath() string {
	return fmt.Sprintf("orgs/%s/cost_provider_tests", s.Client.Config.OrgName)
}

// TestExisting - GET /orgs/:orgName/cost_providers/:costProviderId/test
func (s CostProviderTests) TestExisting(ctx context.Context, costProviderId int64) error {
	res, err := s.Client.Do(ctx, http.MethodGet, s.testExistingPath(costProviderId), nil, nil, nil)
	if err != nil {
		return err
	}
	return response.Verify(res)
}

// TestNew - POST /orgs/:orgName/cost_provider_tests
func (s CostProviderTests) TestNew(ctx context.Context, providerId int64) error {
	rawPayload, _ := json.Marshal(map[string]any{"providerId": providerId})
	res, err := s.Client.Do(ctx, http.MethodPost, s.testNewPath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}
	return response.Verify(res)
}
