package response

import "net/http"

// Verify returns a specific error interface
// If status code < 400, no error is emitted
func Verify(res *http.Response) error {
	if res.StatusCode < 400 {
		return nil
	}

	switch res.StatusCode {
	case http.StatusBadRequest: // 400
		return BadRequestErrorFromResponse(res)
	case http.StatusNotFound: // 404
		return NotFoundErrorFromResponse(res)
	case http.StatusUnprocessableEntity: // 422
		return InvalidRequestErrorFromResponse(res)
	default: // normally 500, but anything above 400
		return ApiErrorFromResponse(res)
	}
}
