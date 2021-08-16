package response

import "net/http"

// Verify returns a specific error interface
// If status code < 400, no error is emitted
func Verify(res *http.Response) error {
	ae := ApiErrorFromResponse(res)

	switch res.StatusCode {
	case http.StatusBadRequest: // 400
		return ApiBadRequestError{
			ApiError: ae,
			Details:  parseBadRequestDetails(res),
		}
	case http.StatusNotFound: // 404
		return NotFoundError{ApiError: ae}
	case http.StatusUnprocessableEntity: // 422
		return ApiValidationError{
			ApiError:         ae,
			ValidationErrors: parseApiValidationErrors(res),
		}
	}
	if res.StatusCode >= 400 {
		return ApiErrorFromResponse(res)
	}
	return nil
}
