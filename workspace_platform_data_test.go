package api_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nullstone-io/module/platformdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/mocks"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func TestWorkspacePlatformData_Get(t *testing.T) {
	stackId := int64(42)
	workspaceUid := uuid.MustParse("6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b")
	stateVersionUid := uuid.MustParse("11111111-2222-4333-8444-555555555555")
	wantPath := "/orgs/nullstone/stacks/42/workspaces/6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b/platform-data/env"

	okBody := `{
  "kind": "env", "version": 1,
  "state_version": {"uid": "11111111-2222-4333-8444-555555555555", "serial": 12, "created_at": "2026-09-25T10:00:00Z"},
  "overlay": {"deploy_id": 123, "actor": "brad", "updated_at": "2026-09-25T11:00:00Z"},
  "validated": true,
  "data": {
    "variables": {
      "DATABASE_HOST": {"template": "{{ NULLSTONE_ENV }}-db", "value": "dev-db"},
      "NULLSTONE_VERSION": {"value": "1.2.3"},
      "DATABASE_PASSWORD": {"sensitive": true, "ref": {"type": "secret", "id": "arn:aws:secretsmanager:us-east-1:123:secret:db"}},
      "POD_IP": {"ref": {"type": "k8s_field", "field_path": "status.podIP"}}
    }
  },
  "sources": {"DATABASE_HOST": "apply", "NULLSTONE_VERSION": "deploy"}
}`

	tests := []struct {
		name       string
		status     int
		body       string
		want       *types.PlatformData
		wantEnv    *platformdataEnv
		wantErr    func(t *testing.T, err error)
		wantNilErr bool
	}{
		{
			name:   "200 decodes the record and Env() parses the env payload",
			status: http.StatusOK,
			body:   okBody,
			want: &types.PlatformData{
				Kind:    "env",
				Version: 1,
				StateVersion: types.PlatformDataStateVersion{
					Uid:       stateVersionUid,
					Serial:    12,
					CreatedAt: mustTime(t, "2026-09-25T10:00:00Z"),
				},
				Overlay: &types.PlatformDataOverlay{
					DeployId:  123,
					Actor:     "brad",
					UpdatedAt: mustTime(t, "2026-09-25T11:00:00Z"),
				},
				Validated: true,
				Sources:   map[string]string{"DATABASE_HOST": "apply", "NULLSTONE_VERSION": "deploy"},
			},
			wantEnv: &platformdataEnv{
				variables:  map[string]string{"DATABASE_HOST": "dev-db", "NULLSTONE_VERSION": "1.2.3"},
				templates:  map[string]string{"DATABASE_HOST": "{{ NULLSTONE_ENV }}-db"},
				sensitive:  []string{"DATABASE_PASSWORD"},
				refs: map[string]platformdata.EnvV1Ref{
					"DATABASE_PASSWORD": {Type: platformdata.RefTypeSecret, Id: "arn:aws:secretsmanager:us-east-1:123:secret:db"},
					"POD_IP":            {Type: platformdata.RefTypeK8sField, FieldPath: "status.podIP"},
				},
			},
			wantNilErr: true,
		},
		{
			name:       "404 returns nil, nil",
			status:     http.StatusNotFound,
			body:       `{"error": "not found"}`,
			want:       nil,
			wantNilErr: true,
		},
		{
			name:   "403 returns an unauthorized error",
			status: http.StatusForbidden,
			body:   `{"error": "forbidden"}`,
			wantErr: func(t *testing.T, err error) {
				assert.True(t, isUnauthorizedError(err), "expected UnauthorizedError, got %T: %v", err, err)
			},
		},
		{
			name:   "500 returns an api error",
			status: http.StatusInternalServerError,
			body:   `{"error": "boom"}`,
			wantErr: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.False(t, response.IsNotFoundError(err))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var gotMethod, gotPath string
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath = r.Method, r.URL.Path
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			client := mocks.Client(t, "nullstone", handler)

			got, err := client.WorkspacePlatformData().Get(context.Background(), stackId, workspaceUid, "env")

			assert.Equal(t, http.MethodGet, gotMethod)
			assert.Equal(t, wantPath, gotPath)
			if test.wantErr != nil {
				test.wantErr(t, err)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			if test.want == nil {
				assert.Nil(t, got)
				return
			}
			require.NotNil(t, got)
			// Data is compared through Env() below; clear it for the struct comparison
			gotData := got.Data
			got.Data = nil
			assert.Equal(t, test.want, got)
			got.Data = gotData

			env, err := got.Env()
			require.NoError(t, err)
			for key, value := range test.wantEnv.variables {
				assert.Equal(t, value, env.Variables[key].Value, "variables[%s].value", key)
			}
			for key, template := range test.wantEnv.templates {
				assert.Equal(t, template, env.Variables[key].Template, "variables[%s].template", key)
			}
			for _, key := range test.wantEnv.sensitive {
				assert.True(t, env.Variables[key].IsSensitive(), "variables[%s] sensitive", key)
				assert.Empty(t, env.Variables[key].Value, "variables[%s] must not carry a value", key)
			}
			for key, ref := range test.wantEnv.refs {
				require.NotNil(t, env.Variables[key].Ref, "variables[%s].ref", key)
				assert.Equal(t, ref, *env.Variables[key].Ref, "variables[%s].ref", key)
			}
		})
	}
}

