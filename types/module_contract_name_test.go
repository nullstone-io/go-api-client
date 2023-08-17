package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModuleContractName_Match(t *testing.T) {
	tests := []struct {
		name      string
		mustMatch ModuleContractName
		candidate ModuleContractName
	}{
		{
			name: "match capability",
			mustMatch: ModuleContractName{
				Category: "capability",
				Provider: "aws",
				Platform: "*",
			},
			candidate: ModuleContractName{
				Category:    "capability",
				Subcategory: "ingress",
				Provider:    "aws",
				Platform:    "alb",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.mustMatch.Match(test.candidate)
			assert.True(t, result)
		})
	}
}
