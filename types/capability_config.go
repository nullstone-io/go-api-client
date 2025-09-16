package types

type CapabilityConfig struct {
	// Id is a unique identifier for all Capability objects
	Id int64 `json:"id"`
	// TfId is a unique identifier used for creating unique Terraform resources
	// It is unique among all capabilities in the Application Workspace
	TfId string `json:"tfId"`
	// Name is a unique identifier for the Capability
	// This is for all capabilities on a single Nullstone Application
	Name string `json:"name"`
	// Source refers to the module used for this workspace
	Source string `json:"source"`
	// SourceConstraint is a constraint or desired version for the workspace module
	// Once resolved, SourceVersion contains the effective module version
	SourceConstraint string `json:"sourceConstraint"`
	// SourceVersion refers to the effective module version
	// Variables and Connections on this WorkspaceConfig should match the schema for this module version
	SourceVersion  string      `json:"sourceVersion"`
	Variables      Variables   `json:"variables"`
	Connections    Connections `json:"connections"`
	NeedsDestroyed bool        `json:"needsDestroyed"`
	Namespace      string      `json:"namespace"`
}

func (s CapabilityConfig) ToTf() TfCapability {
	return TfCapability{
		Id:               s.TfId,
		TfId:             s.TfId,
		Name:             s.Name,
		Source:           s.Source,
		SourceConstraint: s.SourceConstraint,
		SourceVersion:    s.SourceVersion,
		Variables:        s.Variables,
		Connections:      s.Connections,
		NeedsDestroyed:   s.NeedsDestroyed,
		Namespace:        s.Namespace,
	}
}

func (c CapabilityConfig) Equal(b CapabilityConfig) bool {
	return c.Source == b.Source &&
		c.SourceVersion == b.SourceVersion &&
		c.Namespace == b.Namespace &&
		c.NeedsDestroyed == b.NeedsDestroyed &&
		c.Connections.Equal(b.Connections) &&
		c.Variables.Equal(b.Variables)
}
