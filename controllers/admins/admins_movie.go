package admins

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"v1/models"
	repositories "v1/repositories/admins"
	"v1/service/admins"
)

type adminMovieController struct {
	service *admins.AdminMovieService
}

func MovieAdminMysqlController(mysql *sql.DB) *adminMovieController {
	return &adminMovieController{
		service: admins.AdminMovieMysqlService(
			repositories.AdminMovieMysqlRepositories(mysql),
		),
	}
}

func (ctr *adminMovieController) Set(c echo.Context) error {
	movie := models.Movie{}

	err := c.Bind(&movie)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ctr.service.Set(movie))
}
