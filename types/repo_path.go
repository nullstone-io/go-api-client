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

// IsRepoRelativePath reports whether value names a location inside the repository: it is not empty,
// not an absolute path, not a drive letter or a url, and does not climb above the repository root.
// It is the rule every build config path (dockerfile, docker_context, build_dir, publish_dir, ...)
// must follow, whether it arrives through IaC, the API or the UI.
func IsRepoRelativePath(value string) bool {
	if strings.TrimSpace(value) == "" {
		return false
	}
	slashed := strings.ReplaceAll(value, `\`, "/")
	if strings.HasPrefix(slashed, "/") || strings.Contains(slashed, ":") {
		// an absolute path, a drive letter or a url
		return false
	}
	cleaned := path.Clean(slashed)
	return cleaned != ".." && !strings.HasPrefix(cleaned, "../")
}
