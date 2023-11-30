package services

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"os"
)

type ContainerHandler interface {
	Create(name string, c *container.Config, ctx context.Context) (string, error)
	GetAll(ctx context.Context) ([]types.Container, error)
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

func (h *containerHandler) Create(name string, c *container.Config, ctx context.Context) (s string, err error) {
	out, err := h.client.ImagePull(ctx, c.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)

	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		return "", err
	}

	cc, err := h.client.ContainerCreate(ctx, c, nil, nil, nil, name)

	if err != nil {
		return "", err
	}

	return cc.ID, nil
}
