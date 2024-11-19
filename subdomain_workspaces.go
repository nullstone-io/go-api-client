package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type SubdomainWorkspaces struct {
	Client *Client
}

func (s SubdomainWorkspaces) subdomainPath(stackId, subdomainId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/subdomains/%d/envs/%d", s.Client.Config.OrgName, stackId, subdomainId, envId)
}

// Get - GET /orgs/:orgName/stacks/:stackId/subdomains/:id/envs/:envId
func (s SubdomainWorkspaces) Get(ctx context.Context, stackId, subdomainId, envId int64) (*types.SubdomainWorkspace, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.subdomainPath(stackId, subdomainId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.SubdomainWorkspace](res)
}
