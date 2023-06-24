package auth

import (
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	fakeJwt = `eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIyYWQxMTljMi00ZTkyLTRiNzgtYWIyMi1hNGY2MDE4MDViOTciLCJodHRwczovL251bGxzdG9uZS5pby9yb2xlcyI6eyJzc2lja2xlcyI6ImFkbWluIiwibnVsbHN0b25lIjoiYWRtaW4iLCJsaXN0ZW4zNjAiOiJhZG1pbiJ9LCJpYXQiOjE2MDI2NDI0ODgsImV4cCI6NDc1ODMxNjA4OH0.IrNRVzpavTt1OzeJM2ZL7mCyHeG_rtIaibjLrqe9VEE6CJM5Bul0pIL-GyvyZq1Cv_eJJhwo-Pls24exNiuvV69ccJohSkB5YPYAq5MZGYH5BiRATdZaupasrBAsrYmp-S-9rOkBlae5BzL97DWsq6pHeFxXfSZglDG5PhKJQ-_d85tWABqsf1PAjmp1Q_zjUPyRSLVk4UhPXl6HM_KgLWkBnFU_K2b5tE_ZHz5OdaKcaR-sOEHPQyy1Ziw02NuE2L6ggalChKzWEa8-a7ooKRCTI2qcqxzSGPgyg8SU62lHcMayqESmq4NEOhhBwkJvqOwq9adMW7txTtpwCjrTJPC6XLr_GkHAc9ey62SJlriOcPxwho3I0rsEb5SAFfl8RKoCRNPuYoNtg_2ZXvpo0aoZTijdFgEK5S4AVa8DDunBcypNqYu2Auqd4z1Hx__LswS48eAW3lzuP5uc2FHhR8sA-t-FqhfaJC_GjyCTHIlHN1yru85gDvYtrQk_giP9ACE_ZkaGCZPZnTCcagAU2sxhGA1UX_CJ10xlNKNzqvv4qKCk5WzeH1OVIssX1Zki5ok64NnYm0aAY6FCg28_YkZ9bEbGSwZwy7NFCNLbAoQDsS-QkAC7B_6et-P8olmLr4EXEIttnfduuJxHijrx4tY6cu74n6EHg71ZLXDsKLE`
)

func TestRunnerAccessTokenSource_ImpersonateRunner(t *testing.T) {
	orgName := "nullstone"
	rkey := &RunnerKey{
		OrgName:                      orgName,
		Context:                      "tester",
		ImpersonationAudience:        []string{"auth-server"},
		ImpersonationExpiresDuration: time.Minute,
	}
	require.NoError(t, rkey.GenerateKeys(), "unexpected error generating jwt key")

	rsaPubKey, err := rkey.RsaPublicKey()
	require.NoError(t, err, "rsa public key")
	verifier, err := jwt.NewVerifierRS(jwt.RS256, rsaPubKey)
	require.NoError(t, err, "verifier")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Manually return environment
		if r.Method == "POST" && r.URL.Path == fmt.Sprintf("/orgs/nullstone/runner_access_tokens") {
			decoder := json.NewDecoder(r.Body)
			m := map[string]string{}
			require.NoError(t, decoder.Decode(&m), "unexpected request body")

			rawToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			_, err := jwt.ParseAndVerifyString(rawToken, verifier)
			assert.NoError(t, err, "invalid input jwt token")

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fakeJwt))
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	rats := RunnerAccessTokenSource{
		OrgName:                orgName,
		AuthServerAddr:         server.URL,
		GetOrCreateRunnerKeyFn: func(orgName string) (*RunnerKey, error) { return rkey, nil },
	}

	gotToken, err := rats.ImpersonateRunner()
	require.NoError(t, err, "unexpected error")
	assert.Equal(t, fakeJwt, gotToken, "result token matches")
}
