package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ApiError
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("[%s] not found", e.Url)
}

func IsNotFoundError(err error) bool {
	nfe := NotFoundError{}
	return errors.As(err, &nfe)
}

func NotFoundErrorFromResponse(res *http.Response) NotFoundError {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	nfe := NotFoundError{ApiError: BaseApiErrorFromResponse(res)}
	decoder.Decode(&nfe)
	return nfe
}
