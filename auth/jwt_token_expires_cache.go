package auth

import (
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"sync"
	"time"
)

var (
	defaultExpiresAtThreshold = 1 * time.Second
)

// JwtTokenExpiresCache keeps a cache of a jwt token
// Every time Refresh() is called, this cache will only acquire a new access token if:
// - access token has not been retrieved yet
// - the access token has expired or close to expiring
// ExpiresAtThreshold is a duration to configure a token refresh before it's expired
// A token is considered expired if `time.Now() + ExpiresAtThreshold > Token.ExpiresAt`
// If ExpiresAtThreshold = 0, will use default of 1s
type JwtTokenExpiresCache struct {
	Token              *jwt.Token
	StdClaims          *jwt.StandardClaims
	ExpiresAtThreshold time.Duration
	sync.Mutex
}

type AcquireFunc func() (*jwt.Token, error)

func (c *JwtTokenExpiresCache) Refresh(acquireFn AcquireFunc) (*jwt.Token, error) {
	c.Lock()
	defer c.Unlock()

	// Retrieve access token if we don't have one yet
	if c.StdClaims == nil {
		if err := c.retrieve(acquireFn); err != nil {
			return nil, err
		}
		return c.Token, nil
	}
	// Retrieve access token if the token expired or is close to expiring
	if c.isExpired() {
		if err := c.retrieve(acquireFn); err != nil {
			return nil, err
		}
		return c.Token, nil
	}

	// We already have a valid token
	return c.Token, nil
}

func (c *JwtTokenExpiresCache) isExpired() bool {
	expiresAt := time.Now()
	if c.StdClaims.ExpiresAt != nil {
		expiresAt = c.StdClaims.ExpiresAt.Time
	}
	threshold := defaultExpiresAtThreshold
	if c.ExpiresAtThreshold != 0 {
		threshold = c.ExpiresAtThreshold
	}
	return time.Now().Add(threshold).After(expiresAt)
}

func (c *JwtTokenExpiresCache) retrieve(acquireFn AcquireFunc) error {
	var err error
	c.Token, err = acquireFn()
	if err != nil {
		return err
	}

	var stdClaims jwt.StandardClaims
	if err := json.Unmarshal(c.Token.RawClaims(), &stdClaims); err != nil {
		c.Token = nil
		return fmt.Errorf("error parsing jwt claims: %w", err)
	}
	c.StdClaims = &stdClaims
	return nil
}
