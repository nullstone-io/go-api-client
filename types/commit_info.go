package types

import (
	"fmt"
	"time"
)

const (
	CommitInfoVcsProviderGithub    = "github"
	CommitInfoVcsProviderGitlab    = "gitlab"
	CommitInfoVcsProviderBitbucket = "bitbucket"
)

const (
	CommitInfoTypeBranch = "branch"
	CommitInfoTypePr     = "pr"
	CommitInfoTypePush   = "push"
)

type CommitInfo struct {
	Type string `json:"type"`

	// VcsProvider
	// Deprecated - Use Repo.Provider
	VcsProvider string `json:"vcsProvider"`

	// Repository contains information about the commit repository (e.g. host, owner, name)
	Repository Repo `json:"repository"`
	// RepoOwner
	// Deprecated - Use Repo.Owner
	RepoOwner string `json:"repoOwner"`
	// RepoName
	// Deprecated - Use Repo.Name
	RepoName string `json:"repoName"`
	// Repo is `RepoOwner/RepoName`
	// Deprecated - Use Repo.Owner/Repo.Name
	Repo string `json:"repo"`
	// RepoUrl
	// Deprecated - Use Repo.Url
	RepoUrl string `json:"repoUrl"`

	BranchName string `json:"branchName"`

	CommitSha string `json:"commitSha"`
	// CommitUrl is the HTML URL to browse this commit
	CommitUrl     string `json:"commitUrl"`
	CommitMessage string `json:"commitMessage"`
	// CommitUserId is the user id for the VCS user that created the commit
	// This is not guaranteed to be the same as the AuthorId
	// When using the GitHub UI to merge, the CommitUsername is actually `web-flow`
	CommitUserId string `json:"commitUserId"`
	// CommitUsername is the username for the VCS user that created the commit
	// This is not guaranteed to be the same as the AuthorUsername
	// When using the GitHub UI to merge, the CommitUsername is actually `web-flow`
	CommitUsername string `json:"commitUsername"`
	// CommitUserEmail is the email address for the VCS user that created the commit
	// This is not guaranteed to be the same as the AuthorEmail
	// When using the GitHub UI to merge, the CommitUsername is actually `web-flow`
	CommitUserEmail string `json:"commitUserEmail"`
	// CommitUserAvatarUrl is the avatar url for the VCS user that created the commit
	// This is not guaranteed to be the same as the AuthorAvatarUrl
	// When using the GitHub UI to merge, the CommitUsername is actually `web-flow`
	CommitUserAvatarUrl string `json:"commitUserAvatarUrl"`
	// CommitDate is the date the commit was created
	CommitDate time.Time `json:"commitDate"`

	// AuthorId is the user id for the VCS user that authored the commit
	// This refers to the user that originally made the code changes
	AuthorId string `json:"authorId"`
	// AuthorUsername is the username for the VCS user that authored the commit
	// This refers to the user that originally made the code changes
	AuthorUsername string `json:"authorUsername"`
	// AuthorEmail is the email address for the VCS user that authored the commit
	// This refers to the user that originally made the code changes
	AuthorEmail string `json:"authorEmail"`
	// AuthorAvatarUrl is the avatar url for the VCS user that authored the commit
	// This refers to the user that originally made the code changes
	AuthorAvatarUrl string `json:"authorAvatarUrl"`

	// VcsUsername
	// Deprecated
	VcsUsername string `json:"vcsUsername"`
	PRNumber    int    `json:"prNumber"`
	PRId        int64  `json:"prId"`
	PRTitle     string `json:"prTitle"`
}

// CommitUser is the VCS user that created the commit
// This is not guaranteed to be the same as the Author
// When using the GitHub UI to merge, the CommitUsername is actually `web-flow`
func (i *CommitInfo) CommitUser() ExternalTriggerUser {
	return ExternalTriggerUser{
		Id:        i.CommitUserId,
		Name:      i.CommitUsername,
		Email:     i.CommitUserEmail,
		AvatarUrl: i.CommitUserAvatarUrl,
	}
}

// Author is the VCS user that authored the commit
// This refers to the user that originally made the code changes
func (i *CommitInfo) Author() ExternalTriggerUser {
	return ExternalTriggerUser{
		Id:        i.AuthorId,
		Name:      i.AuthorUsername,
		Email:     i.AuthorEmail,
		AvatarUrl: i.AuthorAvatarUrl,
	}
}

func (i *CommitInfo) InferCommitUrl() {
	switch i.Repository.Provider {
	case CommitInfoVcsProviderGithub:
		i.CommitUrl = fmt.Sprintf("%s/commit/%s", i.Repository.Url, i.CommitSha)
	case CommitInfoVcsProviderGitlab:
		i.CommitUrl = fmt.Sprintf("%s/-/commit/%s", i.Repository.Url, i.CommitSha)
	case CommitInfoVcsProviderBitbucket:
		i.CommitUrl = fmt.Sprintf("%s/commits/%s", i.Repository.Url, i.CommitSha)
	}
}
