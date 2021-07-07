package messagehandler

import (
	"testing"

	satelliteRepository "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite"
	satelliteMessageRepository "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite_mesage"
	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
	"github.com/stretchr/testify/assert"

	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

func TestAddSatelliteMessageHandler(t *testing.T) {

	t.Run("Should return error if adding message from unknow source", func(t *testing.T) {
		satelliteMessageRepositoryinstance := satelliteMessageRepository.NewSatelliteMessageInmemoryRepository([]*models.SatelliteMessage{})
		satelliteRepositoryInstance := satelliteRepository.NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
		handler := AddSatelliteMessageHandler{
			satelliteMessageRepository: satelliteMessageRepositoryinstance,
			satelliteRepository:        satelliteRepositoryInstance,
		}
		_, err := handler.AddSatelliteMessage(*fixtures.SatelliteMessagesUnknowSources[0])
		assert.NotNil(t, err)
		assert.Equal(t, apperrors.ErrUnknowMessageSource.Error(), err.Error())
	})

	t.Run("Should Add satellite message", func(t *testing.T) {
		satelliteMessageRepositoryinstance := satelliteMessageRepository.NewSatelliteMessageInmemoryRepository([]*models.SatelliteMessage{})
		satelliteRepositoryInstance := satelliteRepository.NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
		handler := AddSatelliteMessageHandler{
			satelliteMessageRepository: satelliteMessageRepositoryinstance,
			satelliteRepository:        satelliteRepositoryInstance,
		}
		addedMessage, err := handler.AddSatelliteMessage(*fixtures.SatelliteMessagesSuccess.Messages[0])
		assert.Nil(t, err)
		assert.Equal(t, *fixtures.SatelliteMessagesSuccess.Messages[0], *addedMessage)
	})
}
