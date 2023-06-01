package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/skill"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/d2go/pkg/data/state"
)

func (gd *GameReader) GetPlayerUnitPtr(roster data.Roster) (playerUnitPtr uintptr, corpse data.Corpse) {
	for i := 0; i < 128; i++ {
		unitOffset := gd.offset.UnitTable + uintptr(i*8)
		playerUnitAddr := gd.Process.moduleBaseAddressPtr + unitOffset
		playerUnit := uintptr(gd.Process.ReadUInt(playerUnitAddr, Uint64))
		for playerUnit > 0 {
			pInventory := playerUnit + 0x90
			inventoryAddr := uintptr(gd.Process.ReadUInt(pInventory, Uint64))

			pPath := playerUnit + 0x38
			pathAddress := uintptr(gd.Process.ReadUInt(pPath, Uint64))
			xPos := gd.Process.ReadUInt(pathAddress+0x02, Uint16)
			yPos := gd.Process.ReadUInt(pathAddress+0x06, Uint16)

			// Only current player has inventory
			if inventoryAddr > 0 && xPos > 0 && yPos > 0 {
				expCharPtr := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+gd.offset.Expansion, Uint64))
				expChar := gd.Process.ReadUInt(expCharPtr+0x5C, Uint16)
				baseCheck := gd.Process.ReadUInt(inventoryAddr+0x30, Uint16)
				if expChar > 0 {
					baseCheck = gd.Process.ReadUInt(inventoryAddr+0x70, Uint16)
				}

				isCorpse := gd.Process.ReadUInt(playerUnit+0x1A6, Uint8)
				if isCorpse == 1 {
					unitID := gd.Process.ReadUInt(playerUnit+0x08, Uint32)
					hover := gd.hoveredData()
					corpse = data.Corpse{
						Found:     true,
						IsHovered: hover.IsHovered && hover.UnitID == data.UnitID(unitID) && hover.UnitType == 0,
						Position: data.Position{
							X: int(xPos),
							Y: int(yPos),
						},
					}
				}

				if baseCheck > 0 {
					playerUnitPtr = playerUnit
				} else {
					pUnitData := playerUnit + 0x10
					playerNameAddr := uintptr(gd.Process.ReadUInt(pUnitData, Uint64))
					name := gd.Process.ReadStringFromMemory(playerNameAddr, 0)
					for k, rm := range roster {
						if name != rm.Name {
							continue
						}

						roster[k] = data.RosterMember{
							Name: name,
							Area: rm.Area,
							Position: data.Position{
								X: int(xPos),
								Y: int(yPos),
							},
						}
					}
				}
			}

			playerUnit = uintptr(gd.Process.ReadUInt(playerUnit+0x150, Uint64))
		}
	}

	return
}

func (gd *GameReader) GetPlayerUnit(playerUnit uintptr) data.PlayerUnit {
	unitID := gd.Process.ReadUInt(playerUnit+0x08, Uint32)

	// Read X and Y Positions
	pPath := playerUnit + 0x38
	pathAddress := uintptr(gd.Process.ReadUInt(pPath, Uint64))
	xPos := gd.Process.ReadUInt(pathAddress+0x02, Uint16)
	yPos := gd.Process.ReadUInt(pathAddress+0x06, Uint16)

	// Player name
	pUnitData := playerUnit + 0x10
	playerNameAddr := uintptr(gd.Process.ReadUInt(pUnitData, Uint64))
	name := gd.Process.ReadStringFromMemory(playerNameAddr, 0)

	// Get Stats
	statsListExPtr := uintptr(gd.Process.ReadUInt(playerUnit+0x88, Uint64))
	statPtr := gd.Process.ReadUInt(statsListExPtr+0x30, Uint64)
	statCount := gd.Process.ReadUInt(statsListExPtr+0x38, Uint64)

	stats := map[stat.ID]int{}
	for j := 0; j < int(statCount); j++ {
		statOffset := uintptr(statPtr) + 0x2 + uintptr(j*8)
		statNumber := gd.Process.ReadUInt(statOffset, Uint16)
		statValue := gd.Process.ReadUInt(statOffset+0x02, Uint32)

		switch stat.ID(statNumber) {
		case stat.Life,
			stat.MaxLife,
			stat.Mana,
			stat.MaxMana:
			stats[stat.ID(statNumber)] = int(uint32(statValue) >> 8)
		default:
			stats[stat.ID(statNumber)] = int(statValue)
		}
	}

	// States (Buff, Debuff, Auras)
	states := gd.getStates(statsListExPtr)

	// Skills
	skills := gd.getSkills(playerUnit + 0x100)

	// Class
	class := data.Class(gd.Process.ReadUInt(playerUnit+0x174, Uint32))

	// Level number
	pathPtr := uintptr(gd.Process.ReadUInt(playerUnit+0x38, Uint64))
	room1Ptr := uintptr(gd.Process.ReadUInt(pathPtr+0x20, Uint64))
	room2Ptr := uintptr(gd.Process.ReadUInt(room1Ptr+0x18, Uint64))
	levelPtr := uintptr(gd.Process.ReadUInt(room2Ptr+0x90, Uint64))
	levelNo := gd.Process.ReadUInt(levelPtr+0x1F8, Uint32)

	return data.PlayerUnit{
		Name: name,
		ID:   data.UnitID(unitID),
		Area: area.Area(levelNo),
		Position: data.Position{
			X: int(xPos),
			Y: int(yPos),
		},
		Stats:  stats,
		Skills: skills,
		States: states,
		Class:  class,
	}
}

func (gd *GameReader) getSkills(skillsPtr uintptr) map[skill.Skill]int {
	skills := make(map[skill.Skill]int)
	skillListPtr := uintptr(gd.Process.ReadUInt(skillsPtr, Uint64))

	skillPtr := uintptr(gd.Process.ReadUInt(skillListPtr, Uint64))

	for skillPtr != 0 {
		skillTxtPtr := uintptr(gd.Process.ReadUInt(skillPtr, Uint64))
		skillTxt := uintptr(gd.Process.ReadUInt(skillTxtPtr, Uint16))
		skillLvl := gd.Process.ReadUInt(skillPtr+0x38, Uint16)

		skills[skill.Skill(skillTxt)] = int(skillLvl)

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
