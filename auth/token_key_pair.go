package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/cristalhq/jwt/v3"
)

type TokenKeyPair struct {
	JwtPrivateKey []byte
	JwtPublicKey  []byte
}

// GenerateKeys will produce a new RSA private/public keypair and save on to this model
func (p *TokenKeyPair) GenerateKeys() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("error generating JWT keypair: %w", err)
	}

	pubKeyPemBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("error marshaling JWT public key: %w", err)
	}

	p.JwtPrivateKey = x509.MarshalPKCS1PrivateKey(privateKey)
	p.JwtPublicKey = pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyPemBytes,
	})
	return nil
}

func (p *TokenKeyPair) RsaPublicKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode(p.JwtPublicKey)
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if rpubkey, ok := pubKey.(*rsa.PublicKey); ok {
		return rpubkey, nil
	}
	return nil, errors.New("invalid public key")
}

func (p *TokenKeyPair) CreateToken(claims any) (*jwt.Token, error) {
	builder, err := p.CreateJwtBuilder()
	if err != nil {
		return nil, fmt.Errorf("error initializing JWT token builder: %w", err)
	}

	token, err := builder.Build(claims)
	if err != nil {
		return nil, fmt.Errorf("error building JWT token: %w", err)
	}
	return token, nil
}

func (p *TokenKeyPair) CreateJwtBuilder() (*jwt.Builder, error) {
	key, err := x509.ParsePKCS1PrivateKey(p.JwtPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing JWT private key: %w", err)
	}

	signer, err := jwt.NewSignerRS(jwt.RS256, key)
	if err != nil {
		return nil, fmt.Errorf("error creating RS256 JWT signer: %w", err)
	}
	return jwt.NewBuilder(signer), nil
}
