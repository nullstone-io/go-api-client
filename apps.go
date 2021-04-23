package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Apps struct {
	Client *Client
}

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
