package handlers

import (
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
	"io"
	"net/http"
)

type PackageHandler interface {
	Get(name string, tag string) (*types.GitPackage, error)
	Push(path string) (err error)
	Delete(name string, tag string) (err error)
}

type packageHandler struct {
	PackageHandler
	client      *github.Client
	config      config.GithubConfig
	execHandler ExecHandler
}

func NewPackageHandler(cli *github.Client, c config.GithubConfig, execHandler ExecHandler) PackageHandler {
	return &packageHandler{
		execHandler: execHandler,
		client:      cli,
		config:      c,
	}
}

func (h *packageHandler) Push(path string) (err error) {
	err = h.execHandler.RunMakefile(path, "create")
	err = h.execHandler.RunMakefile(path, "tag")
	err = h.execHandler.RunMakefile(path, "push")

	return
}

func (h *packageHandler) Get(name string, tag string) (*types.GitPackage, error) {
	res, err := h.query("GET", "packages/container/"+name)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	pkg := new(types.GitPackage)
	if errBuf := json.Unmarshal(body, &pkg); err != nil {
		return nil, errBuf
	}

	return pkg, nil
}

func (h *packageHandler) Delete(name string, tag string) error {
	_, err := h.query("DELETE", "packages/container/"+name+"/versions"+tag)
	if err != nil {
		return err
	}

	return nil
}

func (h *packageHandler) query(method string, path string) (*http.Response, error) {
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
