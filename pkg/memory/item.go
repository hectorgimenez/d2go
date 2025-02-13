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

type ItemData struct {
	Item      *data.Item
	ThisPtr   uintptr // Pointer to this item
	ParentPtr uintptr // Parent pointer (only set for socketed items)
}

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

	baseItems := make([]ItemData, 0)     // Items that can have sockets
	socketedItems := make([]ItemData, 0) // Items that are in sockets
	allItems := make([]*data.Item, 0)

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
				itemUnitPtr = uintptr(gd.Process.ReadUInt(itemUnitPtr+0x158, Uint64))
				continue
			}

			unitDataPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x10, Uint64))
			unitDataBuffer := gd.Process.ReadBytesFromMemory(unitDataPtr, 144)
			flags := ReadUIntFromBuffer(unitDataBuffer, 0x18, Uint32)
			invPage := ReadUIntFromBuffer(unitDataBuffer, 0x55, Uint8)
			itemQuality := ReadUIntFromBuffer(unitDataBuffer, 0x00, Uint32)
			itemOwnerNPC := ReadUIntFromBuffer(unitDataBuffer, 0x0C, Uint32)

			// Path and position data
			pathPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x38, Uint64))
			pathBuffer := gd.Process.ReadBytesFromMemory(pathPtr, 144)
			// Item coordinates (X, Y)
			itemX := ReadUIntFromBuffer(pathBuffer, 0x10, Uint16)
			itemY := ReadUIntFromBuffer(pathBuffer, 0x14, Uint16)

			name := item.GetNameByEnum(txtFileNo)
			itemHovered := false
			if hover.IsHovered && hover.UnitType == 4 && hover.UnitID == data.UnitID(unitID) {
				itemHovered = true
			}

			// Read rare affixes (should make rare item name)
			rarePrefix := int16(gd.Process.ReadUInt(unitDataPtr+0x42, Uint16))
			rareSuffix := int16(gd.Process.ReadUInt(unitDataPtr+0x44, Uint16))

			autoAffix := int16(gd.Process.ReadUInt(unitDataPtr+0x46, Uint16))

			// Read magic affixes
			// From prefix1  we can also tell Runeword name : Spirit 20635, cta 20519 , infinity 20566  getlocalestring.txt
			var prefixes [3]int16
			var suffixes [3]int16

			for i := 0; i < 3; i++ {
				prefixes[i] = int16(gd.Process.ReadUInt(unitDataPtr+0x48+uintptr(i*2), Uint16))
				suffixes[i] = int16(gd.Process.ReadUInt(unitDataPtr+0x4E+uintptr(i*2), Uint16))
			}

			itm := &data.Item{
				ID:      int(txtFileNo),
				UnitID:  data.UnitID(unitID),
				Name:    name,
				Quality: item.Quality(itemQuality),
				Position: data.Position{
					X: int(itemX),
					Y: int(itemY),
				},
				IsHovered: itemHovered,
				Affixes: data.ItemAffixes{
					Rare: struct {
						Prefix int16
						Suffix int16
					}{
						Prefix: rarePrefix,
						Suffix: rareSuffix,
					},
					AutoAffix: autoAffix,
					Magic: struct {
						Prefixes [3]int16
						Suffixes [3]int16
					}{
						Prefixes: prefixes,
						Suffixes: suffixes,
					},
				},
				Sockets: make([]data.Item, 0),
			}

			setProperties(itm, uint32(flags))

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
					if invPage == 3 {
						location = item.LocationCube
						invPage = 0
						break
					}
					location = item.LocationStash
					invPage = 0
					break
				}
			case 1:
				isMercItem := (flags & 0x800000) != 0
				if data.UnitID(itemOwnerNPC) == mainPlayer.UnitID || itemOwnerNPC == 1 {
					location = item.LocationEquipped
					if itm.Type().Code == item.TypeBelt {
						belt.Name = itm.Name
					}
				} else if isMercItem {
					location = item.LocationMercenary
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

			bodyLoc := item.LocNone
			equipSlotFlags := uint16(gd.Process.ReadUInt(unitDataPtr+uintptr(0x54), Uint16))
			if equipSlotFlags&0xFF00 == 0xFF00 {
				equipSlot := uint8(equipSlotFlags & 0xFF)
				switch equipSlot {
				case 0x01:
					bodyLoc = item.LocHead
				case 0x02:
					bodyLoc = item.LocNeck
				case 0x03:
					bodyLoc = item.LocTorso
				case 0x04:
					bodyLoc = item.LocLeftArm
				case 0x05:
					bodyLoc = item.LocRightArm
				case 0x06:
					bodyLoc = item.LocLeftRing
				case 0x07:
					bodyLoc = item.LocRightRing
				case 0x08:
					bodyLoc = item.LocBelt
				case 0x09:
					bodyLoc = item.LocFeet
				case 0x0A:
					bodyLoc = item.LocGloves
				case 0x0B:
					bodyLoc = item.LocLeftArmSecondary
				case 0x0C:
					bodyLoc = item.LocRightArmSecondary
				}
			}

			itm.Location = item.Location{
				LocationType: location,
				BodyLocation: bodyLoc,
				Page:         int(invPage),
			}

			// We don't care about the inventory we don't know where they are, probably previous games or random crap
			if location != item.LocationUnknown {
				// Item Stats
				statsListExPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x88, Uint64))
				baseStats, stats := gd.getItemStats(statsListExPtr)
				itm.Stats = stats
				itm.BaseStats = baseStats

				if location == item.LocationSocket {
					// Get parent pointer for socketed items
					itemExtraData := uintptr(gd.Process.ReadUInt(unitDataPtr+0xA0, Uint64))
					parentItemPtr := uintptr(0)
					if itemExtraData != 0 {
						parentItemPtr = uintptr(gd.Process.ReadUInt(itemExtraData+0x08, Uint64))
					}

					socketedItems = append(socketedItems, ItemData{
						Item:      itm,
						ParentPtr: parentItemPtr,
					})
				} else {
					// Store base item pointer if it has sockets
					numSockets, _ := itm.Stats.FindStat(stat.NumSockets, 0)
					if numSockets.Value > 0 {
						baseItems = append(baseItems, ItemData{
							Item:    itm,
							ThisPtr: itemUnitPtr,
						})
					}
				}

				allItems = append(allItems, itm)

				if location == item.LocationBelt {
					belt.Items = append(belt.Items, *itm)
				}
			}

			itemUnitPtr = uintptr(gd.Process.ReadUInt(itemUnitPtr+0x158, Uint64))
		}
	}

	// Link sockets to their base items
	for _, baseItem := range baseItems {
		numSockets, _ := baseItem.Item.Stats.FindStat(stat.NumSockets, 0)
		if numSockets.Value == 0 {
			continue
		}

		var matchingSockets []ItemData
		// Find all sockets that point to this base item
		for _, socketItem := range socketedItems {
			if socketItem.ParentPtr == baseItem.ThisPtr {
				matchingSockets = append(matchingSockets, socketItem)
			}
		}

		// Sort sockets by X position
		sort.Slice(matchingSockets, func(i, j int) bool {
			return matchingSockets[i].Item.Position.X < matchingSockets[j].Item.Position.X
		})

		// If number of sockets matches and positions are valid, assign them
		if len(matchingSockets) == numSockets.Value {
			// Verify positions are consecutive starting from 0
			validPositions := true
			for i, socket := range matchingSockets {
				if socket.Item.Position.X != i {
					validPositions = false
					break
				}
			}

			if validPositions {
				finalSockets := make([]data.Item, len(matchingSockets))
				for i, s := range matchingSockets {
					finalSockets[i] = *s.Item
				}
				baseItem.Item.Sockets = finalSockets
			}
		}
	}

	// Build final inventory
	inventory.AllItems = make([]data.Item, 0)
	for _, itm := range allItems {
		if itm.Location.LocationType != item.LocationSocket &&
			itm.Location.LocationType != item.LocationBelt {
			inventory.AllItems = append(inventory.AllItems, *itm)
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
	// Initial full and base stats extraction
	fullStats := gd.getStatsList(statsListExPtr + 0xA8)
	baseStats := gd.getStatsList(statsListExPtr + 0x30)

	// Flags and last stat list pointers
	flags := gd.Process.ReadUInt(statsListExPtr+0x1C, Uint64)
	lastStatsList := uintptr(gd.Process.ReadUInt(statsListExPtr+0x90, Uint64))

	// If the special flag isn't set, return the current base and full stats
	if (flags & 0x80000000) == 0 {
		return baseStats, fullStats
	}

	// Prepare for additional stats processing
	additionalStats := stat.Stats{}
	statListPtr := lastStatsList

	// Traverse the stat lists to accumulate additional stats
	for statListPtr != 0 {

		statListFlags := gd.Process.ReadUInt(statListPtr+0x1C, Uint64)

		// If we hit a condition where no further traversal is needed, break
		if (0x40 & statListFlags & 0xFFFFDFFF) != 0 {
			break
		}

		// Move to the previous stat list
		statListPtr = uintptr(gd.Process.ReadUInt(statListPtr+0x48, Uint64))
	}

	// If we found a valid previous stat list
	if statListPtr != 0 {
		additionalBaseStats := gd.getStatsList(statListPtr + 0x30)

		// Add only the additional stats that are not already present in fullStats
		for _, statM := range additionalBaseStats {
			if _, found := fullStats.FindStat(statM.ID, statM.Layer); !found {
				additionalStats = append(additionalStats, statM)
			}
		}
	}

	// Merge additional stats into fullStats (but not into baseStats)
	fullStats = append(fullStats, additionalStats...)

	// Handle base stats from potential modifiers and ensure they don't pollute baseStats
	statListPtr = lastStatsList

	for statListPtr != 0 {
		statListFlags := gd.Process.ReadUInt(statListPtr+0x1C, Uint64)

		if statListFlags != 0 {
			modifierBaseStats := gd.getStatsList(statListPtr + 0x30)

			for _, mStat := range modifierBaseStats {
				// Update fullStats
				if existingStat, found := fullStats.FindStat(mStat.ID, mStat.Layer); found {
					existingStat.Value = mStat.Value
				} else {
					fullStats = append(fullStats, mStat)
				}

				// Add to baseStats only if this stat qualifies as a base stat
				if (flags&0x80000000) != 0 && statListFlags == 0x80000000 {
					if _, found := baseStats.FindStat(mStat.ID, mStat.Layer); !found {
						baseStats = append(baseStats, mStat)
					}
				}
			}
		}

		// Move to the previous stat list
		statListPtr = uintptr(gd.Process.ReadUInt(statListPtr+0x48, Uint64))
	}

	return baseStats, fullStats
}

func setProperties(item *data.Item, flags uint32) {
	if 0x4000000&flags != 0 {
		item.IsRuneword = true
	}
	if 0x1000000&flags != 0 {
		item.IsNamed = true
	}
	if 0x400000&flags != 0 {
		item.Ethereal = true
	}
	if 0x20000&flags != 0 {
		item.IsStartItem = true
	}
	if 0x10000&flags != 0 {
		item.IsEar = true
	}
	if 0x2000&flags != 0 {
		item.InTradeOrStoreScreen = true // If item is sold to vendor or placed in trade with another player
	}
	if 0x800&flags != 0 {
		item.HasSockets = true
	}
	if 0x100&flags != 0 {
		item.IsBroken = true
	}
	if 0x10&flags != 0 {
		item.Identified = true
	}
	// Only jewels and runes, gems don't work
	if 0x8&flags != 0 {
		item.IsInSocket = true
	}
	if 0x1&flags != 0 {
		item.IsEquipped = true // Doesnt seem to work will need double check.
	}
}
