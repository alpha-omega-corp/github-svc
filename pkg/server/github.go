package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alpha-omega-corp/github-svc/pkg/services/github"
	pkgTypes "github.com/alpha-omega-corp/github-svc/pkg/types"
	proto "github.com/alpha-omega-corp/github-svc/proto/github"
	"github.com/alpha-omega-corp/services/types"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type GithubServer struct {
	proto.UnimplementedGithubServiceServer

	handler github.Handler
}

func NewGithubServer(env *types.Environment) *GithubServer {
	return &GithubServer{
		handler: github.NewHandler(env.Config),
	}
}

func (s *GithubServer) GetSecretContent(ctx context.Context, req *proto.GetSecretContentRequest) (*proto.GetSecretContentResponse, error) {
	res, err := s.handler.Exec().KvsGet(ctx, strings.ToLower(req.Name))
	if err != nil {
		return nil, err
	}

	return &proto.GetSecretContentResponse{
		Content: res.Kvs[0].Value,
	}, nil
}

func (s *GithubServer) SyncEnvironment(ctx context.Context, req *proto.SyncEnvironmentRequest) (*proto.SyncEnvironmentResponse, error) {
	secrets, err := s.handler.Secrets().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	for _, secret := range secrets {
		key := strings.ToLower(secret.Name)
		res, err := s.handler.Exec().KvsGet(ctx, key)
		if err != nil {
			return nil, err
		}

		if len(res.Kvs) != 0 {
			contentString := string(res.Kvs[0].Value)
			inline := strings.Replace(contentString, "\n", "", -1)
			inline = strings.Replace(inline, ",}", " }, ", -1)

			env[secret.Name] = inline
		}

	}

	buf, err := s.handler.Templates().CreateConfiguration(env)
	if err != nil {
		return nil, err
	}

	if err := s.handler.Exec().WriteConfig(buf); err != nil {
		return nil, err
	}

	return &proto.SyncEnvironmentResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *GithubServer) GetSecrets(ctx context.Context, req *proto.GetSecretsRequest) (*proto.GetSecretsResponse, error) {
	secrets, err := s.handler.Secrets().GetAll(ctx)
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
	if err := s.handler.Exec().KvsPut(ctx, req.Name, string(req.Content)); err != nil {
		return nil, err
	}

	if err := s.handler.Secrets().Create(ctx, req.Name, req.Content); err != nil {
		return nil, err
	}

	return &proto.CreateSecretResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *GithubServer) DeleteSecret(ctx context.Context, req *proto.DeleteSecretRequest) (*proto.DeleteSecretResponse, error) {
	if err := s.handler.Exec().KvsDelete(ctx, strings.ToLower(req.Name)); err != nil {
		return nil, err
	}

	if err := s.handler.Secrets().Delete(ctx, req.Name); err != nil {
		return nil, err
	}

	return &proto.DeleteSecretResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *GithubServer) GetPackageTags(ctx context.Context, req *proto.GetPackageTagsRequest) (*proto.GetPackageTagsResponse, error) {
	res, err := s.handler.Packages().GetVersions(req.Name)
	if err != nil {
		return nil, err
	}

	var resSlice []string
	for _, version := range res {
		resSlice = append(resSlice, version.Metadata.Container.Tags...)
	}

	return &proto.GetPackageTagsResponse{
		Tags: resSlice,
	}, nil
}

