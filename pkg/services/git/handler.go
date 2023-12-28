package git

import (
	"github.com/alpha-omega-corp/services/server"
	"github.com/google/go-github/v56/github"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
	v := viper.New()
	cManager := server.NewConfigManager(v)

	c, err := cManager.GithubConfig()
	if err != nil {
		panic(err)
	}

	client := github.NewClient(nil).WithAuthToken(c.Token)

	return &gitHandler{
		repoHandler: NewRepositoryHandler(client, c),
		pkgHandler:  NewPackageHandler(client, c),
	}
}

func (git *gitHandler) Repositories() RepositoryHandler {
	return git.repoHandler
}

func (git *gitHandler) Packages() PackageHandler {
	return git.pkgHandler
}
