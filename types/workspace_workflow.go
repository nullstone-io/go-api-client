package types

import "time"

type WorkspaceWorkflow struct {
	IdModel
	Actions          []string  `json:"actions"`
	OrgName          string    `json:"orgName"`
	StackId          int64     `json:"stackId"`
	BlockId          int64     `json:"blockId"`
	EnvId            int64     `json:"envId"`
	Status           string    `json:"status"`
	StatusMessage    string    `json:"statusMessage"`
	StatusAt         time.Time `json:"statusAt"`
	IntentWorkflowId *int64    `json:"intentWorkflowId" pg:"intent_workflow_id"`
	Intent           string    `json:"intent"`
}
