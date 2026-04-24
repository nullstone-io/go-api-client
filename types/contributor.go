package types

type Contributor string

const (
	ContributorNullstoneOfficial Contributor = "nullstone-official"
	ContributorMyOrg             Contributor = "my-org"
	ContributorCommunity         Contributor = "community"
)

var AllContributors = []Contributor{
	ContributorNullstoneOfficial,
	ContributorMyOrg,
	ContributorCommunity,
}
