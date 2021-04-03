package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io/ioutil"
	"net/http"
	"path"
)

type Workspaces struct {
	Client *Client
}

func (w Workspaces) Get(stackName, blockName, envName string) (*types.Workspace, error) {
	res, err := w.Client.Do(http.MethodGet, path.Join("stacks", stackName, "blocks", blockName, "envs", envName), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	raw, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting workspace (%d): %s", res.StatusCode, string(raw))
	}

	var workspace types.Workspace
	if err := json.Unmarshal(raw, &workspace); err != nil {
		return nil, fmt.Errorf("invalid response getting workspace: %w", err)
	}
	return &workspace, nil
}
