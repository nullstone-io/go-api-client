package types

import (
	"fmt"
	"strings"
)

func ParseContractName(s string) (ContractName, error) {
	cn := ContractName{}
	tokens := strings.SplitN(s, "/", 3)
	if len(tokens) != 3 {
		return cn, fmt.Errorf("invalid contract format, expected <category>/<provider>/<platform>")
	}
	cn.Category, cn.Provider, cn.Platform = tokens[0], tokens[1], tokens[2]
	if strings.Contains(cn.Category, ":") {
		subtokens := strings.SplitN(cn.Category, ":", 2)
		cn.Category, cn.Subcategory = subtokens[0], subtokens[1]
	}
	if strings.Contains(cn.Platform, ":") {
		subtokens := strings.SplitN(cn.Platform, ":", 2)
		cn.Platform, cn.Subplatform = subtokens[0], subtokens[1]
	}
	return cn, nil
}

type ContractName struct {
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Provider    string `json:"provider"`
	Platform    string `json:"platform"`
	Subplatform string `json:"subplatform"`
}

func (cn ContractName) String() string {
	fullCategory := cn.Category
	if cn.Subcategory != "" {
		fullCategory = fmt.Sprintf("%s:%s", cn.Category, cn.Subcategory)
	}
	fullPlatform := cn.Platform
	if cn.Subplatform != "" {
		fullPlatform = fmt.Sprintf("%s:%s", cn.Platform, cn.Subplatform)
	}
	return fmt.Sprintf("%s/%s/%s", fullCategory, cn.Provider, fullPlatform)
}
