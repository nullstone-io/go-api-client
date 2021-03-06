package types

import (
	"github.com/nullstone-io/module/config"
)

type Connections map[string]Connection

type Connection struct {
	config.Connection

	// Target refers to the block that fulfills the connection
	// If the Target is in the same stack, this is just the block name
	// If the Target is in another stack, this is the fully-qualified block name (i.e. {stack}.{env}.{block})
	Target string `json:"target"`

	// Reference refers to the block that fulfills the connection
	// TODO: Rename to Target once Target is deprecated
	Reference *ConnectionTarget `json:"reference"`

	// Unused signals that the connection is not used by the current module version
	// During promotion of a module into a new workspace, it's possible that the new version removes connections
	// If we removed those connections automatically, a user could face data loss that is unrecoverable
	// Instead, this field was added to signal to the user that they should remove the connection
	Unused bool `json:"unused"`
}
