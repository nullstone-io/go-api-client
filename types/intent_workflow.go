package types

import "time"

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
	Intent  IntentWorkflowIntent `json:"intent"`
	OrgName string               `json:"orgName"`
	StackId int64                `json:"stackId"`
	// BlockId is nil because some intents are targeting an environment while some target a workspace
	BlockId       *int64               `json:"blockId"`
	EnvId         int64                `json:"envId"`
	Status        IntentWorkflowStatus `json:"status"`
	StatusMessage string               `json:"statusMessage"`
	StatusAt      time.Time            `json:"statusAt"`
}
