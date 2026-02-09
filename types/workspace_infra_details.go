package types

type WorkspaceInfraDetails struct {
	WorkspaceDetails `json:",inline"`
	Workspace        Workspace `json:"workspace"`
	// WorkspaceConfig contains the WorkspaceConfig for the last finished run
	WorkspaceConfig WorkspaceConfig `json:"workspaceConfig"`
	// Module contains module information from the last finished run
	Module Module `json:"module"`
}
