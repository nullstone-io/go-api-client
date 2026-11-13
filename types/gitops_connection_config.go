package types

import "time"

// What an env can follow from a repository (GitopsConnectionConfig.TriggerType)
const (
	// GitopsTriggerTypeBranch fires when a commit is pushed to GitBranch
	GitopsTriggerTypeBranch = "branch"
	// GitopsTriggerTypeTag fires when a tag matching TagFilter is created
	GitopsTriggerTypeTag = "tag"
	// GitopsTriggerTypeRelease fires when a release is published
	GitopsTriggerTypeRelease = "release"
	// GitopsTriggerTypePullRequest follows a pull request: its pushes and its head commit.
	// Only a preview env can follow one.
	GitopsTriggerTypePullRequest = "pull_request"
)

// GitopsTrigger is what an env follows from a repository; only the fields of its type are set.
// TriggerType can be left empty when PullRequestId or GitBranch is set: the server derives it.
type GitopsTrigger struct {
	TriggerType string `json:"triggerType"`
	// GitBranch is the branch to follow when TriggerType is branch
	GitBranch string `json:"gitBranch"`
	// PullRequestId is the pull request to follow when TriggerType is pull_request
	PullRequestId *int64 `json:"pullRequestId"`
	// TagFilter is a newline-separated list of tag filter patterns; it is only meaningful when TriggerType is tag.
	// An empty filter matches every tag.
	TagFilter string `json:"tagFilter"`
}

// GitopsConnectionConfig is what one env follows from one repository: its trigger, and whether a push that
// fires the trigger syncs IaC (IsEnabled) and applies it (IsAutoApplyEnabled). The trigger also fires the
// auto-deploy of the repository's apps in the env (see WorkspaceAutomation), whether or not IaC sync is on.
//
// Pipeline and preview envs both have one per repository. A preview env's config has IsPreview set, always has
// IaC sync and auto-apply on, and may follow a pull request.
type GitopsConnectionConfig struct {
	GitopsConnectionId int64 `json:"gitopsConnectionId"`
	EnvId              int64 `json:"envId"`
	IsEnabled          bool  `json:"isEnabled"`
	IsAutoApplyEnabled bool  `json:"isAutoApplyEnabled"`
	IsPreview          bool  `json:"isPreview"`

	TriggerType   string `json:"triggerType"`
	GitBranch     string `json:"gitBranch"`
	PullRequestId *int64 `json:"pullRequestId"`
	TagFilter     string `json:"tagFilter"`

	// CreatedBy is who first chose what the env follows from the repository.
	// UpdatedBy and UpdatedAt are who last changed the config and when.
	// CreatedBy and UpdatedBy are empty on a config written before they were recorded.
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`

	// PullRequest is the pull request named by PullRequestId, when the endpoint loads it
	PullRequest *PullRequest `json:"pullRequest,omitempty"`
}

// Trigger returns what this config follows
func (c GitopsConnectionConfig) Trigger() GitopsTrigger {
	return GitopsTrigger{
		TriggerType:   c.TriggerType,
		GitBranch:     c.GitBranch,
		PullRequestId: c.PullRequestId,
		TagFilter:     c.TagFilter,
	}
}
