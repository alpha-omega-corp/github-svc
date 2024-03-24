package server

import (
	"context"
	"github.com/alpha-omega-corp/github-svc/pkg/services/docker"
	proto "github.com/alpha-omega-corp/github-svc/proto/docker"
	"github.com/alpha-omega-corp/services/types"
	"io"
	"net/http"
	"strings"
)

var repository = "container-images"

type DockerServer struct {
	proto.UnimplementedDockerServiceServer

	handler docker.Handler
}

func NewDockerServer(env *types.Environment) *DockerServer {
	return &DockerServer{
		handler: docker.NewHandler(env.Config),
	}
}

func (s *DockerServer) StopContainer(ctx context.Context, req *proto.StopContainerRequest) (*proto.StopContainerResponse, error) {
	if err := s.handler.Container().Stop(ctx, req.ContainerId); err != nil {
		return nil, err
	}

	return &proto.StopContainerResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *DockerServer) StartContainer(ctx context.Context, req *proto.StartContainerRequest) (*proto.StartContainerResponse, error) {
	if err := s.handler.Container().Start(ctx, req.ContainerId); err != nil {
		return nil, err
	}

	return &proto.StartContainerResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *DockerServer) GetContainers(ctx context.Context, req *proto.GetContainersRequest) (*proto.GetContainersResponse, error) {
	containers, err := s.handler.Container().GetAll(ctx)

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

func (s *DockerServer) DeleteContainer(ctx context.Context, req *proto.DeleteContainerRequest) (*proto.DeleteContainerResponse, error) {
	if err := s.handler.Container().Delete(ctx, req.ContainerId); err != nil {
		return nil, err
	}

	return &proto.DeleteContainerResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *DockerServer) GetContainerLogs(ctx context.Context, req *proto.GetContainerLogsRequest) (*proto.GetContainerLogsResponse, error) {
	logs, err := s.handler.Container().GetLogs(req.ContainerId, ctx)
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

func (s *DockerServer) CreatePackageContainer(ctx context.Context, req *proto.CreatePackageContainerRequest) (*proto.CreatePackageContainerResponse, error) {
	if err := s.handler.Container().CreateFrom(ctx, req.Path, req.Name); err != nil {
		return nil, err
	}

	return &proto.CreatePackageContainerResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *DockerServer) GetPackageVersionContainers(ctx context.Context, req *proto.GetPackageVersionContainersRequest) (*proto.GetPackageVersionContainersResponse, error) {
	res, err := s.handler.Container().GetAllFrom(ctx, req.Path)
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

	return &proto.GetPackageVersionContainersResponse{
		Containers: resSlice,
	}, nil
}
