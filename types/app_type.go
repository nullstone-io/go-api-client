package types

type AppTypeName string

const (
	AppTypeGeneric  AppTypeName = "generic"
	AppTypePackaged AppTypeName = "packaged"
)

var AllAppTypeNames = []string{
	string(AppTypeGeneric),
	string(AppTypePackaged),
}
