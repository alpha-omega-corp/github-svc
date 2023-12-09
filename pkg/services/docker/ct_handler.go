package docker

import (
	"context"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
	"io"
)

type ContainerHandler interface {
	Create(img string, ctx context.Context) error
	GetAll(ctx context.Context) ([]types.Container, error)
	GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error)
}

type containerHandler struct {
	ContainerHandler
	client *client.Client
	config config.Config
	db     *bun.DB
}

func NewContainerHandler(cli *client.Client, db *bun.DB) ContainerHandler {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	return &containerHandler{
		client: cli,
		config: c,
		db:     db,
	}
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

func (h *containerHandler) Create(orgImage string, ctx context.Context) error {
	out, err := h.client.ImagePull(ctx, h.config.GHCR+orgImage, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	fmt.Print(out)

	return nil
}
