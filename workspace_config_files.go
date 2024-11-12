package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"io"
	"net/http"
)

type WorkspaceConfigFiles struct {
	Client *Client
}

func (w WorkspaceConfigFiles) configPath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/config", w.Client.Config.OrgName, stackId, blockId, envId)
}

func (w WorkspaceConfigFiles) overridesPath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/config_variables", w.Client.Config.OrgName, stackId, blockId, envId)
}

// GetConfigFile - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/config
func (w WorkspaceConfigFiles) GetConfigFile(ctx context.Context, stackId, blockId, envId int64, file io.Writer) error {
	res, err := w.Client.Do(ctx, http.MethodGet, w.configPath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return err
	}

	if err := response.ReadFile(res, file); response.IsNotFoundError(err) {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

// GetOverridesFile - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/config_variables
func (w WorkspaceConfigFiles) GetOverridesFile(ctx context.Context, stackId, blockId, envId int64, file io.Writer) error {
	res, err := w.Client.Do(ctx, http.MethodGet, w.overridesPath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return err
	}

	if err := response.ReadFile(res, file); response.IsNotFoundError(err) {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
