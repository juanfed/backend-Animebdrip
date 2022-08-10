package admins

import (
	"v1/models"
	repositories "v1/repositories/admins"
)

type AdminAnimeService struct {
	mysql *repositories.AdminAnimeRepositories
}

func AdminAnimeMysqlService(mysql *repositories.AdminAnimeRepositories) *AdminAnimeService {
	return &AdminAnimeService{
		mysql: mysql,
	}
}

func (s *AdminAnimeService) Set(anime models.Anime) error {
	return s.mysql.Set(anime)
}
