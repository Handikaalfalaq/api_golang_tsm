package routes

import (
	postgres "tsmweb/database"
	"tsmweb/handlers"
	"tsmweb/repositories"

	"github.com/labstack/echo/v4"
)

func VisitorsRoutes(e *echo.Group) {
	VisitorsRepository := repositories.RepositoryVisitors(postgres.DB)
	h := handlers.HandlerVisitors(VisitorsRepository)

	e.GET("/visitors", h.GetVisitors)
	e.POST("/visitors", h.CreateNewVisitors)
}
