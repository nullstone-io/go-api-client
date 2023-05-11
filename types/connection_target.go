package types

import "strings"

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
	if t.EnvId != other.EnvId {
		return false
	}
	if t.BlockId == 0 || other.BlockId == 0 {
		return t.BlockName == other.BlockName
	}
	return t.BlockId == other.BlockId
}
