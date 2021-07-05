package domain

import (
	"fmt"
	"math"

	"github.com/miguelramirez93/quasar_fire_operation/shared/models"
	"github.com/miguelramirez93/quasar_fire_operation/shared/utils/numbers"
	valueobjects "github.com/miguelramirez93/quasar_fire_operation/shared/value_objects"
)

type Radar struct {
	models.RadarProps
}

func NewRadar(satellitesData []*models.Satellite) Radar {
	return Radar{
		models.RadarProps{
			Satellites: satellitesData,
		},
	}
}

func (rd *Radar) GetLocation(distances ...float32) (x, y float32) {
	if len(distances) > (len(rd.Satellites)) {
		panic("Number of distances should be same as number of available satellites")
	}

	r1 := distances[0]
	r2 := distances[1]
	r3 := distances[2]

	x1 := rd.Satellites[0].Coordenates.X
	y1 := rd.Satellites[0].Coordenates.Y
	x2 := rd.Satellites[1].Coordenates.X
	y2 := rd.Satellites[1].Coordenates.Y
	x3 := rd.Satellites[2].Coordenates.X
	y3 := rd.Satellites[2].Coordenates.Y
	//Calculate emisor coordenates with the intersection of the 3 circles that distances and coordenates build.

	// A = (-2x1 + 2x2)
	A := float32((-2 * x1) + (2 * x2))

	//B = (-2y1 + 2y2)
	B := float32((-2 * y1) + (2 * y2))

	//C = (r1^2) - (r2^2) - (x1^2) + (x2^2) - (y1^2) + (y2^2)
	C := (r1 * r1) - (r2 * r2) - float32(x1*x1) + float32(x2*x2) - float32(y1*y1) + float32(y2*y2)
	//D = (-2x2 + 2x3)
	D := float32((-2 * x2) + (2 * x3))
	//E = (-2y2 + 2y3)
	E := float32((-2 * y2) + (2 * y3))
	//F = (r2^2) - (r3^2) - (x2^2) + (x3^2) - (y2^2) + (y3^2)
	F := float32(math.Pow(float64(r2), 2)) - float32(math.Pow(float64(r3), 2)) - float32(math.Pow(float64(x2), 2)) + float32(math.Pow(float64(x3), 2)) - float32(math.Pow(float64(y2), 2)) + float32(math.Pow(float64(y3), 2))

	// Formula: x= (CE - FB) / (EA - BD)
	x = ((C * E) - (F * B)) / ((E * A) - (B * D))

	// Formula: y= (CD - AF) / (BD - AE)
	y = ((C * D) - (A * F)) / ((B * D) - (A * E))

	//check if data is consistent
	resultCoordenates := valueobjects.CoordinatePair{
		X: valueobjects.Coordinate(x),
		Y: valueobjects.Coordinate(y),
	}

	checkResultPointDistance(&resultCoordenates, valueobjects.CoordinatePair{
		X: x1,
		Y: y1,
	}, r1)

	checkResultPointDistance(&resultCoordenates, valueobjects.CoordinatePair{
		X: x2,
		Y: y2,
	}, r2)

	checkResultPointDistance(&resultCoordenates, valueobjects.CoordinatePair{
		X: x3,
		Y: y3,
	}, r3)

	return numbers.RoundDecimals(x, 1), numbers.RoundDecimals(y, 1)
}

func checkResultPointDistance(sourceCoordenates *valueobjects.CoordinatePair, pointToCehck valueobjects.CoordinatePair, distance float32) {
	distanceToX1Y1 := sourceCoordenates.CalculateDistanceTo(pointToCehck)

	if math.Round(float64(distanceToX1Y1)) != math.Round(float64(distance)) {
		panic(fmt.Sprintf("Point (%f,%f) is not consistent with distance %f", pointToCehck.X, pointToCehck.Y, distance))
	}
}
