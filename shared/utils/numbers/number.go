package numbers

import "math"

func RoundDecimals(targetNumber float32, decimalsToRound int) float32 {
	return float32(math.Round(float64(targetNumber*float32(math.Pow(10, float64(decimalsToRound))))) / (math.Pow(10, float64(decimalsToRound))))
}
