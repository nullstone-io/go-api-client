package types

type WorkspaceLaunchStatus struct {
	Workspace    Workspace `json:"workspace"`
	LaunchStatus string    `json:"launchStatus"`
}
