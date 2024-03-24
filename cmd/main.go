package main

import (
	"github.com/alpha-omega-corp/github-svc/pkg/server"
	protoDocker "github.com/alpha-omega-corp/github-svc/proto/docker"
	protoGithub "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/alpha-omega-corp/services/config"
	svc "github.com/alpha-omega-corp/services/server"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {
	env, err := config.NewHandler().Environment("github")
	if err != nil {
		panic(err)
	}

	if err := svc.NewGRPC(env.Host.Url, nil, func(_ *bun.DB, grpc *grpc.Server) {
		protoGithub.RegisterGithubServiceServer(grpc, server.NewGithubServer(env))
		protoDocker.RegisterDockerServiceServer(grpc, server.NewDockerServer(env))
	}); err != nil {
		panic(err)
	}

}
