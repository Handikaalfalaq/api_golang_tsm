package main

import (
	"fmt"
	postgres "tsmweb/database"
	"tsmweb/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Inisialisasi koneksi database
	err := postgres.DatabaseConnection()
	if err != nil {
		fmt.Println("Gagal terhubung ke database")
		return
	}

	e.GET("/tsm/api/v1/buildingData", routes.GetBuildingData)
	e.GET("/tsm/api/v1/visitors", routes.GetVisitors)

	fmt.Println("Server running on localhost:5000")
	e.Logger.Fatal(e.Start(":5000"))
}
