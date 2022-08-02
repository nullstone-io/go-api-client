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

	if moduleSourceVersion == "latest" {
		return module.Versions.FindLatest(), nil
	}

	sort.Sort(sort.Reverse(module.Versions))
	for _, mv := range module.Versions {
		if semver.Compare(types.NormalizedVersion(mv.Version), moduleSourceVersion) == 0 {
			return &mv, nil
		}
	}
	return nil, nil
}
