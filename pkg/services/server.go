package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/docker"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/github"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bun"
	"io"
	"net/http"
	"strings"
)

type Server struct {
	proto.UnimplementedDockerServiceServer

	docker     docker.Handler
	pkg        PackageHandler
	gitHandler github.Handler
}

func NewServer(db *bun.DB) *Server {
	return &Server{
		docker:     docker.NewHandler(db),
		pkg:        NewPackageHandler(db),
		gitHandler: github.NewHandler(),
	}
}

func (s *Server) CreatePackage(ctx context.Context, req *proto.CreatePackageRequest) (*proto.CreatePackageResponse, error) {
	if err := s.pkg.Create(req.Dockerfile, req.Workdir, req.Tag, ctx); err != nil {
		return nil, err
	}

	return &proto.CreatePackageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) GetPackages(ctx context.Context, req *proto.GetPackagesRequest) (*proto.GetPackagesResponse, error) {

	c, err := s.gitHandler.Repositories().GetContents(ctx, "container-images", ".")
	if err != nil {
		return nil, err
	}

	resSlice := make([]*proto.SimplePackage, len(c.Dir))
	for index, pkg := range c.Dir {
		b, err := json.Marshal(pkg)
		if err != nil {
			return nil, err
		}

		if mErr := json.Unmarshal(b, &resSlice[index]); mErr != nil {
			return nil, mErr
		}
	}

	fmt.Print(resSlice)
	return &proto.GetPackagesResponse{
		Packages: resSlice,
	}, nil
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

func (s *Server) GetPackage(ctx context.Context, req *proto.GetPackageRequest) (*proto.GetPackageResponse, error) {
	c, err := s.gitHandler.Repositories().GetContents(ctx, "container-images", req.Name)
	if err != nil {
		return nil, err
	}

	versionSlice := make([]*proto.PackageVersion, len(c.Dir))
	for index, dir := range c.Dir {
		versionSlice[index] = &proto.PackageVersion{
			Name: *dir.Name,
			Path: *dir.Path,
			Sha:  *dir.SHA,
			Link: *dir.HTMLURL,
		}
	}

	return &proto.GetPackageResponse{
		Versions: versionSlice,
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

func (s *Server) DeleteContainer(ctx context.Context, req *proto.DeleteContainerRequest) (*proto.DeleteContainerResponse, error) {
	if err := s.docker.Container().Delete(req.ContainerId, ctx); err != nil {
		return nil, err
	}

	return &proto.DeleteContainerResponse{
		Status: http.StatusOK,
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

	buf, err := s.gitHandler.Templates().CreateMakefile(req.Name, req.Tag)
	if err != nil {
		return nil, err
	}

	path := req.Name + "/Makefile"
	fmt.Print(buf.String())

	if pErr := s.gitHandler.Repositories().
		PutContents(ctx, "container-images", path, buf.Bytes()); pErr != nil {
		return nil, pErr
	}

	return &proto.PushPackageResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) ContainerPackage(ctx context.Context, req *proto.ContainerPackageRequest) (*proto.ContainerPackageResponse, error) {
	if err := s.pkg.CreateContainer(req.Id, req.Name, ctx); err != nil {
		return nil, err
	}

	return &proto.ContainerPackageResponse{
		Status: http.StatusOK,
	}, nil
}
