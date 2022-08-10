package server

import (
	"v1/controllers/admins"
	"v1/controllers/users"
	"v1/dals"
)

func Routes() {
	mysql := dals.OpenConnectionMysql()
	userController := users.UserMysqlController(mysql)
	adminController := admins.MovieAdminMysqlController(mysql)
	adminAnimeController := admins.AdminAnimeMysqlController(mysql)
	userAnimeController := users.UserAnimeMysqlController(mysql)

	// routers for users movies
	e.GET("/movies/:id", userController.Get)

	// routers for users anime's
	e.GET("animeBDRip/listAnime/another/:name", userAnimeController.GetAnime)

	// routes for admins movies
	e.POST("admin/movies", adminController.Set)

	// routers for admins anime's
	e.POST("admin/anime", adminAnimeController.Set)
}
