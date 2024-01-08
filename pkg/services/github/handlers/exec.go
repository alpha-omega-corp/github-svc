package handlers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type ExecHandler interface {
	RunMakefile(path string, act string) error
	WriteConfig(template *bytes.Buffer) error
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

	fmt.Print(string(res))

	return nil
}

func (h *execHandler) WriteConfig(template *bytes.Buffer) error {
	f, err := os.OpenFile("/home/nanstis/.config/act/.test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	if _, err := f.WriteString(template.String() + "\n"); err != nil {
		log.Println(err)
	}

	return nil
}
