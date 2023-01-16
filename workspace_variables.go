package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceVariables struct {
	Client *Client
}

func (wv WorkspaceVariables) basePath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/variables", wv.Client.Config.OrgName, stackId, workspaceUid)
}

func (wv WorkspaceVariables) Update(stackId int64, workspaceUid uuid.UUID, input []types.VariableValue) ([]types.WorkspaceChange, error) {
	raw, _ := json.Marshal(input)
	res, err := wv.Client.Do(http.MethodPut, wv.basePath(stackId, workspaceUid), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	var changes []types.WorkspaceChange
	if err := response.ReadJson(res, &changes); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changes, nil
}
