package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

type ModuleVersions struct {
	Client *Client
}

func (mv ModuleVersions) basePath(orgName, moduleName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s/versions", orgName, moduleName)
}

func (mv ModuleVersions) path(orgName, moduleName, version string) string {
	return fmt.Sprintf("orgs/%s/modules/%s/versions/%s", orgName, moduleName, version)
}

func (mv ModuleVersions) downloadPath(orgName, moduleName, versionName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s/versions/%s/download", orgName, moduleName, versionName)
}

func (mv ModuleVersions) Get(ctx context.Context, orgName, moduleName, version string) (*types.ModuleVersion, error) {
	res, err := mv.Client.Do(ctx, http.MethodGet, mv.path(orgName, moduleName, version), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.ModuleVersion](res)
}

func (mv ModuleVersions) List(ctx context.Context, orgName, moduleName string) ([]types.ModuleVersion, error) {
	res, err := mv.Client.Do(ctx, http.MethodGet, mv.basePath(orgName, moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.ModuleVersion](res)
}

func (mv ModuleVersions) GetDownloadInfo(ctx context.Context, orgName, moduleName string, versionName string) (*types.ModuleDownloadInfo, error) {
	relativePath := mv.downloadPath(orgName, moduleName, versionName)
	res, err := mv.Client.Do(ctx, http.MethodHead, relativePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if err := response.Verify(res); err != nil {
		if response.IsNotFoundError(err) {
			return nil, nil
		}
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

	endpoint, err := mv.Client.Config.ConstructUrl(relativePath, nil)
	if err != nil {
		return nil, err
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

func (mv ModuleVersions) Download(ctx context.Context, orgName, moduleName, versionName string, file io.Writer) error {
	res, err := mv.Client.Do(ctx, http.MethodGet, mv.downloadPath(orgName, moduleName, versionName), nil, nil, nil)
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

func (mv ModuleVersions) Create(ctx context.Context, orgName, moduleName string, manifest types.ModuleManifest, versionName string, file io.Reader) error {
	query := url.Values{}
	query.Set("version", versionName)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add the module file
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s@%s.tgz", moduleName, versionName))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Add the manifest file
	manifestPart, err := writer.CreateFormFile("manifest", ".nullstone/module.json")
	if err != nil {
		return err
	}
	if err := json.NewEncoder(manifestPart).Encode(manifest); err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	var headers = map[string]string{}
	headers["Content-Type"] = writer.FormDataContentType()
	res, err := mv.Client.Do(ctx, http.MethodPost, mv.basePath(orgName, moduleName), query, headers, body)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
