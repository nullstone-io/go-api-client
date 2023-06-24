package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"time"
)

// RunnerKey holds a set of JWT keys that are used to build a trust relationship with the auth server
// During the registration of a runner, the auth server acquires the JwtPublicKey
// During requests to gain an access token, the auth server verifies the incoming jwt token against the JwtPublicKey
type RunnerKey struct {
	OrgName string
	Context string

	ImpersonationAudience        []string
	ImpersonationExpiresDuration time.Duration

	JwtPrivateKey []byte
	JwtPublicKey  []byte
}

// CreateImpersonationToken uses this services private JWT keys to generate a valid jwt token
// The resulting impersonation token is used to acquire an access token from the auth server
func (k *RunnerKey) CreateImpersonationToken() (*jwt.Token, error) {
	builder, err := k.createJwtBuilder()
	if err != nil {
		return nil, fmt.Errorf("error initializing impersonation token: %w", err)
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error creating jwt id")
	}

	claims := jwt.StandardClaims{
		ID:        uid.String(),
		Audience:  k.ImpersonationAudience,
		Issuer:    fmt.Sprintf("https://%s", k.Context),
		Subject:   k.OrgName,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(k.ImpersonationExpiresDuration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token, err := builder.Build(claims)
	if err != nil {
		return nil, fmt.Errorf("error building impersonation token: %w", err)
	}
	return token, nil
}

// createJwtBuilder creates a jwt builder using the jwt private key
// This builder can be used to build a jwt token from claims
func (k *RunnerKey) createJwtBuilder() (*jwt.Builder, error) {
	key, err := x509.ParsePKCS1PrivateKey(k.JwtPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing jwt private key for %q: %w", k.OrgName, err)
	}

	signer, err := jwt.NewSignerRS(jwt.RS256, key)
	if err != nil {
		return nil, fmt.Errorf("error creating jwt signer for %q: %w", k.OrgName, err)
	}
	return jwt.NewBuilder(signer), nil
}

// GenerateKeys will produce a new RSA private/public keypair and save on to this model
func (k *RunnerKey) GenerateKeys() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("error generating JWT keypair: %w", err)
	}

	pubKeyPemBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("error marshaling JWT public key: %w", err)
	}

	k.JwtPrivateKey = x509.MarshalPKCS1PrivateKey(privateKey)
	k.JwtPublicKey = pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyPemBytes,
	})
	return nil
}

func (k *RunnerKey) RsaPublicKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode(k.JwtPublicKey)
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if rpubkey, ok := pubKey.(*rsa.PublicKey); ok {
		return rpubkey, nil
	}
	return nil, errors.New("invalid public key")
}
