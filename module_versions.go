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
	"path"
	"strconv"
)

type ModuleVersions struct {
	Client *Client
}

func (mv ModuleVersions) List(moduleName string) ([]types.ModuleVersion, error) {
	res, err := mv.Client.Do(http.MethodGet, path.Join("modules", moduleName, "versions"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var moduleVersions []types.ModuleVersion
	if err := mv.Client.ReadJsonResponse(res, &moduleVersions); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return moduleVersions, nil
}

func (mv ModuleVersions) GetDownloadInfo(moduleName string, versionName string) (*types.ModuleDownloadInfo, error) {
	endpoint, err := mv.Client.Config.ConstructUrl(path.Join("modules", moduleName, "versions", versionName, "download"), nil)
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

func (mv ModuleVersions) Download(moduleName string, versionName string, file io.Writer) error {
	endpoint := path.Join("modules", moduleName, "versions", versionName, "download")
	res, err := mv.Client.Do(http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return err
	}

	if err := mv.Client.ReadFileResponse(res, file); response.IsNotFoundError(err) {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (mv ModuleVersions) Create(moduleName string, versionName string, file io.Reader) error {
	endpoint := path.Join("modules", moduleName, "versions")

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
	res, err := mv.Client.Do(http.MethodPost, endpoint, query, headers, body)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
