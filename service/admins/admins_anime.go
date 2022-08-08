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

func (s *AdminAnimeService) SetAnime(anime models.Anime) error {
	return s.mysql.SetAnime(anime)
}
