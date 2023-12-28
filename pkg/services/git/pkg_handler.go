package git

import (
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
	"io"
	"net/http"
)

type PackageHandler interface {
	GetOne(pkgName string) (*types.GitPackage, error)
	Delete(pkgName string) error
}

type packageHandler struct {
	PackageHandler
	client    *github.Client
	rawClient *http.Client
	config    config.GithubConfig
}

func NewPackageHandler(cli *github.Client, c config.GithubConfig) PackageHandler {

	return &packageHandler{
		client:    cli,
		rawClient: cli.Client(),
		config:    c,
	}
}

func (h *packageHandler) GetOne(pkgName string) (*types.GitPackage, error) {
	res, err := h.query("GET", "packages/container/"+pkgName)
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

func (h *packageHandler) Delete(pkgName string) error {
	_, err := h.query("DELETE", "packages/container/"+pkgName)
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

	res, err := h.rawClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
