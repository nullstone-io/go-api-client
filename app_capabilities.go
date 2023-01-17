package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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

func (e AppCapabilities) basePath(stackId, appId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/capabilities", e.Client.Config.OrgName, stackId, appId)
}

func (e AppCapabilities) nullfireBasePath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/capabilities", e.Client.Config.OrgName, stackId, workspaceUid)
}

func (e AppCapabilities) capPath(stackId, appId, capId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/capabilities/%d", e.Client.Config.OrgName, stackId, appId, capId)
}

func (e AppCapabilities) nullfireCapPath(stackId int64, workspaceUid uuid.UUID, capId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/capabilities/%d", e.Client.Config.OrgName, stackId, workspaceUid, capId)
}

// List - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities
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

// Get - GET /orgs/:orgName/stacks/:stackId/apps/:app_id/capabilities/:id
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

// Create - POST /orgs/:orgName/stacks/:stackId/workspaces/:workspace_uid/capabilities
func (e AppCapabilities) Create(stackId int64, workspaceUid uuid.UUID, capabilities []*types.Capability, blocks []*types.Block) ([]types.WorkspaceChange, error) {
	input := CreateCapabilitiesInput{
		Capabilities: capabilities,
		Blocks:       blocks,
	}
	rawPayload, _ := json.Marshal(input)
	res, err := e.Client.Do(http.MethodPost, e.nullfireBasePath(stackId, workspaceUid), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var changes []types.WorkspaceChange
	if err := response.ReadJson(res, &changes); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changes, nil
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stackId/workspaces/:workspace_uid/capabilities/:id/variables
func (e AppCapabilities) Update(stackId int64, workspaceUid uuid.UUID, capId int64, variables []*types.VariableInput) ([]types.WorkspaceChange, error) {
	rawPayload, _ := json.Marshal(variables)
	res, err := e.Client.Do(http.MethodPut, fmt.Sprintf("%s/variables", e.nullfireCapPath(stackId, workspaceUid, capId)), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var changes []types.WorkspaceChange
	if err := response.ReadJson(res, &changes); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changes, nil
}

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/workspaces/:workspace_uid/capabilities/:id
func (e AppCapabilities) Destroy(stackId int64, workspaceUid uuid.UUID, capId int64) ([]types.WorkspaceChange, error) {
	res, err := e.Client.Do(http.MethodDelete, e.nullfireCapPath(stackId, workspaceUid, capId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var changes []types.WorkspaceChange
	if err := response.ReadJson(res, &changes); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changes, nil
}
