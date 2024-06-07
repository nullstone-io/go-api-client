package types

import (
	"github.com/google/uuid"
	"time"
)

type WorkspaceWorkflow struct {
	IdModel
	FriendlyAction   string          `json:"friendlyAction"`
	Actions          []string        `json:"actions"`
	OrgName          string          `json:"orgName"`
	WorkspaceUid     uuid.UUID       `json:"workspaceUid"`
	StackId          int64           `json:"stackId"`
	StackName        string          `json:"stackName"`
	BlockId          int64           `json:"blockId"`
	BlockName        string          `json:"blockName"`
	EnvId            int64           `json:"envId"`
	EnvName          string          `json:"envName"`
	Status           string          `json:"status"`
	StatusMessage    string          `json:"statusMessage"`
	StatusAt         time.Time       `json:"statusAt"`
	CommitInfo       CommitInfo      `json:"commitInfo"`
	Trigger          ExternalTrigger `json:"trigger"`
	IntentWorkflowId int64           `json:"intentWorkflowId"`

	// DependencyWorkflowIds is a list of models.WorkspaceWorkflow referring to other dependencies
	// This contains all immediate *and* transitive dependencies
	// All referenced workflow IDs are in the same IntentWorkflow
	DependencyWorkflows []WorkspaceWorkflow `json:"dependencyWorkflows"`

	IntentWorkflow   *IntentWorkflow   `json:"intentWorkflow,omitempty"`
	IacSyncWorkspace *IacSyncWorkspace `json:"iacSyncWorkspace,omitempty"`
	Run              *Run              `json:"run,omitempty"`
}
