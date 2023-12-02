package services

import (
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/config"
	"github.com/google/go-github/v56/github"
)

type GitHandler interface {
	Repositories() RepositoryHandler
}

type gitHandler struct {
	GitHandler
	repoHandler RepositoryHandler
}

func NewGitHandler() GitHandler {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Print(c.GIT)
	client := github.NewClient(nil).WithAuthToken(c.GIT)
	return &gitHandler{
		repoHandler: NewRepositoryHandler(client),
	}
}

func (git *gitHandler) Repositories() RepositoryHandler {
	return git.repoHandler
}
