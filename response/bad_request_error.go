package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type BadRequestError struct {
	ApiError
	Details ErrorDetails `json:"details"`
}

func (e BadRequestError) Error() string {
	buf := bytes.NewBufferString(e.BaseErrorString())
	// sorted so the same response always produces the same error string
	keys := make([]string, 0, len(e.Details))
	for key := range e.Details {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintf(buf, "\n  - %s", e.Details[key])
	}
	return buf.String()
}

func (e BadRequestError) Payload() map[string]any {
	payload := e.ApiError.Payload()
	payload["details"] = e.Details
	return payload
}

func BadRequestErrorFromResponse(res *http.Response) BadRequestError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	bre := BadRequestError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&bre)
	return bre
}
