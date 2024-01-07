package handlers

import (
	"context"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
)

type SecretsHandler interface {
	GetSecrets(ctx context.Context) ([]*github.Secret, error)
}

type secretsHandler struct {
	SecretsHandler

	config config.GithubConfig
	client *github.Client
}

func NewSecretsHandler(config config.GithubConfig, cli *github.Client) SecretsHandler {
	return &secretsHandler{
		config: config,
		client: cli,
	}
}

func (h *secretsHandler) GetSecrets(ctx context.Context) ([]*github.Secret, error) {
	data, _, err := h.client.Actions.ListOrgSecrets(ctx, h.config.Organization.Name, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	return data.Secrets, nil
}
