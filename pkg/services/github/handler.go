package github

import (
	"github.com/alpha-omega-corp/docker-svc/pkg/services/github/handlers"
	"github.com/alpha-omega-corp/services/server"
	"github.com/google/go-github/v56/github"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Handler interface {
	Templates() handlers.TemplateHandler
	Repositories() handlers.RepositoryHandler
	Packages() handlers.PackageHandler
	Exec() handlers.ExecHandler
}

type gitHandler struct {
	Handler

	tmplHandler handlers.TemplateHandler
	repoHandler handlers.RepositoryHandler
	execHandler handlers.ExecHandler
	pkgHandler  handlers.PackageHandler
}

func NewHandler() Handler {
	v := viper.New()
	cManager := server.NewConfigManager(v)
	c, err := cManager.GithubConfig()
	if err != nil {
		panic(err)
	}

	client := github.NewClient(nil).WithAuthToken(c.Token)
	execHandler := handlers.NewExecHandler()

	return &gitHandler{
		tmplHandler: handlers.NewTemplateHandler(c),
		repoHandler: handlers.NewRepositoryHandler(client, c),
		pkgHandler:  handlers.NewPackageHandler(client, c, execHandler),
		execHandler: execHandler,
	}
}

func (git *gitHandler) Templates() handlers.TemplateHandler {
	return git.tmplHandler
}

func (git *gitHandler) Repositories() handlers.RepositoryHandler {
	return git.repoHandler
}

func (git *gitHandler) Packages() handlers.PackageHandler {
	return git.pkgHandler
}

func (git *gitHandler) Exec() handlers.ExecHandler {
	return git.execHandler
}
