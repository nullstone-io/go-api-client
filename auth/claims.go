package auth

import "github.com/cristalhq/jwt/v3"

type Claims struct {
	jwt.StandardClaims
	Email    string            `json:"email"`
	Picture  string            `json:"picture"`
	Username string            `json:"https://nullstone.io/username"`
	Roles    map[string]string `json:"https://nullstone.io/roles"`
}
