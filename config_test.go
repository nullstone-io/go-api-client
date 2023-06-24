package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/nullstone-io/go-api-client.v0/auth"
	"testing"
)

func TestConfig_UseApiKey(t *testing.T) {
	tests := []struct {
		name    string
		initial auth.AccessTokenSource
		apiKey  string
		want    auth.AccessTokenSource
	}{
		{
			name:    "set-empty",
			initial: auth.RawAccessTokenSource{AccessToken: "initial"},
			apiKey:  "",
			want:    nil,
		},
		{
			name:    "keep-empty",
			initial: nil,
			apiKey:  "",
			want:    nil,
		},
		{
			name:    "set-to-real-value",
			initial: nil,
			apiKey:  "new",
			want:    auth.RawAccessTokenSource{AccessToken: "new"},
		},
		{
			name:    "switch-real-value",
			initial: auth.RawAccessTokenSource{AccessToken: "initial"},
			apiKey:  "new",
			want:    auth.RawAccessTokenSource{AccessToken: "new"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := Config{AccessTokenSource: test.initial}
			cfg.UseApiKey(test.apiKey)
			got := cfg.AccessTokenSource
			assert.Equal(t, test.want, got)
		})
	}
}

func TestConfig_ConstructWsEndpoint(t *testing.T) {
	t.Run("no-api-key", func(t *testing.T) {
		cfg := Config{BaseAddress: "ws://localhost:12345"}
		endpoint, headers, err := cfg.ConstructWsEndpoint("/stream")
		require.NoError(t, err, "unexpected error")
		assert.Equal(t, "ws://localhost:12345/stream", endpoint)
		assert.Empty(t, headers.Get("Authorization"))
	})

	t.Run("has-api-key", func(t *testing.T) {
		cfg := Config{BaseAddress: "ws://localhost:12345"}
		cfg.UseApiKey("abc123")
		endpoint, headers, err := cfg.ConstructWsEndpoint("/stream")
		require.NoError(t, err, "unexpected error")
		assert.Equal(t, "ws://localhost:12345/stream", endpoint)
		assert.Equal(t, fmt.Sprint("Bearer abc123"), headers.Get("Authorization"))
	})
}
