package types

type Stack struct {
	IdModel
	Reference    string `json:"reference"`
	Name         string `json:"name"`
	OrgName      string `json:"orgName"`
	Description  string `json:"description"`
	ProviderType string `json:"providerType"`
}
