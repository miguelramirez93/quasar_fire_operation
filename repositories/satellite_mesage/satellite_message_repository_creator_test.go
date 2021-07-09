package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSatelliteMsgRepositoryCreator(t *testing.T) {
	t.Run("Should succesfully create a singleton repository", func(t *testing.T) {
		createdRepo := GetSatelliteMessageRepository()
		assert.NotNil(t, createdRepo)

		createdRepo2 := GetSatelliteMessageRepository()
		assert.Equal(t, createdRepo, createdRepo2)
	})
}
