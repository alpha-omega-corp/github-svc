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

func (s *Server) GetPackageTags(ctx context.Context, req *proto.GetPackageTagsRequest) (*proto.GetPackageTagsResponse, error) {
	res, err := s.gitHandler.Packages().GetVersions(req.Name)
	if err != nil {
		return nil, err
	}

	fmt.Print(res)

	return nil, nil
}

func (s *Server) CreatePackageContainer(ctx context.Context, req *proto.CreatePackageContainerRequest) (*proto.CreatePackageContainerResponse, error) {
	if err := s.dockerHandler.Container().CreateFrom(ctx, req.Path, req.Name); err != nil {
		return nil, err
	}

	return &proto.CreatePackageContainerResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) GetPackageVersionContainers(ctx context.Context, req *proto.GetPackageVersionContainersRequest) (*proto.GetPackageVersionContainersResponse, error) {
	res, err := s.dockerHandler.Container().GetAllFrom(ctx, req.Path)
	if err != nil {
		return nil, err
	}

	resSlice := make([]*proto.Container, len(res))
	for index, container := range res {
		resSlice[index] = &proto.Container{
			Id:      container.ID,
			Names:   container.Names,
			Image:   container.Image,
			Status:  container.Status,
			Command: container.Command,
			State:   container.State,
			Created: container.Created,
		}
	}

	fmt.Print(resSlice)

	return &proto.GetPackageVersionContainersResponse{
		Containers: resSlice,
	}, nil
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
		Status: http.StatusCreated,
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

	if err := s.gitHandler.Repositories().PutContents(ctx, repository, path, file, nil); err != nil {
		return nil, err
	}

	return &proto.CreatePackageVersionResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) DeletePackage(ctx context.Context, req *proto.DeletePackageRequest) (*proto.DeletePackageResponse, error) {
	if err := s.gitHandler.Packages().Delete(req.Name, req.Tag); err != nil {
		return nil, err
	}

	path := req.Name + "/" + req.Tag
	files, err := s.gitHandler.Repositories().GetPackageFiles(ctx, path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if err := s.gitHandler.Repositories().DeleteContents(ctx, repository, path+"/"+file.Name, file.SHA); err != nil {
			return nil, err
		}
	}

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

	return &proto.GetPackagesResponse{
		Packages: resSlice,
	}, nil
}

func (s *Server) GetPackage(ctx context.Context, req *proto.GetPackageRequest) (*proto.GetPackageResponse, error) {
	c, err := s.gitHandler.Repositories().GetContents(ctx, "container-images", req.Name)
	if err != nil {
		return nil, err
	}

	versions, err := s.gitHandler.Packages().GetVersions(req.Name)
	if err != nil {
		return nil, err
	}

	versionMap := make(map[string]int64)
	versionSlice := make([]*proto.PackageVersion, len(c.Dir))
	for _, version := range versions {
		for _, tag := range version.Metadata.Container.Tags {
			versionMap[tag] = version.Id
		}
	}

	for index, dir := range c.Dir {
		vId := versionMap[*dir.Name]

		pkg := &proto.PackageVersion{
			Id:   &vId,
			Name: *dir.Name,
			Path: *dir.Path,
			Sha:  *dir.SHA,
			Link: *dir.HTMLURL,
		}

		versionSlice[index] = pkg
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
	if err := s.dockerHandler.Container().Delete(ctx, req.ContainerId); err != nil {
		return nil, err
	}

	return &proto.DeleteContainerResponse{
		Status: http.StatusOK,
	}, nil
}
