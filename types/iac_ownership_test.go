package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The messages below are verbatim copies of what nullfire produces
// (internal/blocks/errors.go and iac/update_events.go). nullfire's own tests run the same
// matcher over its real errors; these cases pin the parser against the known shapes.
func TestMatchIacOwnershipConflict(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want *IacOwnershipConflict
	}{
		{
			name: "not an ownership conflict",
			msg:  "An error occurred when validating Nullstone IaC files.\nblock \"api\": module_version is invalid",
			want: nil,
		},
		{
			name: "empty",
			msg:  "",
			want: nil,
		},
		{
			name: "single block, single repo",
			msg: "IaC sync from acme/thief cannot update 1 block owned by another repository:\n" +
				"\tacme/infra owns api\n" +
				"A block can only be defined in one repository for IaC sync, remove it from one of them.",
			want: &IacOwnershipConflict{
				Blocks: map[string]string{"api": "acme/infra"},
				Events: map[string]string{},
			},
		},
		{
			name: "several blocks across several repos",
			msg: "IaC sync from acme/thief cannot update 3 blocks owned by other repositories:\n" +
				"\tacme/infra owns api, web\n" +
				"\tacme/other owns worker\n" +
				"A block can only be defined in one repository for IaC sync, remove it from one of them.",
			want: &IacOwnershipConflict{
				Blocks: map[string]string{"api": "acme/infra", "web": "acme/infra", "worker": "acme/other"},
				Events: map[string]string{},
			},
		},
		{
			name: "blank source repo with the CLI hint appended",
			msg: "IaC sync did not identify its source repository, so it cannot update 1 block owned by another repository:\n" +
				"\tacme/infra owns api\n" +
				"Run nullstone iac sync from a git clone with an origin remote, or pass --repo=<owner/name>.",
			want: &IacOwnershipConflict{
				Blocks: map[string]string{"api": "acme/infra"},
				Events: map[string]string{},
			},
		},
		{
			name: "single event (multierror)",
			msg: "1 error occurred:\n" +
				"\t* event \"deploy-notify\" is configured in another repository (https://github.com/acme/infra)\n\n",
			want: &IacOwnershipConflict{
				Blocks: map[string]string{},
				Events: map[string]string{"deploy-notify": "https://github.com/acme/infra"},
			},
		},
		{
			name: "several events, blank source repo, hint appended",
			msg: "IaC sync did not identify its source repository, so it cannot update events owned by a repository:\n" +
				"2 errors occurred:\n" +
				"\t* event \"deploy-notify\" is configured in another repository (https://github.com/acme/infra)\n" +
				"\t* event \"approvals\" is configured in another repository (https://github.com/acme/other)\n\n" +
				"Run nullstone iac sync from a git clone with an origin remote, or pass --repo=<owner/name>.",
			want: &IacOwnershipConflict{
				Blocks: map[string]string{},
				Events: map[string]string{
					"deploy-notify": "https://github.com/acme/infra",
					"approvals":     "https://github.com/acme/other",
				},
			},
		},
		{
			name: "wrapped by a caller",
			msg: "IaC sync failed: IaC sync from acme/thief cannot update 1 block owned by another repository:\n" +
				"\tacme/infra owns api\n" +
				"A block can only be defined in one repository for IaC sync, remove it from one of them.",
			want: &IacOwnershipConflict{
				Blocks: map[string]string{"api": "acme/infra"},
				Events: map[string]string{},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, MatchIacOwnershipConflict(test.msg))
		})
	}
}

func TestIacOwnershipConflict_Names(t *testing.T) {
	c := IacOwnershipConflict{
		Blocks: map[string]string{"web": "acme/infra", "api": "acme/infra"},
		Events: map[string]string{"z": "https://github.com/acme/infra", "a": "https://github.com/acme/infra"},
	}
	assert.Equal(t, []string{"api", "web"}, c.BlockNames())
	assert.Equal(t, []string{"a", "z"}, c.EventNames())
}
