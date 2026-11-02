package api_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/nullstone-io/go-api-client.v0/mocks"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

const workspaceBuildConfigBody = `{
  "workspaceUid": "6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b",
  "version": 3,
  "updatedAt": "2026-11-02T10:00:00Z",
  "updatedBy": "brad",
  "config": {"builder": "nullstone", "packageMode": "dockerfile", "dockerfile": "web/Dockerfile", "dockerContext": "."}
}`

func TestWorkspaceBuildConfigs_Get(t *testing.T) {
	wantPath := "/orgs/nullstone/stacks/42/blocks/7/envs/9/build_config"

	tests := []struct {
		name     string
		status   int
		body     string
		wantNil  bool
		wantErr  func(t *testing.T, err error)
		wantConf *types.BuildConfig
	}{
		{
			name:   "200 decodes the workspace build config",
			status: http.StatusOK,
			body:   workspaceBuildConfigBody,
			wantConf: &types.BuildConfig{
				Builder:       types.BuilderNullstone,
				PackageMode:   types.BuildPackageModeDockerfile,
				Dockerfile:    "web/Dockerfile",
				DockerContext: ".",
			},
		},
		{
			name:    "404 returns nil",
			status:  http.StatusNotFound,
			body:    `{"message": "no build config exists for this workspace"}`,
			wantNil: true,
		},
		{
			name:   "500 returns an api error",
			status: http.StatusInternalServerError,
			body:   `{"message": "boom"}`,
			wantErr: func(t *testing.T, err error) {
				var ae response.ApiError
				require.ErrorAs(t, err, &ae)
				assert.Equal(t, http.StatusInternalServerError, ae.Status)
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

			got, err := client.WorkspaceBuildConfigs().Get(context.Background(), 42, 7, 9)

			assert.Equal(t, http.MethodGet, gotMethod)
			assert.Equal(t, wantPath, gotPath)
			if test.wantErr != nil {
				test.wantErr(t, err)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			if test.wantNil {
				assert.Nil(t, got)
				return
			}
			require.NotNil(t, got)
			assert.Equal(t, uuid.MustParse("6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b"), got.WorkspaceUid)
			assert.Equal(t, int64(3), got.Version)
			assert.Equal(t, "brad", got.UpdatedBy)
			assert.Equal(t, 2026, got.UpdatedAt.Year())
			assert.Equal(t, *test.wantConf, got.Config)
		})
	}
}

func TestWorkspaceBuildConfigs_Update(t *testing.T) {
	wantPath := "/orgs/nullstone/stacks/42/blocks/7/envs/9/build_config"
	input := types.BuildConfig{
		Builder:       types.BuilderNullstone,
		PackageMode:   types.BuildPackageModeDockerfile,
		Dockerfile:    "web/Dockerfile",
		DockerContext: ".",
	}
	wantBody := `{"builder": "nullstone", "packageMode": "dockerfile", "dockerfile": "web/Dockerfile", "dockerContext": "."}`

	var gotMethod, gotPath, gotContentType string
	var gotBody []byte
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath = r.Method, r.URL.Path
		gotContentType = r.Header.Get("Content-Type")
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(workspaceBuildConfigBody))
	})
	client := mocks.Client(t, "nullstone", handler)

	got, err := client.WorkspaceBuildConfigs().Update(context.Background(), 42, 7, 9, input)

	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, gotMethod)
	assert.Equal(t, wantPath, gotPath)
	assert.Equal(t, "application/json", gotContentType)
	assert.JSONEq(t, wantBody, string(gotBody))
	require.NotNil(t, got)
	assert.Equal(t, int64(3), got.Version)
	assert.Equal(t, input, got.Config)
}
