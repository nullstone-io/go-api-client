package types

import (
	"path"
	"strings"
)

// CleanRepoPath normalizes a path written in a build config so the same location is spelled one way:
// backslashes become forward slashes and the result is cleaned (so "./web/" and "web" are equal).
// An empty value cleans to ".".
func CleanRepoPath(value string) string {
	slashed := strings.ReplaceAll(strings.TrimSpace(value), `\`, "/")
	return path.Clean(slashed)
}

// JoinRepoPath joins path segments the way the build service resolves them, with each segment cleaned
// as in CleanRepoPath. Segments are joined in order, so JoinRepoPath("web", "Dockerfile") is "web/Dockerfile".
func JoinRepoPath(segments ...string) string {
	cleaned := make([]string, 0, len(segments))
	for _, segment := range segments {
		cleaned = append(cleaned, CleanRepoPath(segment))
	}
	return CleanRepoPath(path.Join(cleaned...))
}

// IsRelativePath reports whether value is a relative path: not empty, not an absolute path, not a drive
// letter or a url. It may climb with "..": a dockerfile is relative to its docker context and may sit
// beside it (docker_context `web`, dockerfile `../Dockerfile`), so it is checked as a whole with
// IsRepoRelativePath once the context is known (BuildConfig.DockerfilePath).
func IsRelativePath(value string) bool {
	if strings.TrimSpace(value) == "" {
		return false
	}
	slashed := strings.ReplaceAll(value, `\`, "/")
	// an absolute path, a drive letter or a url
	return !strings.HasPrefix(slashed, "/") && !strings.Contains(slashed, ":")
}

// IsRepoRelativePath reports whether value names a location inside the repository: it is a relative
// path (IsRelativePath) that does not climb above the repository root.
// It is the rule every path resolved from the repository root (docker_context, build_dir, publish_dir,
// workflow_filename, and the dockerfile once resolved against its context) must follow, whether it
// arrives through IaC, the API or the UI.
func IsRepoRelativePath(value string) bool {
	if !IsRelativePath(value) {
		return false
	}
	cleaned := CleanRepoPath(value)
	return cleaned != ".." && !strings.HasPrefix(cleaned, "../")
}
