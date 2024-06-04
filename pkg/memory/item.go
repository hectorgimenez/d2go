package memory

import (
	"slices"
	"sort"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/d2go/pkg/data/state"
	"github.com/hectorgimenez/d2go/pkg/utils"
)

func (gd *GameReader) Inventory(rawPlayerUnits RawPlayerUnits, hover data.HoverData) data.Inventory {
	mainPlayer := rawPlayerUnits.GetMainPlayer()
	baseAddr := gd.Process.moduleBaseAddressPtr + gd.offset.UnitTable + (4 * 1024)
	unitTableBuffer := gd.Process.ReadBytesFromMemory(baseAddr, 128*8)

	stashPlayerUnits := make(map[uint]RawPlayerUnit)
	stashPlayerUnitOrder := make([]uint, 0)
	for _, pu := range rawPlayerUnits {
		if pu.States.HasState(state.Sharedstash) {
			order := gd.ReadUInt(pu.Address+0xD8, Uint64)
			stashPlayerUnitOrder = append(stashPlayerUnitOrder, order)
			stashPlayerUnits[order] = pu
		}
	}
	slices.Sort(stashPlayerUnitOrder)

	// Gold
	inventoryGold, _ := mainPlayer.BaseStats.FindStat(stat.Gold, 0)
	mainPlayerStashedGold, _ := mainPlayer.BaseStats.FindStat(stat.StashGold, 0)
	stashedGold := [4]int{}
	stashedGold[0] = mainPlayerStashedGold.Value
	for i, puKey := range stashPlayerUnitOrder {
		if i > 2 {
			break
		}
		stashGold, _ := stashPlayerUnits[puKey].BaseStats.FindStat(stat.StashGold, 0)
		stashedGold[i+1] = stashGold.Value
	}

	inventory := data.Inventory{
		Gold:        inventoryGold.Value,
		StashedGold: stashedGold,
	}
	belt := data.Belt{}
	for i := 0; i < 128; i++ {
		itemOffset := 8 * i
		itemUnitPtr := uintptr(ReadUIntFromBuffer(unitTableBuffer, uint(itemOffset), Uint64))
		for itemUnitPtr > 0 {
			itemDataBuffer := gd.Process.ReadBytesFromMemory(itemUnitPtr, 144)
			itemType := ReadUIntFromBuffer(itemDataBuffer, 0x00, Uint32)
			txtFileNo := ReadUIntFromBuffer(itemDataBuffer, 0x04, Uint32)
			unitID := ReadUIntFromBuffer(itemDataBuffer, 0x08, Uint32)
			// itemLoc = 0 in inventory, 1 equipped, 2 in belt, 3 on ground, 4 cursor, 5 dropping, 6 socketed
			itemLoc := ReadUIntFromBuffer(itemDataBuffer, 0x0C, Uint32)

			if itemType != 4 {
				itemUnitPtr = uintptr(gd.Process.ReadUInt(itemUnitPtr+0x150, Uint64))
				continue
			}

			unitDataPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x10, Uint64))
			unitDataBuffer := gd.Process.ReadBytesFromMemory(unitDataPtr, 144)
			flags := ReadUIntFromBuffer(unitDataBuffer, 0x18, Uint32)
			invPage := ReadUIntFromBuffer(unitDataBuffer, 0x55, Uint8)
			itemQuality := ReadUIntFromBuffer(unitDataBuffer, 0x00, Uint32)
			itemOwnerNPC := ReadUIntFromBuffer(unitDataBuffer, 0x0C, Uint32)

			// Item coordinates (X, Y)
			pathPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x38, Uint64))
			pathBuffer := gd.Process.ReadBytesFromMemory(pathPtr, 144)
			itemX := ReadUIntFromBuffer(pathBuffer, 0x10, Uint16)
			itemY := ReadUIntFromBuffer(pathBuffer, 0x14, Uint16)

			name := item.GetNameByEnum(txtFileNo)
			itemHovered := false
			if hover.IsHovered && hover.UnitType == 4 && hover.UnitID == data.UnitID(unitID) {
				itemHovered = true
			}

			itm := data.Item{
				ID:      int(txtFileNo),
				UnitID:  data.UnitID(unitID),
				Name:    name,
				Quality: item.Quality(itemQuality),
				Position: data.Position{
					X: int(itemX),
					Y: int(itemY),
				},
				IsHovered: itemHovered,
			}
			setProperties(&itm, uint32(flags))

			location := item.LocationUnknown
			switch itemLoc {
			case 0:
				if itemOwnerNPC == 2 || itemOwnerNPC == uint(stashPlayerUnits[stashPlayerUnitOrder[0]].UnitID) {
					location = item.LocationSharedStash
					invPage = 1
					break
				}
				if itemOwnerNPC == 3 || itemOwnerNPC == uint(stashPlayerUnits[stashPlayerUnitOrder[1]].UnitID) {
					location = item.LocationSharedStash
					invPage = 2
					break
				}
				if itemOwnerNPC == 4 || itemOwnerNPC == uint(stashPlayerUnits[stashPlayerUnitOrder[2]].UnitID) {
					location = item.LocationSharedStash
					invPage = 3
					break
				}

				if 0x00002000&flags != 0 && itemOwnerNPC == 4294967295 {
					location = item.LocationVendor
					break
				}
				if data.UnitID(itemOwnerNPC) == mainPlayer.UnitID || itemOwnerNPC == 1 {
					if invPage == 0 {
						location = item.LocationInventory
						break
					}
					location = item.LocationStash
					invPage = 0
					break
				}
			case 1:
				if data.UnitID(itemOwnerNPC) == mainPlayer.UnitID || itemOwnerNPC == 1 {
					location = item.LocationEquipped
					if itm.Type().Code == item.TypeBelt {
						belt.Name = itm.Name
					}
				}
			case 2:
				if data.UnitID(itemOwnerNPC) == mainPlayer.UnitID || itemOwnerNPC == 1 {
					location = item.LocationBelt
				}
			case 3, 5:
				location = item.LocationGround
			case 6:
				location = item.LocationSocket
			case 4:
				location = item.LocationCursor
			}

			itm.Location = item.Location{
				LocationType: location,
				Page:         int(invPage),
			}

			// We don't care about the inventory we don't know where they are, probably previous games or random crap
			if location != item.LocationUnknown {
				// Item Stats
				statsListExPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x88, Uint64))
				baseStats, stats := gd.getItemStats(statsListExPtr)
				itm.Stats = stats
				itm.BaseStats = baseStats

				if location == item.LocationBelt {
					belt.Items = append(belt.Items, itm)
				} else {
					inventory.AllItems = append(inventory.AllItems, itm)
				}
			}

			itemUnitPtr = uintptr(gd.Process.ReadUInt(itemUnitPtr+0x150, Uint64))
		}
	}

	inventory.Belt = belt

	sort.SliceStable(inventory.AllItems, func(i, j int) bool {
		distanceI := utils.DistanceFromPoint(mainPlayer.Position, inventory.AllItems[i].Position)
		distanceJ := utils.DistanceFromPoint(mainPlayer.Position, inventory.AllItems[j].Position)

		return distanceI < distanceJ
	})

	return inventory
}

