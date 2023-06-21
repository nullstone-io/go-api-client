package types

type WorkspaceStatuses struct {
	NeedsLaunched []*Workspace `json:"needs_launched"`
	Updated       []*Workspace `json:"updated"`
}
