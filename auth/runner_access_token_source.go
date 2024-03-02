package auth

import (
	"context"
	"fmt"
	"sync"
)

// RunnerAccessTokenSource coordinates a trust relationship between nullstone auth server and a runner
// The retrieved access token is cached for each org
// The access token is refreshed if the key has expired or is about to expire
type RunnerAccessTokenSource struct {
	AuthServer     RunnerAccessTokenGetter
	RunnerKeyStore RunnerKeyStore

	cache map[string]*Runner
	mu    sync.Mutex
}

func (s *RunnerAccessTokenSource) ensure() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cache == nil {
		s.cache = map[string]*Runner{}
	}
}

func (s *RunnerAccessTokenSource) GetAccessToken(ctx context.Context, orgName string) (string, error) {
	s.ensure()
	runner, err := s.getOrInitialize(ctx, orgName)
	if err != nil {
		return "", err
	}
	token, err := runner.Refresh(s.AuthServer)
	if err != nil {
		return "", fmt.Errorf("error refreshing runner access token: %w", err)
	}
	return token.String(), nil
}

func (s *RunnerAccessTokenSource) getOrInitialize(ctx context.Context, orgName string) (*Runner, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if runner, exists := s.cache[orgName]; exists {
		return runner, nil
	}

	runner, err := NewRunner(ctx, orgName, s.RunnerKeyStore)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize runner: %w", err)
	}
	s.cache[orgName] = runner
	return runner, nil
}
