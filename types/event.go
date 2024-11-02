package types

import "reflect"

type EventAction string

const (
	// EventActionAppDeployed triggers when an app completes a deployment
	EventActionAppDeployed EventAction = "app-deployed"
	// EventActionAppFirstDeployed triggers when an app completes a deployment for the first time
	EventActionAppFirstDeployed EventAction = "app-first-deployed"
	// EventActionBlockLaunched triggers when block infra is initially launched
	EventActionBlockLaunched EventAction = "block-launched"
	// EventActionBlockUpdated triggers when block infra is updated
	EventActionBlockUpdated EventAction = "block-updated"
	// EventActionBlockDestroyed triggers when block infra is destroyed
	EventActionBlockDestroyed EventAction = "block-destroyed"
	// EventActionBlockNeedsApproval triggers when a block requires an approval to proceed
	EventActionBlockNeedsApproval EventAction = "block-needs-approval"
)

type EventStatus string

const (
	EventStatusFailed      EventStatus = "failed"
	EventStatusCompleted   EventStatus = "completed"
	EventStatusCancelled   EventStatus = "cancelled"
	EventStatusDisapproved EventStatus = "disapproved"
)

type EventTarget string

const (
	EventTargetSlack          EventTarget = "slack"
	EventTargetMicrosoftTeams EventTarget = "microsoft-teams"
	EventTargetDiscord        EventTarget = "discord"
	EventTargetWhatsapp       EventTarget = "whatsapp"
	EventTargetWebhook        EventTarget = "webhook"
	EventTargetTask           EventTarget = "task"
)

type ChannelData map[string]any

func IsChannelDataEqual(a, b ChannelData) bool {
	return reflect.DeepEqual(a, b)
}
