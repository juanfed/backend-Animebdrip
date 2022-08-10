package repositories

import (
	"database/sql"
	"fmt"
	"v1/models"
)

type AdminMovieRepositories struct {
	database *sql.DB
}

func AdminMovieMysqlRepositories(mysql *sql.DB) *AdminMovieRepositories {
	return &AdminMovieRepositories{
		database: mysql,
	}
}

func (sq *AdminMovieRepositories) Set(movie models.Movie) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(`insert into movies (name, gender, studio, directtor, departure_date, sequels, duration) values ('%s', '%s', '%s', '%s', '%s', '%s', %d)`,
			movie.Name,
			movie.Gender,
			movie.Studio,
			movie.Director,
			movie.DepartureDate,
			movie.Sequels,
			movie.Duration,
		),
	)

	return err
}

func (sq *AdminMovieRepositories) Update(movie models.Movie) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(
			`update movies set set name='%s', gender='%s', studio='%s', directtor='%s', departure_date='%s', sequels='%s', duration=%d where id=%d`,
			movie.Name,
			movie.Gender,
			movie.Studio,
			movie.Director,
			movie.DepartureDate,
			movie.Sequels,
			movie.Duration,
			movie.Id,
		),
	)

	return err
}

func (sq *AdminMovieRepositories) Delete(id int) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(`delete from movies where id=%d`,
			id,
		),
	)

	return err
}
