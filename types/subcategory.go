package types

type Subcategory struct {
	Name        SubcategoryName `json:"name"`
	Description string          `json:"description"`
}

type SubcategoryName string

const (
	SubcategoryAppContainer  SubcategoryName = "container"
	SubcategoryAppServerless                 = "serverless"
	SubcategoryAppStaticSite                 = "static-site"
	SubcategoryAppServer                     = "server"

	SubcategoryCapabilityIngress    = "ingress"
	SubcategoryCapabilityDatastores = "datastores"
	SubcategoryCapabilitySecrets    = "secrets"
	SubcategoryCapabilitySidecars   = "sidecars"
	SubcategoryCapabilityEvents     = "events"
	SubcategoryCapabilityTelemetry  = "telemetry"
)

var AllSubcategoryNames = map[CategoryName][]string{
	CategoryApp: {
		string(SubcategoryAppContainer),
		SubcategoryAppServerless,
		SubcategoryAppStaticSite,
		SubcategoryAppServer,
	},
	CategoryCapability: {
		SubcategoryCapabilityIngress,
		SubcategoryCapabilityDatastores,
		SubcategoryCapabilitySecrets,
		SubcategoryCapabilitySidecars,
		SubcategoryCapabilityEvents,
		SubcategoryCapabilityTelemetry,
	},
}
