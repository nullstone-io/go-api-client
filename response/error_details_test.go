package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorDetails_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want ErrorDetails
	}{
		{
			name: "bare string (furion DetailedError)",
			raw:  `"Block (network0) can not be updated, it is already owned by acme/other"`,
			want: ErrorDetails{DetailsMessageKey: "Block (network0) can not be updated, it is already owned by acme/other"},
		},
		{
			name: "object of messages (BSick7/go-api)",
			raw:  `{"message":"The body of the request does not match the expected format."}`,
			want: ErrorDetails{DetailsMessageKey: "The body of the request does not match the expected format."},
		},
		{
			name: "object with several keys",
			raw:  `{"name":"can't be blank","source":"is invalid"}`,
			want: ErrorDetails{"name": "can't be blank", "source": "is invalid"},
		},
		{
			name: "non-string values are kept, not dropped",
			raw:  `{"count":3,"nested":{"a":"b"},"missing":null}`,
			want: ErrorDetails{"count": "3", "nested": `{"a":"b"}`, "missing": ""},
		},
		{name: "null", raw: `null`, want: nil},
		{name: "empty string", raw: `""`, want: nil},
		{
			name: "unrecognized shape is preserved as raw json",
			raw:  `["a","b"]`,
			want: ErrorDetails{DetailsMessageKey: `["a","b"]`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var details ErrorDetails
			require.NoError(t, json.Unmarshal([]byte(tt.raw), &details), "details must never fail the decode")
			assert.Equal(t, tt.want, details)
		})
	}
}

// TestBadRequestErrorFromResponse_FurionBody is the regression this type exists for:
// furion sends "details" as a string, and it used to vanish into a nil map because the
// UnmarshalTypeError was discarded.
func TestBadRequestErrorFromResponse_FurionBody(t *testing.T) {
	body := map[string]any{
		"request_id": "abc-123",
		"title":      "Bad Request",
		"type":       "problems/bad-request",
		"code":       http.StatusBadRequest,
		"message":    "Your request could not be processed.",
		"details":    "Block (network0) can not be updated, it is already owned by acme/other",
	}
	res := serveError(t, http.StatusBadRequest, body)

	err := Verify(res)
	require.Error(t, err)

	bre, ok := err.(BadRequestError)
	require.True(t, ok, "a 400 maps to BadRequestError")
	assert.Equal(t,
		"Block (network0) can not be updated, it is already owned by acme/other",
		bre.Details[DetailsMessageKey])
	assert.Equal(t,
		"[abc-123][Bad Request]\n  - Block (network0) can not be updated, it is already owned by acme/other",
		bre.Error())
}

// TestBadRequestErrorFromResponse_GoServiceBody covers the object form the Go services
// send, which worked before and must keep working.
func TestBadRequestErrorFromResponse_GoServiceBody(t *testing.T) {
	body := map[string]any{
		"title":   "Bad Request",
		"code":    http.StatusBadRequest,
		"message": "Your request could not be processed.",
		"details": map[string]string{"message": "Invalid stack id"},
	}
	res := serveError(t, http.StatusBadRequest, body)

	bre, ok := Verify(res).(BadRequestError)
	require.True(t, ok)
	assert.Equal(t, ErrorDetails{DetailsMessageKey: "Invalid stack id"}, bre.Details)
	assert.Equal(t, "[Bad Request]\n  - Invalid stack id", bre.Error())
}

func TestBadRequestError_ErrorIsDeterministic(t *testing.T) {
	bre := BadRequestError{Details: ErrorDetails{"z": "last", "a": "first", "m": "middle"}}
	first := bre.Error()
	for i := 0; i < 50; i++ {
		assert.Equal(t, first, bre.Error(), "details render in sorted key order")
	}
	assert.Equal(t, "[]\n  - first\n  - middle\n  - last", first)
}

func serveError(t *testing.T, status int, body map[string]any) *http.Response {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if id, ok := body["request_id"].(string); ok {
			w.Header().Set("X-Request-Id", id)
		}
		w.WriteHeader(status)
		require.NoError(t, json.NewEncoder(w).Encode(body))
	}))
	t.Cleanup(srv.Close)

	res, err := http.Get(srv.URL)
	require.NoError(t, err)
	return res
}
