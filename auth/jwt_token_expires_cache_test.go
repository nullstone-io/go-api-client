package auth

import (
	"github.com/cristalhq/jwt/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJwtTokenExpiresCache(t *testing.T) {
	fakeJwt := `eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIyYWQxMTljMi00ZTkyLTRiNzgtYWIyMi1hNGY2MDE4MDViOTciLCJodHRwczovL251bGxzdG9uZS5pby9yb2xlcyI6eyJzc2lja2xlcyI6ImFkbWluIiwibnVsbHN0b25lIjoiYWRtaW4iLCJsaXN0ZW4zNjAiOiJhZG1pbiJ9LCJpYXQiOjE2MDI2NDI0ODgsImV4cCI6NDc1ODMxNjA4OH0.IrNRVzpavTt1OzeJM2ZL7mCyHeG_rtIaibjLrqe9VEE6CJM5Bul0pIL-GyvyZq1Cv_eJJhwo-Pls24exNiuvV69ccJohSkB5YPYAq5MZGYH5BiRATdZaupasrBAsrYmp-S-9rOkBlae5BzL97DWsq6pHeFxXfSZglDG5PhKJQ-_d85tWABqsf1PAjmp1Q_zjUPyRSLVk4UhPXl6HM_KgLWkBnFU_K2b5tE_ZHz5OdaKcaR-sOEHPQyy1Ziw02NuE2L6ggalChKzWEa8-a7ooKRCTI2qcqxzSGPgyg8SU62lHcMayqESmq4NEOhhBwkJvqOwq9adMW7txTtpwCjrTJPC6XLr_GkHAc9ey62SJlriOcPxwho3I0rsEb5SAFfl8RKoCRNPuYoNtg_2ZXvpo0aoZTijdFgEK5S4AVa8DDunBcypNqYu2Auqd4z1Hx__LswS48eAW3lzuP5uc2FHhR8sA-t-FqhfaJC_GjyCTHIlHN1yru85gDvYtrQk_giP9ACE_ZkaGCZPZnTCcagAU2sxhGA1UX_CJ10xlNKNzqvv4qKCk5WzeH1OVIssX1Zki5ok64NnYm0aAY6FCg28_YkZ9bEbGSwZwy7NFCNLbAoQDsS-QkAC7B_6et-P8olmLr4EXEIttnfduuJxHijrx4tY6cu74n6EHg71ZLXDsKLE`

	tests := []struct {
		name              string
		initialized       bool
		expiresAtAfterNow time.Duration
		wantHit           bool
	}{
		{
			name:        "initial-acquisition",
			initialized: false,
			wantHit:     true,
		},
		{
			name:              "skip-not-expired",
			initialized:       true,
			expiresAtAfterNow: time.Hour,
			wantHit:           false,
		},
		{
			name:              "refresh-expired",
			initialized:       true,
			expiresAtAfterNow: -time.Hour,
			wantHit:           true,
		},
		{
			name:              "refresh-almost-expired",
			initialized:       true,
			expiresAtAfterNow: -30 * time.Second,
			wantHit:           true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := &JwtTokenExpiresCache{}
			if test.initialized {
				cache.StdClaims = &jwt.StandardClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(test.expiresAtAfterNow)),
				}
			}
			gotHit := false
			_, err := cache.Refresh(func() (*jwt.Token, error) {
				gotHit = true
				return jwt.ParseString(fakeJwt)
			})
			require.NoError(t, err, "unexpected error")
			assert.Equal(t, test.wantHit, gotHit, "hit")
		})
	}
}