func (s *GithubServer) PushPackage(ctx context.Context, req *proto.PushPackageRequest) (*proto.PushPackageResponse, error) {
	buf, err := s.handler.Templates().CreateMakefile(req.Name, req.Tag)
	if err != nil {
		return nil, err
	}

	path := req.Name + "/" + req.Tag + "/Dockerfile"
	content, err := s.handler.Repositories().GetContents(ctx, repository, path)
	if err != nil {
		return nil, err
	}

	dockerfile, cErr := content.File.GetContent()
	if err != nil {
		return nil, cErr
	}

	dir, err := os.MkdirTemp("/tmp", req.VersionSHA)
	if err != nil {
		return nil, err
	}

	defer func(path string) {
		rmErr := os.RemoveAll(path)
		if rmErr != nil {
			panic(rmErr)
		}
	}(dir)

	tmpMakefile := filepath.Join(dir, "Makefile")
	if err := os.WriteFile(tmpMakefile, buf.Bytes(), 0644); err != nil {
		return nil, err
	}

	tmpDockerfile := filepath.Join(dir, "Dockerfile")
	if err := os.WriteFile(tmpDockerfile, []byte(dockerfile), 0644); err != nil {
		return nil, err
	}

	if err := s.handler.Packages().Push(dir); err != nil {
		return nil, err
	}

	return &proto.PushPackageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *GithubServer) CreatePackage(ctx context.Context, req *proto.CreatePackageRequest) (*proto.CreatePackageResponse, error) {
	path := req.Name + "/.gitkeep"

	if err := s.handler.Repositories().PutContents(ctx, repository, path, []byte("."), nil); err != nil {
		return nil, err
	}

	return &proto.CreatePackageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *GithubServer) DeletePackage(ctx context.Context, req *proto.DeletePackageRequest) (*proto.DeletePackageResponse, error) {
	content, err := s.handler.Repositories().GetContents(ctx, repository, req.Name)
	if err != nil {
		return nil, err
	}

	for _, item := range content.Dir {
		if len(content.Dir) == 1 {
			if err := s.handler.Repositories().DeleteContents(ctx, repository, req.Name+"/.gitkeep", *item.SHA); err != nil {
				return nil, err
			}

			return &proto.DeletePackageResponse{
				Status: http.StatusOK,
			}, nil
		}

		path := req.Name + "/" + *item.Name
		files, err := s.handler.Repositories().GetPackageFiles(ctx, path)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if err := s.handler.Repositories().DeleteContents(ctx, repository, path+"/"+file.Name, file.SHA); err != nil {
				return nil, err
			}
		}
	}

	return &proto.DeletePackageResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *GithubServer) CreatePackageVersion(ctx context.Context, req *proto.CreatePackageVersionRequest) (*proto.CreatePackageVersionResponse, error) {
	path := req.Name + "/" + req.Tag + "/Dockerfile"
	file, err := s.handler.Templates().CreateDockerfile(req.Name, req.Tag, req.Content)
	if err != nil {
		return nil, err
	}

	fmt.Print(string(file.Bytes()))

	if err := s.handler.Repositories().PutContents(ctx, repository, path, file.Bytes(), nil); err != nil {
		return nil, err
	}

	return &proto.CreatePackageVersionResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *GithubServer) DeletePackageVersion(ctx context.Context, req *proto.DeletePackageVersionRequest) (*proto.DeletePackageVersionResponse, error) {
	if err := s.handler.Packages().Delete(req.Name, req.Version); err != nil {
		return nil, err
	}

	path := req.Name + "/" + req.Tag
	files, err := s.handler.Repositories().GetPackageFiles(ctx, path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if err := s.handler.Repositories().DeleteContents(ctx, repository, path+"/"+file.Name, file.SHA); err != nil {
			return nil, err
		}
	}

	return &proto.DeletePackageVersionResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *GithubServer) GetPackages(ctx context.Context, req *proto.GetPackagesRequest) (*proto.GetPackagesResponse, error) {
	c, err := s.handler.Repositories().GetContents(ctx, "container-images", ".")
	if err != nil {
		return nil, err
	}

	resSlice := make([]*proto.SimplePackage, len(c.Dir))
	for index, pkg := range c.Dir {
		b, err := json.Marshal(pkg)
		if err != nil {
			return nil, err
		}

		if mErr := json.Unmarshal(b, &resSlice[index]); mErr != nil {
			return nil, mErr
		}
	}

	return &proto.GetPackagesResponse{
		Packages: resSlice,
	}, nil
}

func (s *GithubServer) GetPackage(ctx context.Context, req *proto.GetPackageRequest) (*proto.GetPackageResponse, error) {
	c, err := s.handler.Repositories().GetContents(ctx, "container-images", req.Name)
	versions, err := s.handler.Packages().GetVersions(req.Name)
	if err != nil {
		return nil, err
	}

	versionMap := make(map[string]pkgTypes.GitPackageVersion)
	var versionSlice []*proto.PackageVersion
	for _, version := range versions {
		for _, tag := range version.Metadata.Container.Tags {
			versionMap[tag] = version
		}
	}

	for _, dir := range c.Dir {
		if *dir.Type == "dir" {
			v := versionMap[*dir.Name]
			pkg := &proto.PackageVersion{
				RepoName:    *dir.Name,
				RepoPath:    *dir.Path,
				RepoSha:     *dir.SHA,
				RepoLink:    *dir.HTMLURL,
				VersionId:   &v.Id,
				VersionSha:  &v.Name,
				VersionLink: &v.PackageHtmlUrl,
			}

			versionSlice = append(versionSlice, pkg)
		}
	}

	return &proto.GetPackageResponse{
		Versions: versionSlice,
	}, nil
}

func (s *GithubServer) GetPackageFile(ctx context.Context, req *proto.GetPackageFileRequest) (*proto.GetPackageFileResponse, error) {
	c, err := s.handler.Repositories().GetContents(ctx, "container-images", req.Path+"/"+req.Name)
	if err != nil {
		return nil, err
	}

	file, err := c.File.GetContent()
	if err != nil {
		return nil, err
	}

	return &proto.GetPackageFileResponse{
		Content: []byte(file),
	}, nil
}
