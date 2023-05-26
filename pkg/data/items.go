package data

import (
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

type Items struct {
	Belt     Belt
	AllItems []Item
}

func (i Items) Find(name item.Name, locations ...item.Location) (Item, bool) {
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

func (i Items) ByLocation(locations ...item.Location) []Item {
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
	UnitID
	Name       item.Name
	Quality    item.Quality
	Position   Position
	Location   item.Location
	Ethereal   bool
	IsHovered  bool
	Stats      map[stat.ID]stat.Data
	Identified bool
}

func (i Item) Type() string {
	t, _ := item.TypeForItemName(string(i.Name))

	return t
}

func (i Item) IsPotion() bool {
	return i.IsHealingPotion() || i.IsManaPotion() || i.IsRejuvPotion()
}

func (i Item) IsHealingPotion() bool {
	return strings.Contains(string(i.Name), string(HealingPotion))
}

func (i Item) IsManaPotion() bool {
	return strings.Contains(string(i.Name), string(ManaPotion))
}
func (i Item) IsRejuvPotion() bool {
	return strings.Contains(string(i.Name), string(RejuvenationPotion))
}
