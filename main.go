package main

import (
	"fmt"
	postgres "tsmweb/database"
	"tsmweb/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	postgres.DatabaseConnection()

	routes.RouteInit(e.Group("tsm/api/v1"))

	fmt.Println("Server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
