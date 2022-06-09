package types

type Category struct {
	Name        CategoryName `json:"name"`
	Description string       `json:"description"`
}

type CategoryName string

const (
	CategoryApp        CategoryName = "app"
	CategoryCapability CategoryName = "capability"
	CategoryDatastore  CategoryName = "datastore"
	CategorySubdomain  CategoryName = "subdomain"
	CategoryDomain     CategoryName = "domain"
	CategoryCluster    CategoryName = "cluster"
	CategoryNetwork    CategoryName = "network"
	CategoryBlock      CategoryName = "block"
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
