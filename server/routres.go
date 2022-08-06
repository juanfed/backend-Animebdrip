package server

import (
	"v1/controllers"
	"v1/dals"
)

func Routes() {
	mysql := dals.OpenConectionMysql()
	userController := controllers.UserMysqlController(mysql)

	adminController

	e.GET("/movies/:id", userController.Get)

	e.SET("admin/movies", adminController.Set)
}
