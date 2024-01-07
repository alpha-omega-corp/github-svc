package github

import (
	"github.com/alpha-omega-corp/github-svc/pkg/services/github/handlers"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
	_ "github.com/spf13/viper/remote"
)

type Handler interface {
	Repositories() handlers.RepositoryHandler
	Packages() handlers.PackageHandler
	Secrets() handlers.SecretsHandler
	Templates() handlers.TemplateHandler
	Exec() handlers.ExecHandler
}

type gitHandler struct {
	Handler

	repoHandler    handlers.RepositoryHandler
	pkgHandler     handlers.PackageHandler
	secretsHandler handlers.SecretsHandler
	tmplHandler    handlers.TemplateHandler
	execHandler    handlers.ExecHandler
}

func NewHandler(config config.GithubConfig) Handler {
	client := github.NewClient(nil).WithAuthToken(config.Token)
	execHandler := handlers.NewExecHandler()

	return &gitHandler{
		repoHandler:    handlers.NewRepositoryHandler(config, client),
		pkgHandler:     handlers.NewPackageHandler(config, client, execHandler),
		secretsHandler: handlers.NewSecretsHandler(config, client),
		tmplHandler:    handlers.NewTemplateHandler(config),
		execHandler:    execHandler,
	}
}

func (git *gitHandler) Repositories() handlers.RepositoryHandler {
	return git.repoHandler
}

func (git *gitHandler) Packages() handlers.PackageHandler {
	return git.pkgHandler
}

func (git *gitHandler) Secrets() handlers.SecretsHandler {
	return git.secretsHandler
}

func (git *gitHandler) Templates() handlers.TemplateHandler {
	return git.tmplHandler
}

func (git *gitHandler) Exec() handlers.ExecHandler {
	return git.execHandler
}