func (gd *GameReader) getItemStats(statsListExPtr uintptr) (stat.Stats, stat.Stats) {
	fullStats := gd.getStatsList(statsListExPtr + 0x88)
	baseStats := gd.getStatsList(statsListExPtr + 0x30)

	flags := gd.Process.ReadUInt(statsListExPtr+0x1C, Uint64)
	lastStatsList := uintptr(gd.Process.ReadUInt(statsListExPtr+0x70, Uint64))

	if (flags & 0x80000000) == 0 {
		return baseStats, fullStats
	}
	attempts := 0
	statListFlags := uintptr(0)
	statListPrev := lastStatsList
	for (0x40 & statListFlags & 0xFFFFDFFF) == 0 {
		attempts++
		if attempts == 10 {
			return baseStats, fullStats
		}
		statListPrev = uintptr(gd.Process.ReadUInt(statListPrev+0x48, Uint64))
		statListFlags = uintptr(gd.Process.ReadUInt(statListPrev+0x1C, Uint64))
		if statListPrev == 0 {
			return baseStats, fullStats
		}
	}
	modifierBaseStats := gd.getStatsList(statListPrev + 0x30)

	if len(modifierBaseStats) != 0 {
		for _, mStat := range modifierBaseStats {
			if _, found := fullStats.FindStat(mStat.ID, mStat.Layer); !found {
				fullStats = append(fullStats, mStat)
			}
		}

		modifierStatsPrevEx := uintptr(gd.Process.ReadUInt(statListPrev+0x48, Uint64))
		modifierBaseStatsPrev := gd.getStatsList(modifierStatsPrevEx + 0x30)

		for _, mStat := range modifierBaseStatsPrev {
			for i, st := range fullStats {
				if st.ID == mStat.ID && st.Layer == mStat.Layer && st.Value != mStat.Value {
					fullStats[i].Value = st.Value + mStat.Value
				}
			}
			if _, found := fullStats.FindStat(mStat.ID, mStat.Layer); !found {
				fullStats = append(fullStats, mStat)
			}
			if _, found := baseStats.FindStat(mStat.ID, mStat.Layer); !found {
				baseStats = append(baseStats, mStat)
			}
		}
	}

	return baseStats, fullStats
}

func setProperties(item *data.Item, flags uint32) {
	if 0x00400000&flags != 0 {
		item.Ethereal = true
	}
	if 0x00000010&flags != 0 {
		item.Identified = true
	}
}
