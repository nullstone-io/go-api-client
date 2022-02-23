package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AutogenSubdomainDelegation struct {
	Client *Client
}

func (d AutogenSubdomainDelegation) path(stackId, subdomainId, envId int64) string {
	return fmt.Sprintf("orgs/%s/subdomains/%d/envs/%d/autogen_subdomain/delegation", d.Client.Config.OrgName, subdomainId, envId)
}

// Update - PUT /orgs/:orgName/subdomains/:subdomainId/envs/:envId/autogen_subdomain/delegation
func (d AutogenSubdomainDelegation) Update(stackId, subdomainId, envId int64, delegation *types.AutogenSubdomain) (*types.AutogenSubdomain, error) {
	rawPayload, _ := json.Marshal(delegation)
	res, err := d.Client.Do(http.MethodPut, d.path(stackId, subdomainId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDelegation types.AutogenSubdomain
	if err := response.ReadJson(res, &updatedDelegation); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDelegation, nil
}

// Destroy - DELETE /orgs/:orgName/subdomains/:subdomainId/envs/:envId/autogen_subdomain/delegation
func (d AutogenSubdomainDelegation) Destroy(stackId, subdomainId, envId int64) (found bool, err error) {
	res, err := d.Client.Do(http.MethodDelete, d.path(stackId, subdomainId, envId), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := response.Verify(res); response.IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
