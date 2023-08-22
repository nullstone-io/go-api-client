package types

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

type Environment struct {
	IdModel
	Type           EnvironmentType `json:"type"`
	Name           string          `json:"name"`
	OrgName        string          `json:"orgName"`
	StackId        int64           `json:"stackId"`
	Reference      string          `json:"reference"`
	ProviderConfig ProviderConfig  `json:"providerConfig"`
	PipelineOrder  *int            `json:"pipelineOrder,omitempty"`
	ContextKey     string          `json:"contextKey"`
	Status         EnvStatus       `json:"status"`
	IsProd         bool            `json:"isProd"`
}
