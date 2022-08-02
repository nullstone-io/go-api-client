package find

import (
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func ModuleVersion(cfg api.Config, moduleSource, moduleSourceVersion string) (*types.ModuleVersion, error) {
	module, err := Module(cfg, moduleSource)
	if err != nil {
		return nil, err
	}
	if len(module.Versions) < 1 {
		return nil, nil
	}

	return module.Versions.Find(moduleSourceVersion), nil
}
