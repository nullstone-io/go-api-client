package auth

type RawAccessTokenSource struct {
	AccessToken string
}

func (s RawAccessTokenSource) GetAccessToken() (string, error) {
	return s.AccessToken, nil
}
