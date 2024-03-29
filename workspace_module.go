package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceModule struct {
	Client *Client
}

func (wmv WorkspaceModule) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/module-version", wmv.Client.Config.OrgName, stackId, blockId, envId)
}

func (wmv WorkspaceModule) Update(ctx context.Context, stackId, blockId, envId int64, moduleInput types.WorkspaceModuleInput) error {
	raw, _ := json.Marshal(moduleInput)
	res, err := wmv.Client.Do(ctx, http.MethodPut, wmv.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return err
	}

	return response.Verify(res)
}
