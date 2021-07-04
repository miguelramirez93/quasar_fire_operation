package repository

import "github.com/miguelramirez93/quasar_fire_operation/shared/models"

type SatelliteRepository interface {
	GetSatellites() []*models.Satellite
	GetSatelliteByName(string) *models.Satellite
}
