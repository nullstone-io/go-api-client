package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UnauthenticatedError struct {
	ApiError
}

func (e UnauthenticatedError) Error() string {
	return fmt.Sprintf("%s unauthenticated", e.BaseErrorString())
}

func UnauthenticatedErrorFromResponse(res *http.Response) UnauthenticatedError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	uae := UnauthenticatedError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&uae)
	return uae
}
