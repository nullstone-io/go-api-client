package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// WorkspaceBuildConfigs is the API resource for a single app workspace's build configuration.
// It is addressed by block + env (the workspace) and lives beside WorkspaceConfig rather than
// inside it, so an update is effective for the next build without an apply.
type WorkspaceBuildConfigs struct {
	Client *Client
}

func (w WorkspaceBuildConfigs) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/build_config", w.Client.Config.OrgName, stackId, blockId, envId)
}

// Get - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/build_config
// The server materializes a config from the block template or module default when none is stored,
// so a nil result only occurs for a workspace that cannot be built (e.g. a non-app block).
func (w WorkspaceBuildConfigs) Get(ctx context.Context, stackId, blockId, envId int64) (*types.WorkspaceBuildConfig, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.basePath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceBuildConfig](res)
}

// Update - PUT /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/build_config
// Replaces the workspace's build configuration as a whole.
func (w WorkspaceBuildConfigs) Update(ctx context.Context, stackId, blockId, envId int64, config types.BuildConfig) (*types.WorkspaceBuildConfig, error) {
	rawPayload, _ := json.Marshal(config)
	res, err := w.Client.Do(ctx, http.MethodPut, w.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceBuildConfig](res)
}
