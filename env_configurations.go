package api

import (
	"bytes"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"mime/multipart"
	"net/http"
)

type EnvConfigurations struct {
	Client *Client
}

func (ec EnvConfigurations) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/configuration", ec.Client.Config.OrgName, stackId, envId)
}

func (ec EnvConfigurations) Create(stackId, envId int64, file io.Reader) ([]types.Workspace, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("overrides", "previews.yml")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
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
