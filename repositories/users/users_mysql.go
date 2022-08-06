package repositories

import (
	"database/sql"
	"fmt"
	"v1/dals"
	"v1/models"

	// para la conexion con el driver de mysql
	_ "github.com/go-sql-driver/mysql"
)

type UserRepositories struct {
	database *sql.DB
}

func UserMysqlRepositories(mysql *sql.DB) *UserRepositories {
	return &UserRepositories{
		database: dals.OpenConectionMysql(),
	}
}

// CRUD
func (sq *UserRepositories) Get(id int) (models.Movie, error) {
	fmt.Println("Into get for consult mysql")
	value, err := sq.database.Query(
		fmt.Sprintf(
			`select id, name, gender, studio, directtor, departure_date, sequels, duration from movies where id=%d`,
			id,
		),
	)
	if err != nil {
		return models.Movie{}, err
	}

	defer value.Close()
	movie := models.Movie{}
	if value.Next() {
		if err != nil {
			fmt.Println(err)
			return models.Movie{}, err
		}
		err = value.Scan(&movie.Id, &movie.Name, &movie.Gender, &movie.Studio, &movie.Directtor, &movie.Departure_date, &movie.Sequels, &movie.Duracion)
		if err != nil {
			return models.Movie{}, err
		}
	}

	return movie, err
}
