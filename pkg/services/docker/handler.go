package docker

import (
	"github.com/alpha-omega-corp/docker-svc/pkg/services/docker/handlers"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
)

type Handler interface {
	Container() handlers.ContainerHandler
}

type dockerHandler struct {
	Handler
	ctHandler handlers.ContainerHandler
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
		ctHandler: handlers.NewContainerHandler(c, db),
	}
}

func (h *dockerHandler) Container() handlers.ContainerHandler {
	return h.ctHandler
}
