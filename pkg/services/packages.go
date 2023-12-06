package services

import (
	"context"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/uptrace/bun"
	"os"
)

type PackageHandler interface {
	Create(file []byte, workdir string, tag string) error
	GetAll(ctx context.Context) ([]models.ContainerPackage, error)
	GetOne(id int64, ctx context.Context) (*models.ContainerPackage, error)
	Delete(id int64, ctx context.Context) error
	GetByName(name string, ctx context.Context) models.ContainerPackage
}

type packageHandler struct {
	PackageHandler
	db *bun.DB
}

func NewPackageHandler(db *bun.DB) PackageHandler {
	return &packageHandler{
		db: db,
	}
}

func (h *packageHandler) Create(file []byte, workdir string, tag string) error {
	path := "storage/" + workdir + "/Dockerfile"

	fmt.Print(file)
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

func (h *packageHandler) GetAll(ctx context.Context) ([]models.ContainerPackage, error) {
	var packages []models.ContainerPackage
	err := h.db.NewSelect().Model(&packages).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (h *packageHandler) GetOne(id int64, ctx context.Context) (*models.ContainerPackage, error) {
	var pkg models.ContainerPackage
	err := h.db.NewSelect().Model(&pkg).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &pkg, nil
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

	return nil
}

func (h *packageHandler) GetByName(name string, ctx context.Context) models.ContainerPackage {
	var pkg models.ContainerPackage
	err := h.db.NewSelect().Model(&pkg).Where("name = ?", name).Scan(ctx)
	if err != nil {
		panic(err)
	}

	return pkg
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
		padLeft("docker build -t alpha-omega-corp/" + workdir + ":" + tag + " ."),
		"tag:",
		padLeft("docker tag alpha-omega-corp/" + workdir + ":" + tag + " alpha-omega-corp/" + workdir + ":" + tag),
		"push:",
		padLeft("docker push alpha-omega-corp/" + workdir + ":" + tag),
	}

	for _, line := range lines {

		_, err := mFile.WriteString(line + "\n")

		if err != nil {
			return err
		}
	}

	return nil
}

func padLeft(s string) string {
	return fmt.Sprintf("%s"+s, "\t")
}
