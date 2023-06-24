package auth

import (
	"fmt"
	"net/http"
)

type AccessTokenSource interface {
	// GetAccessToken retrieves an API key or JWT key with access to orgName
	GetAccessToken() (string, error)
}

var _ http.RoundTripper = &AccessTokenSourceTransport{}

type AccessTokenSourceTransport struct {
	BaseTransport     http.RoundTripper
	AccessTokenSource AccessTokenSource
}

func (a *AccessTokenSourceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if a.AccessTokenSource != nil {
		accessToken, err := a.AccessTokenSource.GetAccessToken()
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve access token to authenticate http request: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	return a.BaseTransport.RoundTrip(req)
}
