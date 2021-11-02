package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceOutputs struct {
	Client *Client
}

func (w WorkspaceOutputs) GetLatest(stackId, blockId, envId int64) (*types.Outputs, error) {
	endpoint := fmt.Sprintf("stacks/%d/blocks/%d/envs/%d/outputs/latest", stackId, blockId, envId)
	res, err := w.Client.Do(http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var outputs types.Outputs
	if err := w.Client.ReadJsonResponse(res, &outputs); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &outputs, nil
}
