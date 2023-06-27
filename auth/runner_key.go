package auth

import (
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"time"
)

// RunnerKey holds a set of JWT keys that are used to build a trust relationship with the auth server
// During the registration of a runner, the auth server acquires the JwtPublicKey
// During requests to gain an access token, the auth server verifies the incoming jwt token against the JwtPublicKey
type RunnerKey struct {
	OrgName                      string
	ImpersonationContext         string
	ImpersonationAudience        []string
	ImpersonationExpiresDuration time.Duration

	TokenKeyPair
}

// CreateImpersonationToken uses this services private JWT keys to generate a valid jwt token
// The resulting impersonation token is used to acquire an access token from the auth server
func (k *RunnerKey) CreateImpersonationToken() (*jwt.Token, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error creating JWT id")
	}

	claims := jwt.StandardClaims{
		ID:        uid.String(),
		Audience:  k.ImpersonationAudience,
		Issuer:    fmt.Sprintf("https://%s", k.ImpersonationContext),
		Subject:   k.OrgName,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(k.ImpersonationExpiresDuration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	return k.TokenKeyPair.CreateToken(claims)
}
