package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"io"
	"net/http"
)

func (c *Client) ReadFileResponse(res *http.Response, file io.Writer) error {
	if err := response.Verify(res); err != nil {
		return err
	}

	defer res.Body.Close()
	_, err := io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("error reading file body: %w", err)
	}
	return nil
}
