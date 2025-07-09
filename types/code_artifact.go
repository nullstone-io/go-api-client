package types

type CodeArtifact struct {
	IdModel

	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	AppId   int64  `json:"appId"`
	EnvId   int64  `json:"envId"`

	Version    string     `json:"version"`
	CommitInfo CommitInfo `json:"commitInfo"`
}
