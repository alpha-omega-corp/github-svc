package handlers

import (
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
	"net/http"
)

type QueryHandler interface {
	query(method string, path string) (*http.Response, error)
}

type queryHandler struct {
	QueryHandler

	config config.GithubConfig
	client *github.Client
}

func NewQueryHandler(cli *github.Client, config config.GithubConfig) QueryHandler {
	return &queryHandler{
		config: config,
		client: cli,
	}
}

func (h *queryHandler) query(method string, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, h.config.Organization.Url+path, nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+h.config.Token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	if err != nil {
		return nil, err
	}

	res, err := h.client.Client().Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
