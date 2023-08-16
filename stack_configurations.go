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

type StackConfigurations struct {
	Client *Client
}

func (sc StackConfigurations) basePath(stackId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/configuration", sc.Client.Config.OrgName, stackId)
}

func (sc StackConfigurations) Create(stackId int64, config, overrides string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if config != "" {
		part, err := writer.CreateFormFile("config", "config.yml")
		if err != nil {
			return fmt.Errorf("unable to create form file: %w", err)
		}
		configReader := strings.NewReader(config)
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
