package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Apps struct {
	Client *Client
}

func (a Apps) basePath() string {
	return fmt.Sprintf("apps")
}

func (a Apps) appPath(appId int64) string {
	return fmt.Sprintf("apps/%d", appId)
}

// List - GET /orgs/:orgName/apps
func (a Apps) List() ([]types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, a.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var apps []types.Application
	if err := response.ReadJson(res, &apps); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return apps, nil
}

// Get - GET /orgs/:orgName/apps/:id
func (a Apps) Get(appId int64) (*types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, a.appPath(appId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var app types.Application
	if err := response.ReadJson(res, &app); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &app, nil
}

// Create - POST /orgs/:orgName/apps
func (a Apps) Create(app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	res, err := a.Client.Do(http.MethodPost, a.basePath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedApp types.Application
	if err := response.ReadJson(res, &updatedApp); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedApp, nil
}

// Update - PUT/PATCH /orgs/:orgName/apps/:id
func (a Apps) Update(appId int64, app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	res, err := a.Client.Do(http.MethodPut, a.appPath(appId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedApp types.Application
	if err := response.ReadJson(res, &updatedApp); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedApp, nil
}

// Destroy - DELETE /orgs/:orgName/apps/:id
func (a Apps) Destroy(appId int64) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, a.appPath(appId), nil, nil, nil)
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
