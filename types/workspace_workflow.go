package types

import (
	"github.com/google/uuid"
	"slices"
	"time"
)

type WorkspaceWorkflow struct {
	IdModel
	FriendlyAction string          `json:"friendlyAction"`
	Actions        []string        `json:"actions"`
	OrgName        string          `json:"orgName"`
	WorkspaceUid   uuid.UUID       `json:"workspaceUid"`
	StackId        int64           `json:"stackId"`
	StackName      string          `json:"stackName"`
	BlockId        int64           `json:"blockId"`
	BlockName      string          `json:"blockName"`
	EnvId          int64           `json:"envId"`
	EnvName        string          `json:"envName"`
	Status         string          `json:"status"`
	StatusMessage  string          `json:"statusMessage"`
	StatusAt       time.Time       `json:"statusAt"`
	CommitInfo     CommitInfo      `json:"commitInfo"`
	Trigger        ExternalTrigger `json:"trigger"`

	// IntentWorkflowOrder is used to sort connections by order of dependencies
	// Roots of the dependency tree will have a lower order number
	// The order number is calculated by walking through the tree and incrementing a counter using in-order traversal
	IntentWorkflowOrder int `json:"intentWorkflowOrder"`
	// IntentWorkflowId points to the IntentWorkflow that collects a set of workspace workflows into a single execution
	IntentWorkflowId int64 `json:"intentWorkflowId"`

	// PreviousWorkflowId is the ID of the workspace workflow that was created before this workflow within this workspace
	PreviousWorkflowId int64 `json:"previousWorkflowId"`
	// NextWorkflowId is the ID of the workspace workflow that was created after this workflow within this workspace
	NextWorkflowId int64 `json:"nextWorkflowId"`

	// DependencyWorkflowIds is a list of models.WorkspaceWorkflow referring to other dependencies
	// This contains all immediate *and* transitive dependencies
	// All referenced workflow IDs are in the same IntentWorkflow
	DependencyWorkflows []WorkspaceWorkflow `json:"dependencyWorkflows"`

	IntentWorkflow   *IntentWorkflow   `json:"intentWorkflow,omitempty"`
	IacSyncWorkspace *IacSyncWorkspace `json:"iacSyncWorkspace,omitempty"`
	Run              *Run              `json:"run,omitempty"`
}

type WorkspaceWorkflowUpdate struct {
	Id             int64      `json:"id"`
	FriendlyAction *string    `json:"friendlyAction"`
	Actions        []string   `json:"actions,omitempty"`
	Status         *string    `json:"status,omitempty"`
	StatusAt       *time.Time `json:"statusAt,omitempty"`
	StatusMessage  *string    `json:"statusMessage,omitempty"`
	// DependencyWorkflows is not guaranteed to be a full set of all dependencies
	// Instead, each item in this slice should be seen as "create me" or "update me"
	DependencyWorkflows []WorkspaceWorkflow `json:"dependencyWorkflows,omitempty"`
	// NextWorkflowId is the ID of the workspace workflow that was created after this workflow
	NextWorkflowId      *int64  `json:"nextWorkflowId,omitempty"`
	IntentWorkflowOrder *int    `json:"intentWorkflowOrder,omitempty"`
	Run                 *Run    `json:"run,omitempty"`
	Build               *Build  `json:"build,omitempty"`
	Deploy              *Deploy `json:"deploy,omitempty"`
}

func (u WorkspaceWorkflowUpdate) ApplyTo(ww WorkspaceWorkflow) WorkspaceWorkflow {
	if ww.Id != u.Id {
		return ww
	}
	if u.FriendlyAction != nil {
		ww.FriendlyAction = *u.FriendlyAction
	}
	if u.Actions != nil {
		ww.Actions = u.Actions
	}
	if u.Status != nil {
		ww.Status = *u.Status
	}
	if u.StatusAt != nil {
		ww.StatusAt = *u.StatusAt
	}
	if u.StatusMessage != nil {
		ww.StatusMessage = *u.StatusMessage
	}
	if u.NextWorkflowId != nil {
		ww.NextWorkflowId = *u.NextWorkflowId
	}
	if u.IntentWorkflowOrder != nil {
		ww.IntentWorkflowOrder = *u.IntentWorkflowOrder
	}
	if u.Run != nil {
		ww.Run = u.Run
	}
	if u.DependencyWorkflows != nil {
		for _, dw := range u.DependencyWorkflows {
			existingIndex := slices.IndexFunc(ww.DependencyWorkflows, func(w WorkspaceWorkflow) bool {
				return w.Id == dw.Id
			})
			if existingIndex > -1 {
				ww.DependencyWorkflows[existingIndex] = dw
			} else {
				ww.DependencyWorkflows = append(ww.DependencyWorkflows, dw)
			}
		}
	}
	return ww
}
