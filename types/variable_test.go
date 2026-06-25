package types

import (
	"github.com/nullstone-io/module/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariable_ValueEquals(t *testing.T) {
	varDef := config.Variable{
		Type:        "number",
		Description: "",
		Default:     256,
		Sensitive:   false,
	}
	tests := map[string]struct {
		a    Variable
		b    Variable
		want bool
	}{
		"a value does not equal b value": {
			a:    Variable{Variable: varDef, Value: 512},
			b:    Variable{Variable: varDef, Value: 1024},
			want: false,
		},
		"a default equals b value": {
			a:    Variable{Variable: varDef, Value: nil},
			b:    Variable{Variable: varDef, Value: 256},
			want: true,
		},
		"a default equals b default": {
			a:    Variable{Variable: varDef, Value: nil},
			b:    Variable{Variable: varDef, Value: nil},
			want: true,
		},
		"a value equals b default": {
			a:    Variable{Variable: varDef, Value: 256},
			b:    Variable{Variable: varDef, Value: nil},
			want: true,
		},
		"a value equals b value": {
			a:    Variable{Variable: varDef, Value: 512},
			b:    Variable{Variable: varDef, Value: 512},
			want: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			got := test.a.ValueEquals(test.b)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestVariable_ValueEquals_Object(t *testing.T) {
	// Effective config from the API arrives as JSON (numbers => float64),
	// while desired config from IaC may carry numbers as int. These are equal.
	varDef := config.Variable{Type: "object({ timeout_sec = number })"}
	tests := map[string]struct {
		a    Variable
		b    Variable
		want bool
	}{
		"nested object: float64 equals int": {
			a: Variable{Variable: varDef, Value: map[string]any{
				"connection_draining": map[string]any{"timeout_sec": float64(60)},
				"timeout_sec":         float64(600),
			}},
			b: Variable{Variable: varDef, Value: map[string]any{
				"connection_draining": map[string]any{"timeout_sec": 60},
				"timeout_sec":         600,
			}},
			want: true,
		},
		"nested object: differing number": {
			a: Variable{Variable: varDef, Value: map[string]any{"timeout_sec": float64(600)}},
			b: Variable{Variable: varDef, Value: map[string]any{"timeout_sec": 300}},
			want: false,
		},
		"list: float64 equals int": {
			a:    Variable{Variable: varDef, Value: []any{float64(1), float64(2)}},
			b:    Variable{Variable: varDef, Value: []any{1, 2}},
			want: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			got := test.a.ValueEquals(test.b)
			assert.Equal(t, test.want, got)
		})
	}
}
