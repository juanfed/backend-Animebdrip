package admins

import (
	"v1/models"
	repositories "v1/repositories/admins"
)

type AdminMovieService struct {
	mysql *repositories.AdminMovieRepositories
}

func AdminMovieMysqlService(mysql *repositories.AdminMovieRepositories) *AdminMovieService {
	return &AdminMovieService{
		mysql: mysql,
	}
}

func (s *AdminMovieService) Set(movie models.Movie) error {
	return s.mysql.Set(movie)
}

func (s *AdminMovieService) Update(movie models.Movie) error {
	return s.mysql.Update(movie)
}

func (s *AdminMovieService) Delete(id int) error {
	return s.mysql.Delete(id)
}
