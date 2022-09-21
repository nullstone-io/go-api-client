package types

type Outputs map[string]Output

type Output struct {
	Type      any  `json:"type"` // This is typically a string like "string", "number", but could also be ["tuple", "..."]
	Value     any  `json:"value"`
	Sensitive bool `json:"sensitive"`
	Redacted  bool `json:"redacted"`
}
