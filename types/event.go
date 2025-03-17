package types

import (
	"github.com/google/go-cmp/cmp"
)

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

	// EventActionEnvLaunched triggers when an env is launched
	EventActionEnvLaunched EventAction = "env-launched"
	// EventActionEnvDestroyed triggers when an env is destroyed
	EventActionEnvDestroyed EventAction = "env-destroyed"
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
	return cmp.Diff(a, b) != ""
}
