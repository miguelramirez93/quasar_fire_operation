package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	"github.com/stretchr/testify/mock"
)

type SatelliteMessageRepositoryMock struct {
	mock.Mock
}

func (r *SatelliteMessageRepositoryMock) GetSatelliteMessages() (result []*models.SatelliteMessage, err error) {
	args := r.Called()

	result = args.Get(0).([]*models.SatelliteMessage)

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *SatelliteMessageRepositoryMock) AddSatelliteMessage(satelliteMessageData models.SatelliteMessage) (result *models.SatelliteMessage, err error) {
	args := r.Called(satelliteMessageData)
	if args.Get(0) == nil {
		result = nil
	} else {
		result = args.Get(0).(*models.SatelliteMessage)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return
}

func (r *SatelliteMessageRepositoryMock) CleanSatelliteMessages() (err error) {
	args := r.Called()

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return
}
