package api

import (
	"bytes"
	"context"
	"encoding/json"
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

func (ec EnvInfraConfigurations) Create(ctx context.Context, stackId, envId int64, commitInfo types.CommitInfo, config map[string]string) (*types.WorkspaceLaunchNeeds, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	commitInfoJson, err := json.Marshal(commitInfo)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal commitInfo: %w", err)
	}
	part, err := writer.CreateFormField("commitInfo")
	if err != nil {
		return nil, fmt.Errorf("unable to create form field for commitInfo: %w", err)
	}
	_, err = io.Copy(part, bytes.NewReader(commitInfoJson))
	if err != nil {
		return nil, fmt.Errorf("unable to copy commitInfo into form field: %w", err)
	}

	for k, v := range config {
		part, err = writer.CreateFormFile(k, k)
		if err != nil {
			return nil, fmt.Errorf("unable to create form file: %w", err)
		}
		_, err = io.Copy(part, strings.NewReader(v))
		if err != nil {
			return nil, fmt.Errorf("unable to copy file contents into form file: %w", err)
		}
	}
	writer.Close()

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	res, err := ec.Client.Do(ctx, http.MethodPost, ec.basePath(stackId, envId), nil, headers, body)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceLaunchNeeds](res)
}
