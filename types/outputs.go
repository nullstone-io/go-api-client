package types

type Outputs map[string]OutputItem

type OutputItem struct {
	Type      interface{} `json:"type"` // This is typically a string like "string", "number", but could also be ["tuple", "..."]
	Value     interface{} `json:"value"`
	Sensitive bool        `json:"sensitive"`
}
