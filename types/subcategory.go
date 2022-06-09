package types

type Subcategory struct {
	Name        SubcategoryName `json:"name"`
	Description string          `json:"description"`
}

type SubcategoryName string

const (
	SubcategoryAppContainer  SubcategoryName = "container"
	SubcategoryAppServerless SubcategoryName = "serverless"
	SubcategoryAppStaticSite SubcategoryName = "static-site"
	SubcategoryAppServer     SubcategoryName = "server"

	SubcategoryCapabilityIngress    SubcategoryName = "ingress"
	SubcategoryCapabilityDatastores SubcategoryName = "datastores"
	SubcategoryCapabilitySecrets    SubcategoryName = "secrets"
	SubcategoryCapabilitySidecars   SubcategoryName = "sidecars"
	SubcategoryCapabilityEvents     SubcategoryName = "events"
	SubcategoryCapabilityTelemetry  SubcategoryName = "telemetry"
)

var AllSubcategoryNames = map[CategoryName][]string{
	CategoryApp: {
		string(SubcategoryAppContainer),
		string(SubcategoryAppServerless),
		string(SubcategoryAppStaticSite),
		string(SubcategoryAppServer),
	},
	CategoryCapability: {
		string(SubcategoryCapabilityIngress),
		string(SubcategoryCapabilityDatastores),
		string(SubcategoryCapabilitySecrets),
		string(SubcategoryCapabilitySidecars),
		string(SubcategoryCapabilityEvents),
		string(SubcategoryCapabilityTelemetry),
	},
}
