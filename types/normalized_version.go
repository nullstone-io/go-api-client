package types

import (
	"fmt"
	"strings"
)

// NormalizedVersion returns a version with the "v" prefix
// This is necessary to use the functions contained in golang.org/x/mod/semver
func NormalizedVersion(v string) string {
	return fmt.Sprintf("v%s", strings.TrimPrefix(v, "v"))
}
