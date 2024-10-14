package types

import (
	"fmt"
	"strings"
)

var (
	_ fmt.Stringer = Connections{}
)

type Connections map[string]Connection

func (s Connections) Targets() ConnectionTargets {
	result := ConnectionTargets{}
	for k, c := range s {
		if c.Reference != nil {
			result[k] = *c.Reference
		} else {
			result[k] = ConnectionTarget{}
		}
	}
	return result
}

func (s Connections) String() string {
	result := make([]string, 0)
	for name, c := range s {
		result = append(result, fmt.Sprintf("%s=%s", name, c.Reference.Workspace().Id()))
	}
	return strings.Join(result, ",")
}

func (s Connections) Equal(other Connections) bool {
	if s == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	if len(s) != len(other) {
		return false
	}
	for k, conn := range s {
		if otherConn, ok := other[k]; !ok {
			return false
		} else if !conn.Equal(otherConn) {
			return false
		}
	}
	return true
}
