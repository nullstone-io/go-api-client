package types

type Message struct {
	Type    string `json:"type"`
	Context string `json:"context"`
	Content string `json:"content"`
}
