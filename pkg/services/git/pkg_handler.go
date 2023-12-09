package git

import (
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"github.com/google/go-github/v56/github"
	"io"
	"net/http"
)

var orgUrl = "https://api.github.com/orgs/alpha-omega-corp/"

type PackageHandler interface {
	GetOne(pkgName string) (*types.GitPackage, error)
}

type packageHandler struct {
	PackageHandler
	client    *github.Client
	rawClient *http.Client
}

func NewPackageHandler(c *github.Client) PackageHandler {
	return &packageHandler{
		client:    c,
		rawClient: c.Client(),
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

func (h *packageHandler) queryGet(path string) (*http.Response, error) {
	res, err := h.client.Client().Get(orgUrl + path)
	if err != nil {
		return nil, err
	}

	return res, nil
}
