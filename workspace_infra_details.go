package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type WorkspaceInfraDetails struct {
	Client *Client
}

func (d WorkspaceInfraDetails) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_infra_details", d.Client.Config.OrgName, stackId, blockId, envId)
}

func (d WorkspaceInfraDetails) Get(ctx context.Context, stackId, blockId, envId int64, includeArchived bool) (*types.WorkspaceInfraDetails, error) {
	q := url.Values{
		"include_archived": []string{strconv.FormatBool(includeArchived)},
	}
	res, err := d.Client.Do(ctx, http.MethodGet, d.basePath(stackId, blockId, envId), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceInfraDetails](res)
}
