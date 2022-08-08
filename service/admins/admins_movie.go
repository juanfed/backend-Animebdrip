package admins

import (
	"v1/models"
	repositories "v1/repositories/admins"
)

type AdminService struct {
	mysql *repositories.AdminRepositories
}

func AdminMysqlService(mysql *repositories.AdminRepositories) *AdminService {
	return &AdminService{
		mysql: mysql,
	}
}

func (s *AdminService) Set(movie models.Movie) error {
	return s.mysql.Set(movie)
}

func (s *AdminService) Update(movie models.Movie) error {
	return s.mysql.Update(movie)
}

func (s *AdminService) Delete(id int) error {
	return s.mysql.Delete(id)
}
