package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AppCapabilities struct {
	Client *Client
}

type CreateCapabilitiesInput struct {
	Capabilities []*types.Capability `json:"capabilities"`
	Blocks       []*types.Block      `json:"blocks"`
}

func (e AppCapabilities) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/capabilities", e.Client.Config.OrgName, stackId, appId, envId)
}

func (e AppCapabilities) capPath(stackId, appId, envId, capId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/capabilities/%d", e.Client.Config.OrgName, stackId, appId, envId, capId)
}

// List - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities
func (e AppCapabilities) List(stackId, appId, envId int64) ([]types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(stackId, appId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Capability](res)
}

// Get - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities/:id
func (e AppCapabilities) Get(stackId, appId, envId, capId int64) (*types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.capPath(stackId, appId, envId, capId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Capability](res)
}

// Create - POST /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities
func (e AppCapabilities) Create(stackId, appId, envId int64, capabilities []*types.Capability, blocks []*types.Block) ([]*types.Capability, error) {
	input := CreateCapabilitiesInput{
		Capabilities: capabilities,
		Blocks:       blocks,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(http.MethodPost, e.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]*types.Capability](res)
}

// Replace - PUT /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities
func (e AppCapabilities) Replace(stackId, appId, envId int64, capabilities []*types.Capability, blocks []*types.Block) ([]*types.Capability, error) {
	input := CreateCapabilitiesInput{
		Capabilities: capabilities,
		Blocks:       blocks,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(http.MethodPut, e.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]*types.Capability](res)
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities/:id
func (e AppCapabilities) Update(stackId, appId, envId, capId int64, capability *types.Capability) (*types.Capability, error) {
	rawPayload, _ := json.Marshal(capability)
	res, err := e.Client.Do(http.MethodPut, e.capPath(stackId, appId, envId, capId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Capability](res)
}

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/apps/:app_id/envs/:env_id/capabilities/:id
func (e AppCapabilities) Destroy(stackId, appId, envId, capId int64) (bool, error) {
	res, err := e.Client.Do(http.MethodDelete, e.capPath(stackId, appId, envId, capId), nil, nil, nil)
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
