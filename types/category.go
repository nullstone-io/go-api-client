package types

type Category struct {
	Name        CategoryName `json:"name"`
	Description string       `json:"description"`
}

type CategoryName string

const (
	CategoryApp        CategoryName = "app"
	CategoryCapability              = "capability"
	CategoryDatastore               = "datastore"
	CategorySubdomain               = "subdomain"
	CategoryDomain                  = "domain"
	CategoryCluster                 = "cluster"
	CategoryNetwork                 = "network"
	CategoryBlock                   = "block"
)

var AllCategoryNames = []string{
	"app",
	"capability",
	"datastore",
	"subdomain",
	"domain",
	"cluster",
	"network",
	"block",
}
