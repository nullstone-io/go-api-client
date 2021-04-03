package types

import (
	"fmt"
)

type WorkspaceTarget struct {
	OrgName   string `json:"orgName"`
	StackName string `json:"stackName"`
	BlockName string `json:"blockName"`
	EnvName   string `json:"envName"`
}

// Id is a string representation of the workspace target
func (t WorkspaceTarget) Id() string {
	return fmt.Sprintf("%s/%s/%s/%s", t.OrgName, t.StackName, t.BlockName, t.EnvName)
}
