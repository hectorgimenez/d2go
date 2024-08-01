package skill

type ID int16

type Description struct {
	Name    string
	Page    int
	Row     int
	Column  int
	ListRow int
	IconCel int
}

type Skill struct {
	ID
	Name       string
	LeftSkill  bool
	RightSkill bool
}

type Points struct {
	Level    uint
	Quantity uint
	Charges  uint
}

func (sk ID) Desc() Description {
	return Desc[sk]
}

const (
	AttackSkill ID = iota
	Kick
	Throw
	Unsummon
	LeftHandThrow
	LeftHandSwing
	MagicArrow
	FireArrow
	InnerSight
	CriticalStrike
	Jab
	ColdArrow
	MultipleShot
	Dodge
	PowerStrike
	PoisonJavelin
	ExplodingArrow
	SlowMissiles
	Avoid
	Impale
	LightningBolt
	IceArrow
	GuidedArrow
	Penetrate
	ChargedStrike
	PlagueJavelin
	Strafe
	ImmolationArrow
	Decoy
	Evade
	Fend
	FreezingArrow
	Valkyrie
	Pierce
	LightningStrike
	LightningFury
	FireBolt
	Warmth
	ChargedBolt
	IceBolt
	FrozenArmor
	Inferno
	StaticField
	Telekinesis
	FrostNova
	IceBlast
	Blaze
	FireBall
	Nova
	Lightning
	ShiverArmor
	FireWall
	Enchant
	ChainLightning
	Teleport
	GlacialSpike
	Meteor
	ThunderStorm
	EnergyShield
	Blizzard
	ChillingArmor
	FireMastery
	Hydra
	LightningMastery
	FrozenOrb
	ColdMastery
	AmplifyDamage
	Teeth
	BoneArmor
	SkeletonMastery
	RaiseSkeleton
	DimVision
	Weaken
	PoisonDagger
	CorpseExplosion
	ClayGolem
	IronMaiden
	Terror
	BoneWall
	GolemMastery
	RaiseSkeletalMage
	Confuse
	LifeTap
	PoisonExplosion
	BoneSpear
	BloodGolem
	Attract
	Decrepify
	BonePrison
	SummonResist
	IronGolem
	LowerResist
	PoisonNova
	BoneSpirit
	FireGolem
	Revive
	Sacrifice
	Smite
	Might
	Prayer
	ResistFire
	HolyBolt
	HolyFire
	Thorns
	Defiance
	ResistCold
	Zeal
	Charge
	BlessedAim
	Cleansing
	ResistLightning
	Vengeance
	BlessedHammer
	Concentration
	HolyFreeze
	Vigor
	Conversion
	HolyShield
	HolyShock
	Sanctuary
	Meditation
	FistOfTheHeavens
	Fanaticism
	Conviction
	Redemption
	Salvation
	Bash
	BladeMastery
	AxeMastery
	MaceMastery
	Howl
	FindPotion
	Leap
	DoubleSwing
	PolearmMastery
	ThrowingMastery
	SpearMastery
	Taunt
	Shout
	Stun
	DoubleThrow
	IncreasedStamina
	FindItem
	LeapAttack
	Concentrate
	IronSkin
	BattleCry
	Frenzy
	IncreasedSpeed
	BattleOrders
	GrimWard
	Whirlwind
	Berserk
	NaturalResistance
	WarCry
	BattleCommand
	FireHit
	UnHolyBolt
	SkeletonRaise
	MaggotEgg
	ShamanFire
	MagottUp
	MagottDown
	MagottLay
	AndrialSpray
	Jump
	SwarmMove
	Nest
	QuickStrike
	VampireFireball
	VampireFirewall
	VampireMeteor
	GargoyleTrap
	SpiderLay
	VampireHeal
	VampireRaise
	Submerge
	FetishAura
	FetishInferno
	ZakarumHeal
	Emerge
	Resurrect
	Bestow
	MissileSkill1
	MonTeleport
	PrimeLightning
	PrimeBolt
	PrimeBlaze
	PrimeFirewall
	PrimeSpike
	PrimeIceNova
	PrimePoisonball
	PrimePoisonNova
	DiabLight
	DiabCold
	DiabFire
	FingerMageSpider
	DiabWall
	DiabRun
	DiabPrison
	PoisonBallTrap
	AndyPoisonBolt
	HireableMissile
	DesertTurret
	ArcaneTower
	MonBlizzard
	Mosquito
	CursedBallTrapRight
	CursedBallTrapLeft
	MonFrozenArmor
	MonBoneArmor
	MonBoneSpirit
	MonCurseCast
	HellMeteor
	RegurgitatorEat
	MonFrenzy
	QueenDeath
	ScrollOfIdentify
	TomeOfIdentify
	ScrollOfTownPortal
	TomeOfTownPortal
	Raven
	PoisonCreeper
	Werewolf
	Lycanthropy
	Firestorm
	OakSage
	SummonSpiritWolf
	Werebear
	MoltenBoulder
	ArcticBlast
	CarrionVine
	FeralRage
	Maul
	Fissure
	CycloneArmor
	HeartOfWolverine
	SummonDireWolf
	Rabies
	FireClaws
	Twister
	SolarCreeper
	Hunger
	ShockWave
	Volcano
	Tornado
	SpiritOfBarbs
	SummonGrizzly
	Fury
	Armageddon
	Hurricane
	FireBlast
	ClawMastery
	PsychicHammer
	TigerStrike
	DragonTalon
	ShockWeb
	BladeSentinel
	BurstOfSpeed
	FistsOfFire
	DragonClaw
	ChargedBoltSentry
	WakeOfFire
	WeaponBlock
	CloakOfShadows
	CobraStrike
	BladeFury
	Fade
	ShadowWarrior
	ClawsOfThunder
	DragonTail
	LightningSentry
	WakeOfInferno
	MindBlast
	BladesOfIce
	DragonFlight
	DeathSentry
	BladeShield
	Venom
	ShadowMaster
	PhoenixStrike
	WakeOfDestructionSentry
	ImpInferno
	ImpFireball
	BaalTaunt
	BaalCorpseExplode
	BaalMonsterSpawn
	CatapultChargedBall
	CatapultSpikeBall
	SuckBlood
	CryHelp
	HealingVortex
	Teleport2
	SelfResurrect
	VineAttack
	OverseerWhip
	BarbsAura
	WolverineAura
	OakSageAura
	ImpFireMissile
	Impregnate
	SiegeBeastStomp
	MinionSpawner
	CatapultBlizzard
	CatapultPlague
	CatapultMeteor
	BoltSentry
	CorpseCycler
	DeathMaul
	DefenseCurse
	BloodMana
	monInfernoSentry
	monDeathSentry
	sentryLightning
	fenrisRage
	BaalTentacle
	BaalNova
	BaalInferno
	BaalColdMissiles
	MegademonInferno
	EvilHutSpawner
	CountessFirewall
	ImpBolt
	HorrorArcticBlast
	deathSentryLtng
	VineCycler
	BearSmite
	Resurrect2
	BloodLordFrenzy
	BaalTeleport
	ImpTeleport
	BaalCloneTeleport
	ZakarumLightning
	VampireMissile
	MephistoMissile
	DoomKnightMissile
	RogueMissile
	HydraMissile
	NecromageMissile
	MonBow
	MonFireArrow
	MonColdArrow
	MonExplodingArrow
	MonFreezingArrow
	MonPowerStrike
	SuccubusBolt
	MephFrostNova
	MonIceSpear
	ShamanIce
	Diablogeddon
	DeleriumChange
	NihlathakCorpseExplosion
	SerpentCharge
	TrapNova
	UnHolyBoltEx
	ShamanFireEx
	ImpFireMissileEx
	FixedSiegeBeastStomp
	Unset ID = -1
)

