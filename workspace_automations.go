package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// WorkspaceAutomations is the API resource for workspace automation: the per-workspace flags that
// decide whether a block participates in an env, whether the env launches on its own, and whether a
// push deploys an app. Block-level defaults seed each workspace record the first time it is read.
type WorkspaceAutomations struct {
	Client *Client
}

func (w WorkspaceAutomations) defaultsPath(stackId, blockId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/automation", w.Client.Config.OrgName, stackId, blockId)
}

func (w WorkspaceAutomations) stackDefaultsPath(stackId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/automation_defaults", w.Client.Config.OrgName, stackId)
}

func (w WorkspaceAutomations) workspacePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/automation", w.Client.Config.OrgName, stackId, blockId, envId)
}

func (w WorkspaceAutomations) envPath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/automations", w.Client.Config.OrgName, stackId, envId)
}

func (w WorkspaceAutomations) sourcePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/automations/source", w.Client.Config.OrgName, stackId, envId)
}

// UpdateBlockAutomationDefaultsInput replaces the writable fields of a block's automation defaults.
// Repo and RepoUrl are not writable here: they follow the block's repository.
type UpdateBlockAutomationDefaultsInput struct {
	IsEnabled     bool   `json:"isEnabled"`
	AutoDeploy    bool   `json:"autoDeploy"`
	DefaultBranch string `json:"defaultBranch"`
}

// UpdateWorkspaceAutomationInput replaces the writable fields of one workspace's automation.
// BlockId is only read by ReplaceByEnv, where each entry names its block; the single-workspace
// Update takes the block from the path and ignores it.
// The branch or pull request a workspace follows is not writable here: see SetSource.
type UpdateWorkspaceAutomationInput struct {
	BlockId    int64 `json:"blockId,omitempty"`
	IsEnabled  bool  `json:"isEnabled"`
	AutoDeploy bool  `json:"autoDeploy"`
}

// SetAutomationSourceInput names what every application of one repository follows in an env:
// a branch or a pull request, exactly one of the two.
type SetAutomationSourceInput struct {
	RepoUrl       string `json:"repoUrl"`
	TrackedBranch string `json:"trackedBranch"`
	PullRequestId *int64 `json:"pullRequestId"`
}

// GetDefaults - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/automation
// The server seeds defaults for a block that has none, so nil only means the block does not exist.
func (w WorkspaceAutomations) GetDefaults(ctx context.Context, stackId, blockId int64) (*types.BlockAutomationDefaults, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.defaultsPath(stackId, blockId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.BlockAutomationDefaults](res)
}

// ListDefaults - GET /orgs/:orgName/stacks/:stackId/automation_defaults
// Returns the automation defaults of every block in the stack, ordered by block id, seeding any that do not exist yet.
func (w WorkspaceAutomations) ListDefaults(ctx context.Context, stackId int64) ([]types.BlockAutomationDefaults, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.stackDefaultsPath(stackId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.BlockAutomationDefaults](res)
}

// UpdateDefaults - PUT /orgs/:orgName/stacks/:stackId/blocks/:blockId/automation
func (w WorkspaceAutomations) UpdateDefaults(ctx context.Context, stackId, blockId int64, input UpdateBlockAutomationDefaultsInput) (*types.BlockAutomationDefaults, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := w.Client.Do(ctx, http.MethodPut, w.defaultsPath(stackId, blockId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.BlockAutomationDefaults](res)
}

// Get - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/automation
// The server seeds the workspace record from the block defaults when none is stored,
// so nil only means the block or env does not exist.
func (w WorkspaceAutomations) Get(ctx context.Context, stackId, blockId, envId int64) (*types.WorkspaceAutomation, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.workspacePath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceAutomation](res)
}

// Update - PUT /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/automation
func (w WorkspaceAutomations) Update(ctx context.Context, stackId, blockId, envId int64, input UpdateWorkspaceAutomationInput) (*types.WorkspaceAutomation, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := w.Client.Do(ctx, http.MethodPut, w.workspacePath(stackId, blockId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceAutomation](res)
}

// ListByEnv - GET /orgs/:orgName/stacks/:stackId/envs/:envId/automations
// Returns one record per block in the stack, seeding any that do not exist yet.
func (w WorkspaceAutomations) ListByEnv(ctx context.Context, stackId, envId int64) ([]types.WorkspaceAutomation, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.envPath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceAutomation](res)
}

// ReplaceByEnv - PUT /orgs/:orgName/stacks/:stackId/envs/:envId/automations
// Writes every entry in the payload (each must name its BlockId) and leaves blocks that are
// not in the payload untouched. The response is the full list for the env, as ListByEnv returns it.
func (w WorkspaceAutomations) ReplaceByEnv(ctx context.Context, stackId, envId int64, input []UpdateWorkspaceAutomationInput) ([]types.WorkspaceAutomation, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := w.Client.Do(ctx, http.MethodPut, w.envPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceAutomation](res)
}

// SetSource - PUT /orgs/:orgName/stacks/:stackId/envs/:envId/automations/source
// Sets the branch or pull request that every application of input.RepoUrl follows in the env.
// The response is the full list for the env, as ListByEnv returns it.
func (w WorkspaceAutomations) SetSource(ctx context.Context, stackId, envId int64, input SetAutomationSourceInput) ([]types.WorkspaceAutomation, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := w.Client.Do(ctx, http.MethodPut, w.sourcePath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceAutomation](res)
}
