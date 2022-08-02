package main

import (
	// para la conexion con el driver

	"v1/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server.Start()
}
