package repositories

import (
	"database/sql"
	"fmt"
	"v1/models"
)

type AdminAnimeRepositories struct {
	database *sql.DB
}

func AdminAnimeMysqlRepositories(mysql *sql.DB) *AdminAnimeRepositories {
	return &AdminAnimeRepositories{
		database: mysql,
	}
}

func (sq AdminAnimeRepositories) Set(anime models.Anime) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(
			`insert into animes (name, seasons, chapters, specials, movies, gender, start_date, end_date, state, duration, audio, subtitle, password, manga, sipnosis)
					values('%s', %d, %d, '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`,
			anime.Name,
			anime.Seasons,
			anime.Chapters,
			anime.Specials,
			anime.Movies,
			anime.Gender,
			anime.StartDate,
			anime.EndDate,
			anime.State,
			anime.Duration,
			anime.Audio,
			anime.Subtitle,
			anime.Password,
			anime.Manga,
			anime.Sipnosis,
		),
	)
	if err != nil {
		return err
	}

	return err
}
