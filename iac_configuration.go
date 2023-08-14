package api

import (
	"bytes"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type IacConfiguration struct {
	Client *Client
}

func (ic IacConfiguration) basePath() string {
	return "/iac_config"
}

func (ic IacConfiguration) Create(config string) (*types.EnvConfiguration, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if config != "" {
		part, err := writer.CreateFormFile("config", "config.yml")
		if err != nil {
			return nil, fmt.Errorf("unable to create form file: %w", err)
		}
		configReader := strings.NewReader(config)
		_, err = io.Copy(part, configReader)
		if err != nil {
			return nil, fmt.Errorf("unable to copy file contents into form file: %w", err)
		}
	}
	writer.Close()

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	res, err := ic.Client.Do(http.MethodPost, ic.basePath(), nil, headers, body)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.EnvConfiguration](res)
}
