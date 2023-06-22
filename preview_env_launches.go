package api

import (
	"fmt"
	"net/http"
)

type PreviewEnvLaunches struct {
	Client *Client
}

func (pe PreviewEnvLaunches) envPath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/preview_envs/%d", pe.Client.Config.OrgName, stackId, envId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/preview_envs/:id/launches
func (pe PreviewEnvLaunches) Create(stackId, envId int64) error {
	_, err := pe.Client.Do(http.MethodPost, pe.envPath(stackId, envId), nil, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
