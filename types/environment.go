package types

import "time"

type EnvironmentType string

const (
	EnvTypePipeline       EnvironmentType = "PipelineEnv"
	EnvTypePreview        EnvironmentType = "PreviewEnv"
	EnvTypePreviewsShared EnvironmentType = "PreviewsSharedEnv"
	EnvTypeGlobal         EnvironmentType = "GlobalEnv"
)

type EnvStatus string

const (
	EnvStatusActive   = "active"
	EnvStatusArchived = "archived"
)

type EnvInfraStatus string

const (
	EnvInfraStatusLive      EnvInfraStatus = "live"
	EnvInfraStatusUpdating  EnvInfraStatus = "updating"
	EnvInfraStatusLaunching EnvInfraStatus = "launching"
	EnvInfraStatusDegraded  EnvInfraStatus = "degraded"
	EnvInfraStatusStopping  EnvInfraStatus = "stopping"
	EnvInfraStatusStopped   EnvInfraStatus = "stopped"
)

type Environment struct {
	IdModel
	Type             EnvironmentType `json:"type"`
	Name             string          `json:"name"`
	OrgName          string          `json:"orgName"`
	StackId          int64           `json:"stackId"`
	Reference        string          `json:"reference"`
	ProviderConfig   ProviderConfig  `json:"providerConfig"`
	PipelineOrder    *int            `json:"pipelineOrder,omitempty"`
	CreatedBy        string          `json:"createdBy"`
	ContextKey       string          `json:"contextKey"`
	Status           EnvStatus       `json:"status"`
	IsProd           bool            `json:"isProd"`
	LatestActivityAt time.Time       `json:"latestActivityAt"`
}

type EnvironmentWithStack struct {
	Environment `json:",inline"`
	StackName   string `json:"stackName"`
}

type EnvironmentWithSummary struct {
	EnvironmentWithStack `json:",inline"`

	NumApps     int              `json:"numApps"`
	NumLiveApps int              `json:"numLiveApps"`
	InfraStatus EnvInfraStatus   `json:"infraStatus"`
	Repos       []EnvSummaryRepo `json:"repos"`
}

type EnvSummaryRepo struct {
	Repo                  `json:",inline"`
	PullRequestId         int64  `json:"pullRequestId"`
	PullRequestNumber     int    `json:"pullRequestNumber"`
	PullRequestHeadBranch string `json:"pullRequestHeadBranch"`
	PullRequestBaseBranch string `json:"pullRequestBaseBranch"`
}
