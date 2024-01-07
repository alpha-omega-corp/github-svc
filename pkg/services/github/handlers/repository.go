package handlers

import (
	"context"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
)

type Content struct {
	File     *github.RepositoryContent
	Dir      []*github.RepositoryContent
	Response *github.Response
}

type PackageFile struct {
	Name    string
	SHA     string
	Content []byte
}

type RepositoryHandler interface {
	GetPackageFiles(ctx context.Context, name string) ([]*PackageFile, error)
	GetContents(ctx context.Context, repo string, path string) (content *Content, err error)
	PutContents(ctx context.Context, repo string, path string, content []byte, sha *string) error
	DeleteContents(ctx context.Context, repo string, path string, sha string) error
	GetAll(ctx context.Context) ([]*github.Repository, error)
}

type repositoryHandler struct {
	RepositoryHandler
	config config.GithubConfig
	client *github.Client
}

func NewRepositoryHandler(config config.GithubConfig, cli *github.Client) RepositoryHandler {
	return &repositoryHandler{
		config: config,
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

func (h *repositoryHandler) GetPackageFiles(ctx context.Context, name string) ([]*PackageFile, error) {
	_, dir, _, err := h.client.Repositories.GetContents(ctx, h.config.Organization.Name, "container-images", name, nil)
	if err != nil {
		return nil, err
	}

	files := make([]*PackageFile, len(dir))
	for index, file := range dir {
		f, _, _, err := h.client.Repositories.GetContents(ctx, h.config.Organization.Name, "container-images", name+"/"+*file.Name, nil)
		if err != nil {
			return nil, err
		}

		content, err := f.GetContent()
		if err != nil {
			return nil, err
		}

		files[index] = &PackageFile{
			SHA:     *file.SHA,
			Name:    *file.Name,
			Content: []byte(content),
		}
	}

	return files, nil
}

func (h *repositoryHandler) GetContents(ctx context.Context, repo string, path string) (*Content, error) {
	file, dir, res, err := h.client.Repositories.GetContents(ctx, h.config.Organization.Name, repo, path, nil)

	if err != nil {
		return nil, err
	}

	return &Content{
		File:     file,
		Dir:      dir,
		Response: res,
	}, nil
}

func (h *repositoryHandler) PutContents(ctx context.Context, repo string, path string, content []byte, sha *string) error {
	_, _, err := h.client.Repositories.CreateFile(ctx, h.config.Organization.Name, repo, path, &github.RepositoryContentFileOptions{
		Message: github.String("Webhook: Action"),
		Content: content,
		SHA:     sha,
	})
	if err != nil {
		return err
	}

	return nil
}

func (h *repositoryHandler) DeleteContents(ctx context.Context, repo string, path string, sha string) error {
	_, _, err := h.client.Repositories.DeleteFile(ctx, h.config.Organization.Name, repo, path, &github.RepositoryContentFileOptions{
		Message: github.String("Webhook: Action"),
		SHA:     github.String(sha),
	})

	if err != nil {
		return err
	}

	return nil
}
