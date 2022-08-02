package types

import (
	"github.com/nullstone-io/module/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModuleVersions_Find(t *testing.T) {
	mvs := ModuleVersions{
		ModuleVersion{
			Version:  "0.0.1",
			Manifest: config.Manifest{},
		},
		ModuleVersion{
			Version:  "0.0.2",
			Manifest: config.Manifest{},
		},
		ModuleVersion{
			Version:  "0.0.3+759234ac",
			Manifest: config.Manifest{},
		},
		ModuleVersion{
			Version:  "0.0.3",
			Manifest: config.Manifest{},
		},
		ModuleVersion{
			Version:  "0.1.0",
			Manifest: config.Manifest{},
		},
		ModuleVersion{
			Version:  "0.1.0+7d34ac3f",
			Manifest: config.Manifest{},
		},
	}

	tests := []struct {
		constraint string
		// Use wantVersion = "" if expecting to not find version
		wantVersion string
	}{
		{
			constraint:  "latest",
			wantVersion: "0.1.0",
		},
		{
			constraint:  "0.1.0",
			wantVersion: "0.1.0",
		},
		{
			constraint:  "0.0.3+759234ac",
			wantVersion: "0.0.3+759234ac",
		},
		{
			constraint:  "0.10.0",
			wantVersion: "",
		},
	}

	for _, test := range tests {
		t.Run(test.constraint, func(t *testing.T) {
			got := mvs.Find(test.constraint)
			if test.wantVersion == "" {
				assert.Nil(t, got)
			} else if assert.NotNil(t, got) {
				assert.Equal(t, test.wantVersion, got.Version)
			}
		})
	}
}
