package repositories

import (
	"tsmweb/models"

	"gorm.io/gorm"
)

type BuildingDataRepository interface {
	FindBuildingData() ([]models.BuildingData, error)
	CreateBuildingData(buildingData models.BuildingData) (models.BuildingData, error)
}

func RepositoryBuildingData(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindBuildingData() ([]models.BuildingData, error) {
	var buildingData []models.BuildingData
	err := r.db.Find(&buildingData).Error

	return buildingData, err
}

func (r *repository) CreateBuildingData(buildingData models.BuildingData) (models.BuildingData, error) {
	err := r.db.Create(&buildingData).Error

	return buildingData, err
}
