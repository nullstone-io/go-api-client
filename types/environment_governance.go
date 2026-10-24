package types

// EnvironmentGovernance is the closed set of change-control policies an environment
// enforces. Like EnvironmentMetadata it is persisted as one jsonb column, so a new policy
// is a new field here rather than a migration. Every policy defaults to off (zero value),
// which is the behavior from before the policy existed.
//
// Policies are set through the environment API and UI by stack architects and owners;
// they are not declared in IaC files.
type EnvironmentGovernance struct {
	// DisableManualIacSync rejects `nullstone iac sync` run by a person (a CLI invocation
	// with no CI automation tool detected) against this environment. Syncs triggered by a
	// repository webhook or by the CLI inside a CI tool are unaffected, so IaC changes can
	// only reach the environment through the sanctioned automated paths.
	//
	// CI detection reads the CI tool's environment variables, so this is a guardrail
	// against habitual bypasses rather than a security boundary.
	DisableManualIacSync bool `json:"disableManualIacSync"`
}
