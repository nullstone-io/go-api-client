package api

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error (%d): %s => %s", e.StatusCode, e.Status, e.Body)
}

func IsNotFoundError(err error) bool {
	if he, ok := err.(*HttpError); ok {
		return he.StatusCode == http.StatusNotFound
	}
	return false
}
