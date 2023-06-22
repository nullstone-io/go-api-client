package types

type WorkspaceLaunchNeeds struct {
	NeedsLaunched []Workspace `json:"needsLaunched"`
	Updated       []Workspace `json:"updated"`
}
