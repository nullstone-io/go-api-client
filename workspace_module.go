package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type WorkspaceModule struct {
	Client *Client
}

func (wmv WorkspaceModule) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/module", wmv.Client.Config.OrgName, stackId, blockId, envId)
}

type UpdateWorkspaceModuleInput struct {
	Module        string                  `json:"module,omitempty"`
	ModuleVersion string                  `json:"moduleVersion,omitempty"`
	Connections   types.ConnectionTargets `json:"connections,omitempty"`
}

func (wmv WorkspaceModule) Update(ctx context.Context, stackId, blockId, envId int64, moduleInput UpdateWorkspaceModuleInput) error {
	raw, _ := json.Marshal(moduleInput)
	res, err := wmv.Client.Do(ctx, http.MethodPut, wmv.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return err
	}

	return response.Verify(res)
}
