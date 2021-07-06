package messagehandler

import (
	"errors"

	"github.com/miguelramirez93/quasar_fire_operation/domain"
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"
)

type DecodeMessageAndSourceHandler struct {
	satelliteRepository repository.SatelliteRepository
}

func NewDecodeMessageAndSourceHandler(satelliteRepo repository.SatelliteRepository) DecodeMessageAndSourceHandler {
	return DecodeMessageAndSourceHandler{
		satelliteRepository: satelliteRepo,
	}
}

func (d *DecodeMessageAndSourceHandler) areMessagesFromRegisteredSatellites(satelliteMessages []models.SatelliteMessage) bool {
	for _, satelliteMessage := range satelliteMessages {
		if d.satelliteRepository.GetSatelliteByName(satelliteMessage.SatelliteName) == nil {
			return false
		}
	}
	return true
}

func (d *DecodeMessageAndSourceHandler) DecodeMessageAndSource(satelliteMessages []models.SatelliteMessage) (msg models.DecodedMessage, err error) {

	if !d.areMessagesFromRegisteredSatellites(satelliteMessages) {
		err := apperrors.ErrUnknowMessageSource
		return models.DecodedMessage{}, err
	}

	radar := domain.NewRadar(d.satelliteRepository.GetSatellites())

	messages := [][]string{}
	distances := []float32{}

	for _, satelliteMessage := range satelliteMessages {
		messages = append(messages, satelliteMessage.Message)
		distances = append(distances, float32(satelliteMessage.Distance))
	}

	defer func() {
		if r := recover(); r != nil {
			switch recoverError := r.(type) {
			case string:
				err = errors.New(recoverError)
			default:
				err = errors.New("unknow error")
			}
		}
	}()

	x, y := radar.GetLocation(distances...)

	decodedMsg := domain.GetMessage(messages...)

	return models.DecodedMessage{
		Position: valueobjects.CoordinatePair{
			X: valueobjects.Coordinate(x),
			Y: valueobjects.Coordinate(y),
		},
		Message: decodedMsg,
	}, nil
}
