package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Runs struct {
	Client *Client
}

// RunCreateResult contains the result of Runs Create
// The result can be one of types.Run, types.IntentWorkflow, or a skipped response.
// When the server short-circuits (e.g. when CreateRunInput.If is not satisfied),
// Skipped is true and SkipReason explains why; Run and IntentWorkflow are nil.
type RunCreateResult struct {
	Run            *types.Run
	IntentWorkflow *types.IntentWorkflow
	Skipped        bool
	SkipReason     string
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

type CreateRunInput struct {
	CommitSha string `json:"commitSha"`

	// Create a run that destroys this workspace
	IsDestroy bool `json:"isDestroy"`

	// DestroyDependencies allows the user to identify which dependencies to destroy along with the block
	// IsDestroy must be enabled for this field to have an effect
	// `*` indicates attempt to destroy the workspace and its dependencies
	// ``  indicates attempt to destroy only the specified workspace
	// `<stack-id>/<block-id>/<env-id>,...` indicates a comma-delimited list of dependencies to destroy with the workspace
	DestroyDependencies string `json:"destroyDependencies"`

	IsApproved     *bool     `json:"isApproved"`
	LatestUpdateAt time.Time `json:"latestUpdateAt"`

	// AppVersion forces a specific app version to be used during the apply.
	// Only valid for application blocks; causes a Deploy record to be created with this version
	// without triggering a build or the enigma deploy workflow.
	AppVersion *string `json:"appVersion,omitempty"`

	// If is an optional condition the server evaluates before creating the run.
	// Empty string means "always create".
	// "any-changes" means "only create when there are unapplied workspace changes".
	If string `json:"if"`

	// Deprecated
	Version *int64 `json:"version"`
}

func (r Runs) Create(ctx context.Context, stackId int64, workspaceUid uuid.UUID, input CreateRunInput) (*RunCreateResult, error) {
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
		// Try skipped envelope first ({"skipped": true, "reason": "..."})
		var skipEnv struct {
			Skipped bool   `json:"skipped"`
			Reason  string `json:"reason"`
		}
		if err := json.Unmarshal(raw, &skipEnv); err == nil && skipEnv.Skipped {
			result.Skipped = true
			result.SkipReason = skipEnv.Reason
			return result, nil
		}
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
