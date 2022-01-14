package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UnauthorizedError struct {
	ApiError
}

func (e UnauthorizedError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s][%s] unauthorized", e.Url, e.RequestId)
	return buf.String()
}

func UnauthorizedErrorFromResponse(res *http.Response) UnauthorizedError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	uae := UnauthorizedError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&uae)
	return uae
}
