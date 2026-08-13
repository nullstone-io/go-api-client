package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// UpdateEnvironmentTagsInput is a per-key patch of an environment's tags. It
// deliberately avoids whole-map replacement, which would force every caller into
// a read-modify-write and make concurrent single-key writes clobber each other.
type UpdateEnvironmentTagsInput struct {
	// Tags applies a per-key patch: a key mapped to a value sets/updates it,
	// a key mapped to nil clears it, and any key not present is left untouched.
	// Note that setting a key to the empty string is distinct from clearing it —
	// the key remains present with an empty value.
	Tags map[string]*string `json:"tags"`
}

// ApplyTo merges the patch onto an existing tag map and returns the result. The
// input map is never mutated.
func (i UpdateEnvironmentTagsInput) ApplyTo(existing map[string]string) map[string]string {
	result := make(map[string]string, len(existing)+len(i.Tags))
	for k, v := range existing {
		result[k] = v
	}
	for k, v := range i.Tags {
		if v == nil {
			delete(result, k)
			continue
		}
		result[k] = *v
	}
	return result
}

func (s Environments) envTagsPath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/tags", s.Client.Config.OrgName, stackId, envId)
}

// UpdateTags - PATCH /orgs/:orgName/stacks/:stack_id/envs/:id/tags
// Applies a per-key patch to the environment's tags and returns the updated environment.
// This is a dedicated route rather than a field on Update so that tag writes are
// atomic server-side and can be authorized separately from the rest of the env.
func (s Environments) UpdateTags(ctx context.Context, stackId, envId int64, input UpdateEnvironmentTagsInput) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPatch, s.envTagsPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Environment](res)
}
