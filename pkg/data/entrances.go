package data

import (
	"github.com/hectorgimenez/d2go/pkg/data/entrance"
)

type Entrance struct {
	ID UnitID
	entrance.Name
	IsHovered  bool
	Selectable bool
	Position   Position
}

type Entrances []Entrance

func (e Entrances) FindOne(name entrance.Name) (Entrance, bool) {
	for _, ent := range e {
		if ent.Name == name {
			return ent, true
		}
	}
	return Entrance{}, false
}

func (e Entrances) FindByID(id UnitID) (Entrance, bool) {
	for _, ent := range e {
		if ent.ID == id {
			return ent, true
		}
	}
	return Entrance{}, false
}
