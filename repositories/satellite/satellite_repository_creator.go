package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
)

func GetSatelliteRepository() *repository.SatelliteRepository {
	repo := NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
	return &repo
}
