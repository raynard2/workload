package main

import (
	"Proto/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	e.POST("/clients", controller.Options)

	e.POST("/client", controller.Client2)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
