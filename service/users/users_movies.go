package users

import (
	"v1/models"
	"v1/repositories/users"
)

type UserService struct {
	mysql *repositories.UserMovieRepositories
}

func UserMysqlService(mysql *repositories.UserMovieRepositories) *UserService {
	return &UserService{
		mysql: mysql,
	}
}

func (s *UserService) Get(id int) (models.Movie, error) {
	return s.mysql.Get(id)
}
