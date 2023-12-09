package docker

import (
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
)

type Handler interface {
	Container() ContainerHandler
}

type dockerHandler struct {
	Handler
	ctHandler ContainerHandler
}

func NewHandler(db *bun.DB) Handler {
	c, err := client.NewClientWithOpts(
		client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	defer func(c *client.Client) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)

	return &dockerHandler{
		ctHandler: NewContainerHandler(c, db),
	}
}

func (h *dockerHandler) Container() ContainerHandler {
	return h.ctHandler
}
