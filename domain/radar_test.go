package domain

import (
	"testing"

	"github.com/miguelramirez93/quasar_fire_operation/tests/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestRadarDomain(t *testing.T) {
	radar := NewRadar(fixtures.SatelliteDataSample)
	t.Run("Should panic if  # of distances provided are greater than available satellites", func(t *testing.T) {

		assert.Panics(t, func() {
			radar.GetLocation(1, 2, 3, 4)
		}, "Should panic as expected")
	})

	t.Run("Should get expected coordenates for provided distances and samples", func(t *testing.T) {
		x, y := radar.GetLocation(485.67, 266.08, 600.5)
		assert.Equal(t, float32(-100), x)
		assert.Equal(t, float32(75.6), y)
	})

	t.Run("Should panic if one or more points are unconsistent", func(t *testing.T) {
		assert.Panics(t, func() {
			radar.GetLocation(100, 115.5, 142.7)
		}, "Should panic as expected")
	})
}
