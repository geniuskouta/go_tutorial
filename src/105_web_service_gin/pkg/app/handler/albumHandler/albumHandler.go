package albumHandler

import (
	"net/http"
	"tutorial_105/pkg/domain/entity"
	"tutorial_105/pkg/infra/repository/albumRepository"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	repo albumRepository.AlbumRepository
}

func NewAlbumHandler(repo albumRepository.AlbumRepository) *AlbumHandler {
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
