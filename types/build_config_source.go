package types

// BuildConfigSource names the IaC file that last wrote a build config.
// It is nil on a build config that was last written through the UI or API, or that was never written.
// The next IaC sync of that file overwrites the build config again.
type BuildConfigSource struct {
	// Repo is the name of the repository that holds the file (e.g. acme/widgets)
	Repo string `json:"repo"`
	// RepoUrl is the url of the repository that holds the file (e.g. https://github.com/acme/widgets)
	RepoUrl string `json:"repoUrl"`
	// File is the IaC file that declared `build:` (e.g. .nullstone/config.yml, .nullstone/prod.yml)
	File string `json:"file"`
}
