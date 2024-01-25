package types

import (
	"strings"
)

type ConnectionTarget struct {
	StackId   int64  `json:"stackId,omitempty" yaml:"stack_id,omitempty"`
	StackName string `json:"stackName,omitempty" yaml:"stack_name,omitempty"`
	BlockId   int64  `json:"blockId,omitempty" yaml:"block_id,omitempty"`
	BlockName string `json:"blockName,omitempty" yaml:"block_name,omitempty"`
	EnvId     *int64 `json:"envId,omitempty" yaml:"env_id,omitempty"`
	EnvName   string `json:"envName,omitempty" yaml:"env_name,omitempty"`
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

// Normalize
// Deprecated
func (t ConnectionTarget) Normalize(stackId int64, blocks Blocks, sharedEnvId int64) ConnectionTarget {
	result := t
	if result.StackId == 0 {
		result.StackId = stackId
	}
	if block := blocks.FindByName(result.BlockName); block != nil {
		result.BlockId = block.Id
		if block.IsShared && sharedEnvId != 0 {
			result.EnvId = &sharedEnvId
		}
	}
	return result
}

func (t ConnectionTarget) Match(other ConnectionTarget) bool {
	if t.StackId != other.StackId {
		return false
	}
	if t.BlockName != other.BlockName {
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
