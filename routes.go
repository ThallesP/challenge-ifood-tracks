package main

import (
	"challenge-ifood/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/tracks", controllers.HandlerTracksShow)
}