package memory

import (
	"slices"
	"sort"
	"strings"

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

	// Process shared stash data
	stashPlayerUnits := make(map[uint]RawPlayerUnit)
	stashPlayerUnitOrder := make([]uint, 0, 3) // Pre-allocate with expected size
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
	stashedGold := [4]int{mainPlayerStashedGold.Value, 0, 0, 0}

	for i, puKey := range stashPlayerUnitOrder {
		if i > 2 {
			break
		}
		if stashGold, found := stashPlayerUnits[puKey].BaseStats.FindStat(stat.StashGold, 0); found {
			stashedGold[i+1] = stashGold.Value
		}
	}

	inventory := data.Inventory{
		Gold:        inventoryGold.Value,
		StashedGold: stashedGold,
	}
	belt := data.Belt{}

	type socketInfo struct {
		item     *data.Item
		position int // Store X position
	}
	socketedItemsMap := make(map[data.UnitID][]socketInfo, 120) // Up to ~120 items could be socketable
	baseItemsMap := make(map[data.UnitID]*data.Item, 120)       // Same number of potential base items
	allItems := make([]*data.Item, 0, 480)                      // max capacity: 400 (stashes) + 40 (inv) + 12 (cube) + 12 (equipped) + 16 (belt)

	// Pre-allocate buffers for repeated use
	var itemDataBuffer = make([]byte, 144)
	var unitDataBuffer = make([]byte, 144)
	var pathBuffer = make([]byte, 144)

	// Process all items in a single pass
	for i := 0; i < 128; i++ {
		itemOffset := 8 * i
		itemUnitPtr := uintptr(ReadUIntFromBuffer(unitTableBuffer, uint(itemOffset), Uint64))

		for itemUnitPtr > 0 {
			nextItemPtr := uintptr(gd.Process.ReadUInt(itemUnitPtr+0x158, Uint64))

			// Read basic item data into pre-allocated buffer
			if err := gd.Process.ReadIntoBuffer(itemUnitPtr, itemDataBuffer); err != nil {
				itemUnitPtr = nextItemPtr
				continue
			}

			itemType := ReadUIntFromBuffer(itemDataBuffer, 0x00, Uint32)

			// Skip non-item entries early
			if itemType != 4 {
				itemUnitPtr = nextItemPtr
				continue
			}

			txtFileNo := ReadUIntFromBuffer(itemDataBuffer, 0x04, Uint32)
			unitID := ReadUIntFromBuffer(itemDataBuffer, 0x08, Uint32)

			// itemLoc = 0 in inventory, 1 equipped, 2 in belt, 3 on ground, 4 cursor, 5 dropping, 6 socketed
			itemLoc := ReadUIntFromBuffer(itemDataBuffer, 0x0C, Uint32)

			unitDataPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x10, Uint64))
			if err := gd.Process.ReadIntoBuffer(unitDataPtr, unitDataBuffer); err != nil {
				itemUnitPtr = nextItemPtr
				continue
			}

			flags := ReadUIntFromBuffer(unitDataBuffer, 0x18, Uint32)
			invPage := ReadUIntFromBuffer(unitDataBuffer, 0x55, Uint8)
			itemQuality := ReadUIntFromBuffer(unitDataBuffer, 0x00, Uint32)
			itemOwnerNPC := ReadUIntFromBuffer(unitDataBuffer, 0x0C, Uint32)

			// Link to uniqueitems.txt, setitems.txt
			txtUniqueSet := int32(gd.Process.ReadUInt(unitDataPtr+0x34, Uint32))

			pathPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x38, Uint64))
			if err := gd.Process.ReadIntoBuffer(pathPtr, pathBuffer); err != nil {
				itemUnitPtr = nextItemPtr
				continue
			}
			// Item coordinates (X, Y)
			itemX := ReadUIntFromBuffer(pathBuffer, 0x10, Uint16)
			itemY := ReadUIntFromBuffer(pathBuffer, 0x14, Uint16)

			// Create item structure
			itm := &data.Item{
				ID:      int(txtFileNo),
				UnitID:  data.UnitID(unitID),
				Name:    item.GetNameByEnum(txtFileNo),
				Quality: item.Quality(itemQuality),
				Position: data.Position{
					X: int(itemX),
					Y: int(itemY),
				},
				IsHovered:   hover.IsHovered && hover.UnitType == 4 && hover.UnitID == data.UnitID(unitID),
				Sockets:     make([]data.Item, 0),
				UniqueSetID: txtUniqueSet,
			}

			// Set item properties
			setProperties(itm, uint32(flags))

			// Read rare affixes
			rarePrefix := int16(gd.Process.ReadUInt(unitDataPtr+0x42, Uint16))
			rareSuffix := int16(gd.Process.ReadUInt(unitDataPtr+0x44, Uint16))
			//autoAffix := int16(gd.Process.ReadUInt(unitDataPtr+0x46, Uint16))

			// Read magic affixes
			var prefixes [3]int16
			var suffixes [3]int16
			for i := 0; i < 3; i++ {
				prefixes[i] = int16(gd.Process.ReadUInt(unitDataPtr+0x48+uintptr(i*2), Uint16))
				suffixes[i] = int16(gd.Process.ReadUInt(unitDataPtr+0x4E+uintptr(i*2), Uint16))
			}

			itm.Affixes = data.ItemAffixes{
				Rare: struct {
					Prefix int16
					Suffix int16
				}{
					Prefix: rarePrefix,
					Suffix: rareSuffix,
				},
				Magic: struct {
					Prefixes [3]int16
					Suffixes [3]int16
				}{
					Prefixes: prefixes,
					Suffixes: suffixes,
				},
			}

			maxAffixReq := 0
			if itm.Identified {
				switch itm.Quality {
				case item.QualityUnique:
					// find matching item (uniqueitems.txt)
					for _, uniqueInfo := range item.UniqueItems {
						if uniqueInfo.ID == int(txtUniqueSet) {
							itm.IdentifiedName = uniqueInfo.Name
							itm.LevelReq = uniqueInfo.LevelReq
							break
						}
					}
				case item.QualitySet:
					// find matching item (setitems.txt)
					for setItemName, setItemInfo := range item.SetItems {
						if setItemInfo.ID == int(txtUniqueSet) {
							itm.IdentifiedName = string(setItemName)
							itm.LevelReq = setItemInfo.LevelReq
							break
						}
					}
				case item.QualityRare:
					// Set item name from rare affixes
					if prefix, exists := item.RarePrefixDesc[int(rarePrefix)]; exists {
						if suffix, exists := item.RareSuffixDesc[int(rareSuffix)]; exists {
							itm.IdentifiedName = prefix.Name + " " + suffix.Name
						}
					}
					// Get level requirements from magic affixes
					for _, prefixID := range prefixes {
						if prefix, exists := item.MagicPrefixDesc[int(prefixID)]; exists && prefixID != 0 {
							if prefix.LevelReq > maxAffixReq {
								maxAffixReq = prefix.LevelReq
							}
						}
					}
					for _, suffixID := range suffixes {
						if suffix, exists := item.MagicSuffixDesc[int(suffixID)]; exists && suffixID != 0 {
							if suffix.LevelReq > maxAffixReq {
								maxAffixReq = suffix.LevelReq
							}
						}
					}
				case item.QualityMagic:
					var prefixParts []string
					var suffixParts []string

					// Get all prefixes
					for _, prefixID := range prefixes {
						if prefix, exists := item.MagicPrefixDesc[int(prefixID)]; exists && prefixID != 0 {
							prefixParts = append(prefixParts, prefix.Name)
							if prefix.LevelReq > maxAffixReq {
								maxAffixReq = prefix.LevelReq
							}
						}
					}

					// Get all suffixes
					for _, suffixID := range suffixes {
						if suffix, exists := item.MagicSuffixDesc[int(suffixID)]; exists && suffixID != 0 {
							suffixParts = append(suffixParts, suffix.Name)
							if suffix.LevelReq > maxAffixReq {
								maxAffixReq = suffix.LevelReq
							}
						}
					}

					// Construct name: prefixes + base name + suffixes
					var nameParts []string
					if len(prefixParts) > 0 {
						nameParts = append(nameParts, prefixParts...)
					}
					nameParts = append(nameParts, string(itm.Name))
					if len(suffixParts) > 0 {
						nameParts = append(nameParts, suffixParts...)
					}
					itm.IdentifiedName = strings.Join(nameParts, " ")
				}
			}

			// Set runeword name if applicable
			if itm.IsRuneword {
				if runeword, exists := item.RunewordIDMap[prefixes[0]]; exists {
					itm.RunewordName = runeword
				}
			}
			// Determine item location
			location := item.LocationUnknown
			switch itemLoc {
			case 0:
				if itemOwnerNPC == 2 || itemOwnerNPC == uint(stashPlayerUnits[stashPlayerUnitOrder[0]].UnitID) {
					location = item.LocationSharedStash
					invPage = 1
				} else if itemOwnerNPC == 3 || itemOwnerNPC == uint(stashPlayerUnits[stashPlayerUnitOrder[1]].UnitID) {
					location = item.LocationSharedStash
					invPage = 2
				} else if itemOwnerNPC == 4 || itemOwnerNPC == uint(stashPlayerUnits[stashPlayerUnitOrder[2]].UnitID) {
					location = item.LocationSharedStash
					invPage = 3
				} else if 0x00002000&flags != 0 && itemOwnerNPC == 4294967295 {
					location = item.LocationVendor
				} else if data.UnitID(itemOwnerNPC) == mainPlayer.UnitID || itemOwnerNPC == 1 {
					if invPage == 0 {
						location = item.LocationInventory
					} else if invPage == 3 {
						location = item.LocationCube
						invPage = 0
					} else {
						location = item.LocationStash
						invPage = 0
					}
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

			// Set body location if equipped
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
			if location == item.LocationUnknown {
				itemUnitPtr = nextItemPtr
				continue
			}

			// Read item stats
			statsListExPtr := uintptr(ReadUIntFromBuffer(itemDataBuffer, 0x88, Uint64))
			itm.BaseStats, itm.Stats = gd.getItemStats(statsListExPtr)

			// Process socket information
			if location == item.LocationSocket {
				if itm.Desc().Code == "jew" {
					// Base requirement for jewels
					itm.LevelReq = item.Desc[itm.ID].RequiredLevel

					// For magic/rare jewels, check affixes
					if itm.Quality == item.QualityMagic || itm.Quality == item.QualityRare {
						itm.LevelReq = updateMaxReqFromAffixes(itm.LevelReq, itm.Affixes)
					}
					// Rainbow facets
				} else if itm.Quality == item.QualityUnique {
					for _, uniqueInfo := range item.UniqueItems {
						if uniqueInfo.Code == itm.Desc().Code {
							itm.LevelReq = uniqueInfo.LevelReq
							break
						}
					}
				} else {
					// Normal socketed items (runes,gems) just use base requirement
					itm.LevelReq = item.Desc[itm.ID].RequiredLevel
				}
				itemExtraData := uintptr(gd.Process.ReadUInt(unitDataPtr+0xA0, Uint64))
				if itemExtraData != 0 {
					parentInfoPtr := uintptr(gd.Process.ReadUInt(itemExtraData+0x08, Uint64))
					if parentInfoPtr != 0 {
						// Read parent unit ID directly from base item memory structure
						if err := gd.Process.ReadIntoBuffer(parentInfoPtr, itemDataBuffer); err == nil {
							parentUnitID := data.UnitID(ReadUIntFromBuffer(itemDataBuffer, 0x08, Uint32))
							socketedItemsMap[parentUnitID] = append(socketedItemsMap[parentUnitID], socketInfo{
								item:     itm,
								position: itm.Position.X,
							})
						}
					}
				}
			} else {
				// Check if item has sockets
				if numSockets, _ := itm.Stats.FindStat(stat.NumSockets, 0); numSockets.Value > 0 {
					baseItemsMap[itm.UnitID] = itm
				}
			}

			// Add to appropriate collections
			if location == item.LocationBelt {
				belt.Items = append(belt.Items, *itm)
			} else if location != item.LocationSocket {
				allItems = append(allItems, itm)
			}

			itemUnitPtr = nextItemPtr
		}
	}

	// Link sockets to base items
	for baseUnitID, baseItem := range baseItemsMap {
		numSockets, _ := baseItem.Stats.FindStat(stat.NumSockets, 0)
		if numSockets.Value == 0 {
			continue
		}

		// Get socket list for this base item
		sockets := socketedItemsMap[baseUnitID]
		if len(sockets) != numSockets.Value {
			continue
		}

		// Sort by position if we have multiple sockets
		if len(sockets) > 1 {
			sort.Slice(sockets, func(i, j int) bool {
				return sockets[i].position < sockets[j].position
			})
		}

		// Validate socket positions and build final list in one pass
		baseItem.Sockets = make([]data.Item, 0, numSockets.Value)
		for i, socket := range sockets {
			if socket.position != i {
				baseItem.Sockets = nil // Reset if positions are invalid
				break
			}
			baseItem.Sockets = append(baseItem.Sockets, *socket.item)
		}
	}

	// Build final inventory
	inventory.AllItems = make([]data.Item, len(allItems))
	for i, itm := range allItems {
		baseDesc := item.Desc[itm.ID]
		maxReq := calculateItemLevelReq(itm, baseDesc)

		// Check magic/rare affixes
		if itm.Identified && (itm.Quality == item.QualityMagic || itm.Quality == item.QualityRare) {
			maxReq = updateMaxReqFromAffixes(maxReq, itm.Affixes)
		}

		// Check socketed items
		for _, socketItem := range itm.Sockets {
			socketBaseDesc := item.Desc[socketItem.ID]
			socketReq := calculateItemLevelReq(&socketItem, socketBaseDesc)

			if socketReq > maxReq {
				maxReq = socketReq
			}

			// Check affixes on magic socketed items
			if socketItem.Quality == item.QualityMagic {
				maxReq = updateMaxReqFromAffixes(maxReq, socketItem.Affixes)
			}
		}

		itm.LevelReq = maxReq
		// Update stat 92 (LevelRequire) so we can use with .nip
		for it, s := range itm.Stats {
			if s.ID == stat.LevelRequire {
				itm.Stats[it].Value = maxReq
				break
			}
		}
		inventory.AllItems[i] = *itm
	}

	// Sort items by distance if needed, using pre-calculated distances
	if len(inventory.AllItems) > 0 {
		// Pre-calculate distances to avoid recalculating in sort comparisons
		distances := make([]int, len(inventory.AllItems))
		for i, invItem := range inventory.AllItems {
			distances[i] = utils.DistanceFromPoint(mainPlayer.Position, invItem.Position)
		}

		sort.SliceStable(inventory.AllItems, func(i, j int) bool {
			return distances[i] < distances[j]
		})
	}

	inventory.Belt = belt
	return inventory
}

