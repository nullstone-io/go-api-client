package types

import "time"

const (
	BuildStatusQueued       = "queued"
	BuildStatusInitializing = "initializing"
	BuildStatusRunning      = "running"
	BuildStatusCompleted    = "completed"
	BuildStatusFailed       = "failed"
	BuildStatusCancelled    = "cancelled"
)

type Build struct {
	IdModel
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	AppId   int64  `json:"appId"`
	EnvId   int64  `json:"envId"`

	Phase         string    `json:"phase"`
	Status        string    `json:"status"`
	StatusMessage string    `json:"statusMessage"`
	StatusAt      time.Time `json:"statusAt"`

	// ContextKey is <stack-name>/<app-name>
	// This provides a unique identifier for reporting updates to VCS
	ContextKey string `json:"contextKey"`

	// CommittedBy is the Nullstone user associated with the commit author
	CommittedBy string `json:"committedBy"`

	Version     string      `json:"version"`
	CommitInfo  CommitInfo  `json:"commitInfo"`
	BuildConfig BuildConfig `json:"buildConfig"`
}
