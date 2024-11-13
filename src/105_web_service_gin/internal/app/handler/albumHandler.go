package handler

import (
	"net/http"
	"tutorial_105/internal/domain/entity"
	"tutorial_105/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	repo repository.AlbumRepository
}

func NewAlbumHandler(repo repository.AlbumRepository) *AlbumHandler {
	return &AlbumHandler{repo: repo}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums := h.repo.FindAll()
	c.IndentedJSON(http.StatusOK, albums)
}

func (h *AlbumHandler) PostAlbums(c *gin.Context) {
	var newAlbum entity.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.repo.Create(newAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (h *AlbumHandler) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := h.repo.GetById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
