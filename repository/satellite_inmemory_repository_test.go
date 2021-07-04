package repository

import (
	"testing"

	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"
	"github.com/stretchr/testify/assert"
)

var (
	satelliteDataSample = []*models.Satellite{
		{
			Name:             "kenobi",
			DistanceToEmisor: valueobjects.Distance(100.0),
			MessageData:      valueobjects.Message([]string{"este", "", "", "mensaje", ""}),
		}, {
			Name:             "skywalker",
			DistanceToEmisor: valueobjects.Distance(115.5),
			MessageData:      valueobjects.Message([]string{"este", "", "", "mensaje", ""}),
		}, {
			Name:             "sato",
			DistanceToEmisor: valueobjects.Distance(142.7),
			MessageData:      valueobjects.Message([]string{"este", "", "", "mensaje", ""}),
		},
	}
)

func TestRepositoryCreation(t *testing.T) {
	t.Run("Should create satellite in memory repository with expected data", func(t *testing.T) {
		repositoryinstance := NewSatelliteInmemoryRepository(satelliteDataSample)

		allRepoSatellites := repositoryinstance.GetSatellites()

		assert.Equal(t, len(satelliteDataSample), len(allRepoSatellites), "Should create an array of satelites with same len as samples")

		for _, satelliteSample := range satelliteDataSample {
			assert.NotNil(t, repositoryinstance.GetSatelliteByName(satelliteSample.Name), "Should map all samples")
		}

	})

}
