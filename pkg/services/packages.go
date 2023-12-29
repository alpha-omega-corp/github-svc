package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/docker"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/github"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/template"
	"github.com/uptrace/bun"
	"os"
)

type PackageHandler interface {
	Create(file []byte, name string, tag string, ctx context.Context) error
	GetAll(ctx context.Context) ([]models.ContainerPackage, error)
	GetOne(name string, ctx context.Context) (*models.ContainerPackage, error)
	Delete(id int64, ctx context.Context) error
	Push(id int64, ctx context.Context) error
	CreateContainer(id int64, ctName string, ctx context.Context) error
}

type packageHandler struct {
	PackageHandler

	db     *bun.DB
	docker docker.Handler
	store  template.Handler
	git    github.Handler
}

func NewPackageHandler(db *bun.DB) PackageHandler {
	return &packageHandler{
		db:     db,
		docker: docker.NewHandler(db),
		store:  template.NewHandler(),
		git:    github.NewHandler(),
	}
}

func (h *packageHandler) Create(file []byte, name string, tag string, ctx context.Context) (err error) {
	file = bytes.Trim(file, "\x00")

	err = h.git.Repositories().PutContents(ctx, "container-images", name+"/"+tag+"/Dockerfile", file)
	return
}

func (h *packageHandler) GetOne(name string, ctx context.Context) (*models.ContainerPackage, error) {

	return nil, nil
}

func (h *packageHandler) GetAll(ctx context.Context) ([]models.ContainerPackage, error) {
	var packages []models.ContainerPackage
	err := h.db.NewSelect().Model(&packages).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (h *packageHandler) Delete(id int64, ctx context.Context) error {

	pkg := new(models.ContainerPackage)
	if err := h.db.NewSelect().Model(pkg).Where("id = ?", id).Scan(ctx); err != nil {
		return err
	}

	if err := os.RemoveAll("template/" + pkg.Name); err != nil {
		return err
	}
	if err := h.git.Packages().Delete(pkg.Name); err != nil {
		return err
	}

	_, err := h.db.NewDelete().Model(&models.ContainerPackage{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *packageHandler) Push(id int64, ctx context.Context) error {
	pkg := new(models.ContainerPackage)
	if err := h.db.NewSelect().Model(pkg).Where("id = ?", id).Scan(ctx); err != nil {
		return err
	}

	if err := h.store.PushPackage(pkg.Name); err != nil {
		return err
	}

	pkg.Pushed = true
	_, err := h.db.NewUpdate().Model(pkg).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h *packageHandler) CreateContainer(id int64, ctName string, ctx context.Context) error {
	pkg := new(models.ContainerPackage)
	if err := h.db.NewSelect().Model(pkg).Where("id = ?", id).Scan(ctx); err != nil {
		return err
	}

	if err := h.docker.Container().CreateFrom(pkg, ctName, ctx); err != nil {
		return err
	}

	return nil
}

func writeMakefile(name string, tag string) []byte {
	lines := []string{
		"create:",
		padLeft("docker build -t alpha-omega-corp/" + name + ":" + tag + " ."),
		"tag:",
		padLeft("docker tag alpha-omega-corp/" + name + ":" + tag + " ghcr.io/alpha-omega-corp/" + name + ":" + tag),
		"push:",
		padLeft("docker push ghcr.io/alpha-omega-corp/" + name + ":" + tag),
	}

	fileBytes := []byte(nil)

	for _, line := range lines {
		fileBytes = append(fileBytes, []byte(line+"\n")...)
	}

	return fileBytes
}

func padLeft(s string) string {
	return fmt.Sprintf("%s"+s, "\t")
}
