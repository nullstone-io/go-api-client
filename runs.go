package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"net/http"
)

type Runs struct {
	Client *Client
}

// RunCreateResult contains the result of Runs Create
// The result can be one of types.Run or types.IntentWorkflow
type RunCreateResult struct {
	Run            *types.Run
	IntentWorkflow *types.IntentWorkflow
}

func (r Runs) basePath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/runs", r.Client.Config.OrgName, stackId, workspaceUid)
}

func (r Runs) runPath(stackId int64, runUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/runs/%s", r.Client.Config.OrgName, stackId, runUid)
}

func (r Runs) List(ctx context.Context, stackId int64, workspaceUid uuid.UUID) ([]types.Run, error) {
	res, err := r.Client.Do(ctx, http.MethodGet, r.basePath(stackId, workspaceUid), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Run](res)
}

func (r Runs) Get(ctx context.Context, stackId int64, runUid uuid.UUID) (*types.Run, error) {
	res, err := r.Client.Do(ctx, http.MethodGet, r.runPath(stackId, runUid), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Run](res)
}

func (r Runs) Create(ctx context.Context, stackId int64, workspaceUid uuid.UUID, input types.CreateRunInput) (*RunCreateResult, error) {
	raw, _ := json.Marshal(input)
	res, err := r.Client.Do(ctx, http.MethodPost, r.basePath(stackId, workspaceUid), nil, nil, json.RawMessage(raw))
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
	result := &RunCreateResult{}
	if raw, err := io.ReadAll(res.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	} else {
		// Try to parse into IntentWorkflow; if it doesn't match, parse into Deploy
		if err := json.Unmarshal(raw, &result.IntentWorkflow); err != nil || result.IntentWorkflow.Intent == "" {
			result.IntentWorkflow = nil
			if err := json.Unmarshal(raw, &result.Run); err != nil {
				return result, fmt.Errorf("unknown response body: %w", err)
			}
		}
	}
	return result, nil
}
