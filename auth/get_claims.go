package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cristalhq/jwt/v3"
)

var (
	NotValidJwtTokenErr = errors.New("access token is not a valid JWT Token")
)

// GetClaims attempts to read the access token and parse jwt.StandardClaims
// If the access token is not a valid JWT, this will return NotValidJwtTokenErr
func GetClaims(ctx context.Context, source AccessTokenSource, orgName string) (*jwt.StandardClaims, error) {
	accessToken, err := source.GetAccessToken(ctx, orgName)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseString(accessToken)
	if err != nil {
		return nil, NotValidJwtTokenErr
	}

	var claims jwt.StandardClaims
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		return nil, NotValidJwtTokenErr
	}
	return &claims, nil
}

// GetCustomClaims attempts to read the access token and parse a custom Claims json
// If the access token is not a valid JWT, this will return NotValidJwtTokenErr
func GetCustomClaims(ctx context.Context, source AccessTokenSource, orgName string) (*Claims, error) {
	accessToken, err := source.GetAccessToken(ctx, orgName)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseString(accessToken)
	if err != nil {
		return nil, NotValidJwtTokenErr
	}

	var claims Claims
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		return nil, NotValidJwtTokenErr
	}
	return &claims, nil
}
