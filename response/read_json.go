package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadJson(res *http.Response, obj interface{}) error {
	if err := Verify(res); err != nil {
		return err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(obj); err != nil {
		return fmt.Errorf("error decoding json body: %w", err)
	}
	return nil
}
