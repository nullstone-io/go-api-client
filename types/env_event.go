package types

import (
	"github.com/google/uuid"
	"maps"
	"slices"
)

type EnvEvent struct {
	Uid     uuid.UUID `json:"uid"`
	OrgName string    `json:"orgName"`
	StackId int64     `json:"stackId"`
	EnvId   int64     `json:"envId"`
	// Name is used to uniquely identify an env event in an IaC file
	// It is unique within an org/stack/env
	Name string `json:"name"`
	// OwningRepoUrl identifies the repo that configured this event
	// If empty, then a user manually configured this event (and not via iac-sync)
	OwningRepoUrl string `json:"owningRepoUrl"`
	// Actions is a list of event actions (e.g. app-deployed)
	// This event will only trigger on these event actions
	Actions []EventAction `json:"actions"`
	// Status is a list of event statuses (e.g. failed, completed)
	// This event will only trigger on these event statuses
	Statuses []EventStatus `json:"statuses"`
	// Blocks is a list of block ids
	// The event will only trigger for these blocks
	Blocks []int64 `json:"blocks"`
	// Channels represents the channel data for each integration tool
	// For example, this is how to configure which Slack channels to send notifications
	Channels map[IntegrationTool]ChannelData `json:"channels"`
}

func (e *EnvEvent) Normalize() {
	if e.Actions == nil {
		e.Actions = []EventAction{}
	}
	slices.Sort(e.Actions)
	if e.Statuses == nil {
		e.Statuses = []EventStatus{}
	}
	slices.Sort(e.Statuses)
	if e.Blocks == nil {
		e.Blocks = []int64{}
	}
	slices.Sort(e.Blocks)
}

func (e *EnvEvent) IsEqual(existing EnvEvent) bool {
	e.Normalize()
	existing.Normalize()

	return slices.Compare(e.Actions, existing.Actions) == 0 &&
		slices.Compare(e.Statuses, existing.Statuses) == 0 &&
		slices.Compare(e.Blocks, existing.Blocks) == 0 &&
		maps.EqualFunc(e.Channels, existing.Channels, IsChannelDataEqual)
}
