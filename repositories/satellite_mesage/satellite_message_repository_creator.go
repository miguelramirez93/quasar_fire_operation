package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

var (
	satelliteMsgRepoInstance *repository.SatelliteMessageRepository
)

func GetSatelliteMessageRepository() *repository.SatelliteMessageRepository {
	if satelliteMsgRepoInstance != nil {
		return satelliteMsgRepoInstance
	}
	repo := NewSatelliteMessageInmemoryRepository([]*models.SatelliteMessage{})
	satelliteMsgRepoInstance = &repo
	return satelliteMsgRepoInstance
}
