package handlers

import (
	"net/http"
	"strconv"
	buildingDatadto "tsmweb/dto/buildingData"
	resultdto "tsmweb/dto/result"
	"tsmweb/models"
	"tsmweb/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerBuildingData struct {
	BuildingDataRepository repositories.BuildingDataRepository
}

func HandlerBuildingData(BuildingDataRepository repositories.BuildingDataRepository) *handlerBuildingData {
	return &handlerBuildingData{BuildingDataRepository}
}

func (h *handlerBuildingData) GetBuildingData(c echo.Context) error {
	buildingData, err := h.BuildingDataRepository.FindBuildingData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data dari database"})
	}

	return c.JSON(http.StatusOK, buildingData)
}

func (h *handlerBuildingData) CreateNewBuildingData(c echo.Context) error {
	request := new(buildingDatadto.BuildingDataRequest)

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
	allowedCapacity, err := strconv.Atoi(c.FormValue("allowedCapacity"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid allowedCapacity value",
		})
	}

	longitude, err := strconv.ParseFloat(c.FormValue("longitude"), 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid longitude value",
		})
	}

	latitude, err := strconv.ParseFloat(c.FormValue("latitude"), 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid latitude value",
		})
	}

	BuildingData := models.BuildingData{
		Name:            c.FormValue("name"),
		PhoneNumber:     c.FormValue("phoneNumber"),
		Address:         c.FormValue("address"),
		AllowedCapacity: allowedCapacity,
		Longitude:       longitude,
		Latitude:        latitude,
	}

	data, err := h.BuildingDataRepository.CreateBuildingData(BuildingData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseBuildingData(data)})
}

func convertResponseBuildingData(u models.BuildingData) buildingDatadto.BuildingDataResponse {
	return buildingDatadto.BuildingDataResponse{
		ID:              u.ID,
		Name:            u.Name,
		PhoneNumber:     u.PhoneNumber,
		Address:         u.Address,
		AllowedCapacity: u.AllowedCapacity,
		Longitude:       u.Longitude,
		Latitude:        u.Latitude,
	}
}
