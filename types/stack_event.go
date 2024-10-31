package types

import "github.com/google/uuid"

type StackEvent struct {
	Uid     uuid.UUID `json:"uid"`
	OrgName string    `json:"orgName"`
	StackId int64     `json:"stackId"`
	// Actions is a list of event actions (e.g. app-deployed)
	// This event will only trigger on these event actions
	Actions []string `json:"actions"`
	// Status is a list of event statuses (e.g. failed, completed)
	// This event will only trigger on these event statuses
	Statuses []string `json:"statuses"`
	// Envs is a list of environment ids
	// The event will only trigger for these environments
	// Normally, this refers to the env id; however, there are 2 exceptions;
	//   - "-2" => any pipeline env
	//   - "-1" => any preview env
	Envs []int64 `json:"envs"`
	// Channels represents the channel data for each integration tool
	// For example, this is how to configure which Slack channels to send notifications
	Channels map[IntegrationTool]ChannelData `json:"channels"`
}
