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

type Tree map[Skill]Position
type Position struct {
	Tab    int
	Row    int
	Column int
}
