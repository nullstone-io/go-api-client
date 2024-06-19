package types

type WorkspaceWorkflowActivities struct {
	WorkspaceWorkflowId int64   `json:"workspaceWorkflowId"`
	Build               *Build  `json:"build"`
	Deploy              *Deploy `json:"deploy"`
}
