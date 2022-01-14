package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UnauthenticatedError struct {
	ApiError
}

func (e UnauthenticatedError) Error() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "[%s][%s] unauthenticated", e.Url, e.RequestId)
	return buf.String()
}

func UnauthenticatedErrorFromResponse(res *http.Response) UnauthenticatedError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	uae := UnauthenticatedError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&uae)
	return uae
}
