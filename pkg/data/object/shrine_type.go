package object

type ShrineType uint

type ShrinePosition struct {
	X int
	Y int
}

type ShrineData struct {
	ShrineName string
	ShrineType ShrineType
	Position   ShrinePosition
}

const (
	RefillShrine          ShrineType = 0x01
	HealthShrine          ShrineType = 0x02
	ManaShrine            ShrineType = 0x03
	HPXChangeShrine       ShrineType = 0x04
	ManaXChangeShrine     ShrineType = 0x05
	ArmorShrine           ShrineType = 0x06
	CombatShrine          ShrineType = 0x07
	ResistFireShrine      ShrineType = 0x08
	ResistColdShrine      ShrineType = 0x09
	ResistLightningShrine ShrineType = 0x0A
	ResistPoisonShrine    ShrineType = 0x0B
	SkillShrine           ShrineType = 0x0C
	ManaRegenShrine       ShrineType = 0x0D
	StaminaShrine         ShrineType = 0x0E
	ExperienceShrine      ShrineType = 0x0F
	UnknownShrine         ShrineType = 0x10
	PortalShrine          ShrineType = 0x11
	GemShrine             ShrineType = 0x12
	FireShrine            ShrineType = 0x13
	MonsterShrine         ShrineType = 0x14
	ExplosiveShrine       ShrineType = 0x15
	PoisonShrine          ShrineType = 0x16
)

var ShrineTypeNames = map[ShrineType]string{
	RefillShrine:          "Refill Shrine",
	HealthShrine:          "Health Shrine",
	ManaShrine:            "Mana Shrine",
	HPXChangeShrine:       "HP XChange Shrine",
	ManaXChangeShrine:     "Mana XChange Shrine",
	ArmorShrine:           "Armor Shrine",
	CombatShrine:          "Combat Shrine",
	ResistFireShrine:      "Resist Fire Shrine",
	ResistColdShrine:      "Resist Cold Shrine",
	ResistLightningShrine: "Resist Lightning Shrine",
	ResistPoisonShrine:    "Resist Poison Shrine",
	SkillShrine:           "Skill Shrine",
	ManaRegenShrine:       "Mana Regen Shrine",
	StaminaShrine:         "Stamina Shrine",
	ExperienceShrine:      "Experience Shrine",
	UnknownShrine:         "Unknown Shrine",
	PortalShrine:          "Portal Shrine",
	GemShrine:             "Gem Shrine",
	FireShrine:            "Fire Shrine",
	MonsterShrine:         "Monster Shrine",
	ExplosiveShrine:       "Explosive Shrine",
	PoisonShrine:          "Poison Shrine",
}
