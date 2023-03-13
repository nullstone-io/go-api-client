package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AppPipelineCapabilities struct {
	Client *Client
}

func (e AppPipelineCapabilities) basePath(stackId, appId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/capabilities", e.Client.Config.OrgName, stackId, appId)
}

func (e AppPipelineCapabilities) capPath(stackId, appId, capId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/capabilities/%d", e.Client.Config.OrgName, stackId, appId, capId)
}

// List - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities
func (e AppPipelineCapabilities) List(stackId, appId int64) ([]types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(stackId, appId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Capability](res)
}

// Get - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities/:id
func (e AppPipelineCapabilities) Get(stackId, appId, capId int64) (*types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.capPath(stackId, appId, capId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[*types.Capability](res)
}

// Create - POST /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities
func (e AppPipelineCapabilities) Create(stackId, appId int64, capabilities []*types.Capability, blocks []*types.Block) ([]*types.Capability, error) {
	input := CreateCapabilitiesInput{
		Capabilities: capabilities,
		Blocks:       blocks,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(http.MethodPost, e.basePath(stackId, appId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]*types.Capability](res)
}

// Replace - PUT /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities
func (e AppPipelineCapabilities) Replace(stackId, appId int64, capabilities []*types.Capability, blocks []*types.Block) ([]*types.Capability, error) {
	input := CreateCapabilitiesInput{
		Capabilities: capabilities,
		Blocks:       blocks,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(http.MethodPut, e.basePath(stackId, appId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]*types.Capability](res)
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities/:id
func (e AppPipelineCapabilities) Update(stackId, appId, capId int64, capability *types.Capability) (*types.Capability, error) {
	rawPayload, _ := json.Marshal(capability)
	res, err := e.Client.Do(http.MethodPut, e.capPath(stackId, appId, capId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[*types.Capability](res)
}

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities/:id
func (e AppPipelineCapabilities) Destroy(stackId, appId, capId int64) (bool, error) {
	res, err := e.Client.Do(http.MethodDelete, e.capPath(stackId, appId, capId), nil, nil, nil)
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
