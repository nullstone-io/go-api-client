package types

import (
	"github.com/nullstone-io/module/config"
)

type Connection struct {
	config.Connection `json:",inline"`

	// DesiredTarget refers to the ConnectionTarget that fulfills this connection
	// This value is input by the user via UI or IaC and is not fully qualified
	// This usually contains StackId, StackName, BlockId, and BlockName
	// It *can* contain EnvId and EnvName if different from the owning workspace
	DesiredTarget *ConnectionTarget `json:"desiredTarget"`
	// EffectiveTarget refers to the ConnectionTarget that fulfills this connection
	// This value is a fully normalized, effective version of DesiredTarget
	// All fields must be specified
	EffectiveTarget *ConnectionTarget `json:"effectiveTarget"`
	// Unused signals that the connection is not used by the current module version
	// During promotion of a module into a new workspace, it's possible that the new version removes connections
	// If we removed those connections automatically, a user could face data loss that is unrecoverable
	// Instead, this field was added to signal to the user that they should remove the connection
	Unused bool `json:"unused"`

	// OldTarget refers to the ConnectionTarget that fulfills this connection
	// This value is input by the user via UI or IaC and is not normalized
	// Deprecated
	OldTarget string `json:"target"`
	// OldReference refers to the old Reference field
	// Deprecated
	OldReference *ConnectionTarget `json:"reference"`
}

func (c *Connection) Equal(other Connection) bool {
	return c.SchemaEquals(other) &&
		c.Unused == other.Unused &&
		isConnectionTargetEqual(c.DesiredTarget, other.DesiredTarget) &&
		isConnectionTargetEqual(c.EffectiveTarget, other.EffectiveTarget)
}

func (c *Connection) SchemaEquals(other Connection) bool {
	if c == nil {
		return false
	}
	s1 := c.Connection
	s2 := other.Connection
	return s1.Type == s2.Type &&
		s1.Contract == s2.Contract &&
		s1.Optional == s2.Optional
}

func (c *Connection) TargetEquals(other Connection) bool {
	return isConnectionTargetEqual(c.DesiredTarget, other.DesiredTarget) &&
		isConnectionTargetEqual(c.EffectiveTarget, other.EffectiveTarget)
}

type ConnectionInput struct {
	Name          string            `json:"name"`
	DesiredTarget *ConnectionTarget `json:"desiredTarget"`
}
