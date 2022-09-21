package api

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
)

type WorkspaceOutputs struct {
	Client *Client
}

func (w WorkspaceOutputs) path(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/current-outputs", w.Client.Config.OrgName, stackId, workspaceUid)
}

// GetCurrent - GET /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/current-outputs
func (w WorkspaceOutputs) GetCurrent(stackId int64, workspaceUid uuid.UUID, showSensitive bool) (types.Outputs, error) {
	q := url.Values{}
	if showSensitive {
		q.Set("show_sensitive", "true")
	}
	res, err := w.Client.Do(http.MethodGet, w.path(stackId, workspaceUid), q, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[types.Outputs](res)
}
