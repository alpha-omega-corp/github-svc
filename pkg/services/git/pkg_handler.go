package git

import (
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"github.com/google/go-github/v56/github"
	"io"
	"net/http"
)

type PackageHandler interface {
	Push(cp *models.ContainerPackage) error
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

func (h *packageHandler) Push(cp *models.ContainerPackage) error {
	res, err := h.ClientQuery("packages/container/" + cp.Name)
	if err != nil {
		return err
	}

	fmt.Print(res)

	return nil
}

func (h *packageHandler) GetOne(pkgName string) (*types.GitPackage, error) {
	httpClient := h.client.Client()
	res, err := httpClient.Get("https://api.github.com/orgs/alpha-omega-corp/packages/container/" + pkgName)
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

func (h *packageHandler) ClientQuery(path string) (res *http.Response, err error) {
	client := h.rawClient

	return client.Get("https://api.github.com/orgs/alpha-omega-corp/" + path)
}
