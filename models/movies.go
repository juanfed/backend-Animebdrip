package models

type Movie struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Studio        string `json:"studio"`
	Director      string `json:"director"`
	DepartureDate string `json:"departure_date"`
	Sequels       string `json:"sequels"`
	Duration      int    `json:"duration"`
}
