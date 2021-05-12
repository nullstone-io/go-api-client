package types

type Environment struct {
	IdModel
	Name           string         `json:"name"`
	OrgName        string         `json:"orgName"`
	StackName      string         `json:"stackName"`
	PipelineOrder  int            `json:"pipelineOrder"`
	ProviderConfig ProviderConfig `json:"providerConfig"`
}
