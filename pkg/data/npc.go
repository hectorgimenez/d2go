package data

import (
	"github.com/hectorgimenez/d2go/pkg/data/mode"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/d2go/pkg/data/state"
)

type NPC struct {
	ID        npc.ID
	Name      string
	Positions []Position
}

type MonsterType string

type Monster struct {
	UnitID
	Name      npc.ID
	IsHovered bool
	Position  Position
	Stats     map[stat.ID]int
	Type      MonsterType
	States    state.States
	Mode      mode.NpcMode
}

type Monsters []Monster
type NPCs []NPC

func (n NPCs) FindOne(npcid npc.ID) (NPC, bool) {
	for _, np := range n {
		if np.ID == npcid {
			return np, true
		}
	}

	return NPC{}, false
}

func (m Monsters) FindOne(id npc.ID, t MonsterType) (Monster, bool) {
	for _, monster := range m {
		if monster.Name == id {
			if t == MonsterTypeNone || t == monster.Type {
				return monster, true
			}
		}
	}

	return Monster{}, false
}

func (m Monsters) Enemies(filters ...MonsterFilter) []Monster {
	monsters := make([]Monster, 0)
	for _, mo := range m {
		if !mo.IsMerc() && !mo.IsSkip() && !mo.IsGoodNPC() && !mo.IsPet() && mo.Stats[stat.Life] > 0 {
			monsters = append(monsters, mo)
		}
	}

	for _, f := range filters {
		monsters = f(monsters)
	}

	return monsters
}

type MonsterFilter func(m Monsters) []Monster

func MonsterEliteFilter() MonsterFilter {
	return func(m Monsters) []Monster {
		var filteredMonsters []Monster
		for _, mo := range m {
			if mo.IsElite() {
				filteredMonsters = append(filteredMonsters, mo)
			}
		}

		return filteredMonsters
	}
}

func MonsterAnyFilter() MonsterFilter {
	return func(m Monsters) []Monster {
		return m
	}
}

func (m Monsters) FindByID(id UnitID) (Monster, bool) {
	for _, monster := range m {
		if monster.UnitID == id {
			return monster, true
		}
	}

	return Monster{}, false
}

func (m Monster) IsImmune(resist stat.Resist) bool {
	for st, value := range m.Stats {
		// We only want max resistance
		if value < 100 {
			continue
		}
		if resist == stat.ColdImmune && st == stat.ColdResist {
			return true
		}
		if resist == stat.FireImmune && st == stat.FireResist {
			return true
		}
		if resist == stat.LightImmune && st == stat.LightningResist {
			return true
		}
		if resist == stat.PoisonImmune && st == stat.PoisonResist {
			return true
		}
		if resist == stat.MagicImmune && st == stat.MagicResist {
			return true
		}
	}
	return false
}

func (m Monster) IsMerc() bool {
	if m.Name == npc.Guard || m.Name == npc.Act5Hireling1Hand || m.Name == npc.Act5Hireling2Hand || m.Name == npc.IronWolf || m.Name == npc.Rogue2 {
		return true
	}

	return false
}

func (m Monster) IsPet() bool {
	// Necro revive
	if m.States.HasState(state.Revive) {
		return true
	}

	switch m.Name {
	case npc.DruHawk, npc.DruSpiritWolf, npc.DruFenris, npc.HeartOfWolverine,
		npc.OakSage, npc.DruBear, npc.DruPlaguePoppy, npc.VineCreature,
		npc.DruCycleOfLife, npc.ClayGolem, npc.BloodGolem, npc.IronGolem,
		npc.FireGolem, npc.NecroSkeleton, npc.NecroMage, npc.Valkyrie, npc.Decoy,
		npc.ShadowWarrior, npc.ShadowMaster:
		return true
	default:
		return false
	}
}

func (m Monster) IsGoodNPC() bool {
	switch m.Name {
	case 146, 154, 147, 150, 155, 148, 244, 210, 175, 199, 198, 177, 178, 201, 202, 200, 331, 245, 264, 255, 176,
		252, 254, 253, 297, 246, 251, 367, 521, 257, 405, 265, 520, 512, 518, 527, 515, 513, 511, 514, 266, 408, 406:
		return true
	}

	return false
}

func (m Monster) IsElite() bool {
	return m.Type == MonsterTypeMinion || m.Type == MonsterTypeUnique || m.Type == MonsterTypeChampion || m.Type == MonsterTypeSuperUnique
}

// IsMonsterRaiser returns true if the monster is able to spawn new monsters.
func (m Monster) IsMonsterRaiser() bool {
	switch m.Name {
	case npc.FallenShaman,
		npc.CarverShaman,
		npc.DevilkinShaman,
		npc.DarkShaman,
		npc.WarpedShaman:
		return true
	}

	return false
}

// IsSkip monster can not be killed as a normal enemy, for example can not be targeted
func (m Monster) IsSkip() bool {
	switch m.Name {
	case npc.WaterWatcherLimb, npc.WaterWatcherHead, npc.BaalTaunt, npc.Act5Combatant, npc.Act5Combatant2, npc.BarricadeTower, npc.DarkWanderer, npc.POW:
		return true
	}

	return false
}
func (m Monster) IsSealBoss() bool {

	return (m.Type == MonsterTypeSuperUnique || m.Type == MonsterTypeMinion) && (m.Name == npc.OblivionKnight || (m.Name == npc.DoomKnight) || // Lord De Seis
		m.Name == npc.VenomLord || // Infector of Souls
		m.Name == npc.StormCaster) // Grand Vizier of Chaos
}

// IsEscapingType Monster cannot be attacked when airborne or hide in water (NpcMode 8)
func (m Monster) IsEscapingType() bool {
	switch m.Name {
	case npc.CarrionBird, npc.CarrionBird2, npc.WaterWatcherLimb, npc.RiverStalkerLimb, npc.StygianWatcherLimb,
		npc.WaterWatcherHead, npc.RiverStalkerHead, npc.StygianWatcherHead, npc.CloudStalker, npc.Sucker, npc.UndeadScavenger, npc.FoulCrow:

		return true
	}
	return false
}
