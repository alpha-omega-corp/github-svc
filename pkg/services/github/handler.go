package github

import (
	"github.com/alpha-omega-corp/github-svc/pkg/services/github/handlers"
	"github.com/alpha-omega-corp/services/types"
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

func NewHandler(c types.ConfigGithubService) Handler {
	client := github.NewClient(nil).WithAuthToken(c.Organization.Token)
	execHandler := handlers.NewExecHandler()

	return &gitHandler{
		repoHandler:    handlers.NewRepositoryHandler(c, client),
		pkgHandler:     handlers.NewPackageHandler(c, client, execHandler),
		secretsHandler: handlers.NewSecretsHandler(c, client),
		tmplHandler:    handlers.NewTemplateHandler(c),
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
