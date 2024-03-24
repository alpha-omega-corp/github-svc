package github

import (
	"github.com/alpha-omega-corp/github-svc/pkg/services/github/handlers"
	"github.com/alpha-omega-corp/services/types"
	"github.com/google/go-github/v56/github"
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

func NewHandler(c types.Config) Handler {
	client := github.NewClient(nil).WithAuthToken(c.Viper.GetString("token"))

	exec := handlers.NewExecHandler()
	tmpl := handlers.NewTemplateHandler(c)
	repo := handlers.NewRepositoryHandler(client, c)
	secret := handlers.NewSecretsHandler(client, c)
	query := handlers.NewQueryHandler(client, c)
	pkg := handlers.NewPackageHandler(query, exec)

	return &gitHandler{
		tmplHandler:    tmpl,
		repoHandler:    repo,
		secretsHandler: secret,
		pkgHandler:     pkg,
		execHandler:    exec,
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
