package git

import (
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"github.com/google/go-github/v56/github"
	"io"
)

type PackageHandler interface {
	Get(pkgName string) (*types.GitPackage, error)
}

type packageHandler struct {
	PackageHandler
	client *github.Client
}

func NewPackageHandler(c *github.Client) PackageHandler {
	return &packageHandler{
		client: c,
	}
}

func (h *packageHandler) Get(pkgName string) (*types.GitPackage, error) {
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
