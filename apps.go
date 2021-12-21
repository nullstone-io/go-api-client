package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Apps struct {
	Client *Client
}

// List - GET /orgs/:orgName/apps
func (a Apps) List() ([]types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, path.Join("apps"), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.JsonArray[types.Application](res)
}

// Get - GET /orgs/:orgName/apps/:id
func (a Apps) Get(appId int) (*types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, path.Join("apps", strconv.Itoa(appId)), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Application](res)
}

// Create - POST /orgs/:orgName/apps
func (a Apps) Create(app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	res, err := a.Client.Do(http.MethodPost, path.Join("apps"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Application](res)
}

// Update - PUT/PATCH /orgs/:orgName/apps/:id
func (a Apps) Update(appId int, app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	endpoint := path.Join("apps", strconv.Itoa(appId))
	res, err := a.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Application](res)
}

// Destroy - DELETE /orgs/:orgName/apps/:id
func (a Apps) Destroy(appId int) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, path.Join("apps", strconv.Itoa(appId)), nil, nil, nil)
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
