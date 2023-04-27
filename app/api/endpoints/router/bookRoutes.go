package router

import (
	"github.com/labstack/echo/v4"
	"treinamento/app/api/endpoints/dicontainer"
)

func loadBookRoutes(api *echo.Group) {
	bookGroup := api.Group("/book")

	bookHandlers := dicontainer.GetBookHandlers()

	bookGroup.POST("/new", bookHandlers.PostBook)
}
