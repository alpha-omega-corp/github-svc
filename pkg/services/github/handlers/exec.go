package handlers

import (
	"bytes"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type ExecHandler interface {
	KvsGet(ctx context.Context, key string) (*clientv3.GetResponse, error)
	KvsPut(ctx context.Context, key string, value string) (err error)
	KvsDelete(ctx context.Context, key string) (err error)
	WriteConfig(template *bytes.Buffer) error
	RunMakefile(path string, act string) error
}

type execHandler struct {
	ExecHandler

	etcdClient *clientv3.Client
}

func NewExecHandler() ExecHandler {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	return &execHandler{
		etcdClient: cli,
	}
}

func (h *execHandler) KvsGet(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	return h.etcdClient.Get(ctx, key)
}

func (h *execHandler) KvsPut(ctx context.Context, key string, value string) (err error) {
	_, err = h.etcdClient.Put(ctx, strings.ToLower(key), value)
	if err != nil {
		return
	}

	return
}

func (h *execHandler) KvsDelete(ctx context.Context, key string) (err error) {
	_, err = h.etcdClient.Delete(ctx, key)
	if err != nil {
		return
	}

	return
}

func (h *execHandler) WriteConfig(template *bytes.Buffer) error {
	return os.WriteFile("/home/nanstis/.config/act/.secrets", template.Bytes(), 0644)
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
