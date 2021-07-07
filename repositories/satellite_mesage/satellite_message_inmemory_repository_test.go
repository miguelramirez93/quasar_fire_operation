package repositories

import (
	"testing"

	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestSateliteMessageInMemoryRepo(t *testing.T) {
	t.Run("Should create satellite in memory repository with expected data and get it", func(t *testing.T) {
		repositoryinstance := NewSatelliteMessageInmemoryRepository([]*models.SatelliteMessage{})

		_, err := repositoryinstance.AddSatelliteMessage(models.SatelliteMessage{
			SatelliteName: "kenobi",
			Distance:      valueobjects.Distance(100),
			Message:       valueobjects.Message{"", "", "message"},
		})

		assert.Nil(t, err)
		allSatelliteMessages, _ := repositoryinstance.GetSatelliteMessages()
		assert.Equal(t, 1, len(allSatelliteMessages))

	})

}
