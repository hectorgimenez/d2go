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

func (i Inventory) Find(name item.Name, locations ...item.LocationType) (Item, bool) {
	for _, it := range i.AllItems {
		if strings.EqualFold(string(it.Name), string(name)) {
			// If no locations are specified, return the first item found
			if len(locations) == 0 {
				return it, true
			}

			for _, l := range locations {
				if it.Location.LocationType == l {
					return it, true
				}
			}
		}
	}

	return Item{}, false
}

func (i Inventory) FindByID(unitID UnitID) (Item, bool) {
	for _, it := range i.AllItems {
		if it.UnitID == unitID {
			return it, true
		}
	}

	return Item{}, false
}

func (i Inventory) ByLocation(locations ...item.LocationType) []Item {
	var items []Item

	for _, it := range i.AllItems {
		for _, l := range locations {
			if it.Location.LocationType == l {
				items = append(items, it)
			}
		}
	}

	return items
}

func (i Inventory) Matrix() [4][10]bool {
	invMatrix := [4][10]bool{} // false = empty, true = occupied
	for _, itm := range i.ByLocation(item.LocationInventory) {
		for k := range itm.Desc().InventoryWidth {
			for j := range itm.Desc().InventoryHeight {
				invMatrix[itm.Position.Y+j][itm.Position.X+k] = true
			}
		}
	}

	return invMatrix
}

type UnitID int

type ItemAffixes struct {
	Rare struct {
		Prefix int16
		Suffix int16
	}
	AutoAffix int16
	Magic     struct {
		Prefixes [3]int16 // Prefix1, Prefix2, Prefix3
		Suffixes [3]int16 // Suffix1, Suffix2, Suffix3
	}
}

type Item struct {
	ID int
	UnitID
	Name                 item.Name
	Quality              item.Quality
	Position             Position
	Location             item.Location
	IsRuneword           bool
	Ethereal             bool
	IsHovered            bool
	BaseStats            stat.Stats
	Stats                stat.Stats
	Affixes              ItemAffixes
	Sockets              []Item
	Identified           bool
	IsNamed              bool
	IsStartItem          bool
	IsEar                bool
	IsBroken             bool
	IsEquipped           bool
	HasSockets           bool
	InTradeOrStoreScreen bool
	IsInSocket           bool
}

type Drop struct {
	Item         Item
	Rule         string
	RuleFile     string
	DropLocation string
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

// HasPrefix checks if the item has a specific prefix ID
// Works with both magic and rare prefixes
func (i Item) HasPrefix(id int16) bool {
	// Check rare prefix
	if i.Affixes.Rare.Prefix == id {
		return true
	}

	// Check magic prefixes
	for _, prefix := range i.Affixes.Magic.Prefixes {
		if prefix == id {
			return true
		}
	}
	return false
}

// HasSuffix checks if the item has a specific suffix ID
// Works with both magic and rare suffixes
func (i Item) HasSuffix(id int16) bool {
	// Check rare suffix
	if i.Affixes.Rare.Suffix == id {
		return true
	}

	// Check magic suffixes
	for _, suffix := range i.Affixes.Magic.Suffixes {
		if suffix == id {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the item has any of the provided prefix IDs
func (i Item) HasAnyPrefix(ids ...int16) bool {
	for _, id := range ids {
		if i.HasPrefix(id) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the item has any of the provided suffix IDs
func (i Item) HasAnySuffix(ids ...int16) bool {
	for _, id := range ids {
		if i.HasSuffix(id) {
			return true
		}
	}
	return false
}

// HasAllPrefixes checks if the item has all the provided prefix IDs
func (i Item) HasAllPrefixes(ids ...int16) bool {
	for _, id := range ids {
		if !i.HasPrefix(id) {
			return false
		}
	}
	return true
}

// HasAllSuffixes checks if the item has all provided suffix IDs
func (i Item) HasAllSuffixes(ids ...int16) bool {
	for _, id := range ids {
		if !i.HasSuffix(id) {
			return false
		}
	}
	return true
}

// GetMagicPrefixes returns a slice of all non-zero magic prefixes
func (i Item) GetMagicPrefixes() []int16 {
	var prefixes []int16
	for _, prefix := range i.Affixes.Magic.Prefixes {
		if prefix > 0 {
			prefixes = append(prefixes, prefix)
		}
	}
	return prefixes
}

// GetMagicSuffixes returns a slice of all non-zero magic suffixes
func (i Item) GetMagicSuffixes() []int16 {
	var suffixes []int16
	for _, suffix := range i.Affixes.Magic.Suffixes {
		if suffix > 0 {
			suffixes = append(suffixes, suffix)
		}
	}
	return suffixes
}

// HasMagicAffixes returns true if the item has any magic affixes
func (i Item) HasMagicAffixes() bool {
	return len(i.GetMagicPrefixes()) > 0 || len(i.GetMagicSuffixes()) > 0
}

// HasRareAffixes returns true if the item has rare affixes
func (i Item) HasRareAffixes() bool {
	return i.Affixes.Rare.Prefix > 0 || i.Affixes.Rare.Suffix > 0
}

// GetNumMagicAffixes returns the total number of magic affixes present
func (i Item) GetNumMagicAffixes() int {
	return len(i.GetMagicPrefixes()) + len(i.GetMagicSuffixes())
}

// HasAutoAffix checks if the item has a specific auto affix ID
func (i Item) HasAutoAffix(id int16) bool {
	return i.Affixes.AutoAffix == id
}

// HasAnyAutoAffix checks if the item has any of the provided auto affix IDs
func (i Item) HasAnyAutoAffix(ids ...int16) bool {
	for _, id := range ids {
		if i.HasAutoAffix(id) {
			return true
		}
	}
	return false
}

// GetAutoAffix returns the auto affix ID if it exists, or -1 if not present
func (i Item) GetAutoAffix() int16 {
	if i.Affixes.AutoAffix > 0 {
		return i.Affixes.AutoAffix
	}
	return -1
}

// GetSocketedItems returns all items socketed in this base item
func (i Item) GetSocketedItems() []Item {
	return i.Sockets
}

// FindSocketedItem finds the first socketed item matching the given name
func (i Item) FindSocketedItem(name item.Name) (Item, bool) {
	for _, socket := range i.Sockets {
		if socket.Name == name {
			return socket, true
		}
	}
	return Item{}, false
}

// CountSocketedItems returns the total number of items socketed
func (i Item) CountSocketedItems() int {
	return len(i.Sockets)
}

// HasSocketedItem checks if a specific item is socketed in this base
func (i Item) HasSocketedItem(name item.Name) bool {
	_, found := i.FindSocketedItem(name)
	return found
}

// HasAllSocketedItems checks if all of the specified items are socketed
// (e.g., checking runeword recipe) if item.HasAllSocketedItems("Jah", "Ith", "Ber")
func (i Item) HasAllSocketedItems(names ...item.Name) bool {
	for _, name := range names {
		if !i.HasSocketedItem(name) {
			return false
		}
	}
	return true
}
