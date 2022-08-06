package users

import (
	"database/sql"
	"net/http"
	"strconv"
	"v1/repositories/users"
	"v1/service/users"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service *users.UserService
}

func UserMysqlController(mysql *sql.DB) *userController {
	return &userController{
		service: users.UserMysqlService(
			repositories.UserMysqlRepositories(mysql),
		),
	}
}

func (ctr *userController) Get(c echo.Context) error {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	value, err := ctr.service.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, value)
}
