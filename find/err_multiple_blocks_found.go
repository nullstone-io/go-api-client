package find

import (
	"fmt"
	"strings"
)

type ErrMultipleBlocksFound struct {
	BlockName  string
	StackNames []string
}

func (e ErrMultipleBlocksFound) Error() string {
	return fmt.Sprintf("found multiple blocks named %q located in the following stacks: %s\n use the stack param to select a specific block", e.BlockName, strings.Join(e.StackNames, ","))
}
