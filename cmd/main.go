package main

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"github.com/alpha-omega-corp/docker-svc/pkg/services"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/database"
	"github.com/alpha-omega-corp/services/server"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {

	v := viper.New()
	cManager := server.NewConfigManager(v)
	c, err := cManager.HostsConfig()
	if err != nil {
		panic(err)
	}

	dbHandler := database.NewHandler(c.Docker.Dsn)

	if err := server.NewGRPC(c.Docker.Host, dbHandler, func(db *bun.DB, grpc *grpc.Server) {
		s := services.NewServer(db)
		proto.RegisterDockerServiceServer(grpc, s)
	}); err != nil {
		panic(err)
	}

}
