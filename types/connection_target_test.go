package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection_TargetEquals(t *testing.T) {
	ptr := func(id int64) *int64 {
		return &id
	}

	tests := map[string]struct {
		a    *ConnectionTarget
		b    *ConnectionTarget
		want bool
	}{
		"a stack-id does not equal b stack-id": {
			a:    &ConnectionTarget{StackId: 1, BlockName: "block"},
			b:    &ConnectionTarget{StackId: 2, BlockName: "block"},
			want: false,
		},
		"a env-id does not equal b env-id": {
			a:    &ConnectionTarget{BlockName: "block", EnvId: ptr(1)},
			b:    &ConnectionTarget{BlockName: "block", EnvId: ptr(2)},
			want: false,
		},
		"a missing env-id equals b on env-name": {
			a:    &ConnectionTarget{BlockName: "block", EnvId: nil, EnvName: "dev"},
			b:    &ConnectionTarget{BlockName: "block", EnvId: ptr(2), EnvName: "dev"},
			want: true,
		},
		"a does not equal missing b": {
			a:    &ConnectionTarget{BlockId: 10, BlockName: "block"},
			b:    nil,
			want: false,
		},
		"missing a does not equal b": {
			a:    nil,
			b:    &ConnectionTarget{BlockId: 10, BlockName: "block"},
			want: false,
		},
		"a equals b": {
			a:    &ConnectionTarget{StackId: 1, StackName: "core", BlockId: 10, BlockName: "block", EnvId: ptr(110), EnvName: "dev"},
			b:    &ConnectionTarget{StackId: 1, StackName: "core", BlockId: 10, BlockName: "block", EnvId: ptr(110), EnvName: "dev"},
			want: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			aconn := &Connection{EffectiveTarget: test.a}
			bconn := &Connection{EffectiveTarget: test.b}
			got := aconn.TargetEquals(*bconn)
			assert.Equal(t, test.want, got)
		})
	}
}
