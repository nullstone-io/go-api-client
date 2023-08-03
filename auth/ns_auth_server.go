package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"io"
	"net/http"
	"net/url"
	"path"
)

type NsAuthServer struct {
	Address string
}

func (s *NsAuthServer) GetRunnerAccessToken(orgName, context string, impersonationToken *jwt.Token) (*jwt.Token, error) {
	u, err := url.Parse(s.Address)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "orgs", orgName, "runner_access_tokens")

	rawPayload, _ := json.Marshal(map[string]any{"context": context})
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", impersonationToken.String()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error attempting to create runner access token with status code %d", res.StatusCode)
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return jwt.Parse(raw)
}
