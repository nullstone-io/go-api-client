package types

import "time"

type CreateWorkspaceWorkflowPayload struct {
	Actions       []string  `json:"actions"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	Status        string    `json:"status"`
	StatusMessage string    `json:"statusMessage"`
	StatusAt      time.Time `json:"statusAt"`
}
