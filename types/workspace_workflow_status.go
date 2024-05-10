package types

import "time"

type WorkspaceWorkflowStatus struct {
	WorkspaceWorkflowId int64     `json:"workspaceWorkflowId"`
	CreatedAt           time.Time `json:"createdAt"`
	CreatedBy           string    `json:"createdBy"`
	Status              string    `json:"status"`
	StatusMessage       string    `json:"statusMessage"`
	StatusAt            time.Time `json:"statusAt"`
	Intent              string    `json:"intent"`
	IntentWorkflowId    *int64    `json:"intentWorkflowId"`
	Actions             []string  `json:"actions"`
}
