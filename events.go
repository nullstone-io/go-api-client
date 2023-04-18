package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EventPayload struct {
	Event string `json:"event"`
}

type Events struct {
	Client *Client
}

func (e Events) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/events", e.Client.Config.OrgName, stackId, blockId, envId)
}

func (e Events) Create(stackId, blockId, envId int64, event string) error {
	input := EventPayload{
		Event: event,
	}
	raw, _ := json.Marshal(input)
	_, err := e.Client.Do(http.MethodPost, e.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return err
	}
	return nil
}
