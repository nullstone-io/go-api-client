package main

import (
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/auth"
	"log"
	"time"
)

func main() {
	username := "brad.sickles"
	email := "brad.sickles@acme.com"
	roles := map[string]string{}
	audience := []string{"go-api-client"}
	issuer := "nullstone"
	expirationDuration := 24 * time.Hour

	tkp := &auth.TokenKeyPair{}
	must(tkp.GenerateKeys())

	orgName := "acme"
	now := time.Now()
	token, err := tkp.CreateToken(auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ID:        uuid.New().String(),
			Audience:  audience,
			Issuer:    issuer,
			Subject:   orgName,
			ExpiresAt: jwt.NewNumericDate(now.Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
		Username: username,
		Email:    email,
		Roles:    roles,
		Picture:  "",
	})
	must(err)

	fmt.Println(token.String())
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
