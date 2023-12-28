package git

import (
	"context"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
)

type RepositoryHandler interface {
	GetContents(ctx context.Context, repo string, path string) (file *github.RepositoryContent, dir []*github.RepositoryContent, err error)
	PutContents(ctx context.Context, repo string, path string, content []byte) error
	GetAll(ctx context.Context) ([]*github.Repository, error)
}

type repositoryHandler struct {
	RepositoryHandler
	config config.GithubConfig
	client *github.Client
}

func NewRepositoryHandler(cli *github.Client, c config.GithubConfig) RepositoryHandler {
	return &repositoryHandler{
		config: c,
		client: cli,
	}
}

func (h *repositoryHandler) GetAll(ctx context.Context) ([]*github.Repository, error) {
	opt := &github.RepositoryListByOrgOptions{}
	packages, _, err := h.client.Repositories.ListByOrg(ctx, h.config.Organization.Name, opt)

	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (h *repositoryHandler) GetContents(ctx context.Context, repo string, path string) (file *github.RepositoryContent, dir []*github.RepositoryContent, err error) {
	file, dir, _, err = h.client.Repositories.GetContents(ctx, h.config.Organization.Name, repo, path, nil)

	if err != nil {
		return
	}

	return
}

func (h *repositoryHandler) PutContents(ctx context.Context, repo string, path string, content []byte) error {
	_, _, err := h.client.Repositories.CreateFile(ctx, h.config.Organization.Name, repo, path, &github.RepositoryContentFileOptions{
		Message: github.String("Create package"),
		Content: content,
	})
	if err != nil {
		return err
	}

	return nil
}
