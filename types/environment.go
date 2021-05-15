package types

type Environment struct {
	IdModel
	Name           string         `json:"name"`
	Reference      string         `json:"reference"`
	OrgName        string         `json:"orgName"`
	StackName      string         `json:"stackName"`
	ProviderConfig ProviderConfig `json:"providerConfig"`
	PipelineOrder  int            `json:"pipelineOrder"`
}
