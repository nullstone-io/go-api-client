package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type AutogenSubdomainDelegation struct {
	Client *Client
}

// Update - PUT /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain/delegation
func (d AutogenSubdomainDelegation) Update(subdomainId int, envName string, delegation *types.AutogenSubdomain) (*types.AutogenSubdomain, error) {
	rawPayload, _ := json.Marshal(delegation)
	endpoint := path.Join("subdomains", strconv.Itoa(subdomainId), "envs", envName, "autogen_subdomain", "delegation")
	res, err := d.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
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
func (d AutogenSubdomainDelegation) Destroy(subdomainId int, envName string) (found bool, err error) {
	res, err := d.Client.Do(http.MethodDelete, path.Join("subdomains", strconv.Itoa(subdomainId), "envs", envName, "autogen_subdomain", "delegation"), nil, nil, nil)
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
