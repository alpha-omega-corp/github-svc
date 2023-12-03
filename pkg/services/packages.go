package services

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/uptrace/bun"
)

type PackageHandler interface {
	GetAll(ctx context.Context) ([]models.ContainerPackage, error)
	GetFile(t string) ([]byte, error)
	Package(name string, ctx context.Context) *models.ContainerPackage
}

type packageHandler struct {
	PackageHandler
	dbHandler *bun.DB
}

func NewPackageHandler(db *bun.DB) PackageHandler {
	return &packageHandler{
		dbHandler: db,
	}
}

func (h *packageHandler) GetAll(ctx context.Context) ([]models.ContainerPackage, error) {
	var packages []models.ContainerPackage
	err := h.dbHandler.NewSelect().Model(&packages).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (h *packageHandler) GetByName(name string, ctx context.Context) models.ContainerPackage {
	var pkg models.ContainerPackage
	err := h.dbHandler.NewSelect().Model(&pkg).Where("name = ?", name).Scan(ctx)
	if err != nil {
		panic(err)
	}

	return pkg
}
