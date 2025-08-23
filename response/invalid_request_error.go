package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InvalidRequestError struct {
	ApiError
	ValidationErrors map[string][]string `json:"validation_errors"`
}

func (e InvalidRequestError) Error() string {
	buf := bytes.NewBufferString(e.BaseErrorString())
	for field, errs := range e.ValidationErrors {
		for _, err := range errs {
			fmt.Fprintf(buf, "\n  - %s: %s", field, err)
		}
	}
	return buf.String()
}

func InvalidRequestErrorFromResponse(res *http.Response) InvalidRequestError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	ire := InvalidRequestError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&ire)
	return ire
}
