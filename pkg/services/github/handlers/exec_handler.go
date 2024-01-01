package handlers

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type ExecHandler interface {
	RunMakefile(path string, act string) error
}

type execHandler struct {
	ExecHandler
}

func NewExecHandler() ExecHandler {
	return &execHandler{}
}

func (h *execHandler) RunMakefile(path string, act string) error {
	path, err := filepath.Abs(path)
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
