package services

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/git"
	"github.com/alpha-omega-corp/docker-svc/pkg/services/storage"
	"github.com/uptrace/bun"
	"os"
)

type PackageHandler interface {
	Create(file []byte, name string, tag string, ctx context.Context) error
	GetAll(ctx context.Context) ([]models.ContainerPackage, error)
	GetOne(id int64, ctx context.Context) (*models.ContainerPackage, error)
	Delete(id int64, ctx context.Context) error
	Push(id int64, ctx context.Context) error
	Container(id int64, ctx context.Context) error
}

type packageHandler struct {
	PackageHandler
	db    *bun.DB
	git   git.Handler
	store storage.Handler
}

func NewPackageHandler(db *bun.DB) PackageHandler {
	return &packageHandler{
		db:    db,
		git:   git.NewHandler(),
		store: storage.NewHandler(),
	}
}

func (h *packageHandler) Create(file []byte, name string, tag string, ctx context.Context) error {
	pkg := new(models.ContainerPackage)
	pkg.Name = name
	pkg.Tag = tag

	_, err := h.db.NewInsert().Model(pkg).Exec(ctx)
	if err != nil {
		return err
	}

	return h.store.CreatePackage(pkg, file)
}

func (h *packageHandler) GetOne(id int64, ctx context.Context) (*models.ContainerPackage, error) {
	var pkg models.ContainerPackage
	err := h.db.NewSelect().Model(&pkg).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	gitPkg, err := h.git.Packages().GetOne(pkg.Name)
	if err != nil {
		return nil, err
	}

	pkg.Dockerfile = h.store.GetPackageFile(pkg.Name + "/Dockerfile")
	pkg.Makefile = h.store.GetPackageFile(pkg.Name + "/Makefile")
	pkg.Git = gitPkg

	return &pkg, nil
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

	if err := os.RemoveAll("storage/" + pkg.Name); err != nil {
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

	return nil
}

func (h *packageHandler) Container(id int64, ctx context.Context) error {
	pkg := new(models.ContainerPackage)
	if err := h.db.NewSelect().Model(pkg).Where("id = ?", id).Scan(ctx); err != nil {
		return err
	}

	return nil
}
