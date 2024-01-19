package main

import (
	"github.com/alpha-omega-corp/github-svc/pkg/server"
	protoDocker "github.com/alpha-omega-corp/github-svc/proto/docker"
	protoGithub "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/alpha-omega-corp/services/config"
	svc "github.com/alpha-omega-corp/services/server"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {
	cManager := config.NewHandler()
	cHost, err := cManager.Manager().Hosts()
	if err != nil {
		panic(err)
	}

	if err := svc.NewGRPC(cHost.Github.Host, nil, func(_ *bun.DB, grpc *grpc.Server) {
		c, cErr := cManager.Manager().GithubService()
		if cErr != nil {
			panic(cErr)
		}

		githubServer := server.NewGithubServer(c)
		dockerServer := server.NewDockerServer(c)
		protoGithub.RegisterGithubServiceServer(grpc, githubServer)
		protoDocker.RegisterDockerServiceServer(grpc, dockerServer)
	}); err != nil {
		panic(err)
	}

}
