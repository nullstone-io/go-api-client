package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiValidationError struct {
	ApiError
	ValidationErrors map[string][]string `json:"validation_errors"`
}

func (e ApiValidationError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s][%s] invalid request:", e.Url, e.RequestId)
	for field, errs := range e.ValidationErrors {
		for _, err := range errs {
			fmt.Fprintf(buf, "\n  %s: %s", field, err)
		}
	}
	return buf.String()
}

func parseApiValidationErrors(res *http.Response) map[string][]string {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	ave := ApiValidationError{}
	decoder.Decode(&ave)
	return ave.ValidationErrors
}
