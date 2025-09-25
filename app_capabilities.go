package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type AppCapabilities struct {
	Client *Client
}

func (e AppCapabilities) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/capabilities", e.Client.Config.OrgName, stackId, appId, envId)
}

func (e AppCapabilities) capPath(stackId, appId, envId int64, capName string) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/capabilities/%s", e.Client.Config.OrgName, stackId, appId, envId, capName)
}

// List - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities
func (e AppCapabilities) List(ctx context.Context, stackId, appId, envId int64) (types.CapabilityConfigs, error) {
	res, err := e.Client.Do(ctx, http.MethodGet, e.basePath(stackId, appId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[types.CapabilityConfigs](res)
}

// Get - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities/:name
func (e AppCapabilities) Get(ctx context.Context, stackId, appId, envId int64, capName string) (*types.CapabilityConfig, error) {
	res, err := e.Client.Do(ctx, http.MethodGet, e.capPath(stackId, appId, envId, capName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.CapabilityConfig](res)
}

type CreateCapabilitiesInput struct {
	Capabilities []CreateCapabilityInput         `json:"capabilities"`
	Updates      []UpdateCapabilityOnCreateInput `json:"updates"`
	Blocks       []types.Block                   `json:"blocks"`
}

type CreateCapabilityInput struct {
	Name                string                  `json:"name"`
	ModuleSource        string                  `json:"moduleSource"`
	ModuleSourceVersion string                  `json:"moduleSourceVersion"`
	Namespace           string                  `json:"namespace"`
	Connections         types.ConnectionTargets `json:"connections"`
}

type UpdateCapabilityOnCreateInput struct {
	// Name is used to identify and cannot be changed
	Name string `json:"name"`
	// Namespace is modified in the create API endpoint
	Namespace string `json:"namespace"`
}

// Create - POST /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities
func (e AppCapabilities) Create(ctx context.Context, stackId, appId, envId int64, capabilities []CreateCapabilityInput, updates []UpdateCapabilityOnCreateInput, blocks []types.Block) (types.CapabilityConfigs, *http.Response, error) {
	input := CreateCapabilitiesInput{
		Capabilities: capabilities,
		Updates:      updates,
		Blocks:       blocks,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(ctx, http.MethodPost, e.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, res, err
	}

	result, err := response.ReadJsonVal[types.CapabilityConfigs](res)
	return result, res, err
}

type UpdateCapabilityInput struct {
	Namespace        string                  `json:"namespace"`
	ModuleSource     string                  `json:"moduleSource"`
	ModuleConstraint string                  `json:"moduleConstraint"`
	Connections      types.ConnectionTargets `json:"connections"`
}

// Update - PUT /orgs/:orgName/stacks/:stackId/apps/:appId/envs/:envId/capabilities/:name
func (e AppCapabilities) Update(ctx context.Context, stackId, appId, envId int64, capName string, input UpdateCapabilityInput) (*types.CapabilityConfig, *http.Response, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(ctx, http.MethodPut, e.capPath(stackId, appId, envId, capName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, res, err
	}

	result, err := response.ReadJsonPtr[types.CapabilityConfig](res)
	return result, res, err
}

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities/:name
func (e AppCapabilities) Destroy(ctx context.Context, stackId, appId, envId int64, capName string) (bool, error) {
	res, err := e.Client.Do(ctx, http.MethodDelete, e.capPath(stackId, appId, envId, capName), nil, nil, nil)
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
