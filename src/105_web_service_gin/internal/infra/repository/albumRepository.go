package repository

import (
	"fmt"
	"tutorial_105/internal/domain/entity"
)

type AlbumRepository interface {
	FindAll() []entity.Album
	Create(album entity.Album) error
	GetById(id string) (*entity.Album, error)
}

type InMemoryAlbumRepository struct {
	albums []entity.Album
}

func NewAlbumRepository() *InMemoryAlbumRepository {
	return &InMemoryAlbumRepository{
		albums: []entity.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (r *InMemoryAlbumRepository) FindAll() []entity.Album {
	return r.albums
}

func (r *InMemoryAlbumRepository) Create(album entity.Album) error {
	r.albums = append(r.albums, album)
	return nil
}

func (r *InMemoryAlbumRepository) GetById(id string) (*entity.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, fmt.Errorf("album with ID %s not found", id)
}
