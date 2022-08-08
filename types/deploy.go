package types

import "time"

const (
	DeployPhaseInit        = "init"
	DeployPhaseCheckout    = "checkout"
	DeployPhaseBuild       = "build"
	DeployPhasePush        = "push"
	DeployPhaseDeploy      = "deploy"
	DeployPhaseWaitHealthy = "wait-healthy"
)

const (
	DeployStatusQueued       = "queued"
	DeployStatusInitializing = "initializing"
	DeployStatusRunning      = "running"
	DeployStatusCompleted    = "completed"
	DeployStatusFailed       = "failed"
	DeployStatusCancelled    = "cancelled"
)

const (
	DeployTypeBuildPushDeploy = "build-push-deploy"
	DeployTypeDeployOnly      = "deploy-only"
	DeployTypeRollback        = "rollback"
)

const (
	DeployPackageModeNone       = "none"
	DeployPackageModeDockerfile = "dockerfile"
	DeployPackageModeBuildpacks = "buildpacks"
	DeployPackageModeZip        = "zip"
)

type Deploy struct {
	IdModel
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	AppId   int64  `json:"appId"`
	EnvId   int64  `json:"envId"`

	Reference     string    `json:"reference"`
	Phase         string    `json:"phase"`
	Status        string    `json:"status"`
	StatusMessage string    `json:"statusMessage"`
	StatusAt      time.Time `json:"statusAt"`

	RepoUrl     string `json:"repoUrl"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	PackageMode string `json:"packageMode"`

	App *Application `json:"app,omitempty"`
	Env *Environment `json:"env,omitempty"`
}
