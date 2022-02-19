package types

import (
	"golang.org/x/mod/semver"
	"sort"
)

type ModuleVersions []ModuleVersion

var _ sort.Interface = ModuleVersions{}

func (s ModuleVersions) Len() int           { return len(s) }
func (s ModuleVersions) Less(i, j int) bool { return semver.Compare(s[i].Version, s[j].Version) < 0 }
func (s ModuleVersions) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
