package auth

import (
	"fmt"
	"github.com/cristalhq/jwt/v3"
)

func NewRunner(orgName string, store RunnerKeyStore) (*Runner, error) {
	runnerKey, err2 := store.GetOrCreate(orgName)
	if err2 != nil {
		return nil, fmt.Errorf("error retrieving or creating runner key: %w", err2)
	}
	if runnerKey == nil {
		return nil, fmt.Errorf("could not create new runner for %q", orgName)
	}

	return &Runner{
		RunnerKey:     runnerKey,
		JwtTokenCache: &JwtTokenExpiresCache{},
	}, nil
}

// Runner provides a mechanism for a runner acquiring an access token from the auth server
// Refresh acquires
type Runner struct {
	RunnerKey     *RunnerKey
	JwtTokenCache *JwtTokenExpiresCache
}

func (r *Runner) Refresh(authServer RunnerAccessTokenGetter) (*jwt.Token, error) {
	return r.JwtTokenCache.Refresh(func() (*jwt.Token, error) {
		token, err := r.RunnerKey.CreateImpersonationToken()
		if err != nil {
			return nil, err
		}

		jwtToken, err := authServer.GetRunnerAccessToken(r.RunnerKey.OrgName, r.RunnerKey.Context, token)
		if err != nil {
			return nil, fmt.Errorf("error acquiring runner access token from auth server: %w", err)
		}
		return jwtToken, nil
	})
}
