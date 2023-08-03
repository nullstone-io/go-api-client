package auth

import (
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRunnerKey_CreateImpersonationToken(t *testing.T) {
	rkey := &RunnerKey{
		OrgName:                      "nullstone",
		ImpersonationContext:         "tester",
		ImpersonationAudience:        []string{"auth-server"},
		ImpersonationExpiresDuration: time.Minute,
	}
	require.NoError(t, rkey.GenerateKeys(), "generate keys")
	rsaPubKey, err := rkey.RsaPublicKey()
	require.NoError(t, err, "rsa public key")
	verifier, err := jwt.NewVerifierRS(jwt.RS256, rsaPubKey)
	require.NoError(t, err, "verifier")

	got, err := rkey.CreateImpersonationToken()
	require.NoError(t, err, "unexpected error")
	if assert.NotNil(t, got, "result token") {
		decoded, err := jwt.ParseAndVerifyString(got.String(), verifier)
		assert.NoError(t, err, "parse/verify token")
		var newClaims jwt.StandardClaims
		assert.NoError(t, json.Unmarshal(decoded.RawClaims(), &newClaims), "decoded token claims")
		assert.Equal(t, newClaims.Subject, "nullstone", "subject is org name")
	}
}
