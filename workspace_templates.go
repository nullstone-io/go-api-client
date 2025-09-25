package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type WorkspaceTemplates struct {
	Client *Client
}

func (wt WorkspaceTemplates) basePath(stackId, blockId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/workspace_template", wt.Client.Config.OrgName, stackId, blockId)
}

func (wt WorkspaceTemplates) connectionPath(stackId, blockId int64, connectionName string) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/workspace_template/connections/%s", wt.Client.Config.OrgName, stackId, blockId, connectionName)
}

func (wt WorkspaceTemplates) capabilitiesPath(stackId, blockId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/workspace_template/capabilities", wt.Client.Config.OrgName, stackId, blockId)
}

func (wt WorkspaceTemplates) capabilityPath(stackId, blockId int64, capabilityName string) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/workspace_template/capabilities/%s", wt.Client.Config.OrgName, stackId, blockId, capabilityName)
}

// Get - GET /orgs/:orgName/stacks/:stack_id/blocks/:id/workspace_template
func (wt WorkspaceTemplates) Get(ctx context.Context, stackId, blockId int64) (*types.WorkspaceTemplate, error) {
	res, err := wt.Client.Do(ctx, http.MethodGet, wt.basePath(stackId, blockId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceTemplate](res)
}

// Update - PUT /orgs/:orgName/stacks/:stack_id/blocks/:id/workspace_template
func (wt WorkspaceTemplates) Update(ctx context.Context, stackId, blockId int64, config types.WorkspaceTemplateConfig) (*types.WorkspaceTemplate, error) {
	rawPayload, _ := json.Marshal(config)
	res, err := wt.Client.Do(ctx, http.MethodPut, wt.basePath(stackId, blockId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceTemplate](res)
}

type UpdateTemplateConnectionInput struct {
	Target *types.ConnectionTarget `json:"target"`
}

// UpdateConnection - PUT /orgs/:orgName/stacks/:stack_id/blocks/:id/workspace_template/connections/:connectionName
func (wt WorkspaceTemplates) UpdateConnection(ctx context.Context, stackId, blockId int64, connectionName string, input UpdateTemplateConnectionInput) (bool, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := wt.Client.Do(ctx, http.MethodPut, wt.connectionPath(stackId, blockId, connectionName), nil, nil, json.RawMessage(rawPayload))
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

type CreateTemplateCapabilityInput struct {
	Capability          types.WorkspaceCapabilityTemplateConfig   `json:"capability"`
	UpdatedCapabilities []types.WorkspaceCapabilityTemplateConfig `json:"updatedCapabilities"`
	BlocksToCreate      []CreateBlockInput                        `json:"blocksToCreate"`
}

// CreateCapability - POST /orgs/:orgName/stacks/:stack_id/blocks/:id/workspace_template/capabilities
func (wt WorkspaceTemplates) CreateCapability(ctx context.Context, stackId, blockId int64, input CreateTemplateCapabilityInput) (bool, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := wt.Client.Do(ctx, http.MethodPost, wt.capabilitiesPath(stackId, blockId), nil, nil, json.RawMessage(rawPayload))
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

// RemoveCapability - DELETE /orgs/:orgName/stacks/:stack_id/blocks/:id/workspace_template/capabilities/:capabilityName
func (wt WorkspaceTemplates) RemoveCapability(ctx context.Context, stackId, blockId int64, capabilityName string) (bool, error) {
	res, err := wt.Client.Do(ctx, http.MethodDelete, wt.capabilityPath(stackId, blockId, capabilityName), nil, nil, nil)
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
