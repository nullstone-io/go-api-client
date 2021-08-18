package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BadRequestError struct {
	ApiError
	Details map[string]string `json:"details"`
}

func (e BadRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s][%s] bad request:", e.Url, e.RequestId)
	for key, value := range e.Details {
		fmt.Fprintf(buf, "\n  %s: %s", key, value)
	}
	return buf.String()
}

func BadRequestErrorFromResponse(res *http.Response) BadRequestError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	bre := BadRequestError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&bre)
	return bre
}
