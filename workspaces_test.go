package api

import (
	"fmt"
	"testing"
)

func TestWorkspaces(t *testing.T) {
	w := Workspaces{Client: &Client{Config: DefaultConfig()}}
	workspace, err := w.Get(1, 1, 1)
	fmt.Println(workspace, err)
}
