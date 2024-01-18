package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CapabilityCopies struct {
	Client *Client
}

type CapabilityCopiesPayload struct {
	AppIds []int64 `json:"appIds"`
	EnvId  int64   `json:"envId"`
}

func (c CapabilityCopies) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/capability_copies", c.Client.Config.OrgName, stackId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/capability_copies
func (c CapabilityCopies) Create(stackId int64, appIds []int64, envId int64) error {
	payload := CapabilityCopiesPayload{
		AppIds: appIds,
		EnvId:  envId,
	}
	rawPayload, _ := json.Marshal(payload)
	_, err := c.Client.Do(http.MethodPost, c.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}

	return nil
}
