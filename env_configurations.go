package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"io"
	"io/ioutil"
	"net/http"
)

type EnvConfigurations struct {
	Client *Client
}

func (ec EnvConfigurations) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/runs", ec.Client.Config.OrgName, stackId, envId)
}

func (ec EnvConfigurations) Create(stackId, envId int64, file io.Reader) error {
	raw, err := ioutil.ReadAll(file)
	headers := map[string]string{
		"Content-Type": "application/multipart-form-data",
	}
	res, err := ec.Client.Do(http.MethodPost, ec.basePath(stackId, envId), nil, headers, raw)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
