package types

type ProviderTypes []string

type Module struct {
	UidCreatedModel
	OrgName       string        `json:"orgName"`
	Name          string        `json:"name"`
	FriendlyName  string        `json:"friendlyName"`
	Description   string        `json:"description"`
	IsPublic      bool          `json:"isPublic"`
	Category      CategoryName  `json:"category"`
	Layer         Layer         `json:"layer"`
	Type          string        `json:"type"`
	ProviderTypes ProviderTypes `json:"providerTypes"`
	Status        ModuleStatus  `json:"status"`

	Versions ModuleVersions `json:"versions,omitempty"`
}
