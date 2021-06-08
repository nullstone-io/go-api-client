package types

type ProviderTypes []string

type Module struct {
	UidCreatedModel
	OrgName       string        `json:"orgName"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	IsPublic      bool          `json:"isPublic"`
	Layer         Layer         `json:"layer"`
	Category      CategoryName  `json:"category"`
	Type          string        `json:"type"`
	ProviderTypes ProviderTypes `json:"providerTypes"`

	Versions []ModuleVersion `json:"versions,omitempty"`
}
