package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"net/http"
)

type EnvRuns struct {
	Client *Client
}

func (er EnvRuns) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/runs", er.Client.Config.OrgName, stackId, envId)
}

// EnvRunsCreateResult contains the result of EnvRuns Create
// The result can be one of []types.Run or types.IntentWorkflow
type EnvRunsCreateResult struct {
	Runs           []types.Run
	IntentWorkflow types.IntentWorkflow
}

func (er EnvRuns) Create(ctx context.Context, stackId, envId int64, input types.CreateEnvRunInput) (*EnvRunsCreateResult, error) {
	raw, _ := json.Marshal(input)
	res, err := er.Client.Do(ctx, http.MethodPost, er.basePath(stackId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	if err := response.Verify(res); err != nil {
		if response.IsNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	defer res.Body.Close()
	result := &EnvRunsCreateResult{}
	if raw, err := io.ReadAll(res.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	} else {
		// Try to parse into Runs, then try to parse into IntentWorkflow
		if err := json.Unmarshal(raw, &result.Runs); err != nil {
			if err := json.Unmarshal(raw, &result.IntentWorkflow); err != nil {
				return result, fmt.Errorf("unknown response body: %w", err)
			}
		}
		return result, nil
	}
}
