package api

import (
	"bytes"
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

func (mv ModuleVersions) path(moduleName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s/versions", mv.Client.Config.OrgName, moduleName)
}

func (mv ModuleVersions) downloadPath(moduleName, versionName string) string {
	return fmt.Sprintf("orgs/%s/modules/%s/versions/%s/download", mv.Client.Config.OrgName, moduleName, versionName)
}

func (mv ModuleVersions) List(moduleName string) ([]types.ModuleVersion, error) {
	res, err := mv.Client.Do(http.MethodGet, mv.path(moduleName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.ModuleVersion](res)
}

func (mv ModuleVersions) GetDownloadInfo(moduleName string, versionName string) (*types.ModuleDownloadInfo, error) {
	relativePath := mv.downloadPath(moduleName, versionName)
	res, err := mv.Client.Do(http.MethodHead, relativePath, nil, nil, nil)
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

func (mv ModuleVersions) Download(moduleName string, versionName string, file io.Writer) error {
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

func (mv ModuleVersions) Create(moduleName string, versionName string, file io.Reader) error {
	query := url.Values{}
	query.Set("version", versionName)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s@%s.tgz", moduleName, versionName))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	var headers = map[string]string{}
	headers["Content-Type"] = writer.FormDataContentType()
	res, err := mv.Client.Do(http.MethodPost, mv.path(moduleName), query, headers, body)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
