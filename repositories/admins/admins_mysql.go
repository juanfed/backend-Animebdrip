package admins

import (
	"database/sql"
	"fmt"
	"v1/dals"
	"v1/models"
)

type AdminRepositories struct {
	database *sql.DB
}

func AdminMysqlRepositories(mysql *sql.DB) *AdminRepositories {
	return &AdminRepositories{
		database: dals.OpenConectionMysql(),
	}
}

func (sq *AdminRepositories) Set(movie models.Movie) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(`insert into movies (name, gender, studio, directtor, departure_date, sequels, duration) values (%s, %s, %s, %s, %s, %s, %d)`,
			movie.Name,
			movie.Gender,
			movie.Studio,
			movie.Directtor,
			movie.Departure_date,
			movie.Sequels,
			movie.Duracion,
		),
	)
	if err != nil {
		return err
	}

	return err
}

func (sq *AdminRepositories) Update(movie models.Movie) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(
			``,
		),
	)
	if err != nil {
		return err
	}

	return err
}
