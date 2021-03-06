package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/artifacts"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Module(cfg api.Config, moduleSource string) (*types.Module, error) {
	ms, err := artifacts.ParseSource(moduleSource)
	if err != nil {
		return nil, err
	}

	client := api.Client{Config: cfg}
	module, err := client.Org(ms.OrgName).Modules().Get(ms.ModuleName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving module: %w", err)
	} else if module == nil {
		return nil, fmt.Errorf("module %q does not exist in organization %q", moduleSource, ms.OrgName)
	}
	return module, nil
}
