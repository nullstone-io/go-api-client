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

type PipelineInfraConfigurations struct {
	Client *Client
}

func (sc PipelineInfraConfigurations) basePath(stackId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/configuration", sc.Client.Config.OrgName, stackId)
}

func (sc PipelineInfraConfigurations) Create(ctx context.Context, stackId int64, commitInfo types.CommitInfo, config map[string]string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	commitInfoJson, err := json.Marshal(commitInfo)
	if err != nil {
		return fmt.Errorf("unable to marshal commitInfo: %w", err)
	}
	part, err := writer.CreateFormField("commitInfo")
	if err != nil {
		return fmt.Errorf("unable to create form field for commitInfo: %w", err)
	}
	_, err = io.Copy(part, bytes.NewReader(commitInfoJson))
	if err != nil {
		return fmt.Errorf("unable to copy commitInfo into form field: %w", err)
	}

	for k, v := range config {
		part, err = writer.CreateFormFile(k, k)
		if err != nil {
			return fmt.Errorf("unable to create form file: %w", err)
		}
		_, err = io.Copy(part, strings.NewReader(v))
		if err != nil {
			return fmt.Errorf("unable to copy file contents into form file: %w", err)
		}
	}
	writer.Close()

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	res, err := sc.Client.Do(ctx, http.MethodPost, sc.basePath(stackId), nil, headers, body)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
