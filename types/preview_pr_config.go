package types

type PreviewPRConfig struct {
	OrgName            string      `json:"orgName"`
	StackId            int64       `json:"stackId"`
	AppId              int64       `json:"appId"`
	IsEnabled          bool        `json:"isEnabled"`
	AutoLaunch         bool        `json:"autoLaunch"`
	AutoDeploy         bool        `json:"autoDeploy"`
	Status             string      `json:"status"`
	Repo               string      `json:"repo"`
	RepoUrl            string      `json:"repoUrl"`
	BuildConfig        BuildConfig `json:"buildConfig"`
	GitopsConnectionId *int64      `json:"gitopsConnectionId"`
}
