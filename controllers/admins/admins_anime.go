package admins

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"v1/models"
	repositories "v1/repositories/admins"
	"v1/service/admins"
)

type adminAnimeController struct {
	service *admins.AdminAnimeService
}

func AdminAnimeMysqlController(mysql *sql.DB) *adminAnimeController {
	return &adminAnimeController{
		service: admins.AdminAnimeMysqlService(
			repositories.AdminAnimeMysqlRepositories(mysql),
		),
	}
}

func (ctr *adminAnimeController) Set(c echo.Context) error {
	anime := models.Anime{}

	err := c.Bind(&anime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ctr.service.Set(anime))
}
