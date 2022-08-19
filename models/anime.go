package models

type Anime struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Seasons   int    `json:"seasons"`
	Chapters  int    `json:"chapters"`
	Specials  string `json:"specials"`
	Movies    int    `json:"movies"`
	Gender    string `json:"gender"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	State     string `json:"state"`
	Duration  string `json:"duration"`
	Audio     string `json:"audio"`
	Subtitle  string `json:"subtitle"`
	Password  string `json:"password"`
	Manga     string `json:"manga"`
	Sipnosis  string `json:"sipnosis"`
	Photos    string `json:"photos"`
}
