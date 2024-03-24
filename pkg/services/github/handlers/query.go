package handlers

import (
	"github.com/alpha-omega-corp/services/types"
	"github.com/google/go-github/v56/github"
	"net/http"
)

type QueryHandler interface {
	query(method string, path string) (*http.Response, error)
}

type queryHandler struct {
	QueryHandler

	client *github.Client
	config types.Config
}

func NewQueryHandler(cli *github.Client, c types.Config) QueryHandler {
	return &queryHandler{
		client: cli,
		config: c,
	}
}

func (h *queryHandler) query(method string, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, h.config.Viper.GetString("url")+path, nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+h.config.Viper.GetString("token"))
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
