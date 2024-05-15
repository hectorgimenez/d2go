package data

import (
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

type Inventory struct {
	Belt        Belt
	AllItems    []Item
	Gold        int
	StashedGold [4]int
}

func (i Inventory) Find(name item.Name, locations ...item.Location) (Item, bool) {
	for _, it := range i.AllItems {
		if strings.EqualFold(string(it.Name), string(name)) {
			// If no locations are specified, return the first item found
			if len(locations) == 0 {
				return it, true
			}

			for _, l := range locations {
				if it.Location == l {
					return it, true
				}
			}
		}
	}

	return Item{}, false
}

func (i Inventory) ByLocation(locations ...item.Location) []Item {
	var items []Item

	for _, it := range i.AllItems {
		for _, l := range locations {
			if it.Location == l {
				items = append(items, it)
			}
		}
	}

	return items
}

type UnitID int

type Item struct {
	ID int
	UnitID
	Name       item.Name
	Quality    item.Quality
	Position   Position
	Location   item.Location
	Page       int // Used for shared stash
	Ethereal   bool
	IsHovered  bool
	BaseStats  stat.Stats
	Stats      stat.Stats
	Identified bool
}

func (i Item) Desc() item.Description {
	return item.Desc[i.ID]
}

func (i Item) Type() item.Type {
	return i.Desc().GetType()
}

func (i Item) IsPotion() bool {
	return i.IsHealingPotion() || i.IsManaPotion() || i.IsRejuvPotion()
}

func (i Item) IsHealingPotion() bool {
	return i.Type().IsType(item.TypeHealingPotion)
}

func (i Item) IsManaPotion() bool {
	return i.Type().IsType(item.TypeManaPotion)
}

func (i Item) IsRejuvPotion() bool {
	return i.Type().IsType(item.TypeRejuvPotion)
}

func (i Item) IsFromQuest() bool {
	return i.Type().IsType(item.TypeQuest)
}

func (i Item) FindStat(id stat.ID, layer int) (stat.Data, bool) {
	st, found := i.Stats.FindStat(id, layer)
	if found {
		return st, true
	}

	return i.BaseStats.FindStat(id, layer)
}
