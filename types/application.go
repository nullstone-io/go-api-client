package types

type Application struct {
	IdModel
	Name      string `json:"name"`
	OrgName   string `json:"orgName"`
	StackName string `json:"stackName"`
	Repo      string `json:"repo"`
	Framework string `json:"framework"`
	BlockId   int    `json:"blockId"`

	Block *Block `json:"block,omitempty"`
}
