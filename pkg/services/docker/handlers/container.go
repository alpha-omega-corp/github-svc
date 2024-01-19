package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/alpha-omega-corp/services/types"
	docker "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	_ "github.com/spf13/viper/remote"
	"io"
	"strings"
)

type ContainerHandler interface {
	CreateFrom(ctx context.Context, path string, name string) error
	GetAll(ctx context.Context) ([]docker.Container, error)
	GetAllFrom(ctx context.Context, path string) ([]docker.Container, error)
	Delete(ctx context.Context, cId string) error
	GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error)
}

type containerHandler struct {
	ContainerHandler
	config types.ConfigGithubService
	client *client.Client
}

func NewContainerHandler(c types.ConfigGithubService, cli *client.Client) ContainerHandler {
	return &containerHandler{
		config: c,
		client: cli,
	}
}

func (h *containerHandler) Delete(ctx context.Context, cId string) error {
	return h.client.ContainerRemove(ctx, cId, docker.ContainerRemoveOptions{
		Force: true,
	})
}

func (h *containerHandler) GetAll(ctx context.Context) ([]docker.Container, error) {
	containers, err := h.client.ContainerList(ctx, docker.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (h *containerHandler) GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error) {
	options := docker.ContainerLogsOptions{
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
	authConfig := docker.AuthConfig{
		Username: "packages",
		Password: h.config.Organization.Token,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}

	authString := base64.URLEncoding.EncodeToString(encodedJSON)
	_, err = h.client.ImagePull(ctx, imgName, docker.ImagePullOptions{RegistryAuth: authString})
	if err != nil {
		return err
	}

	return nil
}

func (h *containerHandler) GetAllFrom(ctx context.Context, path string) ([]docker.Container, error) {
	filter := filters.NewArgs(filters.KeyValuePair{Key: "ancestor", Value: h.imageName(path)})
	return h.client.ContainerList(ctx, docker.ContainerListOptions{
		All:     true,
		Filters: filter,
	})
}

func (h *containerHandler) CreateFrom(ctx context.Context, path string, name string) error {
	imgName := h.imageName(path)
	if err := h.PullImage(imgName, ctx); err != nil {
		return err
	}

	resp, err := h.client.ContainerCreate(ctx, &container.Config{
		Image: imgName,
	}, nil, nil, nil, name)
	if err != nil {
		panic(err)
	}

	if err := h.client.ContainerStart(ctx, resp.ID, docker.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return nil
}

func (h *containerHandler) imageName(path string) string {
	return h.config.Organization.Registry + "/" + h.config.Organization.Name + "/" + strings.Replace(path, "/", ":", 1)
}
