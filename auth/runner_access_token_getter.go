package auth

import (
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

type RunnerAccessTokenGetter interface {
	// GetRunnerAccessToken attempts to impersonate a runner against the auth server
	// payload must be a JWT signed using a private/public keypair from the correctly registered runner
	//
	//	The requested org must match the private/public keypair registered runner
	//	  (A nullstone hosted runner can request any org -> auth server will scope down the access token to the requested org)
	GetRunnerAccessToken(orgName, context string, impersonationToken *jwt.Token) (*jwt.Token, error)
}

var _ RunnerAccessTokenGetter = &FakeAuthServer{}

func NewFakeAuthServer(t *testing.T) *FakeAuthServer {
	rk := &RunnerKey{}
	require.NoError(t, rk.GenerateKeys())

	return &FakeAuthServer{
		T:             t,
		JwtPrivateKey: rk.JwtPrivateKey,
		JwtPublicKey:  rk.JwtPublicKey,
		RunnerKeys:    map[string]*RunnerKey{},
	}
}

type FakeAuthServer struct {
	T             *testing.T
	JwtPrivateKey []byte
	JwtPublicKey  []byte
	RunnerKeys    map[string]*RunnerKey
}

func (m *FakeAuthServer) AddRunnerKey(rk *RunnerKey) {
	m.RunnerKeys[rk.OrgName] = rk
}

func (m *FakeAuthServer) GetRunnerAccessToken(orgName, context string, impersonationToken *jwt.Token) (*jwt.Token, error) {
	rk, ok := m.RunnerKeys[orgName]
	if !ok {
		require.Failf(m.T, "mock auth server does not have runner key for %s", orgName)
	}
	rsaPubKey, err := rk.RsaPublicKey()
	require.NoError(m.T, err)
	verifier, err := jwt.NewVerifierRS(jwt.RS256, rsaPubKey)
	require.NoError(m.T, err)
	if _, err := jwt.ParseAndVerify(impersonationToken.Raw(), verifier); err != nil {
		return nil, fmt.Errorf("invalid imperson token: %w", err)
	}

	return rk.CreateImpersonationToken()
}
