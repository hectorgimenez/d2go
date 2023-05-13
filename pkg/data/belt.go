package data

import (
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data/item"
)

const (
	HealingPotion      PotionType = "HealingPotion"
	ManaPotion         PotionType = "ManaPotion"
	RejuvenationPotion PotionType = "RejuvenationPotion"
)

type Belt struct {
	Items []Item
	Name  item.Name
}

func (b Belt) GetFirstPotion(potionType PotionType) (Position, bool) {
	for _, i := range b.Items {
		// Ensure potion is in row 0 and one of the four columns
		if strings.Contains(string(i.Name), string(potionType)) && i.Position.Y == 0 && (i.Position.X == 0 || i.Position.X == 1 || i.Position.X == 2 || i.Position.X == 3) {
			return i.Position, true
		}
	}

	return Position{}, false
}

func (b Belt) Rows() int {
	switch b.Name {
	case "":
		return 1
	case "Sash", "LightBelt":
		return 2
	case "Belt", "HeavyBelt":
		return 3
	default:
		return 4
	}
}

type PotionType string
