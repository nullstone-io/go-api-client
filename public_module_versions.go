package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"net/http"
	"path"
	"strconv"
)

type PublicModuleVersions struct {
	Client *Client
}

func (mv PublicModuleVersions) path(moduleName string) string {
	return path.Join("orgs", mv.Client.Config.OrgName, "public-modules", moduleName, "versions")
}

func (mv PublicModuleVersions) downloadPath(moduleName, versionName string) string {
	return path.Join("orgs", mv.Client.Config.OrgName, "public-modules", moduleName, "versions", versionName, "download")
}

func (mv PublicModuleVersions) List(moduleName string) ([]types.ModuleVersion, error) {
	res, err := mv.Client.Do(http.MethodGet, mv.path(moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var moduleVersions []types.ModuleVersion
	if err := response.ReadJson(res, &moduleVersions); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return moduleVersions, nil
}

func (mv PublicModuleVersions) GetDownloadInfo(moduleName string, versionName string) (*types.ModuleDownloadInfo, error) {
	endpoint, err := mv.Client.Config.ConstructUrl(mv.downloadPath(moduleName, versionName), nil)
	if err != nil {
		return nil, err
	}
	res, err := http.Head(endpoint.String())
	if err != nil {
		return nil, err
	}

	resErr := res.Header.Get("X-Error")
	if resErr != "" {
		return nil, fmt.Errorf("error getting artifact info: %s", resErr)
	}

	ext := res.Header.Get("X-File-Extension")
	if ext == "" {
		return nil, fmt.Errorf("missing 'X-File-Extension' header")
	}
	rawDataLength := res.Header.Get("X-Data-Length")
	if rawDataLength == "" {
		return nil, fmt.Errorf("missing 'X-Data-Length' header")
	}

	info := &types.ModuleDownloadInfo{
		FileExtension: ext,
		DownloadUrl:   *endpoint,
	}
	if info.DataLength, err = strconv.Atoi(rawDataLength); err != nil {
		return info, fmt.Errorf("invalid 'X-Data-Length' header: %w", err)
	}
	return info, nil
}

func (mv PublicModuleVersions) Download(moduleName string, versionName string, file io.Writer) error {
	res, err := mv.Client.Do(http.MethodGet, mv.downloadPath(moduleName, versionName), nil, nil, nil)
	if err != nil {
		return err
	}

	if err := response.ReadFile(res, file); response.IsNotFoundError(err) {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
