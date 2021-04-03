package types

type Domain struct {
	IdModel
	Name         string
	OrgName      string
	StackName    string
	ModuleSource string
	Registrar    string
	Certificate  string
}
