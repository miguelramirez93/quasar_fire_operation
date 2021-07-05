package valueobjects

import (
	"math"
)

type Distance float32

type Coordinate float32

type CoordinatePair struct {
	X Coordinate
	Y Coordinate
}

func (c *CoordinatePair) CalculateDistanceTo(p2 CoordinatePair) Distance {
	x1 := p2.X
	y1 := p2.Y

	x2 := c.X
	y2 := c.Y

	unformatedResult := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))

	return Distance(unformatedResult)
}
