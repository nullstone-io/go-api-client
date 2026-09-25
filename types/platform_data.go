package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nullstone-io/module/platformdata"
)

// Sources of a platform data value (PlatformData.Sources).
const (
	// PlatformDataSourceApply - the value came from Terraform state (written by an infra apply)
	PlatformDataSourceApply = "apply"
	// PlatformDataSourceDeploy - the value came from the deploy overlay (written by a deploy)
	PlatformDataSourceDeploy = "deploy"
)

// PlatformDataStateVersion identifies the state version the platform data was extracted from.
type PlatformDataStateVersion struct {
	Uid       uuid.UUID `json:"uid"`
	Serial    int64     `json:"serial"`
	CreatedAt time.Time `json:"created_at"`
}

// PlatformDataOverlay describes the deploy overlay merged into Data (nil when no deploy has written one).
type PlatformDataOverlay struct {
	DeployId  int64     `json:"deploy_id"`
	Actor     string    `json:"actor"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PlatformData is the response of GET /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/platform-data/:kind
// Data is the merged payload for (Kind, Version); use Env() to decode the `env` kind.
// Sources maps each key in Data to the source that produced it (PlatformDataSourceApply / PlatformDataSourceDeploy).
type PlatformData struct {
	Kind         string                   `json:"kind"`
	Version      int                      `json:"version"`
	StateVersion PlatformDataStateVersion `json:"state_version"`
	Overlay      *PlatformDataOverlay     `json:"overlay"`
	Validated    bool                     `json:"validated"`
	Data         json.RawMessage          `json:"data"`
	Sources      map[string]string        `json:"sources"`
}

// Env decodes Data as the `env` kind (platformdata.KindEnv).
// Returns an error if Kind is not `env` or if the payload does not satisfy the env v1 schema.
func (p PlatformData) Env() (platformdata.EnvV1, error) {
	if p.Kind != platformdata.KindEnv {
		return platformdata.EnvV1{}, fmt.Errorf("platform data kind is %q, not %q", p.Kind, platformdata.KindEnv)
	}
	return platformdata.ParseEnvV1(p.Data)
}

// PlatformDataEnvOverlay is the `data` payload of a deploy overlay for the `env` kind.
// Variables are merged by key into the applied record's `variables`.
type PlatformDataEnvOverlay struct {
	Variables      map[string]string `json:"variables"`
	OtelAttributes map[string]string `json:"otel_attributes,omitempty"`
}
