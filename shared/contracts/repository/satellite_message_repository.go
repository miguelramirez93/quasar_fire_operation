package repository

import "github.com/miguelramirez93/quasar_fire_operation/shared/models"

type SatelliteMessageRepository interface {
	GetSatelliteMessages() ([]*models.SatelliteMessage, error)
	AddSatelliteMessage(models.SatelliteMessage) (*models.SatelliteMessage, error)
}
