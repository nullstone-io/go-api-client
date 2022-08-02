package types

import (
	"strings"
)

// ValidSemver returns a version with the "v" prefix
// This is necessary to use the functions contained in golang.org/x/mod/semver
func ValidSemver(version string) string {
	if strings.HasPrefix(version, "v") {
		return version
	}
	return "v" + version
}
