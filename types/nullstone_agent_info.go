package types

type NullstoneAgentInfo struct {
	Aws NullstoneAgentInfoAws
	Gcp NullstoneAgentInfoGcp
}

type NullstoneAgentInfoAws struct {
	RoleAccountId string `json:"roleAccountId"`
	RoleName      string `json:"roleName"`
	RoleArn       string `json:"roleArn"`
}

type NullstoneAgentInfoGcp struct {
	ServiceAccountEmail string `json:"serviceAccountEmail"`
}
