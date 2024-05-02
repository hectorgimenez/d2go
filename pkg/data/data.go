package data

import (
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data/quest"

	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/skill"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/d2go/pkg/data/state"
)

const (
	goldPerLevel = 10000

	// Monster Types
	MonsterTypeNone        MonsterType = "None"
	MonsterTypeChampion    MonsterType = "Champion"
	MonsterTypeMinion      MonsterType = "Minion"
	MonsterTypeUnique      MonsterType = "Unique"
	MonsterTypeSuperUnique MonsterType = "SuperUnique"
)

type Data struct {
	AreaOrigin Position
	Corpse     Corpse
	Monsters   Monsters
	// First slice represents X and second Y
	CollisionGrid  [][]bool
	PlayerUnit     PlayerUnit
	NPCs           NPCs
	Items          Items
	Objects        Objects
	AdjacentLevels []Level
	Rooms          []Room
	OpenMenus      OpenMenus
	Roster         Roster
	HoverData      HoverData
	TerrorZones    []area.ID
	Quests         quest.Quests
	KeyBindings    KeyBindings
}

type Room struct {
	Position
	Width  int
	Height int
}

type HoverData struct {
	IsHovered bool
	UnitID
	UnitType int
}

func (r Room) GetCenter() Position {
	return Position{
		X: r.Position.X + r.Width/2,
		Y: r.Position.Y + r.Height/2,
	}
}

func (r Room) IsInside(p Position) bool {
	if p.X >= r.X && p.X <= r.X+r.Width {
		return p.Y >= r.Y && p.Y <= r.Y+r.Height
	}

	return false
}

func (d Data) MercHPPercent() int {
	for _, m := range d.Monsters {
		if m.IsMerc() {
			// Hacky thing to read merc life properly
			maxLife := m.Stats[stat.MaxLife] >> 8
			life := float64(m.Stats[stat.Life] >> 8)
			if m.Stats[stat.Life] <= 32768 {
				life = float64(m.Stats[stat.Life]) / 32768.0 * float64(maxLife)
			}

			return int(life / float64(maxLife) * 100)
		}
	}

	return 0
}

type RosterMember struct {
	Name     string
	Area     area.ID
	Position Position
}
type Roster []RosterMember

func (r Roster) FindByName(name string) (RosterMember, bool) {
	for _, rm := range r {
		if strings.EqualFold(rm.Name, name) {
			return rm, true
		}
	}

	return RosterMember{}, false
}

type Level struct {
	Area       area.ID
	Position   Position
	IsEntrance bool // This means the area can not be accessed just walking through it, needs to be clicked
}

type Class uint

const (
	Amazon Class = iota
	Sorceress
	Necromancer
	Paladin
	Barbarian
	Druid
	Assassin
)

type Corpse struct {
	Found     bool
	IsHovered bool
	Position  Position
}

type Position struct {
	X int
	Y int
}

type PlayerUnit struct {
	Name               string
	ID                 UnitID
	Area               area.ID
	Position           Position
	Stats              map[stat.ID]int
	BaseStats          map[stat.ID]int
	Skills             map[skill.ID]skill.Points
	States             state.States
	Class              Class
	LeftSkill          skill.ID
	RightSkill         skill.ID
	AvailableWaypoints []area.ID // Is only filled when WP menu is open and only for the specific selected tab
}

func (pu PlayerUnit) MaxGold() int {
	return goldPerLevel * pu.Stats[stat.Level]
}

// TotalGold returns the amount of gold, including inventory and stash
func (pu PlayerUnit) TotalGold() int {
	return pu.Stats[stat.Gold] + pu.Stats[stat.StashGold]
}

func (pu PlayerUnit) HPPercent() int {
	return int((float64(pu.Stats[stat.Life]) / float64(pu.Stats[stat.MaxLife])) * 100)
}

func (pu PlayerUnit) MPPercent() int {
	return int((float64(pu.Stats[stat.Mana]) / float64(pu.Stats[stat.MaxMana])) * 100)
}

func (pu PlayerUnit) HasDebuff() bool {
	debuffs := []state.State{
		state.Amplifydamage,
		state.Attract,
		state.Confuse,
		state.Conversion,
		state.Decrepify,
		state.Dimvision,
		state.Ironmaiden,
		state.Lifetap,
		state.Lowerresist,
		state.Terror,
		state.Weaken,
		state.Convicted,
		state.Poison,
		state.Cold,
		state.Slowed,
		state.BloodMana,
		state.DefenseCurse,
	}

	for _, s := range pu.States {
		for _, d := range debuffs {
			if s == d {
				return true
			}
		}
	}

	return false
}

type PointOfInterest struct {
	Name     string
	Position Position
}

type OpenMenus struct {
	Inventory     bool
	LoadingScreen bool
	NPCInteract   bool
	NPCShop       bool
	Stash         bool
	Waypoint      bool
	MapShown      bool
	SkillTree     bool
	Character     bool
	QuitMenu      bool
	Cube          bool
	SkillSelect   bool
	Anvil         bool
}
