package fixtures

import (
	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"
)

//[2]valueobjects.Coordinate{valueobjects.Coordinate(-500), valueobjects.Coordinate(-200)}
var (
	SatelliteDataSample = []*models.Satellite{
		{
			Name: "kenobi",
			Coordenates: valueobjects.CoordinatePair{
				X: valueobjects.Coordinate(-500),
				Y: valueobjects.Coordinate(-200),
			},
		}, {
			Name: "skywalker",
			Coordenates: valueobjects.CoordinatePair{
				X: valueobjects.Coordinate(100),
				Y: valueobjects.Coordinate(-100),
			},
		}, {
			Name: "sato",
			Coordenates: valueobjects.CoordinatePair{
				X: valueobjects.Coordinate(500),
				Y: valueobjects.Coordinate(100),
			},
		},
	}
)
