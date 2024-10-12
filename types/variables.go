package types

import (
	"github.com/nullstone-io/module/config"
)

type Variables map[string]Variable

type Variable struct {
	config.Variable `json:",inline"`

	// Value is the exact value set for this variable
	// This Value can be nearly any data type and is determined by Variable Type
	Value interface{} `json:"value"`

	// Unused signals that the variable is not used by the current module version
	// During promotion of a module into a new workspace, it's possible that the new version removes variables
	// If we removed those variables automatically, a user could face data loss that is unrecoverable
	// Instead, this field was added to signal to the user that they should remove the variable
	Unused bool `json:"unused"`
}

func (v Variable) HasValue() bool {
	return v.Value != nil && !v.Unused
}

type VariableInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
