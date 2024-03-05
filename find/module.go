package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/artifacts"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Module(ctx context.Context, cfg api.Config, moduleSource string) (*types.Module, error) {
	ms, err := artifacts.ParseSource(moduleSource)
	if err != nil {
		return nil, err
	}
	ms.OverrideBaseAddress(&cfg)

	client := api.Client{Config: cfg}
	module, err := client.Modules().Get(ctx, ms.OrgName, ms.ModuleName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving module: %w", err)
	} else if module == nil {
		return nil, fmt.Errorf("module %s does not exist in organization %s", ms.ModuleName, ms.OrgName)
	}
	return module, nil
}
