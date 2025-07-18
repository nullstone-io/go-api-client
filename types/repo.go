package types

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	RepoProviderGithub = "github"
)

const (
	DefaultRepoHost = "github.com"
)

type Repo struct {
	Provider string `json:"provider"`
	Host     string `json:"host"`
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

func RepoFromUrl(repoUrl string) (Repo, error) {
	repoUri, err := url.Parse(repoUrl)
	if err != nil {
		return Repo{}, fmt.Errorf("invalid repository url %q: %s", repoUrl, err)
	}

	repoTokens := strings.SplitN(strings.TrimPrefix(repoUri.Path, "/"), "/", 3)
	if host := repoUri.Host; host != "" {
		// <scheme>://<repo-host>/<repo-owner>/<repo-name>
		if len(repoTokens) != 2 {
			return Repo{}, fmt.Errorf("invalid repository url %q: must be [<repo-host>/]<repo-owner>/<repo-name>", repoUrl)
		}
		return Repo{
			Host:  host,
			Owner: repoTokens[0],
			Name:  repoTokens[1],
			Url:   repoUri.String(),
		}, nil
	}

	switch len(repoTokens) {
	case 2:
		// <repo-owner>/<repo-name>
		return Repo{
			Host:  DefaultRepoHost,
			Owner: repoTokens[0],
			Name:  repoTokens[1],
			Url:   repoUri.String(),
		}, nil
	case 3:
		// <repo-host>/<repo-owner>/<repo-name>
		return Repo{
			Host:  repoTokens[0],
			Owner: repoTokens[1],
			Name:  repoTokens[2],
			Url:   repoUri.String(),
		}, nil
	default:
		return Repo{}, fmt.Errorf("invalid repository url %q: must be [<repo-host>/]<repo-owner>/<repo-name>", repoUrl)
	}
}

func (r *Repo) InferVcsProvider() {
	switch r.Host {
	case "github.com":
		r.Provider = CommitInfoVcsProviderGithub
	case "gitlab.com":
		r.Provider = CommitInfoVcsProviderGitlab
	case "bitbucket.org":
		r.Provider = CommitInfoVcsProviderBitbucket
	}
}
