package main

import (
	"github.com/alpha-omega-corp/services/config"
	"github.com/alpha-omega-corp/services/database"
	"github.com/alpha-omega-corp/services/server"
	"github.com/alpha-omega-org/docker-svc/pkg/proto"
	"github.com/alpha-omega-org/docker-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c := config.Get("dev")

	err := server.NewGRPC(c.AUTH, c, func(h *database.Handler, grpc *grpc.Server) {
		s := services.NewServer(h.Database())
		proto.RegisterDockerServiceServer(grpc, s)
	})

	if err != nil {
		panic(err)
	}
}
