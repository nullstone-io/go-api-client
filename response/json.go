package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json[T any](res *http.Response) (*T, error) {
	var t T
	if err := Verify(res); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&t); err != nil {
		if IsNotFoundError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("error decoding json body: %w", err)
	}
	return &t, nil
}

func JsonArray[T any](res *http.Response) ([]T, error) {
	t := make([]T, 0)
	if err := Verify(res); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&t); err != nil {
		if IsNotFoundError(err) {
			return t, nil
		}
		return nil, fmt.Errorf("error decoding json body: %w", err)
	}
	return t, nil
}
