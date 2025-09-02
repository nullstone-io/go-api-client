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
	ErrorCode int    `json:"error_code"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%s: %s", e.BaseErrorString(), e.Message)
}

func (e ApiError) BaseErrorString() string {
	if e.RequestId != "" {
		return fmt.Sprintf("[%s][%s]", e.RequestId, http.StatusText(e.Status))
	}
	return fmt.Sprintf("[%s]", http.StatusText(e.Status))
}

func (e ApiError) StatusCode() int {
	return e.Status
}

func (e ApiError) Payload() map[string]any {
	return map[string]any{
		"title":      e.Title,
		"type":       e.Type,
		"code":       e.Status,
		"message":    e.Message,
		"error_code": e.ErrorCode,
	}
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
