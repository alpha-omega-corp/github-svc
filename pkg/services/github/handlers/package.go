package handlers

import (
	"encoding/json"
	"github.com/alpha-omega-corp/github-svc/pkg/types"
	"github.com/alpha-omega-corp/services/config"
	"github.com/google/go-github/v56/github"
	"io"
	"strconv"
)

type PackageHandler interface {
	GetVersions(name string) ([]types.GitPackageVersion, error)
	GetVersion(name string, vId int64) (*types.GitPackageVersion, error)
	Push(path string) (err error)
	Delete(name string, vId *int64) error
}

type packageHandler struct {
	PackageHandler
	config       config.GithubConfig
	execHandler  ExecHandler
	queryHandler QueryHandler
}

func NewPackageHandler(config config.GithubConfig, cli *github.Client, execHandler ExecHandler) PackageHandler {
	return &packageHandler{
		config:       config,
		queryHandler: NewQueryHandler(cli, config),
		execHandler:  execHandler,
	}
}

func (h *packageHandler) Push(path string) (err error) {
	err = h.execHandler.RunMakefile(path, "create")
	err = h.execHandler.RunMakefile(path, "tag")
	err = h.execHandler.RunMakefile(path, "push")

	return
}

func (h *packageHandler) GetVersions(name string) ([]types.GitPackageVersion, error) {
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

	var versions []types.GitPackageVersion
	if errBuf := json.Unmarshal(body, &versions); err != nil {
		return nil, errBuf
	}

	return versions, nil
}

func (h *packageHandler) GetVersion(name string, vId int64) (*types.GitPackageVersion, error) {
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

	pkg := new(types.GitPackageVersion)
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
