package api

import (
	"bytes"
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type AutogenSubdomainsDelegation struct {
	Client *Client
}

// GET /orgs/autogen_subdomains/:subdomainName/delegation
func (d *AutogenSubdomainsDelegation) Get(subdomainName string) (*types.AutogenSubdomainDelegation, error) {
	res, err := d.Client.Do(http.MethodGet, path.Join("autogen_subdomains", subdomainName, "delegation"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var delegation types.AutogenSubdomainDelegation
	if err := d.Client.ReadJsonResponse(res, &delegation); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &delegation, nil
}

// PUT /orgs/autogen_subdomains/:subdomainId/delegation ...
func (d *AutogenSubdomainsDelegation) UpdateAutogenSubdomainDelegation(subdomainName string, delegation *types.AutogenSubdomainDelegation) (*types.AutogenSubdomainDelegation, error) {
	rawPayload, _ := json.Marshal(delegation)
	endpoint := path.Join("autogen_subdomains", subdomainName, "delegation")
	headers := map[string]string{"Content-Type": "application/json"}
	res, err := d.Client.Do(http.MethodPut, endpoint, nil, headers, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDelegation types.AutogenSubdomainDelegation
	if err := d.Client.ReadJsonResponse(res, &updatedDelegation); err != nil {
		return nil, err
	}
	return &updatedDelegation, nil
}

// DELETE /orgs/autogen_subdomains/:subdomainId/delegation ...
func (d *AutogenSubdomainsDelegation) DestroyAutogenSubdomainDelegation(subdomainName string) (found bool, err error) {
	res, err := d.Client.Do(http.MethodDelete, path.Join("autogen_subdomains", subdomainName, "delegation"), nil, nil, nil)
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
