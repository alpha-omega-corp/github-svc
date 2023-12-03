package services

import (
	"bytes"
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bun"
	"io"
	"net/http"
	"strings"
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

	packageHandler := NewPackageHandler(s.db)
	packages, err := packageHandler.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var resSlice []*proto.Package
	for _, pkg := range packages {
		resSlice = append(resSlice, &proto.Package{
			Id:         pkg.ID,
			Tag:        pkg.Tag,
			Name:       pkg.Name,
			Dockerfile: pkg.GetFile("Dockerfile"),
			Makefile:   pkg.GetFile("Makefile"),
		})
	}
	return &proto.GetPackagesResponse{
		Packages: resSlice,
	}, nil
}

func (s *Server) CreatePackage(ctx context.Context, req *proto.CreatePackageRequest) (*proto.CreatePackageResponse, error) {
	req.Dockerfile = bytes.Trim(req.Dockerfile, "\x00")

	err := s.docker.Container().CreatePackage(req.Dockerfile, req.Workdir, req.Tag, ctx)
	if err != nil {
		return nil, err
	}

	pkg := new(models.ContainerPackage)
	pkg.Name = req.Workdir
	pkg.Tag = req.Tag

	_, err = s.db.NewInsert().Model(pkg).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.CreatePackageResponse{
		Status:    http.StatusCreated,
		Container: "2",
	}, nil

}
