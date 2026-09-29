package api_test

import (
	"context"
	"errors"
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

func TestWorkspaceMappings_Put(t *testing.T) {
	stackId := int64(42)
	workspaceUid := uuid.MustParse("6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b")
	arcanaWorkspaceUid := uuid.MustParse("11111111-2222-4333-8444-555555555555")
	wantPath := "/orgs/nullstone/stacks/42/workspaces/6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b/mapping"
	input := types.WorkspaceMappingInput{BlockId: 7, EnvId: 9, WorkspaceUid: workspaceUid}
	wantBody := `{"blockId": 7, "envId": 9, "workspaceUid": "6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b"}`

	tests := []struct {
		name    string
		status  int
		body    string
		want    *types.WorkspaceMapping
		wantErr func(t *testing.T, err error)
	}{
		{
			name:   "200 decodes the mapping",
			status: http.StatusOK,
			body: `{
  "arcanaWorkspaceUid": "11111111-2222-4333-8444-555555555555",
  "nullfireWorkspaceUid": "6f1a2b3c-4d5e-4f60-8a9b-0c1d2e3f4a5b",
  "orgName": "nullstone", "stackId": 42, "blockId": 7, "envId": 9,
  "createdAt": "2026-09-29T10:00:00Z"
}`,
			want: &types.WorkspaceMapping{
				ArcanaWorkspaceUid:   arcanaWorkspaceUid,
				NullfireWorkspaceUid: workspaceUid,
				OrgName:              "nullstone",
				StackId:              42,
				BlockId:              7,
				EnvId:                9,
				CreatedAt:            mustTime(t, "2026-09-29T10:00:00Z"),
			},
		},
		{
			name:   "400 returns a bad request error",
			status: http.StatusBadRequest,
			body:   `{"message": "invalid", "details": {"blockId": "must be a positive integer"}}`,
			wantErr: func(t *testing.T, err error) {
				var bre response.BadRequestError
				assert.True(t, errors.As(err, &bre), "expected BadRequestError, got %T: %v", err, err)
			},
		},
		{
			name:   "403 returns an unauthorized error",
			status: http.StatusForbidden,
			body:   `{"type": "problems/workspace-mapping-mismatch", "message": "mismatch"}`,
			wantErr: func(t *testing.T, err error) {
				assert.True(t, isUnauthorizedError(err), "expected UnauthorizedError, got %T: %v", err, err)
			},
		},
		{
			name:   "404 returns a not found error",
			status: http.StatusNotFound,
			body:   `{"message": "Workspace not found"}`,
			wantErr: func(t *testing.T, err error) {
				assert.True(t, response.IsNotFoundError(err), "expected NotFoundError, got %T: %v", err, err)
			},
		},
		{
			name:   "502 returns an api error",
			status: http.StatusBadGateway,
			body:   `{"type": "problems/workspace-verification-failed", "message": "Unable to verify the workspace with Nullstone."}`,
			wantErr: func(t *testing.T, err error) {
				var ae response.ApiError
				require.True(t, errors.As(err, &ae), "expected ApiError, got %T: %v", err, err)
				assert.Equal(t, http.StatusBadGateway, ae.Status)
				assert.Equal(t, "problems/workspace-verification-failed", ae.Type)
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
				_, _ = w.Write([]byte(test.body))
			})
			client := mocks.Client(t, "nullstone", handler)

			got, err := client.WorkspaceMappings().Put(context.Background(), stackId, workspaceUid, input)

			assert.Equal(t, http.MethodPut, gotMethod)
			assert.Equal(t, wantPath, gotPath)
			assert.Equal(t, "application/json", gotContentType)
			assert.JSONEq(t, wantBody, string(gotBody))
			if test.wantErr != nil {
				test.wantErr(t, err)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
