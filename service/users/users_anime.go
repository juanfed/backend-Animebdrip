package users

import (
	"v1/models"
	repositories "v1/repositories/users"
)

type UserAnimeService struct {
	mysql *repositories.UserAnimeRepositories
}

func UserAnimeMysqlService(mysql *repositories.UserAnimeRepositories) *UserAnimeService {
	return &UserAnimeService{
		mysql: mysql,
	}
}

func (s *UserAnimeService) GetAnime(name string) (models.Anime, error) {
	return s.mysql.Get(name)
}

func (s *UserAnimeService) GetPhotos(id int) (models.Photo, error) {
	return s.mysql.GetPhotos(id)
}
