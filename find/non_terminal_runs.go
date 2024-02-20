package find

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"sort"
)

// NonTerminalRuns finds all runs for a given workspace that have not completed
// This includes queued, resolving, initializing, awaiting-dependencies, running, and needs-approval
// The result is sorted from oldest to newest
func NonTerminalRuns(cfg api.Config, stackId int64, workspaceUid uuid.UUID) ([]types.Run, error) {
	client := api.Client{Config: cfg}
	runs, err := client.Runs().List(stackId, workspaceUid)
	if err != nil {
		return nil, fmt.Errorf("error retrieving runs for application workspace: %w", err)
	}

	ntRuns := make([]types.Run, 0)
	for _, run := range runs {
		if !types.IsTerminalRunStatus(run.Status) {
			ntRuns = append(ntRuns, run)
		}
	}

	sort.SliceStable(ntRuns, func(i, j int) bool {
		a, b := ntRuns[i], ntRuns[j]
		return a.CreatedAt.Before(b.CreatedAt)
	})

	return ntRuns, nil
}
