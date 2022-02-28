package types

type LiveLogMessage struct {
	Source  string `json:"source"`
	Context string `json:"context"`
	Content string `json:"content"`
}
