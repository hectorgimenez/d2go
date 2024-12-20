package memory

import (
	"sort"

	"github.com/hectorgimenez/d2go/pkg/data/mode"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/d2go/pkg/utils"
)

func (gd *GameReader) Monsters(playerPosition data.Position, hover data.HoverData) data.Monsters {
	baseAddr := gd.Process.moduleBaseAddressPtr + gd.offset.UnitTable + 1024
	unitTableBuffer := gd.Process.ReadBytesFromMemory(baseAddr, 128*8)

	monsters := data.Monsters{}
	for i := 0; i < 128; i++ {
		monsterOffset := 8 * i
		monsterUnitPtr := uintptr(ReadUIntFromBuffer(unitTableBuffer, uint(monsterOffset), Uint64))
		for monsterUnitPtr > 0 {
			// Quick corpse check first
			isCorpse := gd.Process.ReadUInt(monsterUnitPtr+0x1AE, Uint8)
			if isCorpse != 0 {
				monsterUnitPtr = uintptr(gd.Process.ReadUInt(monsterUnitPtr+0x158, Uint64))
				continue
			}

			monsterDataBuffer := gd.Process.ReadBytesFromMemory(monsterUnitPtr, 144)
			txtFileNo := ReadUIntFromBuffer(monsterDataBuffer, 0x04, Uint32)
			unitID := ReadUIntFromBuffer(monsterDataBuffer, 0x08, Uint32)

			// Get stats early for filtering
			statsListExPtr := uintptr(ReadUIntFromBuffer(monsterDataBuffer, 0x88, Uint64))
			statPtr := uintptr(gd.Process.ReadUInt(statsListExPtr+0x30, Uint64))
			statCount := gd.Process.ReadUInt(statsListExPtr+0x38, Uint64)
			stats := gd.getMonsterStats(statCount, statPtr)

			if !gd.shouldBeIgnored(txtFileNo) || stats[stat.Experience] > 0 {
				monsterMode := mode.NpcMode(gd.Process.ReadUInt(monsterUnitPtr+0x0c, Uint32))
				unitDataPtr := uintptr(ReadUIntFromBuffer(monsterDataBuffer, 0x10, Uint64))
				flag := gd.Process.ReadBytesFromMemory(unitDataPtr+0x1A, Uint8)[0]
				//unitDataBuffer := gd.Process.ReadBytesFromMemory(unitDataPtr, 144)

				// Coordinates (X, Y)
				pathPtr := uintptr(gd.Process.ReadUInt(monsterUnitPtr+0x38, Uint64))
				posX := gd.Process.ReadUInt(pathPtr+0x02, Uint16)
				posY := gd.Process.ReadUInt(pathPtr+0x06, Uint16)

				states := gd.GetStates(statsListExPtr)

				monsters = append(monsters, data.Monster{
					UnitID:    data.UnitID(unitID),
					Name:      npc.ID(int(txtFileNo)),
					IsHovered: hover.IsHovered && hover.UnitType == 1 && hover.UnitID == data.UnitID(unitID),
					Position: data.Position{
						X: int(posX),
						Y: int(posY),
					},
					Stats:  stats,
					Type:   getMonsterType(flag),
					States: states,
					Mode:   monsterMode,
				})
			}

			monsterUnitPtr = uintptr(gd.Process.ReadUInt(monsterUnitPtr+0x158, Uint64))
		}
	}

	if len(monsters) > 0 {
		sort.SliceStable(monsters, func(i, j int) bool {
			distanceI := utils.DistanceFromPoint(playerPosition, monsters[i].Position)
			distanceJ := utils.DistanceFromPoint(playerPosition, monsters[j].Position)
			return distanceI < distanceJ
		})
	}

	return monsters
}

func (gd *GameReader) Corpses(playerPosition data.Position, hover data.HoverData) data.Monsters {
	baseAddr := gd.Process.moduleBaseAddressPtr + gd.offset.UnitTable + 1024
	unitTableBuffer := gd.Process.ReadBytesFromMemory(baseAddr, 128*8)

	corpses := data.Monsters{}

	for i := 0; i < 128; i++ {
		monsterOffset := 8 * i
		monsterUnitPtr := uintptr(ReadUIntFromBuffer(unitTableBuffer, uint(monsterOffset), Uint64))
		for monsterUnitPtr > 0 {
			isCorpse := gd.Process.ReadUInt(monsterUnitPtr+0x1AE, Uint8)
			if isCorpse == 0 {
				monsterUnitPtr = uintptr(gd.Process.ReadUInt(monsterUnitPtr+0x158, Uint64))
				continue
			}

			monsterDataBuffer := gd.Process.ReadBytesFromMemory(monsterUnitPtr, 144)
			txtFileNo := ReadUIntFromBuffer(monsterDataBuffer, 0x04, Uint32)

			statsListExPtr := uintptr(ReadUIntFromBuffer(monsterDataBuffer, 0x88, Uint64))
			statPtr := uintptr(gd.Process.ReadUInt(statsListExPtr+0x30, Uint64))
			statCount := gd.Process.ReadUInt(statsListExPtr+0x38, Uint64)
			stats := gd.getMonsterStats(statCount, statPtr)

			if !gd.shouldBeIgnored(txtFileNo) || stats[stat.Experience] > 0 {
				unitID := ReadUIntFromBuffer(monsterDataBuffer, 0x08, Uint32)
				unitDataPtr := uintptr(ReadUIntFromBuffer(monsterDataBuffer, 0x10, Uint64))
				flag := gd.Process.ReadBytesFromMemory(unitDataPtr+0x1A, Uint8)[0]

				pathPtr := uintptr(gd.Process.ReadUInt(monsterUnitPtr+0x38, Uint64))
				posX := gd.Process.ReadUInt(pathPtr+0x02, Uint16)
				posY := gd.Process.ReadUInt(pathPtr+0x06, Uint16)

				states := gd.GetStates(statsListExPtr)

				corpses = append(corpses, data.Monster{
					UnitID:    data.UnitID(unitID),
					Name:      npc.ID(int(txtFileNo)),
					IsHovered: hover.IsHovered && hover.UnitType == 1 && hover.UnitID == data.UnitID(unitID),
					Position: data.Position{
						X: int(posX),
						Y: int(posY),
					},
					Stats:  stats,
					Type:   getMonsterType(flag),
					States: states,
				})
			}

			monsterUnitPtr = uintptr(gd.Process.ReadUInt(monsterUnitPtr+0x158, Uint64))
		}
	}

	if len(corpses) > 0 {
		sort.SliceStable(corpses, func(i, j int) bool {
			distanceI := utils.DistanceFromPoint(playerPosition, corpses[i].Position)
			distanceJ := utils.DistanceFromPoint(playerPosition, corpses[j].Position)
			return distanceI < distanceJ
		})
	}

	return corpses
}

