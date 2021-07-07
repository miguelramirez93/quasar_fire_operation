package messagehandler

import (
	"errors"
	"testing"

	satellitemsgrepo "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite_mesage"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSplitedMessageAndSource(t *testing.T) {
	t.Run("Should return error if there is an error getting satellite msg data", func(t *testing.T) {
		satelliteMessageMockedRepo := new(satellitemsgrepo.SatelliteMessageRepositoryMock)
		testErr := errors.New("test-error")
		satelliteMessageMockedRepo.On("GetSatelliteMessages").Return([]*models.SatelliteMessage{}, testErr)
		handlerInstance := DecodeSplitedMessageAndSourceHandler{
			satelliteMessageRepository: satelliteMessageMockedRepo,
			decodeMsgAndSourceHandler:  NewDecodeMessageAndSourceHandler(),
		}

		_, err := handlerInstance.DecodeSplitedMessageAndSource()

		assert.Equal(t, testErr, err)
	})

	t.Run("Should return error if data is incomplete", func(t *testing.T) {
		satelliteMessageMockedRepo := new(satellitemsgrepo.SatelliteMessageRepositoryMock)
		satelliteMessageMockedRepo.On("GetSatelliteMessages").Return([]*models.SatelliteMessage{}, nil)
		handlerInstance := DecodeSplitedMessageAndSourceHandler{
			satelliteMessageRepository: satelliteMessageMockedRepo,
			decodeMsgAndSourceHandler:  NewDecodeMessageAndSourceHandler(),
		}

		_, err := handlerInstance.DecodeSplitedMessageAndSource()

		assert.Equal(t, apperrors.ErrUncompleteMsg, err)
	})

	t.Run("Should clean messages data after success calculation", func(t *testing.T) {
		satelliteMessageMockedRepo := new(satellitemsgrepo.SatelliteMessageRepositoryMock)
		satelliteMessageMockedRepo.On("GetSatelliteMessages").Return(fixtures.SatelliteMessagesSuccess.Messages, nil)
		satelliteMessageMockedRepo.On("CleanSatelliteMessages").Return(nil)
		handlerInstance := DecodeSplitedMessageAndSourceHandler{
			satelliteMessageRepository: satelliteMessageMockedRepo,
			decodeMsgAndSourceHandler:  NewDecodeMessageAndSourceHandler(),
		}

		_, err := handlerInstance.DecodeSplitedMessageAndSource()
		assert.Nil(t, err)
	})

}
