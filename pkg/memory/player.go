package memory

import (
	"encoding/binary"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/skill"
	"github.com/hectorgimenez/d2go/pkg/data/state"
)

func (gd *GameReader) GetRawPlayerUnits() RawPlayerUnits {
	rawPlayerUnits := make(RawPlayerUnits, 0)
	hover := gd.hoveredData()
	for i := 0; i < 128; i++ {
		unitOffset := gd.offset.UnitTable + uintptr(i*8)
		playerUnitAddr := gd.Process.moduleBaseAddressPtr + unitOffset
		playerUnit := uintptr(gd.Process.ReadUInt(playerUnitAddr, Uint64))
		for playerUnit > 0 {
			unitID := gd.Process.ReadUInt(playerUnit+0x08, Uint32)
			pInventory := playerUnit + 0x90
			inventoryAddr := uintptr(gd.Process.ReadUInt(pInventory, Uint64))

			pPath := playerUnit + 0x38
			pathAddress := uintptr(gd.Process.ReadUInt(pPath, Uint64))
			room1Ptr := uintptr(gd.Process.ReadUInt(pathAddress+0x20, Uint64))
			room2Ptr := uintptr(gd.Process.ReadUInt(room1Ptr+0x18, Uint64))
			levelPtr := uintptr(gd.Process.ReadUInt(room2Ptr+0x90, Uint64))
			levelNo := gd.Process.ReadUInt(levelPtr+0x1F8, Uint32)

			xPos := gd.Process.ReadUInt(pathAddress+0x02, Uint16)
			yPos := gd.Process.ReadUInt(pathAddress+0x06, Uint16)
			pUnitData := playerUnit + 0x10
			playerNameAddr := uintptr(gd.Process.ReadUInt(pUnitData, Uint64))
			name := gd.Process.ReadStringFromMemory(playerNameAddr, 0)

			expCharPtr := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+gd.offset.Expansion, Uint64))
			expChar := gd.Process.ReadUInt(expCharPtr+0x5C, Uint16)
			isMainPlayer := gd.Process.ReadUInt(inventoryAddr+0x30, Uint16)
			if expChar > 0 {
				isMainPlayer = gd.Process.ReadUInt(inventoryAddr+0x70, Uint16)
			}
			isCorpse := gd.Process.ReadUInt(playerUnit+0x1A6, Uint8)

			statsListExPtr := uintptr(gd.Process.ReadUInt(playerUnit+0x88, Uint64))
			states := gd.getStates(statsListExPtr)

			rawPlayerUnits = append(rawPlayerUnits, RawPlayerUnit{
				UnitID:       data.UnitID(unitID),
				Address:      playerUnit,
				Name:         name,
				IsMainPlayer: isMainPlayer > 0,
				IsCorpse:     isCorpse == 1 && inventoryAddr > 0 && xPos > 0 && yPos > 0,
				Area:         area.ID(levelNo),
				Position: data.Position{
					X: int(xPos),
					Y: int(yPos),
				},
				IsHovered: hover.IsHovered && hover.UnitID == data.UnitID(unitID) && hover.UnitType == 0,
				States:    states,
			})
			playerUnit = uintptr(gd.Process.ReadUInt(playerUnit+0x150, Uint64))
		}
	}

	return rawPlayerUnits
}

func (gd *GameReader) GetPlayerUnit(mainPlayerUnit RawPlayerUnit) data.PlayerUnit {
	// Get Stats
	statsListExPtr := uintptr(gd.Process.ReadUInt(mainPlayerUnit.Address+0x88, Uint64))
	baseStats := gd.getStatsList(statsListExPtr + 0x30)
	stats := gd.getStatsList(statsListExPtr + 0x88)

	// Skills
	skillListPtr := uintptr(gd.Process.ReadUInt(mainPlayerUnit.Address+0x100, Uint64))
	skills := gd.getSkills(skillListPtr)

	leftSkillPtr := gd.Process.ReadUInt(skillListPtr+0x08, Uint64)
	leftSkillTxtPtr := uintptr(gd.Process.ReadUInt(uintptr(leftSkillPtr), Uint64))
	leftSkillId := uintptr(gd.Process.ReadUInt(leftSkillTxtPtr, Uint16))

	rightSkillPtr := gd.Process.ReadUInt(skillListPtr+0x10, Uint64)
	rightSkillTxtPtr := uintptr(gd.Process.ReadUInt(uintptr(rightSkillPtr), Uint64))
	rightSkillId := uintptr(gd.Process.ReadUInt(rightSkillTxtPtr, Uint16))

	// Class
	class := data.Class(gd.Process.ReadUInt(mainPlayerUnit.Address+0x174, Uint32))

	availableWPs := make([]area.ID, 0)
	// Probably there is a better place to pick up those values, since this seems to be very tied to the UI
	wpList := gd.Process.ReadBytesFromMemory(gd.moduleBaseAddressPtr+0x21AD220, 0x48)
	for i := 0; i < 0x48; i = i + 8 {
		a := binary.LittleEndian.Uint32(wpList[i : i+4])
		available := binary.LittleEndian.Uint32(wpList[i+4 : i+8])
		if available == 1 || mainPlayerUnit.Area == area.ID(a) {
			availableWPs = append(availableWPs, area.ID(a))
		}
	}

	d := data.PlayerUnit{
		Name:               mainPlayerUnit.Name,
		ID:                 mainPlayerUnit.UnitID,
		Area:               mainPlayerUnit.Area,
		Position:           mainPlayerUnit.Position,
		Stats:              stats,
		BaseStats:          baseStats,
		Skills:             skills,
		States:             mainPlayerUnit.States,
		Class:              class,
		LeftSkill:          skill.ID(leftSkillId),
		RightSkill:         skill.ID(rightSkillId),
		AvailableWaypoints: availableWPs,
	}

	return d
}

