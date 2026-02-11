package types

type Secret struct {
	Id       string `json:"id"`
	Value    string `json:"value"`
	Redacted bool   `json:"redacted"`
}
