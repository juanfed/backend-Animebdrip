package server

import "github.com/labstack/echo/v4"

var e = echo.New()

func Start() {
	Routes()
	e.Logger.Fatal(e.Start(":8080"))
}
