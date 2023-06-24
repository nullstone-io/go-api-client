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
	"sync"
)

type CreateRunnerKeyFunc func(orgName string) (*RunnerKey, error)

// RunnerAccessTokenSource coordinates a trust relationship between nullstone auth server and a runner (a robot user for an org)
// The retrieval is cached for repeated use and refreshed if the key is about to expire
type RunnerAccessTokenSource struct {
	OrgName                string
	AuthServerAddr         string
	GetOrCreateRunnerKeyFn CreateRunnerKeyFunc

	cache *JwtTokenExpiresCache
	mu    sync.Mutex
}

func (s *RunnerAccessTokenSource) GetAccessToken() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cache == nil {
		s.cache = &JwtTokenExpiresCache{}
	}
	err := s.cache.Refresh()
	if err != nil {
		return "", err
	}
	return s.cache.Raw, nil

}

// ImpersonateRunner allows us to impersonate a runner on behalf of an organization
// A local JWT key (per organization) defines a pre-established trust relationship between runner and auth servers
// If a local JWT key is not yet initialized for an org, this function will initialize in the local runner's cache and emit to the auth server
func (s *RunnerAccessTokenSource) ImpersonateRunner() (string, error) {
	runnerKey, err := s.GetOrCreateRunnerKeyFn(s.OrgName)
	if err != nil {
		return "", fmt.Errorf("error retrieving or creating runner key: %w", err)
	}
	if runnerKey == nil {
		return "", fmt.Errorf("could not create new runner for %q", s.OrgName)
	}
	token, err := runnerKey.CreateImpersonationToken()
	if err != nil {
		return "", nil
	}

	accessToken, err := s.getRunnerAccessToken(runnerKey, token)
	if err != nil {
		return "", fmt.Errorf("error acquiring runner access token from auth server: %w", err)
	}
	return accessToken, nil
}

// getRunnerAccessToken attempts to impersonate a runner against the auth server
// payload must be a JWT signed using a private/public keypair from the correctly registered runner
//
//	The payload should include which org to impersonate the runner
//	The requested org must match the private/public keypair registered runner
//	  (A nullstone hosted runner can request any org -> void will scope down the access token to the requested org)
func (s *RunnerAccessTokenSource) getRunnerAccessToken(runnerKey *RunnerKey, token *jwt.Token) (string, error) {
	u, err := url.Parse(s.AuthServerAddr)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, "orgs", runnerKey.OrgName, "runner_access_tokens")

	rawPayload, _ := json.Marshal(map[string]any{"context": runnerKey.Context})
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(rawPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.String()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("error attempting to create runner access token with status code %d", res.StatusCode)
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}
