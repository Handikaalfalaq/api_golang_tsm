package routes

import (
	"net/http"
	"strings"
	"time"
	postgres "tsmweb/database"

	"github.com/labstack/echo/v4"
)

type BuildingData struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	PhoneNumber     string  `json:"phone_number"`
	Address         string  `json:"address"`
	AllowedCapacity int     `json:"allowed_capacity"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
}

type Visitors struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	NoKTP       string `json:"no_ktp"`
	TglLahir    string `json:"tgl_lahir"`
	NamaGedung  string `json:"nama_gedung"`
	Suhu        string `json:"suhu"`
	JamMasuk    string `json:"jam_masuk"`
	JamKeluar   string `json:"jam_keluar"`
	PhoneNumber string `json:"phone_number"`
}

func GetBuildingData(c echo.Context) error {
	var buildingData []BuildingData

	result := postgres.DB.Find(&buildingData)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data dari database"})
	}

	return c.JSON(http.StatusOK, buildingData)
}

func GetVisitors(c echo.Context) error {
	var visitorsFromDB []Visitors

	result := postgres.DB.Find(&visitorsFromDB)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data dari database"})
	}

	var visitors []Visitors
	for _, v := range visitorsFromDB {
		tanggalLahirStr := strings.Split(v.TglLahir, "T")[0]

		tanggalLahir, err := time.Parse("2006-01-02", tanggalLahirStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengonversi tanggal lahir: " + err.Error()})
		}

		tanggalLahirFormatted := tanggalLahir.Format("02-01-06")

		visitor := Visitors{
			ID:          v.ID,
			Nama:        v.Nama,
			NoKTP:       v.NoKTP,
			TglLahir:    tanggalLahirFormatted,
			NamaGedung:  v.NamaGedung,
			Suhu:        v.Suhu,
			JamMasuk:    v.JamMasuk,
			JamKeluar:   v.JamKeluar,
			PhoneNumber: v.PhoneNumber,
		}
		visitors = append(visitors, visitor)
	}

	return c.JSON(http.StatusOK, visitors)
}