func (gd *GameReader) getItemStats(statsListExPtr uintptr) (stat.Stats, stat.Stats) {
	// Initial full and base stats extraction
	fullStats := gd.getStatsList(statsListExPtr + 0xA8)
	baseStats := gd.getStatsList(statsListExPtr + 0x30)

	// Create empty LevelRequire stat .We will update it from inventory
	if _, found := fullStats.FindStat(stat.LevelRequire, 0); !found {
		fullStats = append(fullStats, stat.Data{
			ID:    stat.LevelRequire,
			Value: 0,
			Layer: 0,
		})
	}

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
		item.HasBeenEquipped = true // Item was recently equipped (same game)
	}
}

func updateMaxReqFromAffixes(currentMax int, affixes data.ItemAffixes) int {
	maxReq := currentMax
	for _, prefixID := range affixes.Magic.Prefixes {
		if prefix, exists := item.MagicPrefixDesc[int(prefixID)]; exists && prefixID != 0 {
			if prefix.LevelReq > maxReq {
				maxReq = prefix.LevelReq
			}
		}
	}
	for _, suffixID := range affixes.Magic.Suffixes {
		if suffix, exists := item.MagicSuffixDesc[int(suffixID)]; exists && suffixID != 0 {
			if suffix.LevelReq > maxReq {
				maxReq = suffix.LevelReq
			}
		}
	}
	return maxReq
}
func calculateItemLevelReq(itm *data.Item, baseDesc item.Description) int {
	// Start with base item's requirement
	maxReq := baseDesc.RequiredLevel

	if itm.Quality == item.QualityUnique || itm.Quality == item.QualitySet {
		itemCode := baseDesc.Code
		normalCode := baseDesc.NormalCode
		exceptionalCode := baseDesc.UberCode
		eliteCode := baseDesc.UltraCode

		// Find the original item tier by looking up the unique/set item
		var originalCode string
		if itm.Quality == item.QualityUnique {
			for _, uniqueInfo := range item.UniqueItems {
				if uniqueInfo.ID == int(itm.UniqueSetID) {
					originalCode = uniqueInfo.Code
					break
				}
			}
		} else {
			for _, setItemInfo := range item.SetItems {
				if setItemInfo.ID == int(itm.UniqueSetID) {
					originalCode = setItemInfo.Code
					break
				}
			}
		}

		if itemCode == eliteCode {
			if originalCode == normalCode {
				// Double upgraded: Normal -> Exceptional -> Elite
				maxReq += 12 // +5 for exceptional, +7 for elite
			} else if originalCode == exceptionalCode {
				// Single upgraded: Exceptional -> Elite
				maxReq += 7
			}
		} else if itemCode == exceptionalCode && originalCode == normalCode {
			// Single upgraded: Normal -> Exceptional
			maxReq += 5
		}

		// Check unique/set specific level requirement
		if itm.Identified {
			if itm.Quality == item.QualityUnique {
				for _, uniqueInfo := range item.UniqueItems {
					if uniqueInfo.ID == int(itm.UniqueSetID) {
						if uniqueInfo.LevelReq > maxReq {
							maxReq = uniqueInfo.LevelReq
						}
						break
					}
				}
			} else { // Set item
				for _, setItemInfo := range item.SetItems {
					if setItemInfo.ID == int(itm.UniqueSetID) {
						if setItemInfo.LevelReq > maxReq {
							maxReq = setItemInfo.LevelReq
						}
						break
					}
				}
			}
		}
	}

	return maxReq
}
