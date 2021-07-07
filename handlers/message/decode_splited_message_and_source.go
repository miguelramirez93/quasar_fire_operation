package messagehandler

import (
	repositories "github.com/miguelramirez93/quasar_fire_operation/repositories/satellite_mesage"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

type DecodeSplitedMessageAndSourceHandler struct {
	satelliteMessageRepository repository.SatelliteMessageRepository
	decodeMsgAndSourceHandler  DecodeMessageAndSourceHandler
}

func NewDecodeSplitMessageAndSourceHandler() DecodeSplitedMessageAndSourceHandler {
	return DecodeSplitedMessageAndSourceHandler{
		satelliteMessageRepository: *repositories.GetSatelliteMessageRepository(),
		decodeMsgAndSourceHandler:  NewDecodeMessageAndSourceHandler(),
	}
}

func (h *DecodeSplitedMessageAndSourceHandler) DecodeSplitedMessageAndSource() (msg models.DecodedMessage, err error) {
	currentReceivedMsgs, err := h.satelliteMessageRepository.GetSatelliteMessages()

	if err != nil {
		return
	}

	satelliteSources := h.decodeMsgAndSourceHandler.satelliteRepository.GetSatellites()

	if !haveAllMessagesFromSources(satelliteSources, currentReceivedMsgs) {
		err = apperrors.ErrUncompleteMsg
		return
	}

	msg, err = h.decodeMsgAndSourceHandler.DecodeMessageAndSource(currentReceivedMsgs)

	if err == nil {
		err = h.satelliteMessageRepository.CleanSatelliteMessages()
	}

	return
}

func haveAllMessagesFromSources(satelliteSources []*models.Satellite, receivedMessages []*models.SatelliteMessage) bool {
	return len(satelliteSources) == len(receivedMessages)
}
