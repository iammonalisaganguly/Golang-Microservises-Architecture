package main

import (
	"context"
	"fmt"
	"time"

	"github.com/iammonalisaganguly/Golang-Microservises-Architecture/album/domain"
	"github.com/iammonalisaganguly/Golang-Microservises-Architecture/album/repository"
	"github.com/iammonalisaganguly/Golang-Microservises-Architecture/album/service"
)

func main() {
	repo := repository.NewAlbumRepo()
	svc := service.NewAlbumService(repo)
	al := domain.Album{
		ID:          "1",
		Name:        "Monalisa Ganguly",
		ReleaseDate: time.Now(),
		Producer:    "Monalisa Ganguly",
	}
	svc.AddAlbum(context.Background(), al)
	fmt.Println("Done")
}
