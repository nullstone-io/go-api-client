package types

type NullstoneAgent struct {
	Aws NullstoneAgentInfoAws `json:"aws"`
	Gcp NullstoneAgentInfoGcp `json:"gcp"`
}

type NullstoneAgentInfoAws struct {
	AccountId string `json:"accountId"`
	RoleName  string `json:"roleName"`
	RoleArn   string `json:"roleArn"`
}

type NullstoneAgentInfoGcp struct {
	ProjectId           string `json:"projectId"`
	ServiceAccountEmail string `json:"serviceAccountEmail"`
}
