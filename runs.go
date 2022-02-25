package api

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Runs struct {
	Client *Client
}

func (r Runs) basePath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/runs", r.Client.Config.OrgName, stackId, workspaceUid)
}

func (r Runs) Create(stackId int64, workspaceUid uuid.UUID, input types.CreateRunInput) (*types.Run, error) {
	res, err := r.Client.Do(http.MethodPost, r.basePath(stackId, workspaceUid), nil, nil, input)
	if err != nil {
		return nil, err
	}

	var run types.Run
	if err := response.ReadJson(res, &run); response.IsNotFoundError(err) {
		return nil, nil
	}
	return &run, nil
}
