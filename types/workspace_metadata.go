package types

// WorkspaceMetadata is an extensible container for governance/descriptive
// metadata about a workspace. DataClassification is its first member; future
// metadata (cost center, compliance scope, …) lands here too.
//
// It is carried on both WorkspaceTemplateConfig (the canonical, block-level,
// environment-invariant declaration) and WorkspaceConfig (the per-workspace
// materialized value seeded from the template).
type WorkspaceMetadata struct {
	// DataClassification is the sensitivity level of the data the workspace
	// holds (see ClassificationLevel). Empty = unclassified.
	DataClassification ClassificationLevel `json:"dataClassification,omitempty"`
}
