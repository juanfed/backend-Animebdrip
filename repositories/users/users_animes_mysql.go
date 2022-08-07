package repositories

import "database/sql"

type UserAnimeRepositories struct {
	database *sql.DB
}

func UserAnimeMysqlRepositories(mysql *sql.DB) *UserAnimeRepositories {
	return &UserAnimeRepositories{
		database: mysql,
	}
}

func (sq *UserAnimeRepositories) GetAnime(name string) {

}