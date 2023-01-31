package api

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceChanges struct {
	Client *Client
}

func (wc WorkspaceChanges) basePath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/changes", wc.Client.Config.OrgName, stackId, workspaceUid)
}

func (wc WorkspaceChanges) changePath(stackId int64, workspaceUid uuid.UUID, changeId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/changes/%d", wc.Client.Config.OrgName, stackId, workspaceUid, changeId)
}

func (wc WorkspaceChanges) List(stackId int64, workspaceUid uuid.UUID) (*types.WorkspaceChangeset, error) {
	res, err := wc.Client.Do(http.MethodPost, wc.basePath(stackId, workspaceUid), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var changeset *types.WorkspaceChangeset
	if err := response.ReadJson(res, changeset); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changeset, nil
}

func (wc WorkspaceChanges) Destroy(stackId int64, workspaceUid uuid.UUID, changeId int64) (*types.WorkspaceChangeset, error) {
	res, err := wc.Client.Do(http.MethodDelete, wc.changePath(stackId, workspaceUid, changeId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var changeset *types.WorkspaceChangeset
	if err := response.ReadJson(res, changeset); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changeset, nil
}
