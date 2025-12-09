package types

type NullstoneAgent struct {
	Aws NullstoneAgentInfoAws
	Gcp NullstoneAgentInfoGcp
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
