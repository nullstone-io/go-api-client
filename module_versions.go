package api

import (
	"bytes"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
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
	if err := mv.Client.ReadJsonResponse(res, &moduleVersions); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return moduleVersions, nil
}

func (mv ModuleVersions) Download(moduleName string, versionName string, file io.Writer) error {
	endpoint := path.Join("modules", moduleName, "versions", versionName, "download")
	res, err := mv.Client.Do(http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return err
	}

	if err := mv.Client.ReadFileResponse(res, file); IsNotFoundError(err) {
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

	return mv.Client.VerifyResponse(res)
}
