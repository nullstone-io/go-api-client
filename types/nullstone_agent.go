package types

type NullstoneAgent struct {
	Aws   NullstoneAgentInfoAws   `json:"aws"`
	Gcp   NullstoneAgentInfoGcp   `json:"gcp"`
	Azure NullstoneAgentInfoAzure `json:"azure"`
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

type NullstoneAgentInfoAzure struct {
	OidcIssuerUrl string `json:"oidcIssuerUrl"`
	OidcAudience  string `json:"oidcAudience"`
}
