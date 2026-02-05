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

type WorkspaceDetails struct {
	Client *Client
}

func (d WorkspaceDetails) path(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_details", d.Client.Config.OrgName, stackId, blockId, envId)
}

func (d WorkspaceDetails) pathByName(stackName, blockName, envName string) string {
	return fmt.Sprintf("/orgs/%s/stacks_by_name/%s/blocks/%s/envs/%s/workspace_details", d.Client.Config.OrgName, stackName, blockName, envName)
}

func (d WorkspaceDetails) Get(ctx context.Context, stackId, blockId, envId int64, includeArchived bool) (*types.WorkspaceDetails, error) {
	q := url.Values{
		"include_archived": []string{strconv.FormatBool(includeArchived)},
	}
	res, err := d.Client.Do(ctx, http.MethodGet, d.path(stackId, blockId, envId), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceDetails](res)
}

func (d WorkspaceDetails) GetByName(ctx context.Context, stackName, blockName, envName string, includeArchived bool) (*types.WorkspaceDetails, error) {
	q := url.Values{
		"include_archived": []string{strconv.FormatBool(includeArchived)},
	}
	res, err := d.Client.Do(ctx, http.MethodGet, d.pathByName(stackName, blockName, envName), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceDetails](res)
}
