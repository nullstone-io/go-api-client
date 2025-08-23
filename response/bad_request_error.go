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
	if e.RequestId != "" {
		fmt.Fprintf(buf, "[%s] ", e.RequestId)
	}
	fmt.Fprint(buf, "Bad Request: ")
	for _, value := range e.Details {
		fmt.Fprintf(buf, "\n  %s", value)
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
