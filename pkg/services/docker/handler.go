package docker

import (
	"github.com/alpha-omega-corp/github-svc/pkg/services/docker/handlers"
	"github.com/alpha-omega-corp/services/types"
	"github.com/docker/docker/client"
)

type Handler interface {
	Container() handlers.ContainerHandler
}

type dockerHandler struct {
	Handler
	ctHandler handlers.ContainerHandler
}

func NewHandler(c types.Config) Handler {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	defer func(c *client.Client) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(cli)

	return &dockerHandler{
		ctHandler: handlers.NewContainerHandler(cli, c),
	}
}

func (h *dockerHandler) Container() handlers.ContainerHandler {
	return h.ctHandler
}
