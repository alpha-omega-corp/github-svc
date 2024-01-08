package server

import (
	"context"
	"fmt"
	"github.com/alpha-omega-corp/github-svc/pkg/services/github"
	proto "github.com/alpha-omega-corp/github-svc/proto/github"
	svc "github.com/alpha-omega-corp/services/server"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"net/http"
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
	secrets, err := s.gitHandler.Secrets().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resSlice := make([]*proto.Secret, len(secrets))
	for index, secret := range secrets {
		resSlice[index] = &proto.Secret{
			Name:       secret.Name,
			CreatedAt:  secret.CreatedAt.String(),
			UpdatedAt:  secret.UpdatedAt.String(),
			Visibility: secret.Visibility,
		}
	}

	return &proto.GetSecretsResponse{
		Secrets: resSlice,
	}, nil
}

func (s *GithubServer) CreateSecret(ctx context.Context, req *proto.CreateSecretRequest) (*proto.CreateSecretResponse, error) {
	if err := s.gitHandler.Secrets().Create(ctx, req.Name, req.Content); err != nil {
		return nil, err
	}

	buffer, err := s.gitHandler.Templates().CreateConfiguration(req.Name, req.Content)
	if err != nil {
		return nil, err
	}

	if err := s.gitHandler.Exec().WriteConfig(buffer); err != nil {
		return nil, err
	}

	fmt.Print()

	return &proto.CreateSecretResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *GithubServer) DeleteSecret(ctx context.Context, req *proto.DeleteSecretRequest) (*proto.DeleteSecretResponse, error) {
	if err := s.gitHandler.Secrets().Delete(ctx, req.Name); err != nil {
		return nil, err
	}

	return &proto.DeleteSecretResponse{
		Status: http.StatusOK,
	}, nil
}
