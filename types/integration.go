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

type IntegrationStatus struct {
	IsConnected   bool                      `json:"isConnected"`
	Error         string                    `json:"error,omitempty"`
	Data          map[string]string         `json:"data"`
	SlackChannels []IntegrationSlackChannel `json:"slackChannels"`
}

type IntegrationSlackChannel struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	IsPrivate     bool   `json:"isPrivate"`
	IsIM          bool   `json:"isIm"`
	ContextTeamId string `json:"contextTeamId"`
}
