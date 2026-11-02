package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBuildConfig(t *testing.T) {
	tests := []struct {
		name            string
		module          *Module
		wantPackageMode string
	}{
		{name: "nil module", module: nil, wantPackageMode: BuildPackageModeDockerfile},
		{name: "container module", module: &Module{Subcategory: SubcategoryAppContainer, Subplatform: "ecs"}, wantPackageMode: BuildPackageModeDockerfile},
		{name: "zip subplatform", module: &Module{Subcategory: SubcategoryAppServerless, Subplatform: "zip"}, wantPackageMode: BuildPackageModeZip},
		{name: "static site", module: &Module{Subcategory: SubcategoryAppStaticSite, Subplatform: "s3"}, wantPackageMode: BuildPackageModeSiteAssets},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := DefaultBuildConfig(test.module)
			assert.Equal(t, BuilderNullstone, got.Builder)
			assert.Equal(t, test.wantPackageMode, got.PackageMode)
			assert.Equal(t, "Dockerfile", got.Dockerfile)
			assert.Equal(t, ".", got.DockerContext)
			assert.Equal(t, ".github/workflows/build.yml", got.WorkflowFilename)
			assert.True(t, got.InitiateWorkflow)
			assert.False(t, got.PerformPush)
			assert.NotNil(t, got.WorkflowInputs)
		})
	}
}
