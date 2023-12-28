package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/services/config"
	"github.com/alpha-omega-corp/services/server"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bun"
	"io"
)

type ContainerHandler interface {
	CreateFrom(pkg *models.ContainerPackage, ctName string, ctx context.Context) error
	GetAll(ctx context.Context) ([]types.Container, error)
	GetAllFrom(pkg *models.ContainerPackage, ctx context.Context) ([]types.Container, error)
	Delete(containerId string, ctx context.Context) error
	GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error)
}

type containerHandler struct {
	ContainerHandler
	config config.GithubConfig
	client *client.Client
	db     *bun.DB
}

func NewContainerHandler(cli *client.Client, db *bun.DB) ContainerHandler {
	v := viper.New()
	cManager := server.NewConfigManager(v)

	c, err := cManager.GithubConfig()
	if err != nil {
		panic(err)
	}

	return &containerHandler{
		config: c,
		client: cli,
		db:     db,
	}
}

func (h *containerHandler) Delete(containerId string, ctx context.Context) error {
	return h.client.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{
		Force: true,
	})

}

func (h *containerHandler) GetAll(ctx context.Context) ([]types.Container, error) {
	containers, err := h.client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (h *containerHandler) GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error) {
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		Timestamps: true,
		Details:    false,
	}
	logs, err := h.client.ContainerLogs(ctx, containerId, options)

	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (h *containerHandler) PullImage(imgName string, ctx context.Context) error {
	authConfig := types.AuthConfig{
		Username: "packages",
		Password: h.config.Token,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}

	authString := base64.URLEncoding.EncodeToString(encodedJSON)
	_, err = h.client.ImagePull(ctx, imgName, types.ImagePullOptions{RegistryAuth: authString})
	if err != nil {
		return err
	}

	return nil
}

func (h *containerHandler) GetAllFrom(pkg *models.ContainerPackage, ctx context.Context) ([]types.Container, error) {
	f := filters.NewArgs(filters.KeyValuePair{Key: "ancestor", Value: h.imageName(pkg)})
	return h.client.ContainerList(ctx, types.ContainerListOptions{Filters: f})
}

func (h *containerHandler) CreateFrom(pkg *models.ContainerPackage, ctName string, ctx context.Context) error {
	imgName := h.imageName(pkg)
	if err := h.PullImage(imgName, ctx); err != nil {
		return err
	}

	resp, err := h.client.ContainerCreate(ctx, &container.Config{
		Image: imgName,
	}, nil, nil, nil, ctName)
	if err != nil {
		panic(err)
	}

	if err := h.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return nil
}

func (h *containerHandler) imageName(pkg *models.ContainerPackage) string {
	return h.config.Organization.Registry + "/" + h.config.Organization.Name + "/" + pkg.Name + ":" + pkg.Tag
}
