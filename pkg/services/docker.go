package services

import (
	"github.com/docker/docker/client"
)

type DockerHandler interface {
	Container() ContainerHandler
}

type dockerHandler struct {
	DockerHandler
	ctHandler ContainerHandler
}

func NewDockerHandler() DockerHandler {
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
		ctHandler: NewContainerHandler(c),
	}
}

func (h *dockerHandler) Container() ContainerHandler {
	return h.ctHandler
}
