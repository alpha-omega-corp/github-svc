package git

import (
	"context"
	"github.com/google/go-github/v56/github"
)

type RepositoryHandler interface {
	GetAll(ctx context.Context) ([]*github.Repository, error)
}

type repositoryHandler struct {
	RepositoryHandler
	client *github.Client
}

func NewRepositoryHandler(c *github.Client) RepositoryHandler {
	return &repositoryHandler{
		client: c,
	}
}

func (h *repositoryHandler) GetAll(ctx context.Context) ([]*github.Repository, error) {
	opt := &github.RepositoryListByOrgOptions{}
	packages, _, err := h.client.Repositories.ListByOrg(ctx, "alpha-omega-corp", opt)

	if err != nil {
		return nil, err
	}

	return packages, nil
}
