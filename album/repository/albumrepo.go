package repository

import (
	"context"
	"fmt"

	"github.com/iammonalisaganguly/Golang-Microservises-Architecture/album/domain"
)

type AlbumRepoContract interface {
	AddAlbum(ctx context.Context, a domain.Album) (err error)

	GetAlbum(ctx context.Context, id string) (err error, a domain.Album)
}

type repoStruct struct {
}
func NewAlbumRepo()*repoStruct{
	return &repoStruct{}
}
func (r *repoStruct) AddAlbum(ctx context.Context, a domain.Album) (err error) {
	fmt.Println(a)
	return
}

func (r *repoStruct) GetAlbum(ctx context.Context, id string) (err error, a domain.Album) {
	fmt.Println(id)
	return
}
