package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSatelliteRepositoryCreator(t *testing.T) {
	t.Run("Should succesfully create a singleton repository", func(t *testing.T) {
		createdRepo := GetSatelliteRepository()
		assert.NotNil(t, createdRepo)

		createdRepo2 := GetSatelliteRepository()
		assert.Equal(t, createdRepo, createdRepo2)
	})
}
