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
