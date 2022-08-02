package types

import (
	"golang.org/x/mod/semver"
	"sort"
)

type ModuleVersions []ModuleVersion

var _ sort.Interface = ModuleVersions{}

func (s ModuleVersions) Len() int { return len(s) }
func (s ModuleVersions) Less(i, j int) bool {
	return semver.Compare(ValidSemver(s[i].Version), ValidSemver(s[j].Version)) < 0
}
func (s ModuleVersions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Find searches the list of module versions and finds the most appropriate for the input constraint
// constraint="latest"                  -> returns the largest semver that does not contain a build component
// constraint="edge"                    -> (not supported yet)
// constraint="<major>.<minor>.<patch>" -> exact match
func (s ModuleVersions) Find(constraint string) *ModuleVersion {
	sort.Sort(sort.Reverse(s))
	if constraint == "latest" {
		return s.findLatest()
	}
	return s.findExact(ValidSemver(constraint))
}

// findLatest assumes that the module versions are sorted in reverse order (newest version first)
func (s ModuleVersions) findLatest() *ModuleVersion {
	for _, mv := range s {
		curSemver := ValidSemver(mv.Version)
		// Module Versions with build components are ignored from "latest"
		if semver.Build(curSemver) == "" {
			return &mv
		}
	}
	return nil
}

func (s ModuleVersions) findExact(validConstraint string) *ModuleVersion {
	for _, mv := range s {
		if semver.Compare(validConstraint, ValidSemver(mv.Version)) == 0 {
			return &mv
		}
	}
	return nil
}
