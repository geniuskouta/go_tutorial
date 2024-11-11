package albumRepository

import "tutorial_105/pkg/domain/entity"

type AlbumRepository interface {
	FindAll() []entity.Album
	Create(album entity.Album) error
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
