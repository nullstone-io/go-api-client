package types

type GitopsConnectionConfig struct {
	GitopsConnectionId int64  `json:"gitopsConnectionId"`
	EnvId              int64  `json:"envId"`
	IsEnabled          bool   `json:"isEnabled"`
	GitBranch          string `json:"gitBranch"`
	IsAutoApplyEnabled bool   `json:"isAutoApplyEnabled"`
}

func (c GitopsConnectionConfig) IsEnabledForBranch(branchName string) bool {
	return c.IsEnabled && c.GitBranch == branchName
}

type GitopsConnectionConfigs []GitopsConnectionConfig

func (s GitopsConnectionConfigs) EnabledOnBranch(branchName string) GitopsConnectionConfigs {
	result := make(GitopsConnectionConfigs, 0)
	for _, cur := range s {
		if cur.IsEnabledForBranch(branchName) {
			result = append(result, cur)
		}
	}
	return result
}
