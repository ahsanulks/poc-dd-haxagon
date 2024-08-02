package server

import (
	"os"
	"poc/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	route.NewRoute(e)
	e.Start(":" + os.Getenv("APP_PORT"))
}
