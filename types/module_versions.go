package types

import (
	"golang.org/x/mod/semver"
	"sort"
	"strings"
)

type ModuleVersions []ModuleVersion

var _ sort.Interface = ModuleVersions{}

func (s ModuleVersions) Len() int { return len(s) }
func (s ModuleVersions) Less(i, j int) bool {
	return semver.Compare(validSemver(s[i].Version), validSemver(s[j].Version)) < 0
}
func (s ModuleVersions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func validSemver(version string) string {
	if strings.HasPrefix(version, "v") {
		return version
	}
	return "v" + version
}

func (s ModuleVersions) FindLatest() *ModuleVersion {
	sort.Sort(sort.Reverse(s)) // "latest" will be at the beginning now
	for _, mv := range s {
		// Module Versions with build components are ignored from "latest"
		if build := semver.Build(NormalizedVersion(mv.Version)); build == "" {
			return &mv
		}
	}
	return nil
}
