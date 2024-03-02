package auth

import "context"

type RawAccessTokenSource struct {
	AccessToken string
}

func (s RawAccessTokenSource) GetAccessToken(ctx context.Context, orgName string) (string, error) {
	return s.AccessToken, nil
}
