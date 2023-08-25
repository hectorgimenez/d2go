package skill

var SorceressTree = Tree{
	// Cold Spells
	IceBolt:       Position{Tab: 0, Row: 0, Column: 1},
	FrozenArmor:   Position{Tab: 0, Row: 0, Column: 2},
	FrostNova:     Position{Tab: 0, Row: 1, Column: 0},
	IceBlast:      Position{Tab: 0, Row: 1, Column: 1},
	ShiverArmor:   Position{Tab: 0, Row: 2, Column: 2},
	GlacialSpike:  Position{Tab: 0, Row: 3, Column: 1},
	Blizzard:      Position{Tab: 0, Row: 4, Column: 0},
	ChillingArmor: Position{Tab: 0, Row: 4, Column: 2},
	FrozenOrb:     Position{Tab: 0, Row: 5, Column: 0},
	ColdMastery:   Position{Tab: 0, Row: 5, Column: 1},

	// Lightning Spells
	ChargedBolt:      Position{Tab: 1, Row: 0, Column: 1},
	StaticField:      Position{Tab: 1, Row: 1, Column: 0},
	Telekinesis:      Position{Tab: 1, Row: 1, Column: 2},
	Nova:             Position{Tab: 1, Row: 2, Column: 0},
	Lightning:        Position{Tab: 1, Row: 2, Column: 1},
	ChainLightning:   Position{Tab: 1, Row: 3, Column: 1},
	Teleport:         Position{Tab: 1, Row: 3, Column: 2},
	ThunderStorm:     Position{Tab: 1, Row: 4, Column: 0},
	EnergyShield:     Position{Tab: 1, Row: 4, Column: 2},
	LightningMastery: Position{Tab: 1, Row: 5, Column: 1},

	// Fire Spells
	FireBolt:    Position{Tab: 2, Row: 0, Column: 1},
	Warmth:      Position{Tab: 2, Row: 0, Column: 2},
	Inferno:     Position{Tab: 2, Row: 1, Column: 0},
	Blaze:       Position{Tab: 2, Row: 2, Column: 1},
	FireBall:    Position{Tab: 2, Row: 2, Column: 1},
	FireWall:    Position{Tab: 2, Row: 3, Column: 0},
	Enchant:     Position{Tab: 2, Row: 3, Column: 2},
	Meteor:      Position{Tab: 2, Row: 4, Column: 1},
	FireMastery: Position{Tab: 2, Row: 5, Column: 1},
	Hydra:       Position{Tab: 2, Row: 5, Column: 2},
}

var PaladinTree = Tree{
	// Defensive Auras
	Prayer:          Position{Tab: 0, Row: 0, Column: 0},
	ResistFire:      Position{Tab: 0, Row: 0, Column: 2},
	Defiance:        Position{Tab: 0, Row: 1, Column: 1},
	ResistCold:      Position{Tab: 0, Row: 1, Column: 2},
	Cleansing:       Position{Tab: 0, Row: 2, Column: 0},
	ResistLightning: Position{Tab: 0, Row: 2, Column: 2},
	Vigor:           Position{Tab: 0, Row: 3, Column: 1},
	Meditation:      Position{Tab: 0, Row: 4, Column: 0},
	Redemption:      Position{Tab: 0, Row: 5, Column: 1},
	Salvation:       Position{Tab: 0, Row: 5, Column: 2},

	// Offensive Auras
	Might:         Position{Tab: 1, Row: 0, Column: 0},
	HolyFire:      Position{Tab: 1, Row: 1, Column: 1},
	Thorns:        Position{Tab: 1, Row: 1, Column: 2},
	BlessedAim:    Position{Tab: 1, Row: 2, Column: 0},
	Concentration: Position{Tab: 1, Row: 3, Column: 0},
	HolyFreeze:    Position{Tab: 1, Row: 3, Column: 1},
	HolyShock:     Position{Tab: 1, Row: 4, Column: 1},
	Sanctuary:     Position{Tab: 1, Row: 4, Column: 2},
	Fanaticism:    Position{Tab: 1, Row: 5, Column: 0},
	Conviction:    Position{Tab: 1, Row: 5, Column: 2},

	// Combat Skills
	Sacrifice:        Position{Tab: 2, Row: 0, Column: 0},
	Smite:            Position{Tab: 2, Row: 0, Column: 2},
	HolyBolt:         Position{Tab: 2, Row: 1, Column: 1},
	Zeal:             Position{Tab: 2, Row: 2, Column: 0},
	Charge:           Position{Tab: 2, Row: 2, Column: 2},
	Vengeance:        Position{Tab: 2, Row: 3, Column: 0},
	BlessedHammer:    Position{Tab: 2, Row: 3, Column: 1},
	Conversion:       Position{Tab: 2, Row: 4, Column: 0},
	HolyShield:       Position{Tab: 2, Row: 4, Column: 2},
	FistOfTheHeavens: Position{Tab: 2, Row: 5, Column: 1},
}

type Tree map[Skill]Position
type Position struct {
	Tab    int
	Row    int
	Column int
}
