package services

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

type ContainerHandler interface {
	Create(file []byte, workdir string, ctx context.Context) error
	GetAll(ctx context.Context) ([]types.Container, error)
	GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error)
}

type containerHandler struct {
	ContainerHandler
	client *client.Client
}

func NewContainerHandler(c *client.Client) ContainerHandler {
	return &containerHandler{
		client: c,
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

func (h *containerHandler) Create(file []byte, workdir string, ctx context.Context) error {
	path := "storage/" + workdir + "/Dockerfile"

	err := os.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
