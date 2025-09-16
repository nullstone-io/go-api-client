package types

import (
	"fmt"
	"strings"
)

// TfCapability is a copy of the types.CapabilityConfig; however, the Id attribute is a string and uses the TfId backing attribute
// This allows us to upgrade the capability without forcing every Terraform module to upgrade to use TfId
type TfCapability struct {
	// Id represents the old attribute used for Terraform resource identification
	// Deprecated - This will be dropped when all Terraform modules use TfId
	Id string `json:"id"`
	// TfId is the new attribute used for Terraform resource identification
	TfId             string      `json:"tfId"`
	Name             string      `json:"name"`
	Source           string      `json:"source"`
	SourceConstraint string      `json:"sourceConstraint"`
	SourceVersion    string      `json:"sourceVersion"`
	Variables        Variables   `json:"variables"`
	Connections      Connections `json:"connections"`
	NeedsDestroyed   bool        `json:"needsDestroyed"`
	Namespace        string      `json:"namespace"`
}

func (c TfCapability) TfModuleAddr() string {
	return fmt.Sprintf("module.cap_%s", c.TfId)
}

func (c TfCapability) TfModuleName() string {
	return fmt.Sprintf("cap_%s", c.TfId)
}

func (c TfCapability) EnvPrefix() string {
	if c.Namespace == "" {
		return ""
	}
	return fmt.Sprintf("%s_", strings.ToUpper(c.Namespace))
}
