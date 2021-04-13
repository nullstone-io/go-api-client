package types

import (
	"fmt"
	"strings"
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

// FindRelativeConnection returns the PromotionResolveTarget based on the connection target
// The connection target can be defined as <stack>.<env>.<block>, <env>.<block>, or <block>
// If stack or env are not specified in the target value, these are pulled from the source target
func (t WorkspaceTarget) FindRelativeConnection(connection string) WorkspaceTarget {
	result := WorkspaceTarget{
		OrgName:   t.OrgName,
		StackName: t.StackName,
		BlockName: t.BlockName,
		EnvName:   t.EnvName,
	}

	tokens := strings.SplitN(connection, ".", 3)
	switch len(tokens) {
	case 3:
		result.StackName = tokens[0]
		result.EnvName = tokens[1]
		result.BlockName = tokens[2]
	case 2:
		result.EnvName = tokens[0]
		result.BlockName = tokens[1]
	case 1:
		result.BlockName = connection
	}

	return result
}
