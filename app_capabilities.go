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

func (e AppCapabilities) basePath(appId int64) string {
	return fmt.Sprintf("apps/%d/capabilities", appId)
}

func (e AppCapabilities) capPath(appId, capId int64) string {
	return fmt.Sprintf("apps/%d/capabilities/%d", appId, capId)
}

// List - GET /orgs/:orgName/apps/:app_id/capabilities
func (e AppCapabilities) List(appId int64) ([]types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.basePath(appId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.JsonArray[types.Capability](res)
}

// Get - GET /orgs/:orgName/apps/:app_id/capabilities/:id
func (e AppCapabilities) Get(appId, capId int64) (*types.Capability, error) {
	res, err := e.Client.Do(http.MethodGet, e.capPath(appId, capId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Capability](res)
}

// Create - POST /orgs/:orgName/apps/:app_id/capabilities
func (e AppCapabilities) Create(appId int64, capability *types.Capability) (*types.Capability, error) {
	rawPayload, _ := json.Marshal(capability)
	res, err := e.Client.Do(http.MethodPost, e.basePath(appId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Capability](res)
}

// Update - PUT/PATCH /orgs/:orgName/apps/:app_id/capabilities/:id
func (e AppCapabilities) Update(appId, capId int64, capability *types.Capability) (*types.Capability, error) {
	rawPayload, _ := json.Marshal(capability)
	res, err := e.Client.Do(http.MethodPut, e.capPath(appId, capId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Capability](res)
}

// Destroy - DELETE /orgs/:orgName/apps/:app_id/capabilities/:id
func (e AppCapabilities) Destroy(appId, capId int64) (bool, error) {
	res, err := e.Client.Do(http.MethodDelete, e.capPath(appId, capId), nil, nil, nil)
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
