package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/artifacts"
)

func ModuleDownloadUrl(cfg api.Config, source, version string) (string, error) {
	ms, err := artifacts.ParseSource(source)
	if err != nil {
		return "", err
	}
	ms.OverrideBaseAddress(&cfg)

	client := api.Client{Config: cfg}
	info, err := client.ModuleVersions().GetDownloadInfo(ms.OrgName, ms.ModuleName, version)
	if err != nil {
		return "", fmt.Errorf("error retrieving artifact info: %w", err)
	} else if info == nil {
		return "", fmt.Errorf("module version %s@%s does not exist in registry", source, version)
	}
	return info.GetterUrl(), nil

}
