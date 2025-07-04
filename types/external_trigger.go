package types

const (
	ExternalTriggerSourceManual = "manual"
	ExternalTriggerSourceGithub = "github"
)

const (
	ExternalTriggerEventUser                  = "user"
	ExternalTriggerEventAutomation            = "automation"
	ExternalTriggerEventVcsPush               = "vcs-push"
	ExternalTriggerEventVcsPullRequestOpened  = "vcs-pull-request-opened"
	ExternalTriggerEventVcsPullRequestLabeled = "vcs-pull-request-labeled"
	ExternalTriggerEventGitopsEnable          = "gitops-enable"
	ExternalTriggerEventGitopsChangeBranch    = "gitops-change-branch"
	ExternalTriggerEventEnvCreate             = "env-create"
)

// ExternalTrigger represents the source of an external trigger performing an action against Nullstone
type ExternalTrigger struct {
	// Source is the origination of the ExternalTrigger
	// Examples: ExternalTriggerSourceManual, ExternalTriggerSourceGithub
	Source string `json:"source"`
	// Event indicates what type of event caused the ExternalTrigger
	// Examples: ExternalTriggerEventUser, ExternalTriggerEventVcsPush, ExternalTriggerEventVcsPullRequestOpened, ExternalTriggerEventVcsPullRequestLabeled
	Event string `json:"event"`
	// SourceUser indicates the user that caused the ExternalTrigger
	// This user is dependent on the Source
	// For example, if ExternalTriggerSourceGithub, this is the Github user
	SourceUser ExternalTriggerUser `json:"sourceUser"`
	// NullstoneUser is the user in Nullstone that is linked to SourceUser
	NullstoneUser *User `json:"nullstoneUser"`
	// Labels is a list of labels that were added during a "labeled" event
	Labels []string `json:"labels"`
}

type ExternalTriggerUser struct {
	// Id is the platform-dependent identification
	// Many platforms use int64, but string is universal to all platforms
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
}
