package services

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

type ContainerHandler interface {
	CreatePackage(file []byte, workdir string, tag string, ctx context.Context) error
	GetAll(ctx context.Context) ([]types.Container, error)
	GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error)
}

type containerHandler struct {
	ContainerHandler
	client *client.Client
}

func NewContainerHandler(c *client.Client) ContainerHandler {
	return &containerHandler{
		client: c,
	}
}

func (h *containerHandler) GetAll(ctx context.Context) ([]types.Container, error) {
	containers, err := h.client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (h *containerHandler) GetLogs(containerId string, ctx context.Context) (io.ReadCloser, error) {
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		Timestamps: true,
		Details:    false,
	}
	logs, err := h.client.ContainerLogs(ctx, containerId, options)

	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (h *containerHandler) CreatePackage(file []byte, workdir string, tag string, ctx context.Context) error {
	path := "storage/" + workdir + "/Dockerfile"

	fmt.Print(path)
	if err := os.MkdirAll("storage/"+workdir, os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		return err
	}

	go func() {
		err := makeFile(workdir, tag)
		if err != nil {
			panic(err)
		}
	}()

	return nil
}

func PadLeft(s string) string {
	return fmt.Sprintf("%s"+s, "\t")
}

func makeFile(workdir string, tag string) error {
	workDirPath := "storage/" + workdir
	mFile, err := os.Create(workDirPath + "/Makefile")
	if err != nil {
		return err
	}
	defer func(mFile *os.File) {
		err := mFile.Close()
		if err != nil {
			panic(err)
		}
	}(mFile)

	lines := []string{
		"create:",
		PadLeft("docker build -t alpha-omega-corp/" + workdir + ":" + tag + " ."),
		"tag:",
		PadLeft("docker tag alpha-omega-corp/" + workdir + ":" + tag + " alpha-omega-corp/" + workdir + ":" + tag),
		"push:",
		PadLeft("docker push alpha-omega-corp/" + workdir + ":" + tag),
	}

	for _, line := range lines {

		_, err := mFile.WriteString(line + "\n")

		if err != nil {
			return err
		}
	}

	return nil
}
