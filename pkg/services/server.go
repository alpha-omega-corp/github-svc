package services

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/docker"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bun"
	"io"
	"net/http"
	"strings"
)

type Server struct {
	proto.UnimplementedDockerServiceServer

	docker docker.Handler
	pkg    PackageHandler
}

func NewServer(db *bun.DB) *Server {
	return &Server{
		docker: docker.NewHandler(db),
		pkg:    NewPackageHandler(db),
	}
}

func (s *Server) GetContainers(ctx context.Context, req *proto.GetContainersRequest) (*proto.GetContainersResponse, error) {
	containers, err := s.docker.Container().GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var resSlice []*proto.Container
	for _, c := range containers {
		resSlice = append(resSlice, &proto.Container{
			Id:      c.ID,
			Image:   c.Image,
			Status:  c.Status,
			Command: c.Command,
			Created: c.Created,
			State:   c.State,
			Names:   c.Names,
		})
	}

	return &proto.GetContainersResponse{
		Containers: resSlice,
	}, nil
}

func (s *Server) GetContainerLogs(ctx context.Context, req *proto.GetContainerLogsRequest) (*proto.GetContainerLogsResponse, error) {
	logs, err := s.docker.Container().GetLogs(req.ContainerId, ctx)
	if err != nil {
		return nil, err
	}

	logsBuffer := new(strings.Builder)
	_, bufErr := io.Copy(logsBuffer, logs)

	if bufErr != nil {
		return nil, bufErr
	}

	return &proto.GetContainerLogsResponse{
		Logs: logsBuffer.String(),
	}, nil
}

func (s *Server) GetPackages(ctx context.Context, req *proto.GetPackagesRequest) (*proto.GetPackagesResponse, error) {
	packages, err := s.pkg.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var resSlice []*proto.SimplePackage
	for _, pkg := range packages {
		resSlice = append(resSlice, &proto.SimplePackage{
			Id:   pkg.ID,
			Tag:  pkg.Tag,
			Name: pkg.Name,
		})
	}
	return &proto.GetPackagesResponse{
		Packages: resSlice,
	}, nil
}

func (s *Server) GetPackage(ctx context.Context, req *proto.GetPackageRequest) (*proto.GetPackageResponse, error) {
	pkg, err := s.pkg.GetOne(req.Id, ctx)
	if err != nil {
		return nil, err
	}

	return &proto.GetPackageResponse{
		Package: &proto.Package{
			Id:         pkg.ID,
			Tag:        pkg.Tag,
			Name:       pkg.Name,
			Dockerfile: pkg.Dockerfile,
			Makefile:   pkg.Makefile,
			Git: &proto.GitPackage{
				Id:         pkg.Git.Id,
				Name:       pkg.Git.Name,
				Type:       pkg.Git.Type,
				Version:    pkg.Git.Version,
				Visibility: pkg.Git.Visibility,
				Url:        pkg.Git.Url,
				OwnerId:    pkg.Git.Owner.Id,
				OwnerName:  pkg.Git.Owner.Name,
				OwnerNode:  pkg.Git.Owner.NodeId,
				OwnerType:  pkg.Git.Owner.Type,
			},
		},
	}, nil
}

func (s *Server) CreatePackage(ctx context.Context, req *proto.CreatePackageRequest) (*proto.CreatePackageResponse, error) {
	if err := s.pkg.Create(req.Dockerfile, req.Workdir, req.Tag, ctx); err != nil {
		return nil, err
	}

	return &proto.CreatePackageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) DeletePackage(ctx context.Context, req *proto.DeletePackageRequest) (*proto.DeletePackageResponse, error) {
	if err := s.pkg.Delete(req.Id, ctx); err != nil {
		return nil, err
	}

	return &proto.DeletePackageResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) PushPackage(ctx context.Context, req *proto.PushPackageRequest) (*proto.PushPackageResponse, error) {
	if err := s.pkg.Push(req.Id, ctx); err != nil {
		return nil, err
	}

	return &proto.PushPackageResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) ContainerPackage(ctx context.Context, req *proto.ContainerPackageRequest) (*proto.ContainerPackageResponse, error) {
	if err := s.pkg.Container(req.Id, ctx); err != nil {
		return nil, err
	}

	return &proto.ContainerPackageResponse{
		Status: http.StatusOK,
	}, nil
}
