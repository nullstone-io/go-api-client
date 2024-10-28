package types

type IntegrationTool string

const (
	IntegrationToolSlack          IntegrationTool = "slack"
	IntegrationToolMicrosoftTeams IntegrationTool = "microsoft-teams"
	IntegrationToolDiscord        IntegrationTool = "discord"
	IntegrationToolWhatsapp       IntegrationTool = "whatsapp"
)

type Integration struct {
	IdModel `json:",inline"`
	OrgName string          `json:"orgName"`
	Tenant  string          `json:"tenant"`
	Tool    IntegrationTool `json:"tool"`
}
