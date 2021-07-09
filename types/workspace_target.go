package types

import (
	"fmt"
)

type WorkspaceTarget struct {
	StackId int64 `json:"stackId"`
	BlockId int64 `json:"blockId"`
	EnvId   int64 `json:"envId"`
}

// Id is a string representation of the workspace target
func (t WorkspaceTarget) Id() string {
	return fmt.Sprintf("%d/%d/%d", t.StackId, t.BlockId, t.EnvId)
}

// FindRelativeConnection returns the PromotionResolveTarget based on the connection target
func (t WorkspaceTarget) FindRelativeConnection(connection ConnectionTarget) WorkspaceTarget {
	ref := WorkspaceTarget{
		StackId: connection.StackId,
		BlockId: connection.BlockId,
		EnvId:   t.EnvId,
	}
	if connection.EnvId != nil {
		ref.EnvId = *connection.EnvId
	}
	return ref
}
