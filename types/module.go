package types

type ProviderTypes []string

type Module struct {
	UidCreatedModel
	OrgName       string          `json:"orgName"`
	Name          string          `json:"name"`
	FriendlyName  string          `json:"friendlyName"`
	Description   string          `json:"description"`
	IsPublic      bool            `json:"isPublic"`
	Category      CategoryName    `json:"category"`
	Subcategory   SubcategoryName `json:"subcategory"`
	ProviderTypes ProviderTypes   `json:"providerTypes"`
	Platform      string          `json:"platform"`
	Subplatform   string          `json:"subplatform"`
	Type          string          `json:"type"`
	AppCategories []string        `json:"appCategories"`
	SourceUrl     string          `json:"sourceUrl"`
	Status        ModuleStatus    `json:"status"`

	Versions      ModuleVersions `json:"versions,omitempty"`
	LatestVersion *ModuleVersion `json:"latestVersion,omitempty"`
}
