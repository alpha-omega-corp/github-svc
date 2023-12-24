package git

import (
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/config"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"github.com/google/go-github/v56/github"
	"io"
	"net/http"
)

var orgUrl = "https://api.github.com/orgs/alpha-omega-corp/"

type PackageHandler interface {
	GetOne(pkgName string) (*types.GitPackage, error)
	Delete(pkgName string) error
}

type packageHandler struct {
	PackageHandler
	client    *github.Client
	rawClient *http.Client
	config    config.Config
}

func NewPackageHandler(c *github.Client) PackageHandler {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	return &packageHandler{
		client:    c,
		rawClient: c.Client(),
		config:    cfg,
	}
}

func (h *packageHandler) GetOne(pkgName string) (*types.GitPackage, error) {

	res, err := h.queryGet("packages/container/" + pkgName)
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

func (h *packageHandler) queryGet(path string) (*http.Response, error) {
	res, err := h.client.Client().Get(orgUrl + path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *packageHandler) query(method string, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, orgUrl+path, nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+h.config.GIT)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	if err != nil {
		return nil, err
	}

	res, err := h.rawClient.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(res)

	return res, nil
}
