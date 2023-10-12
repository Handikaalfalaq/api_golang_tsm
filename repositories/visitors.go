package repositories

import (
	"tsmweb/models"

	"gorm.io/gorm"
)

type VisitorsRepository interface {
	FindVisitors() ([]models.Visitors, error)
	CreateVisitors(visitors models.Visitors) (models.Visitors, error)
}

func RepositoryVisitors(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindVisitors() ([]models.Visitors, error) {
	var visitors []models.Visitors
	err := r.db.Find(&visitors).Error

	return visitors, err
}

func (r *repository) CreateVisitors(visitors models.Visitors) (models.Visitors, error) {
	err := r.db.Create(&visitors).Error

	return visitors, err
}
