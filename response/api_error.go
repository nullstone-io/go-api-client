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

func BaseApiErrorFromResponse(res *http.Response) ApiError {
	return ApiError{
		Url:       res.Request.RequestURI,
		Status:    res.StatusCode,
		RequestId: res.Header.Get("X-Request-Id"),
	}
}

func ApiErrorFromResponse(res *http.Response) ApiError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	ae := BaseApiErrorFromResponse(res)
	decoder.Decode(&ae)
	return ae
}
