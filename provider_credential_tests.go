package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
)

type ProviderCredentialTests struct {
	Client *Client
}

func (s ProviderCredentialTests) testSavedPath(providerName string) string {
	return fmt.Sprintf("orgs/%s/providers/%s/credential_test", s.Client.Config.OrgName, providerName)
}

func (s ProviderCredentialTests) testNewPath() string {
	return fmt.Sprintf("orgs/%s/provider_credential_tests", s.Client.Config.OrgName)
}

type TestNewProviderCredentialsInput struct {
	ProviderType string          `json:"providerType"`
	ProviderId   string          `json:"providerId"`
	Credentials  json.RawMessage `json:"credentials"`
}

// TestSaved - GET /orgs/:orgName/providers/:providerName/credential_test
func (s ProviderCredentialTests) TestSaved(ctx context.Context, providerName string) error {
	res, err := s.Client.Do(ctx, http.MethodGet, s.testSavedPath(providerName), nil, nil, nil)
	if err != nil {
		return err
	}
	return response.Verify(res)
}

// TestNew - POST /orgs/:orgName/provider_credential_tests
func (s ProviderCredentialTests) TestNew(ctx context.Context, input TestNewProviderCredentialsInput) error {
	rawPayload, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPost, s.testNewPath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}
	return response.Verify(res)
}
