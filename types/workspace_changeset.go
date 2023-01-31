package types

type WorkspaceChangeset struct {
	Version int64             `json:"version"`
	Changes []WorkspaceChange `json:"changes"`
}
