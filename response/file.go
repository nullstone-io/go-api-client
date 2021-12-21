package response

import (
	"fmt"
	"io"
	"net/http"
)

func File(res *http.Response, file io.Writer) error {
	if err := Verify(res); err != nil {
		return err
	}

	defer res.Body.Close()
	_, err := io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("error reading file body: %w", err)
	}
	return nil
}
