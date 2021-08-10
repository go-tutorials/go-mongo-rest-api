package tube_category_service

import (
	"go-service/internal/models"
)

type CategorySync interface {
	GetCagetories(regionCode string) ([]models.DataCategory, error)
}