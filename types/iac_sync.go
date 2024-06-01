package types

import (
	"github.com/google/uuid"
	"time"
)

const (
	IacSyncStatusQueued    = "queued"
	IacSyncStatusFetching  = "fetching"
	IacSyncStatusRunning   = "running"
	IacSyncStatusCompleted = "completed"
	IacSyncStatusCancelled = "cancelled"
	IacSyncStatusFailed    = "failed"
)

// IacSync represents an entire IaC sync operation that is initiated by a commit to a repo.
//
//	Each IacSync record can make updates to more than one block and more than one environment.
//	Multiple IacSyncWorkspace records are created to track the progress of each workspace update.
type IacSync struct {
	Uid              uuid.UUID         `json:"uid"`
	OrgName          string            `json:"orgName"`
	StackId          int64             `json:"stackId"`
	EnvId            *int64            `json:"envId"`
	CommitInfo       CommitInfo        `json:"commitInfo"`
	CreatedBy        string            `json:"createdBy"`
	CreatedAt        time.Time         `json:"createdAt"`
	Status           string            `json:"status"`
	StatusMessage    string            `json:"statusMessage"`
	StatusAt         time.Time         `json:"statusAt"`
	IntentWorkflowId *int64            `json:"intentWorkflowId"`
	YamlConfigFiles  map[string]string `json:"yamlConfigFiles"`
}
