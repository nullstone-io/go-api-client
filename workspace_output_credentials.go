package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
	"strings"
)

type WorkspaceOutputCredentials struct {
	Client *Client
}

func (c WorkspaceOutputCredentials) path(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/output-credentials", c.Client.Config.OrgName, stackId, workspaceUid)
}

// Create - POST /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/output-credentials
func (c WorkspaceOutputCredentials) Create(ctx context.Context, stackId int64, workspaceUid uuid.UUID, provider string, outputNames []string) (*types.OutputCredentials, error) {
	q := url.Values{}
	q.Set("provider", provider)
	q.Set("output_names", strings.Join(outputNames, ","))
	res, err := c.Client.Do(ctx, http.MethodPost, c.path(stackId, workspaceUid), q, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.OutputCredentials](res)
}
