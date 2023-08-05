package types

import (
	"strings"
)

type ConnectionTarget struct {
	StackId   int64  `json:"stackId,omitempty"`
	StackName string `json:"stackName,omitempty"`
	BlockId   int64  `json:"blockId,omitempty"`
	BlockName string `json:"blockName,omitempty"`
	EnvId     *int64 `json:"envId,omitempty"`
	EnvName   string `json:"envName,omitempty"`
}

func ParseConnectionTarget(s string) ConnectionTarget {
	tokens := strings.Split(s, ".")
	switch len(tokens) {
	case 1:
		return ConnectionTarget{
			BlockName: tokens[0],
		}
	case 2:
		return ConnectionTarget{
			StackName: tokens[0],
			BlockName: tokens[1],
		}
	case 3:
		return ConnectionTarget{
			StackName: tokens[0],
			EnvName:   tokens[1],
			BlockName: tokens[2],
		}
	default:
		return ConnectionTarget{}
	}
}

func (t ConnectionTarget) Match(other ConnectionTarget) bool {
	if t.StackId != other.StackId {
		return false
	}
	if t.BlockId != other.BlockId {
		return false
	}
	if t.EnvId == nil {
		return other.EnvId == nil
	}
	if other.EnvId == nil {
		return false
	}
	return *t.EnvId == *other.EnvId
}

func (t ConnectionTarget) Workspace() WorkspaceTarget {
	wt := WorkspaceTarget{
		StackId: t.StackId,
		BlockId: t.BlockId,
	}
	if t.EnvId != nil {
		wt.EnvId = *t.EnvId
	}
	return wt
}
