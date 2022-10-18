package types

type Environment struct {
	IdModel
	Type           string         `json:"type"`
	Name           string         `json:"name"`
	Reference      string         `json:"reference"`
	OrgName        string         `json:"orgName"`
	StackId        int64          `json:"stackId"`
	ProviderConfig ProviderConfig `json:"providerConfig"`
	PipelineOrder  *int           `json:"pipelineOrder"`
}
