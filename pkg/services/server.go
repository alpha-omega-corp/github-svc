package services

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bun"
)

type Server struct {
	proto.UnimplementedDockerServiceServer
	db *bun.DB
}

func NewServer(db *bun.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) CreateContainer(ctx context.Context, req *proto.CreateContainerRequest) (*proto.CreateContainerResponse, error) {
	return &proto.CreateContainerResponse{
		Status:      200,
		ContainerId: 1,
	}, nil

}
