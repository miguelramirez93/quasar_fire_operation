package repositories

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	"github.com/stretchr/testify/mock"
)

type SatelliteRepositoryMock struct {
	mock.Mock
}

func (r *SatelliteRepositoryMock) GetSatellites() []*models.Satellite {
	args := r.Called()
	return args.Get(0).([]*models.Satellite)
}
func (r *SatelliteRepositoryMock) GetSatelliteByName(name string) *models.Satellite {
	args := r.Called(name)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.Satellite)
}
