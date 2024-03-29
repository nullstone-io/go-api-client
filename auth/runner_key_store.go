package auth

import (
	"context"
	"sync"
	"time"
)

type RunnerKeyStore interface {
	GetOrCreate(ctx context.Context, orgName string) (*RunnerKey, error)
}

var _ RunnerKeyStore = &FakeRunnerKeyStore{}

type FakeRunnerKeyStore struct {
	runnerKeys map[string]*RunnerKey
	mu         sync.Mutex
}

func (s *FakeRunnerKeyStore) GetOrCreate(ctx context.Context, orgName string) (*RunnerKey, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.runnerKeys == nil {
		s.runnerKeys = map[string]*RunnerKey{}
	}

	if existing, ok := s.runnerKeys[orgName]; ok {
		return existing, nil
	}

	newRunnerKey := &RunnerKey{
		OrgName:                      orgName,
		ImpersonationContext:         "fake",
		ImpersonationAudience:        []string{"auth-server"},
		ImpersonationExpiresDuration: 24 * time.Hour,
	}
	if err := newRunnerKey.GenerateKeys(); err != nil {
		return nil, err
	}

	s.runnerKeys[orgName] = newRunnerKey
	return newRunnerKey, nil
}
