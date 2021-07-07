package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
)

var (
	satelliteRepoInstance *repository.SatelliteRepository
)

func GetSatelliteRepository() *repository.SatelliteRepository {
	if satelliteRepoInstance != nil {
		return satelliteRepoInstance
	}
	repo := NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
	satelliteRepoInstance = &repo
	return satelliteRepoInstance
}
