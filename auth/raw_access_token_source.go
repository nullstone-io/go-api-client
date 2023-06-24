package auth

type RawAccessTokenSource struct {
	AccessToken string
}

func (s RawAccessTokenSource) GetAccessToken(orgName string) (string, error) {
	return s.AccessToken, nil
}
