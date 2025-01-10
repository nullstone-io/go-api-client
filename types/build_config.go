package types

const (
	BuildPackageModeNone          = "none"
	BuildPackageModeDockerfile    = "dockerfile"
	BuildPackageModeBuildpacks    = "buildpacks"
	BuildPackageModeZip           = "zip"
	BuildPackageModeSiteAssets    = "site-assets"
	BuildPackageModeGithubActions = "github-actions"
)

type BuildConfig struct {
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

	// Environment variables injected at build-time
	BuildEnvVars map[string]string `json:"buildEnvVars,omitempty"`

	// Github Actions configurations
	WorkflowFileName string            `json:"workflowFileName,omitempty"`
	WorkflowInputs   map[string]string `json:"workflowInputs,omitempty"`
}
