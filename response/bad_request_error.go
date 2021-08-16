package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiBadRequestError struct {
	ApiError
	Details map[string]string `json:"details"`
}

func (e ApiBadRequestError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s][%s] bad request:", e.Url, e.RequestId)
	for key, value := range e.Details {
		fmt.Fprintf(buf, "\n  %s: %s", key, value)
	}
	return buf.String()
}

func parseBadRequestDetails(res *http.Response) map[string]string {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	bre := ApiBadRequestError{}
	decoder.Decode(&bre)
	return bre.Details
}
