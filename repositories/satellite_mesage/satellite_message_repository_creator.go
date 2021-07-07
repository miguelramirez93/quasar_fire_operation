package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

func GetSatelliteMessageRepository() *repository.SatelliteMessageRepository {
	repo := NewSatelliteMessageInmemoryRepository([]*models.SatelliteMessage{})
	return &repo
}
