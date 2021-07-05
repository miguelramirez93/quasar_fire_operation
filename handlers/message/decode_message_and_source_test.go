package messagehandler

import (
	"testing"

	repositories "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite"
	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestDecodeMessageAndSource(t *testing.T) {
	t.Run("Should return error if there are messages from unknow source", func(t *testing.T) {
		satelliteMockedRepo := new(repositories.SatelliteRepositoryMock)
		satelliteMockedRepo.On("GetSatelliteByName", "non_exist").Return(nil)
		handler := DecodeMessageAndSourceHandler{
			satelliteRepository: satelliteMockedRepo,
		}
		_, err := handler.DecodeMessageAndSource(fixtures.SatelliteMessagesUnknowSources)
		assert.NotNil(t, err)
		assert.Equal(t, "messages from non registered satellites", err.Error())
	})

	t.Run("Should return success decoded message", func(t *testing.T) {
		satelliteMockedRepo := repositories.NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
		handler := DecodeMessageAndSourceHandler{
			satelliteRepository: satelliteMockedRepo,
		}
		decodedMessage, err := handler.DecodeMessageAndSource(fixtures.SatelliteMessagesSuccess.Messages)
		assert.Nil(t, err)
		assert.Equal(t, fixtures.SatelliteMessagesSuccess.DecodedMessage, decodedMessage.Message)
	})

	t.Run("Should return error at out of range points", func(t *testing.T) {
		satelliteMockedRepo := repositories.NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)
		handler := DecodeMessageAndSourceHandler{
			satelliteRepository: satelliteMockedRepo,
		}
		_, err := handler.DecodeMessageAndSource(fixtures.SatelliteMessagesOutOfRange)
		assert.NotNil(t, err)
	})

}
