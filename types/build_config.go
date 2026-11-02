package types

const (
	BuilderManual        = "manual"
	BuilderNullstone     = "nullstone"
	BuilderGithubActions = "github-actions"
)

const (
	BuildPackageModeNone       = "none"
	BuildPackageModeDockerfile = "dockerfile"
	BuildPackageModeBuildpacks = "buildpacks"
	BuildPackageModeZip        = "zip"
	BuildPackageModeSiteAssets = "site-assets"
)

// DefaultBuildConfig produces the build configuration used when neither a workspace
// nor its block template defines one. The package mode follows the module:
// zip subplatform => zip, static-site subcategory => site-assets, otherwise dockerfile.
// A nil module yields the dockerfile default.
func DefaultBuildConfig(module *Module) BuildConfig {
	packageMode := BuildPackageModeDockerfile
	if module != nil {
		if module.Subplatform == "zip" {
			packageMode = BuildPackageModeZip
		} else if module.Subcategory == SubcategoryAppStaticSite {
			packageMode = BuildPackageModeSiteAssets
		}
	}
	return BuildConfig{
		Builder:          BuilderNullstone,
		PackageMode:      packageMode,
		Dockerfile:       "Dockerfile",
		DockerContext:    ".",
		BuildDir:         "./",
		InstallDepsCmd:   "npm ci",
		BuildCmd:         "npm run build",
		PublishDir:       "./dist",
		WorkflowFilename: ".github/workflows/build.yml",
		InitiateWorkflow: true,
		PerformPush:      false,
		WorkflowInputs:   map[string]string{},
	}
}

type BuildConfig struct {
	// Builder defines which build engine to use (currently Nullstone or Github Actions)
	Builder string `json:"builder"`

	// PackageMode tells the engine which build/package engine to use
	// e.g., dockerfile, buildpacks
	PackageMode string `json:"packageMode"`

	// Docker configurations
	Dockerfile    string `json:"dockerfile,omitempty"`
	DockerContext string `json:"dockerContext,omitempty"`

	// Static Site configurations
	BuildDir       string `json:"buildDir,omitempty"`
	InstallDepsCmd string `json:"installDepsCmd,omitempty"`
	BuildCmd       string `json:"buildCmd,omitempty"`

	// Zip and Static Site configurations
	PublishDir string `json:"publishDir,omitempty"`

	// Github Actions configurations
	WorkflowFilename string            `json:"workflowFilename,omitempty"`
	InitiateWorkflow bool              `json:"initiateWorkflow,omitempty"`
	PerformPush      bool              `json:"performPush,omitempty"`
	WorkflowInputs   map[string]string `json:"workflowInputs,omitempty"`

	// Environment variables injected at build-time
	BuildEnvVars map[string]string `json:"buildEnvVars,omitempty"`
}
