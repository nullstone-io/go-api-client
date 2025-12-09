package types

type NullstoneAgent struct {
	Aws NullstoneAgentInfoAws `json:"aws"`
	Gcp NullstoneAgentInfoGcp `json:"gcp"`
}

type NullstoneAgentInfoAws struct {
	AccountId string `json:"accountId"`
	UserName  string `json:"userName"`
	UserArn   string `json:"userArn"`
}

type NullstoneAgentInfoGcp struct {
	ProjectId           string `json:"projectId"`
	ServiceAccountEmail string `json:"serviceAccountEmail"`
}
