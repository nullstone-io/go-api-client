package types

type Category struct {
	Name        CategoryName `json:"name"`
	Description string       `json:"description"`
}

type CategoryName string

const (
	CategoryAppServer     CategoryName = "app/server"
	CategoryAppContainer  CategoryName = "app/container"
	CategoryAppServerless CategoryName = "app/serverless"
	CategoryAppStaticSite CategoryName = "app/static-site"
	CategorySubdomain     CategoryName = "subdomain"
	CategoryDomain        CategoryName = "domain"
	CategoryCapability    CategoryName = "capability"
	CategoryBlock         CategoryName = "block"
)
