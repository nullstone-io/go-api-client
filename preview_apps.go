package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type PreviewApps struct {
	Client *Client
}

func (p PreviewApps) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/preview_apps", p.Client.Config.OrgName, stackId, envId)
}

// List - GET /orgs/{orgName}/stacks/{stackId}/envs/{envId}/preview_apps
func (p PreviewApps) List(ctx context.Context, stackId, envId int64) ([]types.PreviewApp, error) {
	res, err := p.Client.Do(ctx, http.MethodGet, p.basePath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.PreviewApp](res)
}

// Replace - PUT /orgs/{orgName}/stacks/{stackId}/envs/{envId}/preview_apps
// This has replace semantics: the env's preview app set becomes exactly previewApps,
// and any app not in the list is removed from the env. In a preview env "enabled"
// means "present in this set", so adding or removing an app is a membership change,
// not a field write. Callers wanting to change one app must List first and send the
// full mutated list back.
func (p PreviewApps) Replace(ctx context.Context, stackId, envId int64, previewApps []types.PreviewApp) ([]types.PreviewApp, error) {
	rawPayload, _ := json.Marshal(previewApps)
	res, err := p.Client.Do(ctx, http.MethodPut, p.basePath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.PreviewApp](res)
}
