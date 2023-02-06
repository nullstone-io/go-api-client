package api

import (
	"bytes"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"io"
	"mime/multipart"
	"net/http"
)

type EnvConfigurations struct {
	Client *Client
}

func (ec EnvConfigurations) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/runs", ec.Client.Config.OrgName, stackId, envId)
}

func (ec EnvConfigurations) Create(stackId, envId int64, file io.Reader) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("configuration", "previews.yaml")
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	res, err := ec.Client.Do(http.MethodPost, ec.basePath(stackId, envId), nil, headers, body)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
