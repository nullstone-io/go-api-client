package response

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Url        string `json:"url"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	RequestId  string `json:"requestId"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[%s][%s] http error (%d): %s", e.Url, e.RequestId, e.Status, e.StatusText)
}

func ApiErrorFromResponse(res *http.Response) ApiError {
	requestId := res.Header.Get("X-Request-Id")
	return ApiError{
		Url:        res.Request.RequestURI,
		Status:     res.StatusCode,
		StatusText: res.Status,
		RequestId:  requestId,
	}
}
