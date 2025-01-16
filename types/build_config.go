package types

const (
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
	WorkflowInputs   map[string]string `json:"workflowInputs,omitempty"`

	// Environment variables injected at build-time
	BuildEnvVars map[string]string `json:"buildEnvVars,omitempty"`
}
