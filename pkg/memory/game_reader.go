package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

type GameReader struct {
	offset Offset
	Process
}

func NewGameReader(process Process) *GameReader {
	return &GameReader{
		offset:  calculateOffsets(process),
		Process: process,
	}
}

var previousData *data.Data

func (gd *GameReader) GetData() data.Data {
	if gd.offset.UnitTable == 0 {
		gd.offset = calculateOffsets(gd.Process)
	}

	roster := gd.getRoster()
	playerUnitPtr, corpse := gd.GetPlayerUnitPtr(roster)

	pu := gd.GetPlayerUnit(playerUnitPtr)

	d := data.Data{
		Corpse:     corpse,
		Monsters:   gd.Monsters(pu.Position),
		PlayerUnit: pu,
		Items:      gd.Items(pu.Position),
		Objects:    gd.Objects(pu.Position),
		OpenMenus:  gd.openMenus(),
		Roster:     roster,
	}

	if playerUnitPtr == 0 {
		return *previousData
	}

	previousData = &d

	return d
}

func (gd *GameReader) InGame() bool {
	pu, _ := gd.GetPlayerUnitPtr([]data.RosterMember{})

	return pu > 0
}

func (gd *GameReader) openMenus() data.OpenMenus {
	uiBase := gd.Process.moduleBaseAddressPtr + gd.offset.UI - 0xA

	buffer := gd.Process.ReadBytesFromMemory(uiBase, 0x169)

	isMapShown := gd.Process.ReadUInt(gd.Process.moduleBaseAddressPtr+gd.offset.UI, Uint8)

	return data.OpenMenus{
		Inventory:     buffer[0x01] != 0,
		LoadingScreen: buffer[0x168] != 0,
		NPCInteract:   buffer[0x08] != 0,
		NPCShop:       buffer[0x0B] != 0,
		Stash:         buffer[0x18] != 0,
		Waypoint:      buffer[0x13] != 0,
		MapShown:      isMapShown != 0,
	}
}

func (gd *GameReader) hoveredData() (hoveredUnitID uint, hoveredType uint, isHovered bool) {
	hoverAddressPtr := gd.Process.moduleBaseAddressPtr + gd.offset.Hover
	hoverBuffer := gd.Process.ReadBytesFromMemory(hoverAddressPtr, 12)
	isUnitHovered := ReadUIntFromBuffer(hoverBuffer, 0, Uint16)
	if isUnitHovered > 0 {
		hoveredType = ReadUIntFromBuffer(hoverBuffer, 0x04, Uint32)
		hoveredUnitID = ReadUIntFromBuffer(hoverBuffer, 0x08, Uint32)

		return hoveredUnitID, hoveredType, true
	}

	return 0, 0, false
}

func getStatData(statEnum, statValue uint) (stat.Stat, int) {
	value := int(statValue)
	switch stat.Stat(statEnum) {
	case stat.Life,
		stat.MaxLife,
		stat.Mana,
		stat.MaxMana,
		stat.Stamina,
		stat.LifePerLevel,
		stat.ManaPerLevel:
		value = int(statValue >> 8)
	case stat.ColdLength,
		stat.PoisonLength:
		value = int(statValue / 25)
	}

	return stat.Stat(statEnum), value
}

func setProperties(item *data.Item, flags uint32) {
	if 0x00400000&flags != 0 {
		item.Ethereal = true
	}
	if 0x00000010&flags != 0 {
		item.Identified = true
	}
	if 0x00002000&flags != 0 {
		item.IsVendor = true
	}
}
