package handlers

import (
	"context"
	"fmt"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
)

type SecretsHandler interface {
	GetSecrets(ctx context.Context) ([]*github.Secret, error)
	CreateSecret(ctx context.Context, name string, content []byte) error
	GetEncryptionKey(ctx context.Context) (*github.PublicKey, error)
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

func (h *secretsHandler) CreateSecret(ctx context.Context, name string, content []byte) error {
	key, err := h.GetEncryptionKey(ctx)
	if err != nil {
		return err
	}

	fmt.Print(key)
	return nil
}

func (h *secretsHandler) GetSecrets(ctx context.Context) ([]*github.Secret, error) {
	data, _, err := h.client.Actions.ListOrgSecrets(ctx, h.config.Organization.Name, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	return data.Secrets, nil
}

func (h *secretsHandler) GetEncryptionKey(ctx context.Context) (*github.PublicKey, error) {
	data, _, err := h.client.Actions.GetOrgPublicKey(ctx, h.config.Organization.Name)
	if err != nil {
		return nil, err
	}

	return data, nil
}
