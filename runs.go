package api

import (
	"encoding/json"
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

func (r Runs) runPath(stackId int64, runUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/runs/%s", r.Client.Config.OrgName, stackId, runUid)
}

func (r Runs) Get(stackId int64, runUid uuid.UUID) (*types.Run, error) {
	res, err := r.Client.Do(http.MethodGet, r.runPath(stackId, runUid), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var run types.Run
	if err := response.ReadJson(res, &run); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &run, nil
}

func (r Runs) Create(stackId int64, workspaceUid uuid.UUID, input types.CreateRunInput) (*types.Run, error) {
	raw, _ := json.Marshal(input)
	res, err := r.Client.Do(http.MethodPost, r.basePath(stackId, workspaceUid), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	var run types.Run
	if err := response.ReadJson(res, &run); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &run, nil
}
