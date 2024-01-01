package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/docker"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/github"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/uptrace/bun"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var repository = "container-images"

type Server struct {
	proto.UnimplementedDockerServiceServer

	gitHandler    github.Handler
	dockerHandler docker.Handler
}

func NewServer(db *bun.DB) *Server {
	return &Server{
		gitHandler:    github.NewHandler(),
		dockerHandler: docker.NewHandler(db),
	}
}

func (s *Server) PushPackage(ctx context.Context, req *proto.PushPackageRequest) (*proto.PushPackageResponse, error) {
	buf, err := s.gitHandler.Templates().CreateMakefile(req.Name, req.Tag)
	if err != nil {
		return nil, err
	}

	path := req.Name + "/" + req.Tag + "/Dockerfile"
	content, err := s.gitHandler.Repositories().GetContents(ctx, repository, path)
	if err != nil {
		return nil, err
	}

	dockerfile, cErr := content.File.GetContent()
	if err != nil {
		return nil, cErr
	}

	dir, err := os.MkdirTemp("/tmp", req.VersionSHA)
	if err != nil {
		return nil, err
	}

	defer func(path string) {
		rmErr := os.RemoveAll(path)
		if rmErr != nil {
			panic(rmErr)
		}
	}(dir)

	tmpMakefile := filepath.Join(dir, "Makefile")
	if err := os.WriteFile(tmpMakefile, buf.Bytes(), 0644); err != nil {
		return nil, err
	}

	tmpDockerfile := filepath.Join(dir, "Dockerfile")
	if err := os.WriteFile(tmpDockerfile, []byte(dockerfile), 0644); err != nil {
		return nil, err
	}

	if err := s.gitHandler.Packages().Push(dir); err != nil {
		return nil, err
	}

	return &proto.PushPackageResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) CreatePackage(ctx context.Context, req *proto.CreatePackageRequest) (*proto.CreatePackageResponse, error) {
	path := req.Workdir + "/" + req.Tag + "/Dockerfile"
	file := bytes.Trim(req.Dockerfile, "\x00")

	if err := s.gitHandler.Repositories().PutContents(ctx, repository, path, file, nil); err != nil {
		return nil, err
	}

	return &proto.CreatePackageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) CreatePackageVersion(ctx context.Context, req *proto.CreatePackageVersionRequest) (*proto.CreatePackageVersionResponse, error) {
	path := req.Name + "/" + req.Tag + "/Dockerfile"
	file := bytes.Trim(req.Content, "\x00")

	fmt.Print(file)
	if err := s.gitHandler.Repositories().PutContents(ctx, repository, path, file, nil); err != nil {
		return nil, err
	}

	return &proto.CreatePackageVersionResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) DeletePackageVersion(ctx context.Context, req *proto.DeletePackageRequest) (*proto.DeletePackageResponse, error) {
	if err := s.gitHandler.Packages().Delete(req.PkgID.Name, req.PkgID.Tag); err != nil {
		return nil, err
	}

	ctx.Done()
	return &proto.DeletePackageResponse{
		Status: http.StatusOK,
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

func (s *Server) GetPackage(ctx context.Context, req *proto.GetPackageRequest) (*proto.GetPackageResponse, error) {
	c, err := s.gitHandler.Repositories().GetContents(ctx, "container-images", req.Name)
	if err != nil {
		return nil, err
	}

	versionSlice := make([]*proto.PackageVersion, len(c.Dir))
	for index, dir := range c.Dir {
		pkg, err := s.gitHandler.Packages().Get(req.Name, *dir.Name)
		if err != nil {
			return nil, err
		}

		fmt.Print(pkg)

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

func (s *Server) GetPackageFile(ctx context.Context, req *proto.GetPackageFileRequest) (*proto.GetPackageFileResponse, error) {
	c, err := s.gitHandler.Repositories().GetContents(ctx, "container-images", req.Path+"/"+req.Name)
	if err != nil {
		return nil, err
	}

	file, err := c.File.GetContent()
	if err != nil {
		return nil, err
	}

	return &proto.GetPackageFileResponse{
		Content: []byte(file),
	}, nil
}

func (s *Server) GetContainers(ctx context.Context, req *proto.GetContainersRequest) (*proto.GetContainersResponse, error) {
	containers, err := s.dockerHandler.Container().GetAll(ctx)

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
	logs, err := s.dockerHandler.Container().GetLogs(req.ContainerId, ctx)
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
	if err := s.dockerHandler.Container().Delete(req.ContainerId, ctx); err != nil {
		return nil, err
	}

	return &proto.DeleteContainerResponse{
		Status: http.StatusOK,
	}, nil
}
