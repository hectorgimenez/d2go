package data

import (
	"math"
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data/mode"

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
	AreaOrigin              Position
	Corpse                  Corpse
	Monsters                Monsters
	Corpses                 Monsters
	Game                    OnlineGame
	PlayerUnit              PlayerUnit
	NPCs                    NPCs
	Inventory               Inventory
	Objects                 Objects
	Entrances               Entrances
	AdjacentLevels          []Level
	Rooms                   []Room
	OpenMenus               OpenMenus
	Widgets                 map[string]map[string]interface{}
	Roster                  Roster
	HoverData               HoverData
	TerrorZones             []area.ID
	Quests                  quest.Quests
	KeyBindings             KeyBindings
	LegacyGraphics          bool
	IsOnline                bool
	IsIngame                bool
	IsInCharCreationScreen  bool
	IsInCharSelectionScreen bool
	IsInLobby               bool
	HasMerc                 bool
	ActiveWeaponSlot        int
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

type OnlineGame struct {
	LastGameName     string
	LastGamePassword string
	FPS              int
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
	States    state.States
}

type Position struct {
	X int
	Y int
}

type PlayerUnit struct {
	Address            uintptr
	Name               string
	ID                 UnitID
	Area               area.ID
	Position           Position
	Stats              stat.Stats
	BaseStats          stat.Stats
	Skills             map[skill.ID]skill.Points
	States             state.States
	Class              Class
	LeftSkill          skill.ID
	RightSkill         skill.ID
	AvailableWaypoints []area.ID // Is only filled when WP menu is open and only for the specific selected tab
	Mode               mode.PlayerMode
}

func (pu PlayerUnit) FindStat(id stat.ID, layer int) (stat.Data, bool) {
	st, found := pu.Stats.FindStat(id, layer)
	if found {
		return st, true
	}

	return pu.BaseStats.FindStat(id, layer)
}

func (pu PlayerUnit) MaxGold() int {
	lvl, _ := pu.FindStat(stat.Level, 0)
	return goldPerLevel * lvl.Value
}

// TotalPlayerGold returns the amount of gold, including inventory and player stash (excluding shared stash)
func (pu PlayerUnit) TotalPlayerGold() int {
	gold, _ := pu.FindStat(stat.Gold, 0)
	stashGold, _ := pu.FindStat(stat.StashGold, 0)

	return gold.Value + stashGold.Value
}

func (pu PlayerUnit) HPPercent() int {
	life, _ := pu.FindStat(stat.Life, 0)
	maxLife, _ := pu.FindStat(stat.MaxLife, 0)

	return int((float64(life.Value) / float64(maxLife.Value)) * 100)
}

func (pu PlayerUnit) MPPercent() int {
	mana, _ := pu.FindStat(stat.Mana, 0)
	maxMana, _ := pu.FindStat(stat.MaxMana, 0)

	return int((float64(mana.Value) / float64(maxMana.Value)) * 100)
}

func (pu PlayerUnit) CastingFrames() int {
	// Formula detailed here: https://diablo.fandom.com/wiki/Faster_Cast_Rate
	fcr, _ := pu.FindStat(stat.FasterCastRate, 0)
	baseAnimation := pu.getCastingBaseSpeed()

	ecfr := math.Floor(float64(fcr.Value*120) / float64(fcr.Value+120))
	if ecfr > 75 {
		ecfr = 75
	}

	// Animation speed for druids is different
	cf := math.Ceil(256*baseAnimation/math.Floor(float64(256*(100+ecfr))/100.0)) - 1

	return int(cf)
}

func (pu PlayerUnit) getCastingBaseSpeed() float64 {
	// TODO: Implement logic for Lightning?
	switch pu.Class {
	case Amazon:
		return 20
	case Assassin:
		return 17
	case Barbarian:
		return 14
	case Necromancer, Paladin:
		return 16
	case Druid:
		return 15
	case Sorceress:
		return 14
	}

	return 16
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
	Inventory      bool
	LoadingScreen  bool
	NPCInteract    bool
	NPCShop        bool
	Stash          bool
	Waypoint       bool
	MapShown       bool
	NewSkills      bool
	NewStats       bool
	SkillTree      bool
	Character      bool
	QuitMenu       bool
	Cube           bool
	SkillSelect    bool
	Anvil          bool
	MercInventory  bool
	BeltRows       bool
	QuestLog       bool
	PortraitsShown bool
	ChatOpen       bool
}

func (om OpenMenus) IsMenuOpen() bool {
	return om.Inventory || om.NPCInteract || om.NPCShop || om.Stash || om.Waypoint || om.SkillTree || om.Character || om.QuitMenu || om.Cube || om.SkillSelect || om.Anvil || om.ChatOpen || om.QuestLog || om.BeltRows || om.MercInventory
}
func (c Corpse) StateNotInteractable() bool {
	CorpseStates := []state.State{
		state.CorpseNoselect,
		state.CorpseNodraw,
		state.Revive,
		state.Redeemed,
		state.Shatter,
		state.Freeze,
	}

	for _, s := range c.States {
		for _, d := range CorpseStates {
			if s == d {
				return true
			}
		}
	}

	return false
}
