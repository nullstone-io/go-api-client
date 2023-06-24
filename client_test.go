package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_WithApiKey(t *testing.T) {
	original := Client{Config: Config{}}
	updated := original.WithApiKey("abc123")
	assert.NotSame(t, original, updated)
}
