package math

import "math"

type Point struct {
	X int
	Y int
}

func DistanceFromPoint(from, to Point) int {
	first := math.Pow(float64(to.X-from.X), 2)
	second := math.Pow(float64(to.Y-from.Y), 2)
	return int(math.Sqrt(first + second))
}
