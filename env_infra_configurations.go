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

type EnvInfraConfigurations struct {
	Client *Client
}

func (ec EnvInfraConfigurations) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/configuration", ec.Client.Config.OrgName, stackId, envId)
}

func (ec EnvInfraConfigurations) Create(stackId, envId int64, config map[string]string) (*types.WorkspaceLaunchNeeds, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range config {
		part, err := writer.CreateFormFile(k, k)
		if err != nil {
			return nil, fmt.Errorf("unable to create form file: %w", err)
		}
		configReader := strings.NewReader(v)
		_, err = io.Copy(part, configReader)
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

	return response.ReadJsonPtr[types.WorkspaceLaunchNeeds](res)
}