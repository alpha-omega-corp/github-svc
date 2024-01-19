package handlers

import (
	"bytes"
	"embed"
	pkgTypes "github.com/alpha-omega-corp/github-svc/pkg/types"
	"github.com/alpha-omega-corp/services/types"

	"io/fs"
	"sync"
	"text/template"
)

var (
	//go:embed templates
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

type TemplateHandler interface {
	CreateMakefile(pkgName string, pkgTag string) (*bytes.Buffer, error)
	CreateDockerfile(pkgName string, pkgTag string, content []byte) (*bytes.Buffer, error)
	CreateConfiguration(data map[string]string) (*bytes.Buffer, error)
}

type templateHandler struct {
	TemplateHandler

	fs       fs.FS
	template *template.Template
	config   types.ConfigGithubService
}

func NewTemplateHandler(c types.ConfigGithubService) TemplateHandler {
	fileSys := getFS()
	tmpl, err := template.ParseFS(fileSys, "*.template")

	if err != nil {
		panic(err)
	}

	return &templateHandler{
		fs:       fileSys,
		template: tmpl,
		config:   c,
	}
}

func (h *templateHandler) CreateMakefile(pkgName string, pkgTag string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	if err := h.template.ExecuteTemplate(buf, "makefile.template", &pkgTypes.CreateMakefileDto{
		Registry: h.config.Organization.Registry,
		OrgName:  h.config.Organization.Name,
		Name:     pkgName,
		Tag:      pkgTag,
	}); err != nil {
		return nil, err
	}

	return buf, nil
}

func (h *templateHandler) CreateDockerfile(pkgName string, pkgTag string, content []byte) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	if err := h.template.ExecuteTemplate(buf, "dockerfile.template", &pkgTypes.CreateDockerfileDto{
		Name:    pkgName,
		Tag:     pkgTag,
		Author:  h.config.Organization.Name,
		Content: string(bytes.Trim(content, "\x00")),
	}); err != nil {
		return nil, err
	}

	return buf, nil
}

func (h *templateHandler) CreateConfiguration(data map[string]string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	if err := h.template.ExecuteTemplate(buf, "configuration.template", data); err != nil {
		return nil, err
	}

	return buf, nil
}

func getFS() fs.FS {
	unwrapFSOnce.Do(func() {
		fileSys, err := fs.Sub(embedFS, "templates")
		if err != nil {
			panic(err)
		}
		unwrappedFS = fileSys
	})
	return unwrappedFS
}
