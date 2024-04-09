package types

type Category struct {
	Name        CategoryName `json:"name"`
	Description string       `json:"description"`
}

type CategoryName string

const (
	CategoryApp              CategoryName = "app"
	CategoryCapability       CategoryName = "capability"
	CategoryDatastore        CategoryName = "datastore"
	CategoryIngress          CategoryName = "ingress"
	CategorySubdomain        CategoryName = "subdomain"
	CategoryDomain           CategoryName = "domain"
	CategoryCluster          CategoryName = "cluster"
	CategoryClusterNamespace CategoryName = "cluster-namespace"
	CategoryNetwork          CategoryName = "network"
	CategoryBlock            CategoryName = "block"
)

var AllCategoryNames = []string{
	"app",
	"capability",
	"datastore",
	"ingress",
	"subdomain",
	"domain",
	"cluster-namespace",
	"cluster",
	"network",
	"block",
}
