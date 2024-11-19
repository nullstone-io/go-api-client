package types

type ConnectionTargets map[string]ConnectionTarget

type ConnectionTarget struct {
	StackId   int64  `json:"stackId,omitempty" yaml:"stack_id,omitempty"`
	StackName string `json:"stackName,omitempty" yaml:"stack_name,omitempty"`
	BlockId   int64  `json:"blockId,omitempty" yaml:"block_id,omitempty"`
	BlockName string `json:"blockName,omitempty" yaml:"block_name,omitempty"`
	EnvId     *int64 `json:"envId,omitempty" yaml:"env_id,omitempty"`
	EnvName   string `json:"envName,omitempty" yaml:"env_name,omitempty"`
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

// isConnectionTargetEqual performs equality check on ConnectionTarget
// This assumes that the ConnectionTarget has Id+Name populated for Stack/Block/Env
func isConnectionTargetEqual(target1 *ConnectionTarget, target2 *ConnectionTarget) bool {
	if target1 == nil {
		if target2 == nil {
			return true
		}
		return false
	}
	if target2 == nil {
		return false
	}
	return target1.StackId == target2.StackId &&
		target1.StackName == target2.StackName &&
		target1.BlockId == target2.BlockId &&
		target1.BlockName == target2.BlockName &&
		isConnectionTargetEnvEqual(*target1, *target2)
}

func isConnectionTargetEnvEqual(t1 ConnectionTarget, t2 ConnectionTarget) bool {
	if t1.EnvId != nil && t2.EnvId != nil {
		return *t1.EnvId == *t2.EnvId
	}
	return t1.EnvName == t2.EnvName
}
