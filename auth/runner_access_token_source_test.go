package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunnerAccessTokenSource_GetAccessToken(t *testing.T) {
	// This test verifies GetAccessToken with the following steps
	// - GetAccessToken against "acme1" org (should hit auth server)
	// - GetAccessToken against "acme1" org (should *not* hit auth server)
	// - GetAccessToken against "acme2" org (should hit auth server)
	// This verifies the cache-by-org mechanism is working properly
	// Detailed verification of runner key impersonation and access token retrieval is left to their interfaces

	fakeServer := NewFakeAuthServer(t)
	fakeStore := &FakeRunnerKeyStore{}

	tokenSource := &RunnerAccessTokenSource{
		AuthServer:     fakeServer,
		RunnerKeyStore: fakeStore,
	}

	org1 := "acme1"
	org2 := "acme2"

	rk1, err := fakeStore.GetOrCreate(org1)
	require.NoError(t, err)
	rk2, err := fakeStore.GetOrCreate(org2)
	require.NoError(t, err)

	fakeServer.AddRunnerKey(rk1)
	fakeServer.AddRunnerKey(rk2)

	got1, err := tokenSource.GetAccessToken(org1)
	require.NoError(t, err, "unexpected error")

	got2, err := tokenSource.GetAccessToken(org1)
	require.NoError(t, err, "unexpected error")
	assert.Equal(t, got1, got2, "should be same token")

	got3, err := tokenSource.GetAccessToken(org2)
	require.NoError(t, err, "unexpected error")
	assert.NotEqual(t, got1, got3)
	assert.NotEqual(t, got2, got3)
}
