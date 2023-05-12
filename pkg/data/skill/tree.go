package skill

var SorceressTree = Trees{
	// Cold Spells
	Tree{
		IceBolt:       Position{Row: 0, Column: 1},
		FrozenArmor:   Position{Row: 0, Column: 2},
		FrostNova:     Position{Row: 1, Column: 0},
		IceBlast:      Position{Row: 1, Column: 1},
		ShiverArmor:   Position{Row: 2, Column: 2},
		GlacialSpike:  Position{Row: 3, Column: 1},
		Blizzard:      Position{Row: 4, Column: 0},
		ChillingArmor: Position{Row: 4, Column: 2},
		FrozenOrb:     Position{Row: 5, Column: 0},
		ColdMastery:   Position{Row: 5, Column: 1},
	},
	// Lightning Spells
	Tree{
		ChargedBolt:      Position{Row: 0, Column: 1},
		StaticField:      Position{Row: 1, Column: 0},
		Telekinesis:      Position{Row: 1, Column: 2},
		Nova:             Position{Row: 2, Column: 0},
		Lightning:        Position{Row: 2, Column: 1},
		ChainLightning:   Position{Row: 3, Column: 1},
		Teleport:         Position{Row: 3, Column: 2},
		ThunderStorm:     Position{Row: 4, Column: 0},
		EnergyShield:     Position{Row: 4, Column: 2},
		LightningMastery: Position{Row: 5, Column: 1},
	},
	// Fire Spells
	Tree{
		FireBolt:    Position{Row: 0, Column: 1},
		Warmth:      Position{Row: 0, Column: 2},
		Inferno:     Position{Row: 1, Column: 0},
		Blaze:       Position{Row: 2, Column: 1},
		FireBall:    Position{Row: 2, Column: 2},
		FireWall:    Position{Row: 3, Column: 0},
		Enchant:     Position{Row: 3, Column: 2},
		Meteor:      Position{Row: 4, Column: 1},
		FireMastery: Position{Row: 5, Column: 1},
		Hydra:       Position{Row: 5, Column: 2},
	},
}

type Trees [3]Tree
type Tree map[Skill]Position
type Position struct {
	Row    int
	Column int
}
