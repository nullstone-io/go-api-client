package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type Apps struct {
	Client *Client
}

func (a Apps) Get(appName string) (*types.Application, error) {
	res, err := a.Client.Do(http.MethodGet, path.Join("apps", appName), nil, nil, nil)
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