func (gd *GameReader) getSkills(skillListPtr uintptr) map[skill.ID]skill.Points {
	skills := make(map[skill.ID]skill.Points)

	skillPtr := uintptr(gd.Process.ReadUInt(skillListPtr, Uint64))

	for skillPtr != 0 {
		skillTxtPtr := uintptr(gd.Process.ReadUInt(skillPtr, Uint64))
		skillTxt := uintptr(gd.Process.ReadUInt(skillTxtPtr, Uint16))
		lvl := gd.Process.ReadUInt(skillPtr+0x38, Uint16)
		quantity := gd.Process.ReadUInt(skillPtr+0x40, Uint16)
		charges := gd.Process.ReadUInt(skillPtr+0x48, Uint16)

		skills[skill.ID(skillTxt)] = skill.Points{
			Level:    lvl,
			Quantity: quantity,
			Charges:  charges,
		}

		skillPtr = uintptr(gd.Process.ReadUInt(skillPtr+0x08, Uint64))
	}

	return skills
}

func (gd *GameReader) getStates(statsListExPtr uintptr) []state.State {
	var states []state.State
	for i := 0; i < 6; i++ {
		offset := i * 4
		stateByte := gd.Process.ReadUInt(statsListExPtr+0xAD0+uintptr(offset), Uint32)

		offset = (32 * i) - 1
		states = append(states, calculateStates(stateByte, uint(offset))...)
	}

	return states
}

func calculateStates(stateFlag uint, offset uint) []state.State {
	var states []state.State
	if 0x00000001&stateFlag != 0 {
		states = append(states, state.State(1+offset))
	}
	if 0x00000002&stateFlag != 0 {
		states = append(states, state.State(2+offset))
	}
	if 0x00000004&stateFlag != 0 {
		states = append(states, state.State(3+offset))
	}
	if 0x00000008&stateFlag != 0 {
		states = append(states, state.State(4+offset))
	}
	if 0x00000010&stateFlag != 0 {
		states = append(states, state.State(5+offset))
	}
	if 0x00000020&stateFlag != 0 {
		states = append(states, state.State(6+offset))
	}
	if 0x00000040&stateFlag != 0 {
		states = append(states, state.State(7+offset))
	}
	if 0x00000080&stateFlag != 0 {
		states = append(states, state.State(8+offset))
	}
	if 0x00000100&stateFlag != 0 {
		states = append(states, state.State(9+offset))
	}
	if 0x00000200&stateFlag != 0 {
		states = append(states, state.State(10+offset))
	}
	if 0x00000400&stateFlag != 0 {
		states = append(states, state.State(11+offset))
	}
	if 0x00000800&stateFlag != 0 {
		states = append(states, state.State(12+offset))
	}
	if 0x00001000&stateFlag != 0 {
		states = append(states, state.State(13+offset))
	}
	if 0x00002000&stateFlag != 0 {
		states = append(states, state.State(14+offset))
	}
	if 0x00004000&stateFlag != 0 {
		states = append(states, state.State(15+offset))
	}
	if 0x00008000&stateFlag != 0 {
		states = append(states, state.State(16+offset))
	}
	if 0x00010000&stateFlag != 0 {
		states = append(states, state.State(17+offset))
	}
	if 0x00020000&stateFlag != 0 {
		states = append(states, state.State(18+offset))
	}
	if 0x00040000&stateFlag != 0 {
		states = append(states, state.State(19+offset))
	}
	if 0x00080000&stateFlag != 0 {
		states = append(states, state.State(20+offset))
	}
	if 0x00100000&stateFlag != 0 {
		states = append(states, state.State(21+offset))
	}
	if 0x00200000&stateFlag != 0 {
		states = append(states, state.State(22+offset))
	}
	if 0x00400000&stateFlag != 0 {
		states = append(states, state.State(23+offset))
	}
	if 0x00800000&stateFlag != 0 {
		states = append(states, state.State(24+offset))
	}
	if 0x01000000&stateFlag != 0 {
		states = append(states, state.State(25+offset))
	}
	if 0x02000000&stateFlag != 0 {
		states = append(states, state.State(26+offset))
	}
	if 0x04000000&stateFlag != 0 {
		states = append(states, state.State(27+offset))
	}
	if 0x08000000&stateFlag != 0 {
		states = append(states, state.State(28+offset))
	}
	if 0x10000000&stateFlag != 0 {
		states = append(states, state.State(29+offset))
	}
	if 0x20000000&stateFlag != 0 {
		states = append(states, state.State(30+offset))
	}
	if 0x40000000&stateFlag != 0 {
		states = append(states, state.State(31+offset))
	}
	if 0x80000000&stateFlag != 0 {
		states = append(states, state.State(32+offset))
	}

	return states
}
