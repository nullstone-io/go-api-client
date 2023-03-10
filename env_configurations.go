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

type EnvConfigurations struct {
	Client *Client
}

func (ec EnvConfigurations) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/configuration", ec.Client.Config.OrgName, stackId, envId)
}

func (ec EnvConfigurations) Create(stackId, envId int64, config, overrides string) ([]types.Workspace, error) {
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
	if overrides != "" {
		part, err := writer.CreateFormFile("overrides", "previews.yml")
		if err != nil {
			return nil, fmt.Errorf("unable to create form file: %w", err)
		}
		overridesReader := strings.NewReader(overrides)
		_, err = io.Copy(part, overridesReader)
		if err != nil {
			return nil, fmt.Errorf("unable to copy file contents into form file: %w", err)
		}
	}
	writer.Close()

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	res, err := ec.Client.Do(http.MethodPost, ec.basePath(stackId, envId), nil, headers, body)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Workspace](res)
}
