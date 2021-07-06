package repositories

import (
	"testing"

	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCreation(t *testing.T) {
	t.Run("Should create satellite in memory repository with expected data", func(t *testing.T) {
		repositoryinstance := NewSatelliteInmemoryRepository(fixtures.SatelliteDataSample)

		allRepoSatellites := repositoryinstance.GetSatellites()

		assert.Equal(t, len(fixtures.SatelliteDataSample), len(allRepoSatellites), "Should create an array of satelites with same len as samples")

		for _, satelliteSample := range fixtures.SatelliteDataSample {
			assert.NotNil(t, repositoryinstance.GetSatelliteByName(satelliteSample.Name), "Should map all samples")
		}

	})

}
