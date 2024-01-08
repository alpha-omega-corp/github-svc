package server

import (
	"context"
	"github.com/alpha-omega-corp/github-svc/pkg/services/github"
	proto "github.com/alpha-omega-corp/github-svc/proto/github"
	svc "github.com/alpha-omega-corp/services/server"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type GithubServer struct {
	proto.UnimplementedGithubServiceServer

	gitHandler github.Handler
}

func NewGithubServer() *GithubServer {
	v := viper.New()
	cManager := svc.NewConfigManager(v)
	config, err := cManager.GithubConfig()
	if err != nil {
		panic(err)
	}

	return &GithubServer{
		gitHandler: github.NewHandler(config),
	}
}

func (s *GithubServer) GetSecrets(ctx context.Context, req *proto.GetSecretsRequest) (*proto.GetSecretsResponse, error) {
	secrets, err := s.gitHandler.Secrets().GetSecrets(ctx)
	if err != nil {
		return nil, err
	}

	resSlice := make([]*proto.Secret, len(secrets))
	for index, secret := range secrets {
		resSlice[index] = &proto.Secret{
			Name:       secret.Name,
			CreatedAt:  secret.CreatedAt.GoString(),
			UpdatedAt:  secret.UpdatedAt.GoString(),
			Visibility: secret.Visibility,
		}
	}

	return &proto.GetSecretsResponse{
		Secrets: resSlice,
	}, nil
}
