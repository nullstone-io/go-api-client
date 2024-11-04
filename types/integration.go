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
	IsConnected   bool                      `json:"is_connected"`
	Error         string                    `json:"error,omitempty"`
	Data          map[string]string         `json:"data"`
	SlackChannels []IntegrationSlackChannel `json:"slack_channels"`
}

type IntegrationSlackChannel struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	IsPrivate     bool   `json:"is_private"`
	IsIM          bool   `json:"is_im"`
	ContextTeamId string `json:"context_team_id"`
}
