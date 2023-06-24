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
// ExpiresAtThreshold is a duration to configure a token expired before it's expired
// A token is considered expired if `now() + ExpiresAtThreshold > Token.ExpiresAt`
// If ExpiresAtThreshold = 0, will use default of 1s
type JwtTokenExpiresCache struct {
	AcquireFn          func() (string, error)
	Raw                string
	Token              *jwt.Token
	StdClaims          *jwt.StandardClaims
	ExpiresAtThreshold time.Duration
	sync.Mutex
}

func (c *JwtTokenExpiresCache) Refresh() error {
	c.Lock()
	defer c.Unlock()

	// Retrieve access token if we don't have one yet
	if c.StdClaims == nil {
		return c.retrieve()
	}
	// Retrieve access token if the token expired or is close to expiring
	if c.isExpired() {
		return c.retrieve()
	}

	// We already have a valid token
	return nil
}

func (c *JwtTokenExpiresCache) isExpired() bool {
	expiresAt := c.StdClaims.ExpiresAt.Time
	threshold := defaultExpiresAtThreshold
	if c.ExpiresAtThreshold != 0 {
		threshold = c.ExpiresAtThreshold
	}
	return time.Now().Add(threshold).After(expiresAt)
}

func (c *JwtTokenExpiresCache) retrieve() error {
	var err error
	c.Raw, err = c.AcquireFn()
	if err != nil {
		return err
	}
	c.Token, err = jwt.ParseString(c.Raw)
	if err != nil {
		c.Raw = ""
		return fmt.Errorf("error parsing jwt access token: %w", err)
	}

	var stdClaims jwt.StandardClaims
	if err := json.Unmarshal(c.Token.RawClaims(), &stdClaims); err != nil {
		c.Raw = ""
		c.Token = nil
		return fmt.Errorf("error parsing jwt claims: %w", err)
	}
	c.StdClaims = &stdClaims
	return nil
}
