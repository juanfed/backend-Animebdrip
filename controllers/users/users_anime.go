package users

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	repositories "v1/repositories/users"
	"v1/service/users"
)

type UserAnimeController struct {
	service *users.UserAnimeService
}

func UserAnimeMysqlController(mysql *sql.DB) *UserAnimeController {
	return &UserAnimeController{
		service: users.UserAnimeMysqlService(
			repositories.UserAnimeMysqlRepositories(mysql),
		),
	}
}

func (ctr *UserAnimeController) GetAnime(c echo.Context) error {
	name := c.Param("name")

	value, err := ctr.service.GetAnime(name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, value)
}

func (ctr *UserAnimeController) GetPhotos(c echo.Context) error {
	param := c.Param("animeId")

	animeId, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	value, err := ctr.service.GetPhotos(animeId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, value)
}
