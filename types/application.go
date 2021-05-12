package types

type Application struct {
	Block
	Repo      string `json:"repo"`
	Framework string `json:"framework"`
}
