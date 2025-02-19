package types

import (
	"github.com/nullstone-io/module/config"
)

type Connection struct {
	config.Connection `json:",inline"`

	// Target refers to the ConnectionTarget that fulfills this connection
	// This value is input by the user via UI or IaC and is not normalized
	Target *ConnectionTargetString `json:"target"`

	// EffectiveTarget refers to the ConnectionTarget that fulfills this connection
	// This value is a fully normalized, effective version of Target
	EffectiveTarget *ConnectionTarget `json:"effectiveTarget"`

	// Unused signals that the connection is not used by the current module version
	// During promotion of a module into a new workspace, it's possible that the new version removes connections
	// If we removed those connections automatically, a user could face data loss that is unrecoverable
	// Instead, this field was added to signal to the user that they should remove the connection
	Unused bool `json:"unused"`

	// OldReference refers to the old Reference field
	// Deprecated
	OldReference *ConnectionTarget `json:"reference"`
}

func (c *Connection) Equal(other Connection) bool {
	return c.SchemaEquals(other) &&
		c.Unused == other.Unused &&
		isConnectionTargetEqual(c.target(), other.target()) &&
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
	return isConnectionTargetEqual(c.target(), other.target()) &&
		isConnectionTargetEqual(c.EffectiveTarget, other.EffectiveTarget)
}

func (c *Connection) target() *ConnectionTarget {
	if c.Target == nil {
		return nil
	}
	return &c.Target.ConnectionTarget
}
