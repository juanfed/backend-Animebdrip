package server

import (
	"v1/controllers/admins"
	"v1/controllers/users"
	"v1/dals"
)

func Routes() {
	mysql := dals.OpenConnectionMysql()
	userController := users.UserMysqlController(mysql)
	adminController := admins.AdminMysqlController(mysql)
	userAnimeController := users.UserAnimeMysqlController(mysql)

	// routers for users movies
	e.GET("/movies/:id", userController.Get)

	// routers for users anime's
	e.GET("animeBDRip/listAnime/another/:name", userAnimeController.GetAnime)

	// routes for admins
	e.POST("admin/movies", adminController.Set)
}
