package main

import (
	"github.com/alpha-omega-corp/github-svc/pkg/server"
	protoDocker "github.com/alpha-omega-corp/github-svc/proto/docker"
	protoGithub "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/alpha-omega-corp/services/database"
	s "github.com/alpha-omega-corp/services/server"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {

	v := viper.New()
	cManager := s.NewConfigManager(v)
	cHost, err := cManager.HostsConfig()
	if err != nil {
		panic(err)
	}

	dbHandler := database.NewHandler(cHost.Docker.Dsn)

	if err := s.NewGRPC(cHost.Docker.Host, dbHandler, func(db *bun.DB, grpc *grpc.Server) {
		githubServer := server.NewGithubServer()
		dockerServer := server.NewDockerServer(db)

		protoGithub.RegisterGithubServiceServer(grpc, githubServer)
		protoDocker.RegisterDockerServiceServer(grpc, dockerServer)
	}); err != nil {
		panic(err)
	}

}
