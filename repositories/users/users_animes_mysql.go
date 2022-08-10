package repositories

import (
	"database/sql"
	"fmt"
	"v1/models"
)

type UserAnimeRepositories struct {
	database *sql.DB
}

func UserAnimeMysqlRepositories(mysql *sql.DB) *UserAnimeRepositories {
	return &UserAnimeRepositories{
		database: mysql,
	}
}

func (sq *UserAnimeRepositories) Get(name string) (models.Anime, error) {
	value, err := sq.database.Query(
		fmt.Sprintf(
			`select id, name, seasons, chapters, specials, movies, gender, start_date, end_date, state, duration, audio, subtitle, password, manga, sipnosis from animes where name like "%s"`,
			name,
		),
	)
	if err != nil {
		return models.Anime{}, err
	}

	anime := models.Anime{}
	if value.Next() {
		if err != nil {
			return models.Anime{}, err
		}
		err = value.Scan(&anime.Id, &anime.Name, &anime.Seasons, &anime.Chapters, &anime.Specials, &anime.Movies, &anime.Gender, &anime.StartDate, &anime.EndDate, &anime.State, &anime.Duration, &anime.Audio, &anime.Subtitle, &anime.Password, &anime.Manga, &anime.Sipnosis)
		if err != nil {
			return models.Anime{}, err
		}
	}

	return anime, err
}
