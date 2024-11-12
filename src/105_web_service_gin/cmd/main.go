package main

import (
	"tutorial_105/pkg/app/handler"
	"tutorial_105/pkg/infra/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	albumRepo := repository.NewAlbumRepository()
	albumHdr := handler.NewAlbumHandler(albumRepo)
	router := gin.Default()
	router.GET("/albums", albumHdr.GetAlbums)
	router.POST("/albums", albumHdr.PostAlbums)
	router.GET("/albums/:id", albumHdr.GetAlbumById)
	router.Run("localhost:8080")
}
