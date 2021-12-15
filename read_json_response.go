package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"net/http"
)

func ReadJsonResponse[T any](res *http.Response) (*T, error) {
	var obj T
	if err := response.Verify(res); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&obj); err != nil {
		if response.IsNotFoundError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("error decoding json body: %w", err)
	}
	return &obj, nil
}
