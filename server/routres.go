package server

import (
	"v1/controllers/admins"
	"v1/controllers/users"
	"v1/dals"
)

func Routes() {
	mysql := dals.OpenConectionMysql()
	userController := users.UserMysqlController(mysql)
	adminController := admins.AdminMysqlController(mysql)

	// routers for users
	e.GET("/movies/:id", userController.Get)

	// routes for admins
	e.POST("admin/movies", adminController.Set)
}
