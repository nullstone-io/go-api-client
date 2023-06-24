package auth

type AccessTokenSource interface {
	// GetAccessToken retrieves an API key or JWT key with access to orgName
	GetAccessToken(orgName string) (string, error)
}