var SkillNames = map[ID]string{
	AttackSkill:              "AttackSkill",
	Kick:                     "Kick",
	Throw:                    "Throw",
	Unsummon:                 "Unsummon",
	LeftHandThrow:            "LeftHandThrow",
	LeftHandSwing:            "LeftHandSwing",
	MagicArrow:               "MagicArrow",
	FireArrow:                "FireArrow",
	InnerSight:               "InnerSight",
	CriticalStrike:           "CriticalStrike",
	Jab:                      "Jab",
	ColdArrow:                "ColdArrow",
	MultipleShot:             "MultipleShot",
	Dodge:                    "Dodge",
	PowerStrike:              "PowerStrike",
	PoisonJavelin:            "PoisonJavelin",
	ExplodingArrow:           "ExplodingArrow",
	SlowMissiles:             "SlowMissiles",
	Avoid:                    "Avoid",
	Impale:                   "Impale",
	LightningBolt:            "LightningBolt",
	IceArrow:                 "IceArrow",
	GuidedArrow:              "GuidedArrow",
	Penetrate:                "Penetrate",
	ChargedStrike:            "ChargedStrike",
	PlagueJavelin:            "PlagueJavelin",
	Strafe:                   "Strafe",
	ImmolationArrow:          "ImmolationArrow",
	Decoy:                    "Decoy",
	Evade:                    "Evade",
	Fend:                     "Fend",
	FreezingArrow:            "FreezingArrow",
	Valkyrie:                 "Valkyrie",
	Pierce:                   "Pierce",
	LightningStrike:          "LightningStrike",
	LightningFury:            "LightningFury",
	FireBolt:                 "FireBolt",
	Warmth:                   "Warmth",
	ChargedBolt:              "ChargedBolt",
	IceBolt:                  "IceBolt",
	FrozenArmor:              "FrozenArmor",
	Inferno:                  "Inferno",
	StaticField:              "StaticField",
	Telekinesis:              "Telekinesis",
	FrostNova:                "FrostNova",
	IceBlast:                 "IceBlast",
	Blaze:                    "Blaze",
	FireBall:                 "FireBall",
	Nova:                     "Nova",
	Lightning:                "Lightning",
	ShiverArmor:              "ShiverArmor",
	FireWall:                 "FireWall",
	Enchant:                  "Enchant",
	ChainLightning:           "ChainLightning",
	Teleport:                 "Teleport",
	GlacialSpike:             "GlacialSpike",
	Meteor:                   "Meteor",
	ThunderStorm:             "ThunderStorm",
	EnergyShield:             "EnergyShield",
	Blizzard:                 "Blizzard",
	ChillingArmor:            "ChillingArmor",
	FireMastery:              "FireMastery",
	Hydra:                    "Hydra",
	LightningMastery:         "LightningMastery",
	FrozenOrb:                "FrozenOrb",
	ColdMastery:              "ColdMastery",
	AmplifyDamage:            "AmplifyDamage",
	Teeth:                    "Teeth",
	BoneArmor:                "BoneArmor",
	SkeletonMastery:          "SkeletonMastery",
	RaiseSkeleton:            "RaiseSkeleton",
	DimVision:                "DimVision",
	Weaken:                   "Weaken",
	PoisonDagger:             "PoisonDagger",
	CorpseExplosion:          "CorpseExplosion",
	ClayGolem:                "ClayGolem",
	IronMaiden:               "IronMaiden",
	Terror:                   "Terror",
	BoneWall:                 "BoneWall",
	GolemMastery:             "GolemMastery",
	RaiseSkeletalMage:        "RaiseSkeletalMage",
	Confuse:                  "Confuse",
	LifeTap:                  "LifeTap",
	PoisonExplosion:          "PoisonExplosion",
	BoneSpear:                "BoneSpear",
	BloodGolem:               "BloodGolem",
	Attract:                  "Attract",
	Decrepify:                "Decrepify",
	BonePrison:               "BonePrison",
	SummonResist:             "SummonResist",
	IronGolem:                "IronGolem",
	LowerResist:              "LowerResist",
	PoisonNova:               "PoisonNova",
	BoneSpirit:               "BoneSpirit",
	FireGolem:                "FireGolem",
	Revive:                   "Revive",
	Sacrifice:                "Sacrifice",
	Smite:                    "Smite",
	Might:                    "Might",
	Prayer:                   "Prayer",
	ResistFire:               "ResistFire",
	HolyBolt:                 "HolyBolt",
	HolyFire:                 "HolyFire",
	Thorns:                   "Thorns",
	Defiance:                 "Defiance",
	ResistCold:               "ResistCold",
	Zeal:                     "Zeal",
	Charge:                   "Charge",
	BlessedAim:               "BlessedAim",
	Cleansing:                "Cleansing",
	ResistLightning:          "ResistLightning",
	Vengeance:                "Vengeance",
	BlessedHammer:            "BlessedHammer",
	Concentration:            "Concentration",
	HolyFreeze:               "HolyFreeze",
	Vigor:                    "Vigor",
	Conversion:               "Conversion",
	HolyShield:               "HolyShield",
	HolyShock:                "HolyShock",
	Sanctuary:                "Sanctuary",
	Meditation:               "Meditation",
	FistOfTheHeavens:         "FistOfTheHeavens",
	Fanaticism:               "Fanaticism",
	Conviction:               "Conviction",
	Redemption:               "Redemption",
	Salvation:                "Salvation",
	Bash:                     "Bash",
	BladeMastery:             "BladeMastery",
	AxeMastery:               "AxeMastery",
	MaceMastery:              "MaceMastery",
	Howl:                     "Howl",
	FindPotion:               "FindPotion",
	Leap:                     "Leap",
	DoubleSwing:              "DoubleSwing",
	PolearmMastery:           "PolearmMastery",
	ThrowingMastery:          "ThrowingMastery",
	SpearMastery:             "SpearMastery",
	Taunt:                    "Taunt",
	Shout:                    "Shout",
	Stun:                     "Stun",
	DoubleThrow:              "DoubleThrow",
	IncreasedStamina:         "IncreasedStamina",
	FindItem:                 "FindItem",
	LeapAttack:               "LeapAttack",
	Concentrate:              "Concentrate",
	IronSkin:                 "IronSkin",
	BattleCry:                "BattleCry",
	Frenzy:                   "Frenzy",
	IncreasedSpeed:           "IncreasedSpeed",
	BattleOrders:             "BattleOrders",
	GrimWard:                 "GrimWard",
	Whirlwind:                "Whirlwind",
	Berserk:                  "Berserk",
	NaturalResistance:        "NaturalResistance",
	WarCry:                   "WarCry",
	BattleCommand:            "BattleCommand",
	FireHit:                  "FireHit",
	UnHolyBolt:               "UnHolyBolt",
	SkeletonRaise:            "SkeletonRaise",
	MaggotEgg:                "MaggotEgg",
	ShamanFire:               "ShamanFire",
	MagottUp:                 "MagottUp",
	MagottDown:               "MagottDown",
	MagottLay:                "MagottLay",
	AndrialSpray:             "AndrialSpray",
	Jump:                     "Jump",
	SwarmMove:                "SwarmMove",
	Nest:                     "Nest",
	QuickStrike:              "QuickStrike",
	VampireFireball:          "VampireFireball",
	VampireFirewall:          "VampireFirewall",
	VampireMeteor:            "VampireMeteor",
	GargoyleTrap:             "GargoyleTrap",
	SpiderLay:                "SpiderLay",
	VampireHeal:              "VampireHeal",
	VampireRaise:             "VampireRaise",
	Submerge:                 "Submerge",
	FetishAura:               "FetishAura",
	FetishInferno:            "FetishInferno",
	ZakarumHeal:              "ZakarumHeal",
	Emerge:                   "Emerge",
	Resurrect:                "Resurrect",
	Bestow:                   "Bestow",
	MissileSkill1:            "MissileSkill1",
	MonTeleport:              "MonTeleport",
	PrimeLightning:           "PrimeLightning",
	PrimeBolt:                "PrimeBolt",
	PrimeBlaze:               "PrimeBlaze",
	PrimeFirewall:            "PrimeFirewall",
	PrimeSpike:               "PrimeSpike",
	PrimeIceNova:             "PrimeIceNova",
	PrimePoisonball:          "PrimePoisonball",
	PrimePoisonNova:          "PrimePoisonNova",
	DiabLight:                "DiabLight",
	DiabCold:                 "DiabCold",
	DiabFire:                 "DiabFire",
	FingerMageSpider:         "FingerMageSpider",
	DiabWall:                 "DiabWall",
	DiabRun:                  "DiabRun",
	DiabPrison:               "DiabPrison",
	PoisonBallTrap:           "PoisonBallTrap",
	AndyPoisonBolt:           "AndyPoisonBolt",
	HireableMissile:          "HireableMissile",
	DesertTurret:             "DesertTurret",
	ArcaneTower:              "ArcaneTower",
	MonBlizzard:              "MonBlizzard",
	Mosquito:                 "Mosquito",
	CursedBallTrapRight:      "CursedBallTrapRight",
	CursedBallTrapLeft:       "CursedBallTrapLeft",
	MonFrozenArmor:           "MonFrozenArmor",
	MonBoneArmor:             "MonBoneArmor",
	MonBoneSpirit:            "MonBoneSpirit",
	MonCurseCast:             "MonCurseCast",
	HellMeteor:               "HellMeteor",
	RegurgitatorEat:          "RegurgitatorEat",
	MonFrenzy:                "MonFrenzy",
	QueenDeath:               "QueenDeath",
	ScrollOfIdentify:         "ScrollOfIdentify",
	TomeOfIdentify:           "TomeOfIdentify",
	ScrollOfTownPortal:       "ScrollOfTownPortal",
	TomeOfTownPortal:         "TomeOfTownPortal",
	Raven:                    "Raven",
	PoisonCreeper:            "PoisonCreeper",
	Werewolf:                 "Werewolf",
	Lycanthropy:              "Lycanthropy",
	Firestorm:                "Firestorm",
	OakSage:                  "OakSage",
	SummonSpiritWolf:         "SummonSpiritWolf",
	Werebear:                 "Werebear",
	MoltenBoulder:            "MoltenBoulder",
	ArcticBlast:              "ArcticBlast",
	CarrionVine:              "CarrionVine",
	FeralRage:                "FeralRage",
	Maul:                     "Maul",
	Fissure:                  "Fissure",
	CycloneArmor:             "CycloneArmor",
	HeartOfWolverine:         "HeartOfWolverine",
	SummonDireWolf:           "SummonDireWolf",
	Rabies:                   "Rabies",
	FireClaws:                "FireClaws",
	Twister:                  "Twister",
	SolarCreeper:             "SolarCreeper",
	Hunger:                   "Hunger",
	ShockWave:                "ShockWave",
	Volcano:                  "Volcano",
	Tornado:                  "Tornado",
	SpiritOfBarbs:            "SpiritOfBarbs",
	SummonGrizzly:            "SummonGrizzly",
	Fury:                     "Fury",
	Armageddon:               "Armageddon",
	Hurricane:                "Hurricane",
	FireBlast:                "FireBlast",
	ClawMastery:              "ClawMastery",
	PsychicHammer:            "PsychicHammer",
	TigerStrike:              "TigerStrike",
	DragonTalon:              "DragonTalon",
	ShockWeb:                 "ShockWeb",
	BladeSentinel:            "BladeSentinel",
	BurstOfSpeed:             "BurstOfSpeed",
	FistsOfFire:              "FistsOfFire",
	DragonClaw:               "DragonClaw",
	ChargedBoltSentry:        "ChargedBoltSentry",
	WakeOfFire:               "WakeOfFire",
	WeaponBlock:              "WeaponBlock",
	CloakOfShadows:           "CloakOfShadows",
	CobraStrike:              "CobraStrike",
	BladeFury:                "BladeFury",
	Fade:                     "Fade",
	ShadowWarrior:            "ShadowWarrior",
	ClawsOfThunder:           "ClawsOfThunder",
	DragonTail:               "DragonTail",
	LightningSentry:          "LightningSentry",
	WakeOfInferno:            "WakeOfInferno",
	MindBlast:                "MindBlast",
	BladesOfIce:              "BladesOfIce",
	DragonFlight:             "DragonFlight",
	DeathSentry:              "DeathSentry",
	BladeShield:              "BladeShield",
	Venom:                    "Venom",
	ShadowMaster:             "ShadowMaster",
	PhoenixStrike:            "PhoenixStrike",
	WakeOfDestructionSentry:  "WakeOfDestructionSentry",
	ImpInferno:               "ImpInferno",
	ImpFireball:              "ImpFireball",
	BaalTaunt:                "BaalTaunt",
	BaalCorpseExplode:        "BaalCorpseExplode",
	BaalMonsterSpawn:         "BaalMonsterSpawn",
	CatapultChargedBall:      "CatapultChargedBall",
	CatapultSpikeBall:        "CatapultSpikeBall",
	SuckBlood:                "SuckBlood",
	CryHelp:                  "CryHelp",
	HealingVortex:            "HealingVortex",
	Teleport2:                "Teleport2",
	SelfResurrect:            "SelfResurrect",
	VineAttack:               "VineAttack",
	OverseerWhip:             "OverseerWhip",
	BarbsAura:                "BarbsAura",
	WolverineAura:            "WolverineAura",
	OakSageAura:              "OakSageAura",
	ImpFireMissile:           "ImpFireMissile",
	Impregnate:               "Impregnate",
	SiegeBeastStomp:          "SiegeBeastStomp",
	MinionSpawner:            "MinionSpawner",
	CatapultBlizzard:         "CatapultBlizzard",
	CatapultPlague:           "CatapultPlague",
	CatapultMeteor:           "CatapultMeteor",
	BoltSentry:               "BoltSentry",
	CorpseCycler:             "CorpseCycler",
	DeathMaul:                "DeathMaul",
	DefenseCurse:             "DefenseCurse",
	BloodMana:                "BloodMana",
	monInfernoSentry:         "monInfernoSentry",
	monDeathSentry:           "monDeathSentry",
	sentryLightning:          "sentryLightning",
	fenrisRage:               "fenrisRage",
	BaalTentacle:             "BaalTentacle",
	BaalNova:                 "BaalNova",
	BaalInferno:              "BaalInferno",
	BaalColdMissiles:         "BaalColdMissiles",
	MegademonInferno:         "MegademonInferno",
	EvilHutSpawner:           "EvilHutSpawner",
	CountessFirewall:         "CountessFirewall",
	ImpBolt:                  "ImpBolt",
	HorrorArcticBlast:        "HorrorArcticBlast",
	deathSentryLtng:          "deathSentryLtng",
	VineCycler:               "VineCycler",
	BearSmite:                "BearSmite",
	Resurrect2:               "Resurrect2",
	BloodLordFrenzy:          "BloodLordFrenzy",
	BaalTeleport:             "BaalTeleport",
	ImpTeleport:              "ImpTeleport",
	BaalCloneTeleport:        "BaalCloneTeleport",
	ZakarumLightning:         "ZakarumLightning",
	VampireMissile:           "VampireMissile",
	MephistoMissile:          "MephistoMissile",
	DoomKnightMissile:        "DoomKnightMissile",
	RogueMissile:             "RogueMissile",
	HydraMissile:             "HydraMissile",
	NecromageMissile:         "NecromageMissile",
	MonBow:                   "MonBow",
	MonFireArrow:             "MonFireArrow",
	MonColdArrow:             "MonColdArrow",
	MonExplodingArrow:        "MonExplodingArrow",
	MonFreezingArrow:         "MonFreezingArrow",
	MonPowerStrike:           "MonPowerStrike",
	SuccubusBolt:             "SuccubusBolt",
	MephFrostNova:            "MephFrostNova",
	MonIceSpear:              "MonIceSpear",
	ShamanIce:                "ShamanIce",
	Diablogeddon:             "Diablogeddon",
	DeleriumChange:           "DeleriumChange",
	NihlathakCorpseExplosion: "NihlathakCorpseExplosion",
	SerpentCharge:            "SerpentCharge",
	TrapNova:                 "TrapNova",
	UnHolyBoltEx:             "UnHolyBoltEx",
	ShamanFireEx:             "ShamanFireEx",
	ImpFireMissileEx:         "ImpFireMissileEx",
	FixedSiegeBeastStomp:     "FixedSiegeBeastStomp",
	Unset:                    "Unset",
}
