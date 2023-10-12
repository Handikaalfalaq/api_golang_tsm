package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	resultdto "tsmweb/dto/result"
	visitorsdto "tsmweb/dto/visitors"
	"tsmweb/models"
	"tsmweb/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerVisitors struct {
	VisitorsRepository repositories.VisitorsRepository
}

func HandlerVisitors(VisitorsRepository repositories.VisitorsRepository) *handlerVisitors {
	return &handlerVisitors{VisitorsRepository}
}

func (h *handlerVisitors) GetVisitors(c echo.Context) error {
	Visitors, err := h.VisitorsRepository.FindVisitors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data dari database"})
	}

	var formattedVisitors []models.Visitors
	for _, v := range Visitors {
		tanggalLahirStr := strings.Split(v.TglLahir, "T")[0]
		tanggalLahir, err := time.Parse("2006-01-02", tanggalLahirStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengonversi tanggal lahir: " + err.Error()})
		}
		tanggalLahirFormatted := tanggalLahir.Format("02-01-06")

		suhuFloat, _ := strconv.ParseFloat(v.Suhu, 64)
		suhuFormatted := strconv.FormatFloat(suhuFloat, 'f', -1, 64)

		jamMasukTime, err := time.Parse("15:04:05", v.JamMasuk)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengonversi jam masuk: " + err.Error()})
		}
		jamMasukFormatted := jamMasukTime.Format("15:04")

		jamKeluarTime, err := time.Parse("15:04:05", v.JamKeluar)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengonversi jam Keluar: " + err.Error()})
		}
		jamKeluarFormatted := jamKeluarTime.Format("15:04")

		visitor := models.Visitors{
			ID:          v.ID,
			Nama:        v.Nama,
			NoKTP:       v.NoKTP,
			TglLahir:    tanggalLahirFormatted,
			NamaGedung:  v.NamaGedung,
			Suhu:        suhuFormatted,
			JamMasuk:    jamMasukFormatted,
			JamKeluar:   jamKeluarFormatted,
			PhoneNumber: v.PhoneNumber,
		}
		formattedVisitors = append(formattedVisitors, visitor)
	}

	sort.Slice(formattedVisitors, func(i, j int) bool {
		return formattedVisitors[i].ID < formattedVisitors[j].ID
	})

	return c.JSON(http.StatusOK, formattedVisitors)
}

func (h *handlerVisitors) CreateNewVisitors(c echo.Context) error {
	request := new(visitorsdto.VisitorsRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	Visitors := models.Visitors{
		Nama:        c.FormValue("nama"),
		NoKTP:       c.FormValue("noKTP"),
		TglLahir:    c.FormValue("tglLahir"),
		NamaGedung:  c.FormValue("namaGedung"),
		Suhu:        c.FormValue("suhu"),
		JamMasuk:    c.FormValue("jamMasuk"),
		JamKeluar:   c.FormValue("jamKeluar"),
		PhoneNumber: c.FormValue("phoneNumber"),
	}

	data, err := h.VisitorsRepository.CreateVisitors(Visitors)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseVisitors(data)})
}

func convertResponseVisitors(u models.Visitors) visitorsdto.VisitorsResponse {
	return visitorsdto.VisitorsResponse{
		ID:          u.ID,
		Nama:        u.Nama,
		NoKTP:       u.NoKTP,
		TglLahir:    u.TglLahir,
		NamaGedung:  u.NamaGedung,
		Suhu:        u.Suhu,
		JamMasuk:    u.JamMasuk,
		JamKeluar:   u.JamKeluar,
		PhoneNumber: u.PhoneNumber,
	}
}
