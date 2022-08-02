package models

type Movie struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	Studio         string `json:"studio"`
	Directtor      string `json:"directtor"`
	Departure_date string `json:"departure_date"`
	Sequels        string `json:"sequels"`
	Duracion       int    `json:"duration"`
}
