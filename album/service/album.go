package service

import (
	"context"

	"github.com/iammonalisaganguly/Golang-Microservises-Architecture/album/domain"
	"github.com/iammonalisaganguly/Golang-Microservises-Architecture/album/repository"
)

type AlbumServiceContract interface {
	AddAlbum(ctx context.Context, a domain.Album) (err error)

	GetAlbum(ctx context.Context, id string) (err error, a domain.Album)
}

type service struct {
	repoContract repository.AlbumRepoContract
}

func NewAlbumService(repoContract repository.AlbumRepoContract) *service {
	return &service{repoContract: repoContract}
}
func (svc *service) AddAlbum(ctx context.Context, alb domain.Album) error {
	err := svc.repoContract.AddAlbum(ctx, alb)
	return err
}

func (svc *service) GetAlbum(ctx context.Context, id string) (error, domain.Album) {
	err, album := svc.repoContract.GetAlbum(ctx, id)
	return err, album
}
