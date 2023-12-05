package git

import (
	"github.com/alpha-omega-corp/docker-svc/pkg/config"
	"github.com/google/go-github/v56/github"
)

type Handler interface {
	Repositories() RepositoryHandler
	Packages() PackageHandler
}

type gitHandler struct {
	Handler
	repoHandler RepositoryHandler
	pkgHandler  PackageHandler
}

func NewHandler() Handler {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	client := github.NewClient(nil).WithAuthToken(c.GIT)
	return &gitHandler{
		repoHandler: NewRepositoryHandler(client),
		pkgHandler:  NewPackageHandler(client),
	}
}

func (git *gitHandler) Repositories() RepositoryHandler {
	return git.repoHandler
}

func (git *gitHandler) Packages() PackageHandler {
	return git.pkgHandler
}
