package types

type Category struct {
	Name        CategoryName `json:"name"`
	Description string       `json:"description"`
}

type CategoryName string

const (
	CategoryAppContainer          CategoryName = "app/container"
	CategoryAppServerless         CategoryName = "app/serverless"
	CategoryAppStaticSite         CategoryName = "app/static-site"
	CategoryAppServer             CategoryName = "app/server"
	CategoryCapabilityPublicEntry CategoryName = "capability/public-entry"
	CategoryCapabilityDatastores  CategoryName = "capability/datastores"
	CategoryCapabilitySecrets     CategoryName = "capability/secrets"
	CategoryCapabilitySidecars    CategoryName = "capability/sidecars"
	CategoryCapabilityEvents      CategoryName = "capability/events"
	CategoryCapabilityTelemetry   CategoryName = "capability/telemetry"
	CategoryDatastore             CategoryName = "datastore"
	CategorySubdomain             CategoryName = "subdomain"
	CategoryDomain                CategoryName = "domain"
	CategoryBlock                 CategoryName = "block"
)
