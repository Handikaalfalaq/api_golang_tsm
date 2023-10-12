package routes

import (
	postgres "tsmweb/database"
	"tsmweb/handlers"
	"tsmweb/repositories"

	"github.com/labstack/echo/v4"
)

func BuildingDataRoutes(e *echo.Group) {
	BuildingDataRepository := repositories.RepositoryBuildingData(postgres.DB)
	h := handlers.HandlerBuildingData(BuildingDataRepository)

	e.GET("/buildingData", h.GetBuildingData)
	e.POST("/buildingData", h.CreateNewBuildingData)
}
