package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

type SatelliteInMemoryRepository struct {
	MappedSatelliteData map[string]*models.Satellite
	SatelliteData       []*models.Satellite
}

func SatelliteDataToMap(data []*models.Satellite) (mappedSatelliteData map[string]*models.Satellite) {
	mappedSatelliteData = make(map[string]*models.Satellite)
	for _, satellite := range data {
		mappedSatelliteData[satellite.Name] = satellite
	}
	return
}

func NewSatelliteInmemoryRepository(satelliteData []*models.Satellite) repository.SatelliteRepository {
	mappedSatelliteData := SatelliteDataToMap(satelliteData)
	return &SatelliteInMemoryRepository{
		MappedSatelliteData: mappedSatelliteData,
		SatelliteData:       satelliteData,
	}
}

func (r *SatelliteInMemoryRepository) GetSatellites() []*models.Satellite {
	return r.SatelliteData
}

func (r *SatelliteInMemoryRepository) GetSatelliteByName(name string) *models.Satellite {
	return r.MappedSatelliteData[name]
}
