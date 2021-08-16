package response

import (
	"errors"
	"fmt"
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
