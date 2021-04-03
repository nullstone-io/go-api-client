package types

type Stack struct {
	IdModel
	Name        string `json:"name"`
	OrgName     string `json:"orgName"`
	Description string `json:"description"`
}
