package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"net/http"
)

type ConnectionCopies struct {
	Client *Client
}

type ConnectionCopiesPayload struct {
	EnvId int64 `json:"envId"`
}

func (c ConnectionCopies) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/connection_copies", c.Client.Config.OrgName, stackId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/connection_copies
func (c ConnectionCopies) Create(ctx context.Context, stackId int64, envId int64) error {
	payload := ConnectionCopiesPayload{
		EnvId: envId,
	}
	rawPayload, _ := json.Marshal(payload)
	res, err := c.Client.Do(ctx, http.MethodPost, c.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return err
	}

	return response.Verify(res)
}
