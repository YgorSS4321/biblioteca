package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() *echo.Echo {
	router := echo.New()

	loadMiddleware(router)

	return router
}

func loadMiddleware(router *echo.Echo) {
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
}
