package types

const (
	PreviewPRConfigStatusActive   = "active"
	PreviewPRConfigStatusArchived = "archived"
)

type PreviewPRConfig struct {
	OrgName     string      `json:"orgName"`
	StackId     int64       `json:"stackId"`
	AppId       int64       `json:"appId"`
	IsEnabled   bool        `json:"isEnabled"`
	Repo        string      `json:"repo"`
	AutoLaunch  bool        `json:"autoLaunch"`
	AutoDeploy  bool        `json:"autoDeploy"`
	Status      string      `json:"status"`
	BuildConfig BuildConfig `json:"buildConfig"`
}

func (c PreviewPRConfig) IsAutoDeployEnabled() bool {
	return c.IsEnabled &&
		c.Status == PreviewPRConfigStatusActive &&
		c.AutoDeploy
}
