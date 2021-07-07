package messagehandler

import (
	satelliteRepository "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite"
	satelliteMessageRepository "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite_mesage"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	"github.com/miguelramirez93/quasar_fire_operation/shared/utils/logger"
)

type AddSatelliteMessageHandler struct {
	satelliteMessageRepository repository.SatelliteMessageRepository
	satelliteRepository        repository.SatelliteRepository
}

func NewAddSatelliteMessageHandler() AddSatelliteMessageHandler {
	return AddSatelliteMessageHandler{
		satelliteRepository:        *satelliteRepository.GetSatelliteRepository(),
		satelliteMessageRepository: *satelliteMessageRepository.GetSatelliteMessageRepository(),
	}
}

func (h *AddSatelliteMessageHandler) isMessageFromRegisteredSatellites(satelliteMessage models.SatelliteMessage) bool {
	return h.satelliteRepository.GetSatelliteByName(satelliteMessage.SatelliteName) != nil
}

func (h *AddSatelliteMessageHandler) AddSatelliteMessage(satelliteMessageData models.SatelliteMessage) (*models.SatelliteMessage, error) {
	if !h.isMessageFromRegisteredSatellites(satelliteMessageData) {
		err := apperrors.ErrUnknowMessageSource
		logger.Error("error adding satellite message", err)
		return nil, err
	}

	return h.satelliteMessageRepository.AddSatelliteMessage(satelliteMessageData)

}
