package video

import "go-service/video/models"

type CategorySync interface {
	GetCagetories(regionCode string) (*[]models.DataCategory, error)
}
