package auth

import "context"

type AccessTokenSource interface {
	// GetAccessToken retrieves an API key or JWT key with access to orgName
	GetAccessToken(ctx context.Context, orgName string) (string, error)
}
