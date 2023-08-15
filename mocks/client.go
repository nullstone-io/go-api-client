package mocks

import (
	"gopkg.in/nullstone-io/go-api-client.v0"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Client(t *testing.T, orgName string, handler http.Handler) *api.Client {
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	apiClient := &api.Client{
		Config: api.Config{
			BaseAddress:    server.URL,
			OrgName:        orgName,
			IsTraceEnabled: false,
		},
	}
	return apiClient
}
