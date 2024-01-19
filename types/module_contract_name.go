package types

import (
	"fmt"
	"strings"
)

func ParseModuleContractName(s string) (ModuleContractName, error) {
	cn := ModuleContractName{
		Category:    "",
		Subcategory: "",
		Provider:    "",
		Platform:    "",
		Subplatform: "",
	}
	if s == "" {
		return cn, nil
	}
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

type ModuleContractName struct {
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Provider    string `json:"provider"`
	Platform    string `json:"platform"`
	Subplatform string `json:"subplatform"`
}

func (cn ModuleContractName) String() string {
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

// Match performs a match of one module contract to another
// Typically, a connection or module (`other`) is validated to follow the contract (`cn`)
func (cn ModuleContractName) Match(other ModuleContractName) bool {
	// if the provider is a set of providers, we will check to see if ANY of the providers match
	// if so, we will set the other.provider to cn.Provider so that the matchContractPart below will always match
	// this is essentially moving the provider check here
	if cn.Provider != "*" {
		pts := strings.Split(other.Provider, ",")
		if len(pts) > 1 {
			for _, pt := range pts {
				if pt == cn.Provider {
					other.Provider = cn.Provider
					break
				}
			}
		}
	}

	return matchContractPart(cn.Category, other.Category, false) &&
		matchContractPart(cn.Subcategory, other.Subcategory, true) &&
		matchContractPart(cn.Provider, other.Provider, false) &&
		matchContractPart(cn.Platform, other.Platform, false) &&
		matchContractPart(cn.Subplatform, other.Subplatform, true)
}

func matchContractPart(want, got string, optional bool) bool {
	if want == "*" {
		return true
	}
	if want == "" && optional {
		return true
	}
	return want == got
}

// CompareModuleContractName provides a comparison of two module contract names for sorting by specificity, then alphabetically
// This will walk through category, provider, and platform to ensure that "more-specific" contracts appear first
// A contract with all wildcards is least-specific; a contract with no wildcards is most-specific
// A backup comparison is made on alphabetic order
func CompareModuleContractName(a, b ModuleContractName) bool {
	category := compareContractPart(a.Category, b.Category, false)
	if category != 0 {
		return category < 0
	}
	subcategory := compareContractPart(a.Subcategory, b.Subcategory, true)
	if subcategory != 0 {
		return subcategory < 0
	}
	provider := compareContractPart(a.Provider, b.Provider, false)
	if provider != 0 {
		return provider < 0
	}
	platform := compareContractPart(a.Platform, b.Platform, false)
	if platform != 0 {
		return platform < 0
	}
	subplatform := compareContractPart(a.Subplatform, b.Subplatform, true)
	if subplatform != 0 {
		return subplatform < 0
	}
	return false
}

func compareContractPart(a, b string, optional bool) int {
	if optional {
		if a == "" {
			a = "*"
		}
		if b == "" {
			b = "*"
		}
	}

	if a == b {
		// both * or same value
		return 0
	}
	if a == "*" && b != "*" {
		// a is less-specific
		return 1
	}
	if a != "*" && b == "*" {
		// b is less specific
		return -1
	}
	if a < b {
		// a is before in alphabet
		return -1
	}
	// b is before in alphabet
	return 1
}
