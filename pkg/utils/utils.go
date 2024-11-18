package utils

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/math"
)

func DistanceFromPoint(from data.Position, to data.Position) int {
	return math.DistanceFromPoint(
		math.Point{X: from.X, Y: from.Y},
		math.Point{X: to.X, Y: to.Y},
	)
}
