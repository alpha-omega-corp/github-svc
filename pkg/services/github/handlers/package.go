package handlers

import (
	"encoding/json"
	pkgTypes "github.com/alpha-omega-corp/github-svc/pkg/types"
	"github.com/alpha-omega-corp/services/types"
	"github.com/google/go-github/v56/github"
	"io"
	"strconv"
)

type PackageHandler interface {
	GetVersions(name string) ([]pkgTypes.GitPackageVersion, error)
	GetVersion(name string, vId int64) (*pkgTypes.GitPackageVersion, error)
	Push(path string) (err error)
	Delete(name string, vId *int64) error
}

type packageHandler struct {
	PackageHandler
	config       types.ConfigGithubService
	execHandler  ExecHandler
	queryHandler QueryHandler
}

func NewPackageHandler(c types.ConfigGithubService, cli *github.Client, execHandler ExecHandler) PackageHandler {
	return &packageHandler{
		queryHandler: NewQueryHandler(cli, c),
		execHandler:  execHandler,
		config:       c,
	}
}

func (h *packageHandler) Push(path string) (err error) {
	err = h.execHandler.RunMakefile(path, "create")
	err = h.execHandler.RunMakefile(path, "tag")
	err = h.execHandler.RunMakefile(path, "push")

	return
}

func (h *packageHandler) GetVersions(name string) ([]pkgTypes.GitPackageVersion, error) {
	res, err := h.queryHandler.query("GET", "packages/container/"+name+"/versions")
	if err != nil {
		return nil, err
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

	var versions []pkgTypes.GitPackageVersion
	if errBuf := json.Unmarshal(body, &versions); err != nil {
		return nil, errBuf
	}

	return versions, nil
}

func (h *packageHandler) GetVersion(name string, vId int64) (*pkgTypes.GitPackageVersion, error) {
	path := "packages/container/" + name + "/versions/" + strconv.FormatInt(vId, 10)
	res, err := h.queryHandler.query("GET", path)
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

	pkg := new(pkgTypes.GitPackageVersion)
	if errBuf := json.Unmarshal(body, &pkg); err != nil {
		return nil, errBuf
	}

	return pkg, nil
}

func (h *packageHandler) Delete(name string, vId *int64) error {
	_, err := h.queryHandler.query("DELETE", "packages/container/"+name+"/versions/"+strconv.FormatInt(*vId, 10))
	if err != nil {
		return err
	}

	return nil
}
