package types

// AutomationPurpose constants define security boundaries for performing actions on cloud resources using IAM identities within the cloud accounts.
// Each operation uses these constants to ensure a user or robot is allowed to perform various operations throughout the Nullstone platform.
const (
	AutomationPurposePush        = "push"
	AutomationPurposeDeploy      = "deploy"
	AutomationPurposeRun         = "run"
	AutomationPurposeExecRemote  = "exec-remote"
	AutomationPurposeViewLogs    = "view-logs"
	AutomationPurposeViewMetrics = "view-metrics"
	AutomationPurposeViewStatus  = "view-status"
)
