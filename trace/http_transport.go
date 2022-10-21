package trace

import (
	"log"
	"net/http"
	"net/http/httputil"
)

var _ http.RoundTripper = &HttpTransport{}

type HttpTransport struct {
	BaseTransport http.RoundTripper
}

func (t *HttpTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	raw, _ := httputil.DumpRequestOut(r, true)
	log.Printf("[DEBUG] %s", string(raw))
	return t.BaseTransport.RoundTrip(r)
}
