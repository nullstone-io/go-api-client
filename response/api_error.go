package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	Url       string `json:"url"`
	Status    int    `json:"status"`
	RequestId string `json:"requestId"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[%s][%s] http error (%d): %s", e.Url, e.RequestId, e.Status, e.Message)
}

func ApiErrorFromResponse(res *http.Response) ApiError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	body := ApiError{}
	decoder.Decode(&body)
	// The body doesn't contain url, status code, or request id
	body.Url = res.Request.RequestURI
	body.Status = res.StatusCode
	body.RequestId = res.Header.Get("X-Request-Id")
	return body
}
