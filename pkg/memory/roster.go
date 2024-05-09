package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) getRoster(rawPlayerUnits RawPlayerUnits) (roster []data.RosterMember) {
	partyStruct := uintptr(gd.Process.ReadUInt(gd.Process.moduleBaseAddressPtr+gd.offset.RosterOffset, Uint64))

	// We skip the first position because it's the main player, and we already have the information (+0x148 is the next party member)
	partyStruct = uintptr(gd.Process.ReadUInt(partyStruct+0x148, Uint64))
	for partyStruct > 0 {
		name := gd.Process.ReadStringFromMemory(partyStruct, 16)
		a := area.ID(gd.Process.ReadUInt(partyStruct+0x5C, Uint32))

		xPos := int(gd.Process.ReadUInt(partyStruct+0x60, Uint32))
		yPos := int(gd.Process.ReadUInt(partyStruct+0x64, Uint32))

		// When the player is in town, roster data is not updated, so we need to get the area from the player unit that match the same name
		for _, pu := range rawPlayerUnits {
			if pu.Name == name {
				xPos = pu.Position.X
				yPos = pu.Position.Y
				a = pu.Area
				break
			}
		}

		roster = append(roster, data.RosterMember{
			Name:     name,
			Area:     a,
			Position: data.Position{X: xPos, Y: yPos},
		})
		partyStruct = uintptr(gd.Process.ReadUInt(partyStruct+0x148, Uint64))
	}

	mainPlayerUnit := rawPlayerUnits.GetMainPlayer()

	return append([]data.RosterMember{{
		Name:     mainPlayerUnit.Name,
		Area:     mainPlayerUnit.Area,
		Position: mainPlayerUnit.Position,
	}}, roster...)
}
