package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Modules struct {
	Client *Client
}

func (m Modules) basePath(orgName string) string {
	return fmt.Sprintf("orgs/%s/modules", orgName)
}

func (m Modules) path(orgName, moduleName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s", orgName, moduleName)
}

func (m Modules) List(ctx context.Context, orgName string) ([]types.Module, error) {
	res, err := m.Client.Do(ctx, http.MethodGet, m.basePath(orgName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Module](res)
}

func (m Modules) Get(ctx context.Context, orgName, moduleName string) (*types.Module, error) {
	res, err := m.Client.Do(ctx, http.MethodGet, m.path(orgName, moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Module](res)
}

type CreateModuleInput struct {
	Name          string   `json:"name"`
	FriendlyName  string   `json:"friendlyName"`
	Description   string   `json:"description"`
	IsPublic      bool     `json:"isPublic"`
	Category      string   `json:"category"`
	Subcategory   string   `json:"subcategory"`
	ProviderTypes []string `json:"providerTypes"`
	Platform      string   `json:"platform"`
	Subplatform   string   `json:"subplatform"`
	Type          string   `json:"type"`
	AppCategories []string `json:"appCategories"`
	SourceUrl     string   `json:"sourceUrl"`
	Status        string   `json:"status"`
}

func (m Modules) Create(ctx context.Context, orgName string, input CreateModuleInput) (*types.Module, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := m.Client.Do(ctx, http.MethodPost, m.basePath(orgName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Module](res)
}

type UpdateModuleInput struct {
	IsPublic  *bool   `json:"isPublic,omitempty"`
	SourceUrl *string `json:"sourceUrl,omitempty"`
	Status    *string `json:"status,omitempty"`
}

func (m Modules) Update(ctx context.Context, orgName string, moduleName string, input UpdateModuleInput) (*types.Module, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := m.Client.Do(ctx, http.MethodPut, m.path(orgName, moduleName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Module](res)
}
