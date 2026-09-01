package response

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// DetailsMessageKey is the key a bare-string "details" payload is stored under.
const DetailsMessageKey = "message"

// ErrorDetails is the "details" payload on an api error.
//
// Nullstone services disagree on its shape. Go services (BSick7/go-api) send an object
// of key -> message; furion (Rails) sends a bare string for any DetailedError raised
// with a message, and an object for those raised with a hash. Declaring this as a plain
// map[string]string silently dropped the string form: encoding/json returned an
// UnmarshalTypeError, and the callers that decode api errors discard decode errors, so
// the server's explanation never reached the caller at all.
//
// A bare string is stored under DetailsMessageKey so both shapes read the same way.
type ErrorDetails map[string]string

// UnmarshalJSON accepts either an object or a bare string, and never fails: "details" is
// supplementary, so an unrecognized shape is preserved as raw JSON rather than
// invalidating the error it belongs to.
func (d *ErrorDetails) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}

	// furion: "details": "Block (x) can not be updated, ..."
	var asString string
	if err := json.Unmarshal(trimmed, &asString); err == nil {
		if asString == "" {
			return nil
		}
		*d = ErrorDetails{DetailsMessageKey: asString}
		return nil
	}

	// Go services: "details": {"message": "..."}
	// Decoded as any so a non-string value (a number, a nested object) is kept rather
	// than failing the whole field the way map[string]string did.
	var asObject map[string]any
	if err := json.Unmarshal(trimmed, &asObject); err == nil {
		details := make(ErrorDetails, len(asObject))
		for key, value := range asObject {
			details[key] = detailValueToString(value)
		}
		*d = details
		return nil
	}

	*d = ErrorDetails{DetailsMessageKey: string(trimmed)}
	return nil
}

func detailValueToString(value any) string {
	switch typed := value.(type) {
	case nil:
		return ""
	case string:
		return typed
	default:
		if raw, err := json.Marshal(typed); err == nil {
			return string(raw)
		}
		return fmt.Sprint(typed)
	}
}
