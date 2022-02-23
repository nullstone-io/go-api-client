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

func (e AppCapabilities) basePath(stackId, appId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/capabilities", e.Client.Config.OrgName, stackId, appId)
}

func (e AppCapabilities) capPath(stackId, appId, capId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/capabilities/%d", e.Client.Config.OrgName, stackId, appId, capId)
}

// List - GET /orgs/:orgName/apps/:app_id/capabilities
func (e AppCapabilities) List(stackId, appId int64) ([]types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(stackId, appId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var appCaps []types.Capability
	if err := response.ReadJson(res, &appCaps); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return appCaps, nil
}

// Get - GET /orgs/:orgName/apps/:app_id/capabilities/:id
func (e AppCapabilities) Get(stackId, appId, capId int64) (*types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.capPath(stackId, appId, capId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var appCap types.Capability
	if err := response.ReadJson(res, &appCap); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &appCap, nil
}

// Create - POST /orgs/:orgName/apps/:app_id/capabilities
func (e AppCapabilities) Create(stackId, appId int64, capability *types.Capability) (*types.Capability, error) {
	rawPayload, _ := json.Marshal(capability)
	res, err := e.Client.Do(http.MethodPost, e.basePath(stackId, appId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedCap types.Capability
	if err := response.ReadJson(res, &updatedCap); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedCap, nil
}

// Update - PUT/PATCH /orgs/:orgName/apps/:app_id/capabilities/:id
func (e AppCapabilities) Update(stackId, appId, capId int64, capability *types.Capability) (*types.Capability, error) {
	rawPayload, _ := json.Marshal(capability)
	res, err := e.Client.Do(http.MethodPut, e.capPath(stackId, appId, capId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedCap types.Capability
	if err := response.ReadJson(res, &updatedCap); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedCap, nil
}

// Destroy - DELETE /orgs/:orgName/apps/:app_id/capabilities/:id
func (e AppCapabilities) Destroy(stackId, appId, capId int64) (bool, error) {
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
