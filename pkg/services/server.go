package services

import (
	"context"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/docker/docker/api/types/container"
	"github.com/uptrace/bun"
	"net/http"
)

type Server struct {
	proto.UnimplementedDockerServiceServer

	docker DockerHandler
	db     *bun.DB
}

func NewServer(db *bun.DB) *Server {
	return &Server{
		db:     db,
		docker: NewDockerHandler(),
	}
}

func (s *Server) CreateContainer(ctx context.Context, req *proto.CreateContainerRequest) (*proto.CreateContainerResponse, error) {
	id, err := s.docker.Container().Create(req.Name, &container.Config{
		Image: req.Image,
		Cmd:   []string{"echo", "booted"},
		Tty:   false,
	}, ctx)

	if err != nil {
		return nil, err
	}

	fmt.Print(id)

	return &proto.CreateContainerResponse{
		Status:    http.StatusCreated,
		Container: id,
	}, nil

}
