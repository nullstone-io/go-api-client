package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AutogenSubdomainDelegation struct {
	Client *Client
}

func (d AutogenSubdomainDelegation) path(subdomainId int64, envName string) string {
	return fmt.Sprintf("subdomains/%d/envs/%s/autogen_subdomain/delegation", subdomainId, envName)
}

// Update - PUT /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain/delegation
func (d AutogenSubdomainDelegation) Update(subdomainId int64, envName string, delegation *types.AutogenSubdomain) (*types.AutogenSubdomain, error) {
	rawPayload, _ := json.Marshal(delegation)
	res, err := d.Client.Do(http.MethodPut, d.path(subdomainId, envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDelegation types.AutogenSubdomain
	if err := d.Client.ReadJsonResponse(res, &updatedDelegation); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDelegation, nil
}

// Destroy - DELETE /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain/delegation
func (d AutogenSubdomainDelegation) Destroy(subdomainId int64, envName string) (found bool, err error) {
	res, err := d.Client.Do(http.MethodDelete, d.path(subdomainId, envName), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := d.Client.VerifyResponse(res); IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
