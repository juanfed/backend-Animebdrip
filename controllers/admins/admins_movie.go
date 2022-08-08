package admins

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"v1/models"
	repositories "v1/repositories/admins"
	"v1/service/admins"
)

type adminController struct {
	service *admins.AdminService
}

func AdminMysqlController(mysql *sql.DB) *adminController {
	return &adminController{
		service: admins.AdminMysqlService(
			repositories.AdminMysqlRepositories(mysql),
		),
	}
}

func (ctr *adminController) Set(c echo.Context) error {
	movie := models.Movie{}

	err := c.Bind(&movie)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ctr.service.Set(movie))
}
