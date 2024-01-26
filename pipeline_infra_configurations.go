package api

import (
	"bytes"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type PipelineInfraConfigurations struct {
	Client *Client
}

func (sc PipelineInfraConfigurations) basePath(stackId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/configuration", sc.Client.Config.OrgName, stackId)
}

func (sc PipelineInfraConfigurations) Create(stackId int64, config map[string]string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range config {
		part, err := writer.CreateFormFile(k, k)
		if err != nil {
			return fmt.Errorf("unable to create form file: %w", err)
		}
		configReader := strings.NewReader(v)
		_, err = io.Copy(part, configReader)
		if err != nil {
			return fmt.Errorf("unable to copy file contents into form file: %w", err)
		}
	}
	writer.Close()

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	res, err := sc.Client.Do(http.MethodPost, sc.basePath(stackId), nil, headers, body)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
