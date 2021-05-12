package api

import (
	"encoding/json"
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

	var apps []types.Application
	if err := a.Client.ReadJsonResponse(res, &apps); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return apps, nil
}

// Get - GET /orgs/:orgName/apps/:id
func (a Apps) Get(appId int) (*types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, path.Join("apps", strconv.Itoa(appId)), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var app types.Application
	if err := a.Client.ReadJsonResponse(res, &app); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &app, nil
}

// Create - POST /orgs/:orgName/apps
func (a Apps) Create(app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	res, err := a.Client.Do(http.MethodPost, path.Join("apps"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedApp types.Application
	if err := a.Client.ReadJsonResponse(res, &updatedApp); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedApp, nil
}

// Update - PUT/PATCH /orgs/:orgName/apps/:id
func (a Apps) Update(appId int, app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	endpoint := path.Join("apps", strconv.Itoa(appId))
	res, err := a.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedApp types.Application
	if err := a.Client.ReadJsonResponse(res, &updatedApp); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedApp, nil
}

// Destroy - DELETE /orgs/:orgName/apps/:id
func (a Apps) Destroy(appId int) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, path.Join("apps", strconv.Itoa(appId)), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := a.Client.VerifyResponse(res); IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
