package types

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectionTargetString_UnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		input string
		want  *ConnectionTargetString
	}{
		"missing": {
			input: "{}",
			want:  nil,
		},
		"block": {
			input: `{"target":"network0"}`,
			want: &ConnectionTargetString{
				ConnectionTarget: ConnectionTarget{
					BlockName: "network0",
				},
			},
		},
		"stack-block": {
			input: `{"target":"primary.network0"}`,
			want: &ConnectionTargetString{
				ConnectionTarget: ConnectionTarget{
					StackName: "primary",
					BlockName: "network0",
				},
			},
		},
		"full": {
			input: `{"target":"primary.dev.network0"}`,
			want: &ConnectionTargetString{
				ConnectionTarget: ConnectionTarget{
					StackName: "primary",
					BlockName: "network0",
					EnvName:   "dev",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got struct {
				Target *ConnectionTargetString `json:"target"`
			}
			err := json.Unmarshal([]byte(test.input), &got)
			assert.NoError(t, err, "unexpected error")
			assert.Equal(t, test.want, got.Target)
		})
	}
}

func TestConnectionTargetString_MarshalJSON(t *testing.T) {
	tests := map[string]struct {
		input *ConnectionTargetString
		want  string
	}{
		"missing": {
			input: nil,
			want:  "{}",
		},
		"block": {
			input: &ConnectionTargetString{
				ConnectionTarget: ConnectionTarget{
					BlockName: "network0",
				},
			},
			want: `{"target":"network0"}`,
		},
		"stack-block": {
			input: &ConnectionTargetString{
				ConnectionTarget: ConnectionTarget{
					StackName: "primary",
					BlockName: "network0",
				},
			},
			want: `{"target":"primary.network0"}`,
		},
		"full": {
			input: &ConnectionTargetString{
				ConnectionTarget: ConnectionTarget{
					StackName: "primary",
					BlockName: "network0",
					EnvName:   "dev",
				},
			},
			want: `{"target":"primary.dev.network0"}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			type wrap struct {
				Target *ConnectionTargetString `json:"target,omitempty"`
			}
			got, err := json.Marshal(wrap{Target: test.input})
			assert.NoError(t, err, "unexpected error")
			assert.Equal(t, test.want, string(got))
		})
	}
}
