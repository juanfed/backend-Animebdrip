package server

import (
	"v1/controllers"
	"v1/dals"
)

func Routes() {
	mysql := dals.OpenConectionMysql()
	userController := controllers.UserMysqlController(mysql)

	e.GET("/movies/:id", userController.Get)
}