func TestPlatformData_Env_RejectsOtherKinds(t *testing.T) {
	record := types.PlatformData{Kind: "metrics", Version: 1, Data: json.RawMessage(`{}`)}
	_, err := record.Env()
	assert.ErrorContains(t, err, `platform data kind is "metrics", not "env"`)
}

func TestWorkspacePlatformData_PutOverlay(t *testing.T) {
	stackId := int64(42)
	workspaceUid := uuid.MustParse("6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b")
	wantPath := "/orgs/nullstone/stacks/42/workspaces/6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b/platform-data/env/overlay"
	input := api.PlatformDataOverlayInput{
		DeployId: 123,
		Actor:    "brad",
		Data: types.PlatformDataEnvOverlay{
			Variables:      map[string]string{"NULLSTONE_VERSION": "1.2.3", "NULLSTONE_COMMIT_SHA": "abc123"},
			OtelAttributes: map[string]string{"service.version": "1.2.3", "service.commit.sha": "abc123"},
		},
	}
	wantBody := `{
  "deploy_id": 123,
  "actor": "brad",
  "data": {
    "variables": {"NULLSTONE_VERSION": "1.2.3", "NULLSTONE_COMMIT_SHA": "abc123"},
    "otel_attributes": {"service.version": "1.2.3", "service.commit.sha": "abc123"}
  }
}`

	tests := []struct {
		name    string
		status  int
		body    string
		wantErr func(t *testing.T, err error)
	}{
		{
			name:   "204 succeeds",
			status: http.StatusNoContent,
		},
		{
			name:   "200 succeeds",
			status: http.StatusOK,
			body:   `{"ok": true}`,
		},
		{
			name:   "403 returns an unauthorized error",
			status: http.StatusForbidden,
			body:   `{"error": "runner required"}`,
			wantErr: func(t *testing.T, err error) {
				assert.True(t, isUnauthorizedError(err), "expected UnauthorizedError, got %T: %v", err, err)
			},
		},
		{
			name:   "404 returns a not found error",
			status: http.StatusNotFound,
			body:   `{"error": "no state version"}`,
			wantErr: func(t *testing.T, err error) {
				assert.True(t, response.IsNotFoundError(err), "expected NotFoundError, got %T: %v", err, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var gotMethod, gotPath, gotContentType string
			var gotBody []byte
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath = r.Method, r.URL.Path
				gotContentType = r.Header.Get("Content-Type")
				gotBody, _ = io.ReadAll(r.Body)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.status)
				if test.body != "" {
					_, _ = w.Write([]byte(test.body))
				}
			})
			client := mocks.Client(t, "nullstone", handler)

			err := client.WorkspacePlatformData().PutOverlay(context.Background(), stackId, workspaceUid, "env", input)

			assert.Equal(t, http.MethodPut, gotMethod)
			assert.Equal(t, wantPath, gotPath)
			assert.Equal(t, "application/json", gotContentType)
			assert.JSONEq(t, wantBody, string(gotBody))
			if test.wantErr != nil {
				test.wantErr(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

// platformdataEnv is the subset of platformdata.EnvV1 asserted by the Get test.
type platformdataEnv struct {
	variables  map[string]string
	templates  map[string]string
	sensitive  []string
	refs       map[string]platformdata.EnvV1Ref
}

func isUnauthorizedError(err error) bool {
	var ue response.UnauthorizedError
	return errors.As(err, &ue)
}

func mustTime(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, value)
	require.NoError(t, err)
	return parsed
}
