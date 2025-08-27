package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UnauthorizedError struct {
	ApiError
}

func (e UnauthorizedError) Error() string {
	return fmt.Sprintf("%s unauthorized", e.BaseErrorString())
}

func UnauthorizedErrorFromResponse(res *http.Response) UnauthorizedError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	uae := UnauthorizedError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&uae)
	return uae
}
