package models

type Photo struct {
	Id      int    `json:"id"`
	AnimeId int    `json:"anime_id"`
	Name    string `json:"name"`
	Enlaces string `json:"enlaces"`
}
