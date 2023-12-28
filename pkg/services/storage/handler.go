package storage

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var (
	//go:embed embed
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

type Handler interface {
	GetDirectories(dir string) ([]fs.DirEntry, error)
	GetPackageFile(path string) []byte
	CreatePackage(pkg *models.ContainerPackage, file []byte) error
	PushPackage(name string) error
}

type storageHandler struct {
	Handler
	fileSys fs.FS
}

func NewHandler() Handler {
	return &storageHandler{
		fileSys: getFS(),
	}
}

func (h *storageHandler) GetDirectories(dir string) ([]fs.DirEntry, error) {
	return fs.ReadDir(h.fileSys, dir)
}

func (h *storageHandler) GetPackageFile(path string) []byte {
	file, err := fs.ReadFile(h.fileSys, path)
	if err != nil {
		panic(err)
	}
	return file
}

func (h *storageHandler) CreatePackage(pkg *models.ContainerPackage, file []byte) error {
	path := "pkg/services/storage/embed/" + pkg.Name
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}

	if err := h.writeFiles(path, pkg, file); err != nil {
		return err
	}

	return nil
}

func (h *storageHandler) PushPackage(name string) error {
	if err := runMake(name, "create"); err != nil {
		return err
	}
	if err := runMake(name, "tag"); err != nil {
		return err
	}
	if err := runMake(name, "push"); err != nil {
		return err
	}

	return nil
}

func getFS() fs.FS {
	unwrapFSOnce.Do(func() {
		fileSys, err := fs.Sub(embedFS, "embed")
		if err != nil {
			panic(err)
		}
		unwrappedFS = fileSys
	})
	return unwrappedFS
}

func (h *storageHandler) writeFiles(path string, pkg *models.ContainerPackage, file []byte) error {
	file = bytes.Trim(file, "\x00")

	if err := os.WriteFile(path+"/Dockerfile", file, 0644); err != nil {
		return err
	}

	mFile, err := os.Create(path + "/Makefile")
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
		padLeft("docker build -t alpha-omega-corp/" + pkg.Name + ":" + pkg.Tag + " ."),
		"tag:",
		padLeft("docker tag alpha-omega-corp/" + pkg.Name + ":" + pkg.Tag + " ghcr.io/alpha-omega-corp/" + pkg.Name + ":" + pkg.Tag),
		"push:",
		padLeft("docker push ghcr.io/alpha-omega-corp/" + pkg.Name + ":" + pkg.Tag),
	}

	for _, line := range lines {

		_, err := mFile.WriteString(line + "\n")

		if err != nil {
			return err
		}
	}

	return nil
}

func writeMakefile(pkg string, tag string) error {
	path := "embed/" + pkg
	mFile, err := os.Create(path + "/Makefile")
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
		padLeft("docker build -t alpha-omega-corp/" + pkg + ":" + tag + " ."),
		"tag:",
		padLeft("docker tag alpha-omega-corp/" + pkg + ":" + tag + " ghcr.io/alpha-omega-corp/" + pkg + ":" + tag),
		"push:",
		padLeft("docker push ghcr.io/alpha-omega-corp/" + pkg + ":" + tag),
	}

	fileBytes := []byte(nil)

	for _, line := range lines {
		fileBytes = append(fileBytes, []byte(line+"\n")...)
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

func runMake(pkgName string, act string) error {
	path, err := filepath.Abs("pkg/services/storage/embed/" + pkgName)
	if err != nil {
		return err
	}

	cmd := exec.Command("make", act)
	cmd.Dir = path

	res, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Println(string(res))

	return nil
}
