package service

import (
	"v1/models"
	"v1/repositories"
)

type UserService struct {
	mysql *repositories.UserRepositories
}

func UserMysqlService(mysql *repositories.UserRepositories) *UserService {
	return &UserService{
		mysql: mysql,
	}
}

func (s *UserService) Get() (models.Movie, error) {
	return s.mysql.Get()
}
