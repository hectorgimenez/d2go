package data

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/math"
)

type TransitionState int

const (
	TransitionNone TransitionState = iota
	TransitionInProgress
	TransitionComplete
)

type Transition struct {
	ID         UnitID
	FromArea   area.ID
	ToArea     area.ID
	Position   Position
	IsEntrance bool
}

type Transitions []Transition

func (t Transitions) FindByArea(a area.ID) (Transition, bool) {
	for _, tr := range t {
		if tr.ToArea == a {
			return tr, true
		}
	}
	return Transition{}, false
}

func (t Transitions) FindByID(id UnitID) (Transition, bool) {
	for _, tr := range t {
		if tr.ID == id {
			return tr, true
		}
	}
	return Transition{}, false
}

func (t Transitions) IsValidTransitionTarget(pos Position) bool {
	for _, tr := range t {
		if math.DistanceFromPoint(
			math.Point{X: pos.X, Y: pos.Y},
			math.Point{X: tr.Position.X, Y: tr.Position.Y},
		) < 5 {
			return true
		}
	}
	return false
}
