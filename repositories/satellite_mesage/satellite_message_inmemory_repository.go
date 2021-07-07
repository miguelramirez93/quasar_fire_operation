package repositories

import (
	"errors"

	"github.com/miguelramirez93/quasar_fire_operation/shared/contracts/repository"
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
)

type SatelliteMessageInMemoryRepository struct {
	MappedSatelliteMessageData map[string]*models.SatelliteMessage
	SatelliteMessageData       []*models.SatelliteMessage
}

func SatelliteMessageDataToMap(data []*models.SatelliteMessage) (mappedSatelliteData map[string]*models.SatelliteMessage) {
	mappedSatelliteData = make(map[string]*models.SatelliteMessage)
	for _, message := range data {
		mappedSatelliteData[message.SatelliteName] = message
	}
	return
}

func NewSatelliteMessageInmemoryRepository(satelliteMessageData []*models.SatelliteMessage) repository.SatelliteMessageRepository {
	mappedSatelliteData := SatelliteMessageDataToMap(satelliteMessageData)
	return &SatelliteMessageInMemoryRepository{
		MappedSatelliteMessageData: mappedSatelliteData,
		SatelliteMessageData:       satelliteMessageData,
	}
}

func (r *SatelliteMessageInMemoryRepository) GetSatelliteMessages() ([]*models.SatelliteMessage, error) {
	return r.SatelliteMessageData, nil
}

func (r *SatelliteMessageInMemoryRepository) AddSatelliteMessage(satelliteMessageData models.SatelliteMessage) (result *models.SatelliteMessage, err error) {
	r.MappedSatelliteMessageData[satelliteMessageData.SatelliteName] = &satelliteMessageData
	r.SatelliteMessageData = append(r.SatelliteMessageData, &satelliteMessageData)
	result = r.MappedSatelliteMessageData[satelliteMessageData.SatelliteName]
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
	return
}

func (r *SatelliteMessageInMemoryRepository) CleanSatelliteMessages() (err error) {
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
	r.MappedSatelliteMessageData = make(map[string]*models.SatelliteMessage)
	r.SatelliteMessageData = []*models.SatelliteMessage{}
	return
}
