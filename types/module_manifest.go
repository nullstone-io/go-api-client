package types

type ModuleManifest struct {
	OrgName       string   `yaml:"org_name" json:"orgName"`
	Name          string   `yaml:"name" json:"name"`
	FriendlyName  string   `yaml:"friendly_name" json:"friendlyName"`
	Description   string   `yaml:"description" json:"description"`
	Category      string   `yaml:"category" json:"category"`
	Subcategory   string   `yaml:"subcategory" json:"subcategory"`
	ProviderTypes []string `yaml:"provider_types" json:"providerTypes"`
	Platform      string   `yaml:"platform" json:"platform"`
	Subplatform   string   `yaml:"subplatform" json:"subplatform"`
	Type          string   `yaml:"type" json:"type"`
	AppCategories []string `yaml:"appCategories" json:"appCategories"`
	IsPublic      bool     `yaml:"is_public" json:"isPublic"`
	// ToolName is the name of the tool that this module is for
	// Exmaples: terraform, opentofu
	ToolName string `yaml:"tool_name" json:"toolName"`
	// Includes is a list of filename globs that will add files to the module when packaging
	Includes []string `yaml:"include" json:"includes"`
	// SourceUrl is the repository URL containing the module source code
	SourceUrl string `yaml:"source_url" json:"sourceUrl"`
}
