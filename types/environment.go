package types

type Environment struct {
	IdModel
	Name           string         `json:"name"`
	Reference      string         `json:"reference"`
	OrgName        string         `json:"orgName"`
	StackId        int64          `json:"stackId"`
	StackName      string         `json:"stackName"`
	ProviderConfig ProviderConfig `json:"providerConfig"`
	PipelineOrder  int            `json:"pipelineOrder"`
}
