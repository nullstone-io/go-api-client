package types

import (
	"slices"
	"time"
)

type IntentWorkflowIntent string

const (
	IntentWorkflowIntentEnvUp      IntentWorkflowIntent = "env-up"
	IntentWorkflowIntentEnvDown    IntentWorkflowIntent = "env-down"
	IntentWorkflowIntentBlockDown  IntentWorkflowIntent = "block-down"
	IntentWorkflowIntentBlockApply IntentWorkflowIntent = "block-apply"
	IntentWorkflowIntentAppDeploy  IntentWorkflowIntent = "app-deploy"
	IntentWorkflowIntentRepoSync   IntentWorkflowIntent = "repo-sync"
)

type IntentWorkflowStatus string

const (
	// IntentWorkflowStatusCalculating indicates the workflows was added to the database and has started calculating a plan
	IntentWorkflowStatusCalculating IntentWorkflowStatus = "calculating"
	// IntentWorkflowStatusRunning indicates the workflow is currently running
	IntentWorkflowStatusRunning IntentWorkflowStatus = "running"
	// IntentWorkflowStatusCompleted indicates the workflow completed successfully
	IntentWorkflowStatusCompleted IntentWorkflowStatus = "completed"
	// IntentWorkflowStatusFailed indicates the workflow failed to complete
	IntentWorkflowStatusFailed IntentWorkflowStatus = "failed"
	// IntentWorkflowStatusCancelled indicates the workflow was cancelled
	IntentWorkflowStatusCancelled IntentWorkflowStatus = "cancelled"
)

type IntentWorkflow struct {
	IdModel
	Intent    IntentWorkflowIntent `json:"intent"`
	OrgName   string               `json:"orgName"`
	StackId   int64                `json:"stackId"`
	StackName string               `json:"stackName"`
	// BlockId is nil because some intents are targeting an environment while some target a workspace
	BlockId       *int64               `json:"blockId"`
	BlockName     *string              `json:"blockName"`
	EnvId         int64                `json:"envId"`
	EnvName       string               `json:"envName"`
	Status        IntentWorkflowStatus `json:"status"`
	StatusMessage string               `json:"statusMessage"`
	StatusAt      time.Time            `json:"statusAt"`
	CommitInfo    CommitInfo           `json:"commitInfo"`
	Trigger       ExternalTrigger      `json:"trigger"`

	// PrimaryWorkflow contains the WorkspaceWorkflow if this intent workflow was initiated by a single workspace
	// If BlockId is nil, PrimaryWorkflow is nil
	// If BlockId is not nil, PrimaryWorkflow is not nil
	PrimaryWorkflow *WorkspaceWorkflow `json:"primaryWorkflow"`
	// RootWorkflowIds identifies all workspace workflows that don't have any dependencies
	// If BlockId is nil, RootWorkflowIds is not nil
	// If BlockId is not nil, RootWorkflowIds is nil
	RootWorkflowIds []int64 `json:"rootWorkflowIds"`

	// WorkspaceWorkflows contains all WorkspaceWorkflow in this IntentWorkflow
	// This is not included when Listing many workflows
	WorkspaceWorkflows []WorkspaceWorkflow `json:"workspaceWorkflows"`
}

type IntentWorkflowUpdate struct {
	Id                 int64                 `json:"id"`
	Status             *IntentWorkflowStatus `json:"status,omitempty"`
	StatusAt           *time.Time            `json:"statusAt,omitempty"`
	StatusMessage      *string               `json:"statusMessage,omitempty"`
	WorkspaceWorkflows []WorkspaceWorkflow   `json:"workspaceWorkflows,omitempty"`
	RootWorkflowIds    []int64               `json:"rootWorkflowIds,omitempty"`
}

func (u IntentWorkflowUpdate) ApplyTo(iw IntentWorkflow) IntentWorkflow {
	if iw.Id != u.Id {
		return iw
	}
	if u.Status != nil {
		iw.Status = *u.Status
	}
	if u.StatusMessage != nil {
		iw.StatusMessage = *u.StatusMessage
	}
	if u.StatusAt != nil {
		iw.StatusAt = *u.StatusAt
	}
	if u.RootWorkflowIds != nil {
		iw.RootWorkflowIds = u.RootWorkflowIds
	}
	if u.WorkspaceWorkflows != nil {
		for _, ww := range u.WorkspaceWorkflows {
			existingIndex := slices.IndexFunc(iw.WorkspaceWorkflows, func(w WorkspaceWorkflow) bool {
				return w.Id == ww.Id
			})
			if existingIndex > -1 {
				iw.WorkspaceWorkflows[existingIndex] = ww
			} else {
				iw.WorkspaceWorkflows = append(iw.WorkspaceWorkflows, ww)
			}
		}
	}
	return iw
}
