package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"v1/repositories"
	"v1/service"

	"github.com/labstack/echo/v4"
)

type usercController struct {
	service *service.UserService
}

func UserMysqlController(mysql *sql.DB) *usercController {
	return &usercController{
		service: service.UserMysqlService(
			repositories.UserMysqlRepositories(mysql),
		),
	}
}

func (ctr *usercController) Get(c echo.Context) error {
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
