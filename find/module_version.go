package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/artifacts"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func ModuleVersion(ctx context.Context, cfg api.Config, moduleSource, moduleSourceVersion string) (*types.ModuleVersion, error) {
	ms, err := artifacts.ParseSource(moduleSource)
	if err != nil {
		return nil, err
	}
	ms.OverrideBaseAddress(&cfg)

	client := api.Client{Config: cfg}
	versions, err := client.ModuleVersions().List(ctx, ms.OrgName, ms.ModuleName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving module versions: %w", err)
	} else if versions == nil {
		return nil, fmt.Errorf("module %s does not exist in organization %s", ms.ModuleName, ms.OrgName)
	} else if len(versions) < 1 {
		return nil, nil
	}

	return types.ModuleVersions(versions).Find(moduleSourceVersion), nil
}