func getMonsterType(typeFlag byte) data.MonsterType {
	switch typeFlag {
	case 10:
		return data.MonsterTypeSuperUnique
	case 1 << 2, 12:
		return data.MonsterTypeChampion
	case 1 << 3:
		return data.MonsterTypeUnique
	case 1 << 4:
		return data.MonsterTypeMinion
	}
	return data.MonsterTypeNone
}

func (gd *GameReader) getMonsterStats(statCount uint, statPtr uintptr) map[stat.ID]int {
	stats := map[stat.ID]int{}

	if statCount > 0 {
		statBuffer := gd.Process.ReadBytesFromMemory(statPtr+0x2, statCount*8)
		for i := 0; i < int(statCount); i++ {
			offset := uint(i * 8)
			statEnum := ReadUIntFromBuffer(statBuffer, offset, Uint16)
			statValue := ReadUIntFromBuffer(statBuffer, offset+0x2, Uint32)
			stats[stat.ID(statEnum)] = int(statValue)
		}
	}

	return stats
}

func (gd *GameReader) shouldBeIgnored(txtNo uint) bool {
	switch npc.ID(txtNo) {
	case npc.Chicken,
		npc.Rat,
		npc.Rogue,
		npc.HellMeteor,
		npc.Bird,
		npc.Bird2,
		npc.Bat,
		npc.Act2Male,
		npc.Act2Female,
		npc.Act2Child,
		npc.Cow,
		npc.Camel,
		npc.Act2Guard,
		npc.Act2Vendor,
		npc.Act2Vendor2,
		npc.Maggot,
		npc.Bug,
		npc.Scorpion,
		npc.Rogue2,
		npc.Rogue3,
		npc.Larva,
		npc.Familiar,
		npc.Act3Male,
		npc.Act3Female,
		npc.Snake,
		npc.Parrot,
		npc.Fish,
		npc.EvilHole,
		npc.EvilHole2,
		npc.EvilHole3,
		npc.EvilHole4,
		npc.EvilHole5,
		npc.FireboltTrap,
		npc.HorzMissileTrap,
		npc.VertMissileTrap,
		npc.PoisonCloudTrap,
		npc.LightningTrap,
		npc.InvisoSpawner,
		npc.Guard,
		npc.MiniSper,
		npc.BoneWall,
		npc.Hydra,
		npc.Hydra2,
		npc.Hydra3,
		npc.SevenTombs,
		npc.IronWolf,
		npc.CompellingOrbNpc,
		npc.SpiritMummy,
		npc.Act2Guard4,
		npc.Act2Guard5,
		npc.Window,
		npc.Window2,
		npc.MephistoSpirit,
		npc.WakeOfDestruction,
		npc.ChargedBoltSentry,
		npc.LightningSentry,
		npc.InvisiblePet,
		npc.InfernoSentry,
		npc.DeathSentry,
		npc.BaalThrone,
		npc.BaalTentacle,
                npc.BaalTentacle2,
                npc.BaalTentacle3,
                npc.BaalTentacle4,
                npc.BaalTentacle5,
		npc.InjuredBarbarian,
		npc.InjuredBarbarian2,
		npc.InjuredBarbarian3,
		npc.FrenziedHellSpawn,
		npc.FrenziedIceSpawn,
		npc.CatapultSpotterE,
		npc.CatapultSpotterS,
		npc.CatapultSpotterW,
		npc.CatapultW,
		npc.CatapultE,
		npc.CatapultS,
		npc.CatapultSiege,
		npc.BarricadeWallRight,
		npc.BarricadeWallLeft,
		npc.BarricadeDoor,
		npc.BarricadeDoor2,
		npc.BarricadeTower,
		npc.EvilHut,
		npc.DemonHole:
		return true
	}
	return false
}
