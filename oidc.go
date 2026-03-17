package api

import (
	"context"
	"encoding/json"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Oidc struct {
	Client *Client
}

// GetDiscoveryDocument - GET /.well-known/openid-configuration
func (s Oidc) GetDiscoveryDocument(ctx context.Context) (*types.OidcDiscoveryDocument, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, ".well-known/openid-configuration", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.OidcDiscoveryDocument](res)
}

// GetPublicKeys - GET /.well-known/jwks.json
func (s Oidc) GetPublicKeys(ctx context.Context) (*json.RawMessage, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, ".well-known/jwks.json", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var keys json.RawMessage
	if err := response.ReadJson(res, &keys); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &keys, nil
}

// IssueToken - POST /token
func (s Oidc) IssueToken(ctx context.Context, input types.IssueTokenInput) (*types.IssueTokenOutput, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPost, "token", nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.IssueTokenOutput](res)
}
