package find

import (
	"golang.org/x/mod/semver"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"sort"
)

func ModuleVersion(cfg api.Config, moduleSource, moduleSourceVersion string) (*types.ModuleVersion, error) {
	module, err := Module(cfg, moduleSource)
	if err != nil {
		return nil, err
	}
	if len(module.Versions) < 1 {
		return nil, nil
	}

	sort.Sort(sort.Reverse(module.Versions)) // "latest" will be at the beginning now
	if moduleSourceVersion == "latest" {
		return findLatestInSorted(module.Versions), nil
	}
	for _, mv := range module.Versions {
		if semver.Compare(mv.Version, moduleSourceVersion) == 0 {
			return &mv, nil
		}
	}
	return nil, nil
}

func findLatestInSorted(mvs types.ModuleVersions) *types.ModuleVersion {
	for _, mv := range mvs {
		// Module Versions with build components are ignored from "latest"
		if build := semver.Build(mv.Version); build == "" {
			return &mv
		}
	}
	return nil
}
