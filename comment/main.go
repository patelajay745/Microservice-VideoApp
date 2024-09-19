package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patelajay745/Microservice-VideoApp/comment/config"
	"github.com/patelajay745/Microservice-VideoApp/comment/routes"
)

func main() {

	config.ConnectDb()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.SetUpRouter(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
