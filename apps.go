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

func (a Apps) globalPath() string {
	return fmt.Sprintf("orgs/%s/apps", a.Client.Config.OrgName)
}

func (a Apps) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps", a.Client.Config.OrgName, stackId)
}

func (a Apps) appPath(stackId, appId int64) string {
	return fmt.Sprintf("orgs/%s/apps/%d", a.Client.Config.OrgName, appId)
}

// List - GET /orgs/:orgName/apps
func (a Apps) List() ([]types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, a.globalPath(), nil, nil, nil)
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

// Get - GET /orgs/:orgName/stacks/:stackId/apps/:id
func (a Apps) Get(stackId, appId int64) (*types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, a.appPath(stackId, appId), nil, nil, nil)
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

// Create - POST /orgs/:orgName/stacks/:stackId/apps
func (a Apps) Create(stackId int64, app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	res, err := a.Client.Do(http.MethodPost, a.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
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

// Update - PUT/PATCH /orgs/:orgName/stacks/:stackId/apps/:id
func (a Apps) Update(stackId, appId int64, app *types.Application) (*types.Application, error) {
	rawPayload, _ := json.Marshal(app)
	res, err := a.Client.Do(http.MethodPut, a.appPath(stackId, appId), nil, nil, json.RawMessage(rawPayload))
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

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/apps/:id
func (a Apps) Destroy(stackId, appId int64) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, a.appPath(stackId, appId), nil, nil, nil)
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
