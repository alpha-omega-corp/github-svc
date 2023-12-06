package services

import (
	"bytes"
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/git"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bun"
	"io"
	"net/http"
	"strings"
)

type Server struct {
	proto.UnimplementedDockerServiceServer

	docker DockerHandler
	pkg    PackageHandler
	git    git.Handler
	db     *bun.DB
}

func NewServer(db *bun.DB) *Server {
	return &Server{
		db:     db,
		docker: NewDockerHandler(db),
		pkg:    NewPackageHandler(db),
		git:    git.NewHandler(),
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

	gitPkg, err := s.git.Packages().GetOne(pkg.Name)
	if err != nil {
		return nil, err
	}

	return &proto.GetPackageResponse{
		Package: &proto.Package{
			Id:         pkg.ID,
			Tag:        pkg.Tag,
			Name:       pkg.Name,
			Dockerfile: pkg.GetFile("Dockerfile"),
			Makefile:   pkg.GetFile("Makefile"),
			Git: &proto.GitPackage{
				Id:         gitPkg.Id,
				Name:       gitPkg.Name,
				Type:       gitPkg.Type,
				Version:    gitPkg.Version,
				Visibility: gitPkg.Visibility,
				Url:        gitPkg.Url,
				OwnerId:    gitPkg.Owner.Id,
				OwnerName:  gitPkg.Owner.Name,
				OwnerNode:  gitPkg.Owner.NodeId,
				OwnerType:  gitPkg.Owner.Type,
			},
		},
	}, nil
}

func (s *Server) CreatePackage(ctx context.Context, req *proto.CreatePackageRequest) (*proto.CreatePackageResponse, error) {
	req.Dockerfile = bytes.Trim(req.Dockerfile, "\x00")

	err := s.pkg.Create(req.Dockerfile, req.Workdir, req.Tag)
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

func (s *Server) DeletePackage(ctx context.Context, req *proto.DeletePackageRequest) (*proto.DeletePackageResponse, error) {
	if err := s.pkg.Delete(req.Id, ctx); err != nil {
		return nil, err
	}

	return &proto.DeletePackageResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) PushPackage(ctx context.Context, req *proto.PushPackageRequest) (*proto.PushPackageResponse, error) {
	pkg := new(models.ContainerPackage)
	if err := s.db.NewSelect().Model(&pkg).Where("id = ?", req.Id).Scan(ctx); err != nil {
		return nil, err
	}

	if err := s.git.Packages().Push(pkg); err != nil {
		return nil, err
	}

	return &proto.PushPackageResponse{
		Status: http.StatusOK,
	}, nil
}
