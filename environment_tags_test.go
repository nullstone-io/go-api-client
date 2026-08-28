package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateEnvironmentTagsInput_ApplyTo(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name     string
		existing map[string]string
		input    UpdateEnvironmentTagsInput
		want     map[string]string
	}{
		{
			name:     "sets a new key",
			existing: map[string]string{"tier": "gold"},
			input:    UpdateEnvironmentTagsInput{Tags: map[string]*string{"claim": strPtr("brad")}},
			want:     map[string]string{"tier": "gold", "claim": "brad"},
		},
		{
			name:     "updates an existing key and leaves others untouched",
			existing: map[string]string{"tier": "gold", "claim": "brad"},
			input:    UpdateEnvironmentTagsInput{Tags: map[string]*string{"claim": strPtr("alex")}},
			want:     map[string]string{"tier": "gold", "claim": "alex"},
		},
		{
			name:     "nil value clears only that key",
			existing: map[string]string{"tier": "gold", "claim": "brad"},
			input:    UpdateEnvironmentTagsInput{Tags: map[string]*string{"claim": nil}},
			want:     map[string]string{"tier": "gold"},
		},
		{
			name:     "empty string is distinct from clearing",
			existing: map[string]string{"claim": "brad"},
			input:    UpdateEnvironmentTagsInput{Tags: map[string]*string{"claim": strPtr("")}},
			want:     map[string]string{"claim": ""},
		},
		{
			name:     "clearing a key that does not exist is a no-op",
			existing: map[string]string{"tier": "gold"},
			input:    UpdateEnvironmentTagsInput{Tags: map[string]*string{"claim": nil}},
			want:     map[string]string{"tier": "gold"},
		},
		{
			name:     "empty patch leaves everything untouched",
			existing: map[string]string{"tier": "gold"},
			input:    UpdateEnvironmentTagsInput{},
			want:     map[string]string{"tier": "gold"},
		},
		{
			name:     "applies to nil existing tags",
			existing: nil,
			input:    UpdateEnvironmentTagsInput{Tags: map[string]*string{"claim": strPtr("brad")}},
			want:     map[string]string{"claim": "brad"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.ApplyTo(test.existing)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestUpdateEnvironmentTagsInput_ApplyTo_DoesNotMutateExisting(t *testing.T) {
	strPtr := func(s string) *string { return &s }
	existing := map[string]string{"tier": "gold", "claim": "brad"}

	input := UpdateEnvironmentTagsInput{Tags: map[string]*string{
		"claim": nil,
		"env":   strPtr("preview"),
	}}
	_ = input.ApplyTo(existing)

	assert.Equal(t, map[string]string{"tier": "gold", "claim": "brad"}, existing)
}
