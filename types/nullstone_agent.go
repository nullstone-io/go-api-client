package types

type NullstoneAgent struct {
	Aws   NullstoneAgentInfoAws   `json:"aws"`
	Gcp   NullstoneAgentInfoGcp   `json:"gcp"`
	Azure NullstoneAgentInfoAzure `json:"azure"`
	// IpAddresses is the list of public IP addresses that the Nullstone Agent connects from.
	// Use these to allow the agent through network-level access controls (e.g. security groups, firewall rules).
	IpAddresses []string `json:"ipAddresses"`
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
