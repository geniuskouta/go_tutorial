package main

import (
	"tutorial_105/pkg/app/handler/albumHandler"
	"tutorial_105/pkg/infra/repository/albumRepository"

	"github.com/gin-gonic/gin"
)

func main() {
	albumRepo := albumRepository.NewAlbumRepository()
	albumHdr := albumHandler.NewAlbumHandler(albumRepo)
	router := gin.Default()
	router.GET("/albums", albumHdr.GetAlbums)
	router.POST("/albums", albumHdr.PostAlbums)
	router.Run("localhost:8080")
}
