package artifacts

import (
	"errors"
	"strings"
)

var ErrInvalidModuleSource = errors.New("invalid module source")

type ModuleSource struct {
	Host       string
	OrgName    string
	ModuleName string
}

func ParseSource(source string) (*ModuleSource, error) {
	tokens := strings.Split(source, "/")
	switch len(tokens) {
	case 2:
		// nullstone registry implied
		return &ModuleSource{
			Host:       "",
			OrgName:    tokens[0],
			ModuleName: tokens[1],
		}, nil
	case 3:
		return &ModuleSource{
			Host:       tokens[0],
			OrgName:    tokens[1],
			ModuleName: tokens[2],
		}, nil
	default:
		// this does not match anything resembling a nullstone registry source
		return nil, ErrInvalidModuleSource
	}
}

func (s ModuleSource) String() string {
	tokens := make([]string, 0)
	if s.Host != "" {
		tokens = append(tokens, s.Host)
	}
	if s.OrgName != "" {
		tokens = append(tokens, s.OrgName)
	}
	tokens = append(tokens, s.ModuleName)
	return strings.Join(tokens, "/")
}
