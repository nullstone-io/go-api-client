package types

import (
	"fmt"
	"strings"
)

var (
	_ fmt.Stringer = Connections{}
)

type Connections map[string]Connection

func (s Connections) DesiredTargets() ConnectionTargets {
	result := ConnectionTargets{}
	for k, c := range s {
		if c.DesiredTarget != nil {
			result[k] = *c.DesiredTarget
		} else {
			result[k] = ConnectionTarget{}
		}
	}
	return result
}

func (s Connections) EffectiveTargets() ConnectionTargets {
	result := ConnectionTargets{}
	for k, c := range s {
		if c.EffectiveTarget != nil {
			result[k] = *c.EffectiveTarget
		} else {
			result[k] = ConnectionTarget{}
		}
	}
	return result
}

func (s Connections) String() string {
	result := make([]string, 0)
	for name, c := range s {
		id := "(none)"
		if c.EffectiveTarget != nil {
			id = c.EffectiveTarget.Workspace().Id()
		}
		result = append(result, fmt.Sprintf("%s=%s", name, id))
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
