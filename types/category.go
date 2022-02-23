package types

type Category struct {
	Name        CategoryName `json:"name"`
	Description string       `json:"description"`
}

type CategoryName string

const (
	CategoryAppContainer          CategoryName = "app/container"
	CategoryAppServerless                      = "app/serverless"
	CategoryAppStaticSite                      = "app/static-site"
	CategoryAppServer                          = "app/server"
	CategoryCapabilityPublicEntry              = "capability/public-entry"
	CategoryCapabilityDatastores               = "capability/datastores"
	CategoryCapabilitySecrets                  = "capability/secrets"
	CategoryCapabilitySidecars                 = "capability/sidecars"
	CategoryCapabilityEvents                   = "capability/events"
	CategoryCapabilityTelemetry                = "capability/telemetry"
	CategoryDatastore                          = "datastore"
	CategorySubdomain                          = "subdomain"
	CategoryDomain                             = "domain"
	CategoryBlock                              = "block"
)

var AllCategoryNames = []string{
	"app/container",
	"app/serverless",
	"app/static-site",
	"app/server",
	"capability/public-entry",
	"capability/datastores",
	"capability/secrets",
	"capability/sidecars",
	"capability/events",
	"capability/telemetry",
	"datastore",
	"subdomain",
	"domain",
	"block",
}

var AllAppCategoryNames = []string{
	"app/container",
	"app/serverless",
	"app/static-site",
	"app/server",
}
