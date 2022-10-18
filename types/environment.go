package types

type EnvironmentType string

const (
	EnvTypePipeline EnvironmentType = "PipelineEnv"
	EnvTypePreview  EnvironmentType = "PreviewEnv"
	EnvTypeGlobal                   = "GlobalEnv"
)

type Environment struct {
	IdModel
	Type           string          `json:"type"`
	Name           EnvironmentType `json:"name"`
	Reference      string          `json:"reference"`
	OrgName        string          `json:"orgName"`
	StackId        int64           `json:"stackId"`
	ProviderConfig ProviderConfig  `json:"providerConfig"`
	PipelineOrder  *int            `json:"pipelineOrder"`
}
