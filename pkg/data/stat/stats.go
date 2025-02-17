package stat

import (
	"strconv"
	"strings"
)

type ID int16

type Data struct {
	ID    ID
	Value int
	Layer int
}

func (d Data) String() string {
	return strings.Replace(StatStringMap[int(d.ID)][d.Layer], "#", strconv.Itoa(d.Value), 1)
}

type Stats []Data

func (i Stats) FindStat(id ID, layer int) (Data, bool) {
	for _, s := range i {
		if s.ID == id && s.Layer == layer {
			return s, true
		}
	}

	return Data{}, false
}

const (
	Strength ID = iota
	Energy
	Dexterity
	Vitality
	StatPoints
	SkillPoints
	Life
	MaxLife
	Mana
	MaxMana
	Stamina
	MaxStamina
	Level
	Experience
	Gold
	StashGold
	EnhancedDefense
	EnhancedDamageMin
	EnhancedDamage
	AttackRating
	ChanceToBlock
	MinDamage
	MaxDamage
	TwoHandedMinDamage
	TwoHandedMaxDamage
	DamagePercent
	ManaRecovery
	ManaRecoveryBonus
	StaminaRecoveryBonus
	LastExp
	NextExp
	Defense
	DefenseVsMissiles
	DefenseVsHth
	NormalDamageReduction
	MagicDamageReduction
	DamageReduced
	MagicResist
	MaxMagicResist
	FireResist
	MaxFireResist
	LightningResist
	MaxLightningResist
	ColdResist
	MaxColdResist
	PoisonResist
	MaxPoisonResist
	DamageAura
	FireMinDamage
	FireMaxDamage
	LightningMinDamage
	LightningMaxDamage
	MagicMinDamage
	MagicMaxDamage
	ColdMinDamage
	ColdMaxDamage
	ColdLength
	PoisonMinDamage
	PoisonMaxDamage
	PoisonLength
	LifeSteal
	LifeStealMax
	ManaSteal
	ManaStealMax
	StaminaDrainMinDamage
	StaminaDrainMaxDamage
	StunLength
	VelocityPercent
	AttackRate
	OtherAnimRate
	Quantity
	Value
	Durability
	MaxDurability
	ReplenishLife
	MaxDurabilityPercent
	MaxLifePercent
	MaxManaPercent
	AttackerTakesDamage
	GoldFind
	MagicFind
	Knockback
	TimeDuration
	AddClassSkills
	Unused84
	AddExperience
	LifeAfterEachKill
	ReducePrices
	DoubleHerbDuration
	LightRadius
	LightColor
	Requirements
	LevelRequire
	IncreasedAttackSpeed
	LevelRequirePercent
	LastBlockFrame
	FasterRunWalk
	NonClassSkill
	State
	FasterHitRecovery
	PlayerCount
	PoisonOverrideLength
	FasterBlockRate
	BypassUndead
	BypassDemons
	FasterCastRate
	BypassBeasts
	SingleSkill
	SlainMonstersRestInPeace
	CurseResistance
	PoisonLengthReduced
	NormalDamage
	HitCausesMonsterToFlee
	HitBlindsTarget
	DamageTakenGoesToMana
	IgnoreTargetsDefense
	TargetDefense
	PreventMonsterHeal
	HalfFreezeDuration
	AttackRatingPercent
	MonsterDefensePerHit
	DemonDamagePercent
	UndeadDamagePercent
	DemonAttackRating
	UndeadAttackRating
	Throwable
	FireSkills
	AllSkills
	AttackerTakesLightDamage
	IronMaidenLevel
	LifeTapLevel
	ThornsPercent
	BoneArmor
	BoneArmorMax
	FreezesTarget
	OpenWounds
	CrushingBlow
	KickDamage
	ManaAfterKill
	HealAfterDemonKill
	ExtraBlood
	DeadlyStrike
	AbsorbFirePercent
	AbsorbFire
	AbsorbLightningPercent
	AbsorbLightning
	AbsorbMagicPercent
	AbsorbMagic
	AbsorbColdPercent
	AbsorbCold
	SlowsTarget
	Aura
	Indestructible
	CannotBeFrozen
	SlowerStaminaDrain
	Reanimate
	Pierce
	MagicArrow
	ExplosiveArrow
	ThrowMinDamage
	ThrowMaxDamage
	SkillHandofAthena
	SkillStaminaPercent
	SkillPassiveStaminaPercent
	SkillConcentration
	SkillEnchant
	SkillPierce
	SkillConviction
	SkillChillingArmor
	SkillFrenzy
	SkillDecrepify
	SkillArmorPercent
	Alignment
	Target0
	Target1
	GoldLost
	ConverisonLevel
	ConverisonMaxHP
	UnitDooverlay
	AttackVsMonType
	DamageVsMonType
	Fade
	ArmorOverridePercent
	Unused183
	Unused184
	Unused185
	Unused186
	Unused187
	AddSkillTab
	Unused189
	Unused190
	Unused191
	Unused192
	Unused193
	NumSockets
	SkillOnAttack
	SkillOnKill
	SkillOnDeath
	SkillOnHit
	SkillOnLevelUp
	Unused200
	SkillOnGetHit
	Unused202
	Unused203
	ItemChargedSkill
	Unused205
	Unused206
	Unused207
	Unused208
	Unused209
	Unused210
	Unused211
	Unused212
	Unused213
	DefensePerLevel
	ArmorPercentPerLevel
	LifePerLevel
	ManaPerLevel
	MaxDamagePerLevel
	MaxDamagePercentPerLevel
	StrengthPerLevel
	DexterityPerLevel
	EnergyPerLevel
	VitalityPerLevel
	AttackRatingPerLevel
	AttackRatingPercentPerLevel
	ColdDamageMaxPerLevel
	FireDamageMaxPerLevel
	LightningDamageMaxPerLevel
	PoisonDamageMaxPerLevel
	ResistColdPerLevel
	ResistFirePerLevel
	ResistLightningPerLevel
	ResistPoisonPerLevel
	AbsorbColdPerLevel
	AbsorbFirePerLevel
	AbsorbLightningPerLevel
	AbsorbPoisonPerLevel
	ThornsPerLevel
	ExtraGoldPerLevel
	MagicFindPerLevel
	RegenStaminaPerLevel
	StaminaPerLevel
	DamageDemonPerLevel
	DamageUndeadPerLevel
	AttackRatingDemonPerLevel
	AttackRatingUndeadPerLevel
	CrushingBlowPerLevel
	OpenWoundsPerLevel
	KickDamagePerLevel
	DeadlyStrikePerLevel
	FindGemsPerLevel
	ReplenishDurability
	ReplenishQuantity
	ExtraStack
	FindItem
	SlashDamage
	SlashDamagePercent
	CrushDamage
	CrushDamagePercent
	ThrustDamage
	ThrustDamagePercent
	AbsorbSlash
	AbsorbCrush
	AbsorbThrust
	AbsorbSlashPercent
	AbsorbCrushPercent
	AbsorbThrustPercent
	ArmorByTime
	ArmorPercentByTime
	LifeByTime
	ManaByTime
	MaxDamageByTime
	MaxDamagePercentByTime
	StrengthByTime
	DexterityByTime
	EnergyByTime
	VitalityByTime
	AttackRatingByTime
	AttackRatingPercentByTime
	ColdDamageMaxByTime
	FireDamageMaxByTime
	LightningDamageMaxByTime
	PoisonDamageMaxByTime
	ResistColdByTime
	ResistFireByTime
	ResistLightningByTime
	ResistPoisonByTime
	AbsorbColdByTime
	AbsorbFireByTime
	AbsorbLightningByTime
	AbsorbPoisonByTime
	FindGoldByTime
	MagicFindByTime
	RegenStaminaByTime
	StaminaByTime
	DamageDemonByTime
	DamageUndeadByTime
	AttackRatingDemonByTime
	AttackRatingUndeadByTime
	CrushingBlowByTime
	OpenWoundsByTime
	KickDamageByTime
	DeadlyStrikeByTime
	FindGemsByTime
	PierceCold
	PierceFire
	PierceLightning
	PiercePoison
	DamageVsMonster
	DamagePercentVsMonster
	AttackRatingVsMonster
	AttackRatingPercentVsMonster
	AcVsMonster
	AcPercentVsMonster
	FireLength
	BurningMin
	BurningMax
	ProgressiveDamage
	ProgressiveSteal
	ProgressiveOther
	ProgressiveFire
	ProgressiveCold
	ProgressiveLightning
	ExtraCharges
	ProgressiveAttackRating
	PoisonCount
	DamageFrameRate
	PierceIdx
	FireSkillDamage
	LightningSkillDamage
	ColdSkillDamage
	PoisonSkillDamage
	EnemyFireResist
	EnemyLightningResist
	EnemyColdResist
	EnemyPoisonResist
	PassiveCriticalStrike
	PassiveDodge
	PassiveAvoid
	PassiveEvade
	PassiveWarmth
	PassiveMasteryMeleeAttackRating
	PassiveMasteryMeleeDamage
	PassiveMasteryMeleeCritical
	PassiveMasteryThrowAttackRating
	PassiveMasteryThrowDamage
	PassiveMasteryThrowCritical
	PassiveWeaponBlock
	SummonResist
	ModifierListSkill
	ModifierListLevel
	LastSentHPPercent
	SourceUnitType
	SourceUnitID
	ShortParam1
	QuestItemDifficulty
	PassiveMagicMastery
	PassiveMagicPierce
	SkillCooldown
	SkillMissileDamageScale
)

var StringStats = []string{
	"strength",
	"energy",
	"dexterity",
	"vitality",
	"statpoints",
	"skillpoints",
	"life",
	"maxlife",
	"mana",
	"maxmana",
	"stamina",
	"maxstamina",
	"level",
	"experience",
	"gold",
	"stashgold",
	"enhanceddefense",
	"enhanceddamagemin",
	"enhanceddamage",
	"attackrating",
	"chancetoblock",
	"mindamage",
	"maxdamage",
	"twohandedmindamage",
	"twohandedmaxdamage",
	"damagepercent",
	"manarecovery",
	"manarecoverybonus",
	"staminarecoverybonus",
	"lastexp",
	"nextexp",
	"defense",
	"defensevsmissiles",
	"defensevshth",
	"normaldamagereduction",
	"magicdamagereduction",
	"damagereduced",
	"magicresist",
	"maxmagicresist",
	"fireresist",
	"maxfireresist",
	"lightningresist",
	"maxlightningresist",
	"coldresist",
	"maxcoldresist",
	"poisonresist",
	"maxpoisonresist",
	"damageaura",
	"firemindamage",
	"firemaxdamage",
	"lightningmindamage",
	"lightningmaxdamage",
	"magicmindamage",
	"magicmaxdamage",
	"coldmindamage",
	"coldmaxdamage",
	"coldlength",
	"poisonmindamage",
	"poisonmaxdamage",
	"poisonlength",
	"lifesteal",
	"lifestealmax",
	"manasteal",
	"manastealmax",
	"staminadrainmindamage",
	"staminadrainmaxdamage",
	"stunlength",
	"velocitypercent",
	"attackrate",
	"otheranimrate",
	"quantity",
	"value",
	"durability",
	"maxdurability",
	"replenishlife",
	"maxdurabilitypercent",
	"maxlifepercent",
	"maxmanapercent",
	"attackertakesdamage",
	"goldfind",
	"magicfind",
	"knockback",
	"timeduration",
	"addclassskills",
	"unused84",
	"addexperience",
	"lifeaftereachkill",
	"reduceprices",
	"doubleherbduration",
	"lightradius",
	"lightcolor",
	"requirements",
	"levelrequire",
	"increasedattackspeed",
	"levelrequirepercent",
	"lastblockframe",
	"fasterrunwalk",
	"nonclassskill",
	"state",
	"fasterhitrecovery",
	"playercount",
	"poisonoverridelength",
	"fasterblockrate",
	"bypassundead",
	"bypassdemons",
	"fastercastrate",
	"bypassbeasts",
	"singleskill",
	"slainmonstersrestinpeace",
	"curseresistance",
	"poisonlengthreduced",
	"normaldamage",
	"hitcausesmonstertoflee",
	"hitblindstarget",
	"damagetakengoestomana",
	"ignoretargetsdefense",
	"targetdefense",
	"preventmonsterheal",
	"halffreezeduration",
	"attackratingpercent",
	"monsterdefenseperhit",
	"demondamagepercent",
	"undeaddamagepercent",
	"demonattackrating",
	"undeadattackrating",
	"throwable",
	"fireskills",
	"allskills",
	"attackertakeslightdamage",
	"ironmaidenlevel",
	"lifetaplevel",
	"thornspercent",
	"bonearmor",
	"bonearmormax",
	"freezestarget",
	"openwounds",
	"crushingblow",
	"kickdamage",
	"manaafterkill",
	"healafterdemonkill",
	"extrablood",
	"deadlystrike",
	"absorbfirepercent",
	"absorbfire",
	"absorblightningpercent",
	"absorblightning",
	"absorbmagicpercent",
	"absorbmagic",
	"absorbcoldpercent",
	"absorbcold",
	"slowstarget",
	"aura",
	"indestructible",
	"cannotbefrozen",
	"slowerstaminadrain",
	"reanimate",
	"pierce",
	"magicarrow",
	"explosivearrow",
	"throwmindamage",
	"throwmaxdamage",
	"skillhandofathena",
	"skillstaminapercent",
	"skillpassivestaminapercent",
	"skillconcentration",
	"skillenchant",
	"skillpierce",
	"skillconviction",
	"skillchillingarmor",
	"skillfrenzy",
	"skilldecrepify",
	"skillarmorpercent",
	"alignment",
	"target0",
	"target1",
	"goldlost",
	"converisonlevel",
	"converisonmaxhp",
	"unitdooverlay",
	"attackvsmontype",
	"damagevsmontype",
	"fade",
	"armoroverridepercent",
	"unused183",
	"unused184",
	"unused185",
	"unused186",
	"unused187",
	"addskilltab",
	"unused189",
	"unused190",
	"unused191",
	"unused192",
	"unused193",
	"numsockets",
	"skillonattack",
	"skillonkill",
	"skillondeath",
	"skillonhit",
	"skillonlevelup",
	"unused200",
	"skillongethit",
	"unused202",
	"unused203",
	"itemchargedskill",
	"unused205",
	"unused206",
	"unused207",
	"unused208",
	"unused209",
	"unused210",
	"unused211",
	"unused212",
	"unused213",
	"defenseperlevel",
	"armorpercentperlevel",
	"lifeperlevel",
	"manaperlevel",
	"maxdamageperlevel",
	"maxdamagepercentperlevel",
	"strengthperlevel",
	"dexterityperlevel",
	"energyperlevel",
	"vitalityperlevel",
	"attackratingperlevel",
	"attackratingpercentperlevel",
	"colddamagemaxperlevel",
	"firedamagemaxperlevel",
	"lightningdamagemaxperlevel",
	"poisondamagemaxperlevel",
	"resistcoldperlevel",
	"resistfireperlevel",
	"resistlightningperlevel",
	"resistpoisonperlevel",
	"absorbcoldperlevel",
	"absorbfireperlevel",
	"absorblightningperlevel",
	"absorbpoisonperlevel",
	"thornsperlevel",
	"extragoldperlevel",
	"magicfindperlevel",
	"regenstaminaperlevel",
	"staminaperlevel",
	"damagedemonperlevel",
	"damageundeadperlevel",
	"attackratingdemonperlevel",
	"attackratingundeadperlevel",
	"crushingblowperlevel",
	"openwoundsperlevel",
	"kickdamageperlevel",
	"deadlystrikeperlevel",
	"findgemsperlevel",
	"replenishdurability",
	"replenishquantity",
	"extrastack",
	"finditem",
	"slashdamage",
	"slashdamagepercent",
	"crushdamage",
	"crushdamagepercent",
	"thrustdamage",
	"thrustdamagepercent",
	"absorbslash",
	"absorbcrush",
	"absorbthrust",
	"absorbslashpercent",
	"absorbcrushpercent",
	"absorbthrustpercent",
	"armorbytime",
	"armorpercentbytime",
	"lifebytime",
	"manabytime",
	"maxdamagebytime",
	"maxdamagepercentbytime",
	"strengthbytime",
	"dexteritybytime",
	"energybytime",
	"vitalitybytime",
	"attackratingbytime",
	"attackratingpercentbytime",
	"colddamagemaxbytime",
	"firedamagemaxbytime",
	"lightningdamagemaxbytime",
	"poisondamagemaxbytime",
	"resistcoldbytime",
	"resistfirebytime",
	"resistlightningbytime",
	"resistpoisonbytime",
	"absorbcoldbytime",
	"absorbfirebytime",
	"absorblightningbytime",
	"absorbpoisonbytime",
	"findgoldbytime",
	"magicfindbytime",
	"regenstaminabytime",
	"staminabytime",
	"damagedemonbytime",
	"damageundeadbytime",
	"attackratingdemonbytime",
	"attackratingundeadbytime",
	"crushingblowbytime",
	"openwoundsbytime",
	"kickdamagebytime",
	"deadlystrikebytime",
	"findgemsbytime",
	"piercecold",
	"piercefire",
	"piercelightning",
	"piercepoison",
	"damagevsmonster",
	"damagepercentvsmonster",
	"attackratingvsmonster",
	"attackratingpercentvsmonster",
	"acvsmonster",
	"acpercentvsmonster",
	"firelength",
	"burningmin",
	"burningmax",
	"progressivedamage",
	"progressivesteal",
	"progressiveother",
	"progressivefire",
	"progressivecold",
	"progressivelightning",
	"extracharges",
	"progressiveattackrating",
	"poisoncount",
	"damageframerate",
	"pierceidx",
	"fireskilldamage",
	"lightningskilldamage",
	"coldskilldamage",
	"poisonskilldamage",
	"enemyfireresist",
	"enemylightningresist",
	"enemycoldresist",
	"enemypoisonresist",
	"passivecriticalstrike",
	"passivedodge",
	"passiveavoid",
	"passiveevade",
	"passivewarmth",
	"passivemasterymeleeattackrating",
	"passivemasterymeleedamage",
	"passivemasterymeleecritical",
	"passivemasterythrowattackrating",
	"passivemasterythrowdamage",
	"passivemasterythrowcritical",
	"passiveweaponblock",
	"summonresist",
	"modifierlistskill",
	"modifierlistlevel",
	"lastsenthppercent",
	"sourceunittype",
	"sourceunitid",
	"shortparam1",
	"questitemdifficulty",
	"passivemagicmastery",
	"passivemagicpierce",
	"skillcooldown",
	"skillmissiledamagescale",
}

func (s ID) String() string {
	return StringStats[s]
}

// Map of [Stat ID] -> [layer] -> [Stat String]
var StatStringMap = map[int]map[int]string{
	0:  {0: "+# to Strength"},  // strength
	1:  {0: "+# to Energy"},    // energy
	2:  {0: "+# to Dexterity"}, // dexterity
	3:  {0: "+# to Vitality"},  // vitality
	4:  {0: "Stat Points: #"},
	5:  {0: "New Skills: #"},
	6:  {0: "Hit Points: #"},
	7:  {0: "+# to Life"}, // maxhp
	8:  {0: "Mana: #"},
	9:  {0: "+# to Mana"}, // maxmana
	10: {0: "Stamina: #"},
	11: {0: "+# Maximum Stamina"}, // maxstamina
	12: {0: "Level: #"},
	13: {0: "Experience: #"},
	14: {0: "Gold: #"},
	15: {0: "Gold in Bank: #"},
	16: {0: "+#% Enhanced Defense"},            // enhanceddefense
	17: {0: "+#% Enhanced Damage (Max)"},       // enhanceddamagemax
	18: {0: "+#% Enhanced Damage (Min)"},       // enhanceddamagemin
	19: {0: "+# to Attack Rating"},             // tohit
	20: {0: "#% Increased Chance of Blocking"}, // toblock
	21: {
		0: "#-# Damage",           // mindamage
		1: "+# to Minimum Damage", // plusmindamage
	},
	22: {
		0: "#",                    // maxdamage
		1: "+# to Maximum Damage", // plusmaxdamage
	},
	23: {0: "#"}, //minimum dmg for 2hand weapons - secondarymindamage
	24: {0: "#"}, // maximum dmg for 2hand weapons - secondarymaxdamage
	25: {0: "damagepercent ????"},
	26: {0: "manarecovery ????"},
	27: {0: "Regenerate Mana #%"},   // manarecoverybonus
	28: {0: "Heal Stamina Plus #%"}, // staminarecoverybonus
	29: {0: "lastexp ????"},
	30: {0: "nextexp ????"},
	31: {0: "+# Defense"},                      // plusdefense
	32: {0: "# Defense vs. Missile"},           // armorclassvsmissile
	33: {0: "# Defense vs. Melee"},             // armorclassvsmelee
	34: {0: "Damage Reduced by #"},             // damageresist
	35: {0: "Magic Damage Reduced by #"},       // magicdamageresist
	36: {0: "Damage Reduced by #%"},            // damageresist
	37: {0: "Magic Resist +#%"},                // magicresist
	38: {0: "+#% to Maximum Magic Resist"},     // maxmagicresist
	39: {0: "Fire Resist +#%"},                 // fireresist
	40: {0: "+#% to Maximum Fire Resist"},      // maxfireresist
	41: {0: "Lightning Resist +#%"},            // lightresist
	42: {0: "+#% to Maximum Lightning Resist"}, // maxlightresist
	43: {0: "Cold Resist +#%"},                 // coldresist
	44: {0: "+#% to Maximum Cold Resist"},      // maxcoldresist
	45: {0: "Poison Resist +#%"},               // poisonresist
	46: {0: "+#% to Maximum Poison Resist"},    // maxpoisonresist
	47: {0: "damageaura ????"},
	48: {0: "+# to Minimum Fire Damage"},      // firemindam
	49: {0: "+# to Maximum Fire Damage"},      // firemaxdam
	50: {0: "Adds #-# Lightning Damage"},      // lightmindam
	51: {0: "+# to Maximum Lightning Damage"}, // lightmaxdam
	52: {0: "Adds #-# Magic Damage"},          // magicmindam
	53: {0: "+# to Maximum Magic Damage"},     // magicmaxdam
	54: {0: "Adds #-# Cold Damage"},           // coldmindam
	55: {0: "+# to Maximum Cold Damage"},      // coldmaxdam
	56: {0: "Half Freeze Duration"},           // coldlength
	57: {
		0: "+# to Poison Damage",         //poisonmindam
		1: "+# to Maximum Poison Damage", //poisondamage
	},
	58: {0: "+# to Maximum Poison Damage"}, // poisonmaxdam
	59: {0: "Poison Length Reduced by #%"}, // poisonlength
	60: {0: "#% Life stolen per hit"},      // lifeleech
	61: {0: "lifedrainmaxdam ????"},
	62: {0: "#% Mana stolen per hit"}, // manaleech
	63: {0: "manadrainmaxdam ????"},
	64: {0: "stamdrainmindam ????"},
	65: {0: "stamdrainmaxdam ????"},
	66: {0: "stunlength ????"},
	67: {0: "#% Velocity"},
	68: {0: "#"}, // attackrate
	69: {0: "otheranimrate ????"},
	70: {0: "Quantity: #"},
	71: {0: "Value: #"},
	72: {0: "Durability: # of #"},                      // durability
	73: {0: "# Max Durability"},                        // maxdurability
	74: {0: "Replenish Life +#"},                       // hpregen
	75: {0: "Increase Maximum Durability #%"},          //itemmaxdurabilitypercent
	76: {0: "Increase Maximum Life #%"},                //itemmaxmanapercent
	77: {0: "Increase Maximum Mana #%"},                //itemmaxmanapercent
	78: {0: "Attacker Takes Damage of #"},              //itemattackertakesdamage
	79: {0: "#% Extra Gold from Monsters"},             //itemgoldbonus
	80: {0: "#% Better Chance of Getting Magic Items"}, //itemmagicbonus
	81: {0: "Knockback"},                               //itemknockback
	82: {0: "itemtimeduration ????"},
	83: {
		0: "+# to Amazon Skill Levels",      //amazonskills
		1: "+# to Sorceress Skill Levels",   //sorceressskills
		2: "+# to Necromancer Skill Levels", //necromancerskills
		3: "+# to Paladin Skill Levels",     //paladinskills
		4: "+# to Barbarian Skill Levels",   //barbarianskills
		5: "+# to Druid Skill Levels",       //druidskills
		6: "+# to Assassin Skill Levels",    //assassinskills
	},
	84: {0: "unsentparam1 ????"},
	85: {0: "+#% to Experience Gained"},     //itemaddexperience
	86: {0: "+# Life after each Kill"},      //itemhealafterkill
	87: {0: "Reduces all Vendor Prices #%"}, //itemreducedprices
	88: {0: "itemdoubleherbduration ????"},
	89: {0: "+# to Light Radius"}, //itemlightradius
	90: {0: "itemlightcolor ????"},
	91: {0: "Requirements -#%"},           //itemreqpercent
	92: {0: "Required Level: #"},          //itemlevelreq
	93: {0: "+#% Increased Attack Speed"}, //ias
	94: {0: "itemlevelreqpct ????"},
	95: {0: "lastblockframe ????"},
	96: {0: "+#% Faster Run/Walk"}, //frw
	97: {
		0:   "+# to [Skill]",            // itemnonclassskill
		9:   "+# To Critical Strike",    // plusskillcriticalstrike
		22:  "+# To Guided Arrow",       // plusskillguidedarrow
		54:  "+# To Teleport",           // plusskillteleport
		149: "+# To Battle Orders",      // plusskillbattleorders
		155: "+# To Battle Command",     // plusskillbattlecommand
		146: "+# To Battle Cry",         // plusskillbattlecry
		223: "+# To Werewolf",           // plusskillwerewolf
		224: "+# To Lycanthropy",        // plusskilllycanthropy
		227: "+# To Summon Spirit Wolf", // plusskillsummonspiritwolf
		232: "+# To Feral Rage",         // plusskillferalrage
	},
	98:  {0: "state ????"},
	99:  {0: "+#% Faster Hit Recovery"}, // fhr
	100: {0: "monsterplayercount ????"},
	101: {0: "skillpoisonoverridelength ????"},
	102: {0: "+#% Faster Block Rate"}, //fbr
	103: {0: "skillbypassundead ????"},
	104: {0: "skillbypassdemons ????"},
	105: {0: "+#% Faster Cast Rate"}, //fcr
	106: {0: "skillbypassbeasts ????"},
	107: {
		0:   "+# to [Skill] ([Class] only)",              //itemsingleskill
		6:   "+# to Magic Arrow (Amazon only)",           //itemsmagicarrow
		7:   "+# to Fire Arrow (Amazon only)",            //skillmagicarrow
		8:   "+# to Inner Sight (Amazon only)",           //skillinnersight
		9:   "+# to Critical Strike (Amazon only)",       //skillcriticalstrike
		10:  "+# to Jab (Amazon only)",                   //skilljab
		11:  "+# to Cold Arrow (Amazon only)",            //skillcoldarrow
		12:  "+# to Multiple Shot (Amazon only)",         //skillmultipleshot
		13:  "+# to Dodge (Amazon only)",                 //skilldodge
		14:  "+# to Power Strike (Amazon only)",          //skillpowerstrike
		15:  "+# to Poison Javelin (Amazon only)",        //skillpoisonjavelin
		16:  "+# to Exploding Arrow (Amazon only)",       //skillexplodingarrow
		17:  "+# to Slow Missiles (Amazon only)",         //skillslowmissiles
		18:  "+# to Avoid (Amazon only)",                 //skillavoid
		19:  "+# to Impale (Amazon only)",                //skillimpale
		20:  "+# to Lightning Bolt (Amazon only)",        //skilllightningbolt
		21:  "+# to Ice Arrow (Amazon only)",             //skillicearrow
		22:  "+# to Guided Arrow (Amazon only)",          //skillguidedarrow
		23:  "+# to Penetrate (Amazon only)",             //skillpenetrate
		24:  "+# to Charged Strike (Amazon only)",        //skillchargedstrike
		25:  "+# to Plague Javelin (Amazon only)",        //skillplaguejavelin
		26:  "+# to Strafe (Amazon only)",                //skillstrafe
		27:  "+# to Immolation Arrow (Amazon only)",      //skillimmolationarrow
		28:  "+# to Decoy (Amazon only)",                 //skilldecoy
		29:  "+# to Evade (Amazon only)",                 //skillevade
		30:  "+# to Fend (Amazon only)",                  //skillfend
		31:  "+# to Freezing Arrow (Amazon only)",        //skillfreezingarrow
		32:  "+# to Valkyrie (Amazon only)",              //skillvalkyrie
		33:  "+# to Pierce (Amazon only)",                //skillpierce
		34:  "+# to Lightning Strike (Amazon only)",      //skilllightningstrike
		35:  "+# to Lightning Fury (Amazon only)",        //skilllightningfury
		36:  "+# to Fire Bolt (Sorceress only)",          //skillfirebolt
		37:  "+# to Warmth (Sorceress only)",             //skillwarmth
		38:  "+# to Charged Bolt (Sorceress only)",       //skillchargedbolt
		39:  "+# to Ice Bolt (Sorceress only)",           //skillicebolt
		40:  "+# to Frozen Armor (Sorceress only)",       //skillfrozenarmor
		41:  "+# to Inferno (Sorceress only)",            //skillinferno
		42:  "+# to Static Field (Sorceress only)",       //skillstaticfield
		43:  "+# to Telekinesis (Sorceress only)",        //skilltelekinesis
		44:  "+# to Frost Nova (Sorceress only)",         //skillfrostnova
		45:  "+# to Ice Blast (Sorceress only)",          //skilliceblast
		46:  "+# to Blaze (Sorceress only)",              //skillblaze
		47:  "+# to Fireball (Sorceress only)",           //skillfireball
		48:  "+# to Nova (Sorceress only)",               //skillnova
		49:  "+# to Lightning (Sorceress only)",          //skilllightning
		50:  "+# to Shiver Armor (Sorceress only)",       //skillshiverarmor
		51:  "+# to Firewall (Sorceress only)",           //skillfirewall
		52:  "+# to Enchant (Sorceress only)",            //skillenchant
		53:  "+# to Chain Lightning (Sorceress only)",    //skillchainlightning
		54:  "+# to Teleport (Sorceress only)",           //skillteleport
		55:  "+# to Glacial Spike (Sorceress only)",      //skillglacialspike
		56:  "+# to Meteor (Sorceress only)",             //skillmeteor
		57:  "+# to Thunderstorm (Sorceress only)",       //skillthunderstorm
		58:  "+# to Energy Shield (Sorceress only)",      //skillenergyshield
		59:  "+# to Blizzard (Sorceress only)",           //skillblizzard
		60:  "+# to Chilling Armor (Sorceress only)",     //skillchillingarmor
		61:  "+# to Fire Mastery (Sorceress only)",       //skillfiremastery
		62:  "+# to Hydra (Sorceress only)",              //skillhydra
		63:  "+# to Lightning Mastery (Sorceress only)",  //skilllightningmastery
		64:  "+# to Frozen Orb (Sorceress only)",         //skillfrozenorb
		65:  "+# to Cold Mastery (Sorceress only)",       //skillcoldmastery
		66:  "+# to Amplify Damage (Necromancer only)",   //skillamplifydamage
		67:  "+# to Teeth (Necromancer only)",            //skillteeth
		68:  "+# to Bone Armor (Necromancer only)",       //skillbonearmor
		69:  "+# to Skeleton Mastery (Necromancer only)", //skillskeletonmastery
		70:  "+# to Raise Skeleton (Necromancer only)",   //skillraiseskeleton
		71:  "+# to Dim Vision (Necromancer only)",       //skilldimvision
		72:  "+# to Weaken (Necromancer only)",           //skillweaken
		73:  "+# to Poison Dagger (Necromancer only)",    //skillpoisondagger
		74:  "+# to Corpse Explosion (Necromancer only)", //skillcorpseexplosion
		75:  "+# to Clay Golem (Necromancer only)",       //skillclaygolem
		76:  "+# to Iron Maiden (Necromancer only)",      //skillironmaiden
		77:  "+# to Terror (Necromancer only)",           //skillterror
		78:  "+# to Bone Wall (Necromancer only)",        //skillbonewall
		79:  "+# to Golem Mastery (Necromancer only)",    //skillgolemmastery
		80:  "+# to Skeletal Mage (Necromancer only)",    //skillskeletalmage
		81:  "+# to Confuse (Necromancer only)",          //skillconfuse
		82:  "+# to Life Tap (Necromancer only)",         //skilllifetap
		83:  "+# to Poison Explosion (Necromancer only)", //skillpoisonexplosion
		84:  "+# to Bone Spear (Necromancer only)",       //skillbonespear
		85:  "+# to Blood Golem (Necromancer only)",      //skillbloodgolem
		86:  "+# to Attract (Necromancer only)",          //skillattract
		87:  "+# to Decrepify (Necromancer only)",        //skilldecrepify
		88:  "+# to Bone Prison (Necromancer only)",      //skillboneprison
		89:  "+# to Summon Resist (Necromancer only)",    //skillsummonresist
		90:  "+# to Iron Golem (Necromancer only)",       //skillirongolem
		91:  "+# to Lower Resist (Necromancer only)",     //skilllowerresist
		92:  "+# to Poison Nova (Necromancer only)",      //skillpoisonnova
		93:  "+# to Bone Spirit (Necromancer only)",      //skillbonespirit
		94:  "+# to Fire Golem (Necromancer only)",       //skillfiregolem
		95:  "+# to Revive (Necromancer only)",           //skillrevive
		96:  "+# to Sacrifice (Paladin only)",            //skillsacrifice
		97:  "+# to Smite (Paladin only)",                //skillsmite
		98:  "+# to Might (Paladin only)",                //skillmight
		99:  "+# to Prayer (Paladin only)",               //skillprayer
		100: "+# to Resist Fire (Paladin only)",          //skillresistfire
		101: "+# to Holy Bolt (Paladin only)",            //skillholybolt
		102: "+# to Holy Fire (Paladin only)",            //skillholyfire
		103: "+# to Thorns (Paladin only)",               //skillthorns
		104: "+# to Defiance (Paladin only)",             //skilldefiance
		105: "+# to Resist Cold (Paladin only)",          //skillresistcold
		106: "+# to Zeal (Paladin only)",                 //skillzeal
		107: "+# to Charge (Paladin only)",               //skillcharge
		108: "+# to Blessed Aim (Paladin only)",          //skillblessedaim
		109: "+# to Cleansing (Paladin only)",            //skillcleansing
		110: "+# to Resist Lightning (Paladin only)",     //skillresistlightning
		111: "+# to Vengeance (Paladin only)",            //skillvengeance
		112: "+# to Blessed Hammer (Paladin only)",       //skillblessedhammer
		113: "+# to Concentration (Paladin only)",        //skillconcentration
		114: "+# to Holy Freeze (Paladin only)",          //skillholyfreeze
		115: "+# to Vigor (Paladin only)",                //skillvigor
		116: "+# to Conversion (Paladin only)",           //skillconversion
		117: "+# to Holy Shield (Paladin only)",          //skillholyshield
		118: "+# to Holy Shock (Paladin only)",           //skillholyshock
		119: "+# to Sanctuary (Paladin only)",            //skillsanctuary
		120: "+# to Meditation (Paladin only)",           //skillmeditation
		121: "+# to Fist of the Heavens (Paladin only)",  //skillfistoftheheavens
		122: "+# to Fanaticism (Paladin only)",           //skillfanaticism
		123: "+# to Conviction (Paladin only)",           //skillconviction
		124: "+# to Redemption (Paladin only)",           //skillredemption
		125: "+# to Salvation (Paladin only)",            //skillsalvation
		126: "+# to Bash (Barbarian only)",               //skillbash
		127: "+# to Sword Mastery (Barbarian only)",      //skillswordmastery
		128: "+# to Axe Mastery (Barbarian only)",        //skillaxemastery
		129: "+# to Mace Mastery (Barbarian only)",       //skillmacemastery
		130: "+# to Howl (Barbarian only)",               //skillhowl
		131: "+# to Find Potion (Barbarian only)",        //skillfindpotion
		132: "+# to Leap (Barbarian only)",               //skillleap
		133: "+# to Double Swing (Barbarian only)",       //skilldoubleswing
		134: "+# to Polearm Mastery (Barbarian only)",    //skillpolearmmastery
		135: "+# to Throwing Mastery (Barbarian only)",   //skillthrowingmastery
		136: "+# to Spear Mastery (Barbarian only)",      //skillspearmastery
		137: "+# to Taunt (Barbarian only)",              //skilltaunt
		138: "+# to Shout (Barbarian only)",              //skillshout
		139: "+# to Stun (Barbarian only)",               //skillstun
		140: "+# to Double Throw (Barbarian only)",       //skilldoublethrow
		141: "+# to Increased Stamina (Barbarian only)",  //skillincreasedstamina
		142: "+# to Find Item (Barbarian only)",          //skillfinditem
		143: "+# to Leap Attack (Barbarian only)",        //skillleapattack
		144: "+# to Concentrate (Barbarian only)",        //skillconcentrate
		145: "+# to Iron Skin (Barbarian only)",          //skillironskin
		146: "+# to Battle Cry (Barbarian only)",         //skillbattlecry
		147: "+# to Frenzy (Barbarian only)",             //skillfrenzy
		148: "+# to Increased Speed (Barbarian only)",    //skillincreasedspeed
		149: "+# to Battle Orders (Barbarian only)",      //skillbattleorders
		150: "+# to Grim Ward (Barbarian only)",          //skillgrimward
		151: "+# to Whirlwind (Barbarian only)",          //skillwhirlwind
		152: "+# to Berserk (Barbarian only)",            //skillberserk
		153: "+# to Natural Resistance (Barbarian only)", //skillnaturalresistance
		154: "+# to War Cry (Barbarian only)",            //skillwarcry
		155: "+# to Battle Command (Barbarian only)",     //skillbattlecommand
		221: "+# to Raven (Druid only)",                  //skillraven
		222: "+# to Poison Creeper (Druid only)",         //skillpoisoncreeper
		223: "+# to Werewolf (Druid only)",               //skillwerewolf
		224: "+# to Lycanthropy (Druid only)",            //skilllycanthropy
		225: "+# to Firestorm (Druid only)",              //skillfirestorm
		226: "+# to Oak Sage (Druid only)",               //skilloaksage
		227: "+# to Summon Spirit Wolf (Druid only)",     //skillsummonspiritwolf
		228: "+# to Werebear (Druid only)",               //skillwerebear
		229: "+# to Molten Boulder (Druid only)",         //skillmoltenboulder
		230: "+# to Arctic Blast (Druid only)",           //skillarcticblast
		231: "+# to Carrion Vine (Druid only)",           //skillcarrionvine
		232: "+# to Feral Rage (Druid only)",             //skillferalrage
		233: "+# to Maul (Druid only)",                   //skillmaul
		234: "+# to Fissure (Druid only)",                //skillfissure
		235: "+# to Cyclone Armor (Druid only)",          //skillcyclonearmor
		236: "+# to Heart of Wolverine (Druid only)",     //skillheartofwolverine
		237: "+# to Summon Dire Wolf (Druid only)",       //skillsummondirewolf
		238: "+# to Rabies (Druid only)",                 //skillrabies
		239: "+# to Fire Claws (Druid only)",             //skillfireclaws
		240: "+# to Twister (Druid only)",                //skilltwister
		241: "+# to Solar Creeper (Druid only)",          //skillsolarcreeper
		242: "+# to Hunger (Druid only)",                 //skillhunger
		243: "+# to Shockwave (Druid only)",              //skillshockwave
		244: "+# to Volcano (Druid only)",                //skillvolcano
		245: "+# to Tornado (Druid only)",                //skilltornado
		246: "+# to Spirit of Barbs (Druid only)",        //skillspiritofbarbs
		247: "+# to Summon Grizzly (Druid only)",         //skillsummongrizzly
		248: "+# to Fury (Druid only)",                   //skillfury
		249: "+# to Armageddon (Druid only)",             //skillarmageddon
		250: "+# to Hurricane (Druid only)",              //skillhurricane
		251: "+# to Fire Blast (Assassin only)",          //skillfireblast
		252: "+# to Claw Mastery (Assassin only)",        //skillclawmastery
		253: "+# to Psychic Hammer (Assassin only)",      //skillpsychichammer
		254: "+# to Tiger Strike (Assassin only)",        //skilltigerstrike
		255: "+# to Dragon Talon (Assassin only)",        //skilldragontalon
		256: "+# to Shock Web (Assassin only)",           //skillshockweb
		257: "+# to Blades Sentinel (Assassin only)",     //skillbladesentinel
		258: "+# to Burst of Speed (Assassin only)",      //skillburstofspeed
		259: "+# to Fists of Fire (Assassin only)",       //skillfistsoffire
		260: "+# to Dragon Claw (Assassin only)",         //skilldragonclaw
		261: "+# to Charged Bolt Sentry (Assassin only)", //skillchargedboltsentry
		262: "+# to Wake of Fire (Assassin only)",        //skillwakeoffire
		263: "+# to Weapon Block (Assassin only)",        //skillweaponblock
		264: "+# to Cloak of Shadows (Assassin only)",    //skillcloakofshadows
		265: "+# to Cobra Strike (Assassin only)",        //skillcobrastrike
		266: "+# to Blade Fury (Assassin only)",          //skillbladefury
		267: "+# to Fade (Assassin only)",                //skillfade
		268: "+# to Shadow Warrior (Assassin only)",      //skillshadowwarrior
		269: "+# to Claws of Thunder (Assassin only)",    //skillclawsofthunder
		270: "+# to Dragon Tail (Assassin only)",         //skilldragontail
		271: "+# to Lightning Sentry (Assassin only)",    //skilllightningsentry
		272: "+# to Wake of Inferno (Assassin only)",     //skillwakeofinferno
		273: "+# to Mind Blast (Assassin only)",          //skillmindblast
		274: "+# to Blades of Ice (Assassin only)",       //skillbladesofice
		275: "+# to Dragon Flight (Assassin only)",       //skilldragonflight
		276: "+# to Death Sentry (Assassin only)",        //skilldeathsentry
		277: "+# to Blade Shield (Assassin only)",        //skillbladeshield
		278: "+# to Venom (Assassin only)",               //skillvenom
		279: "+# to Shadow Master (Assassin only)",       //skillshadowmaster
		280: "+# to Phoenix Strike (Assassin only)",      //skillphoenixstrike
	},
	108: {0: "Slain Monsters Rest in Peace"}, //itemrestinpeace
	109: {0: "curseresistance ????"},
	110: {0: "Poison Length Reduced by #%"},        //itempoisonlengthresist
	111: {0: "Damage +#"},                          //itemnormaldamage
	112: {0: "Hit Causes Monster to Flee #%"},      //itemhowl
	113: {0: "Hit Blinds Target +#"},               //itemstupidity
	114: {0: "#% Damage Taken Goes To Mana"},       //itemdamagetomana
	115: {0: "Ignore Target's Defense"},            //itemignoretargetac
	116: {0: "-#% Target Defense"},                 //itemfractionaltargetac
	117: {0: "Prevent Monster Heal"},               //itempreventheal
	118: {0: "Half Freeze Duration"},               //itemhalffreezeduration
	119: {0: " #% Bonus to Attack Rating"},         //itemtohitpercent
	120: {0: "-# to Monster Defense Per Hit"},      //itemdamagetargetac
	121: {0: "+#% Damage to Demons"},               //itemdemondamagepercent
	122: {0: "+#% Damage to Undead"},               //itemundeaddamagepercent
	123: {0: "+# to Attack Rating against Demons"}, //itemdemontohit
	124: {0: "+# to Attack Rating against Undead"}, //itemundeadtohit
	125: {0: "itemthrowable ????"},
	126: {0: "+# to Fire Skills"},                    //itemelemskill
	127: {0: "+# to All Skills"},                     //itemallskills
	128: {0: "Attacker Takes Lightning Damage of #"}, //itemattackertakeslightdamage
	129: {0: "ironmaidenlevel ????"},
	130: {0: "lifetaplevel ????"},
	131: {0: "thornspercent ????"},
	132: {0: "bonearmor ????"},
	133: {0: "bonearmormax ????"},
	134: {0: "Freezes Target +#"},             //itemfreeze
	135: {0: "#% Chance of Open Wounds"},      //itemopenwounds
	136: {0: "#% Chance of Crushing Blow"},    //itemcrushingblow
	137: {0: "+# Kick Damage"},                //itemkickdamage
	138: {0: "+# to Mana after each Kill"},    //itemmanaafterkill
	139: {0: "+# Life after each Demon Kill"}, //itemhealafterdemonkill
	140: {0: "itemextrablood"},
	141: {0: "#% Deadly Strike"},    //itemdeadlystrike
	142: {0: "+# Fire Absorb"},      //itemabsorbfire
	143: {0: "Fire Absorb #%"},      //itemabsorbfirepercent
	144: {0: "+# Lightning Absorb"}, //itemabsorblight
	145: {0: "Lightning Absorb #%"}, //itemabsorblightpercent
	146: {0: "Magic Absorb #%"},     //itemabsorbmagic
	147: {0: "+# Magic Absorb"},     //itemabsorbmagic
	148: {0: "Cold Absorb #%"},      //itemabsorbcoldpercent
	149: {0: "+# Cold Absorb"},      //itemabsorbcold
	150: {0: "Slows Target by #%"},  //itemslow
	151: {
		0:   "Level # [Skill] Aura When Equipped",       // itemaura
		98:  "Level # Might Aura When Equipped",         // mightaura
		102: "Level # Holy Fire Aura When Equipped",     // holyfireaura
		103: "Level # Thorns Aura When Equipped",        // thornsaura
		104: "Level # Defiance Aura When Equipped",      // defianceaura
		113: "Level # Concentration Aura When Equipped", // concentrationaura
		114: "Level # Holy Freeze Aura When Equipped",   // holyfreezeaura
		115: "Level # Vigor Aura When Equipped",         // vigoraura
		118: "Level # Holy Shock Aura When Equipped",    // holyshockaura
		119: "Level # Sanctuary Aura When Equipped",     // sanctuaryaura
		120: "Level # Meditation Aura When Equipped",    // meditationaura
		122: "Level # Fanaticism Aura When Equipped",    // fanaticismaura
		123: "Level # Conviction Aura When Equipped",    // convictionaura
		124: "Level # Redemption Aura When Equipped",    // redemptionaura
	},
	152: {0: "Indestructible"},                  // itemindestructible
	153: {0: "Cannot be Frozen"},                // itemcannotbefrozen
	154: {0: "#% Slower Stamina Drain"},         // itemstaminadrainpct
	155: {0: "Reanimate As: [Returned]"},        // itemreanimate
	156: {0: "Piercing Attack"},                 // itempierce
	157: {0: "Fires Magic Arrows"},              // itemmagicarrow
	158: {0: "Fires Explosive Arrows or Bolts"}, // itemexplosivearrow
	159: {0: "# To Minimum Damage"},
	160: {0: "# Throw Damage"},
	161: {0: "itemskillhandofathena ????"},
	162: {0: "itemskillstaminapercent ????"},
	163: {0: "itemskillpassivestaminapercent ????"},
	164: {0: "itemskillconcentration ????"},
	165: {0: "itemskillenchant ????"},
	166: {0: "itemskillpierce ????"},
	167: {0: "itemskillconviction ????"},
	168: {0: "itemskillchillingarmor ????"},
	169: {0: "itemskillfrenzy ????"},
	170: {0: "itemskilldecrepify ????"},
	171: {0: "itemskillarmorpercent ????"},
	172: {0: "alignment ????"},
	173: {0: "target0 ????"},
	174: {0: "target1 ????"},
	175: {0: "goldlost ????"},
	176: {0: "conversionlevel ????"},
	177: {0: "conversionmaxhp ????"},
	178: {0: "unitdooverlay ????"},
	179: {0: "attackvsmontype ????"},
	180: {0: "damagevsmontype ????"},
	181: {0: "fade ????"},
	182: {0: "armoroverridepercent ????"},
	183: {0: "unused183 ????"},
	184: {0: "unused184 ????"},
	185: {0: "unused185 ????"},
	186: {0: "unused186 ????"},
	187: {0: "Monster Cold Immunity is Sundered"}, //itempiercecoldimmunity
	188: {
		0:  "+# to Bow and Crossbow Skills (Amazon only)",     // bowandcrossbowskilltab
		1:  "+# to Passive and Magic Skills (Amazon only)",    // passiveandmagicskilltab
		2:  "+# to Javelin and Spears Skills (Amazon only)",   // javelinandspearskilltab
		8:  "+# to Fire Skills (Sorceress only)",              // fireskilltab
		9:  "+# to Lightning Skills (Sorceress only)",         // lightningskilltab
		10: "+# to Cold Skills (Sorceress only)",              // coldskilltab
		16: "+# to Curses Skills (Necromancer only)",          // cursesskilltab
		17: "+# to Poison and Bone Skills (Necromancer only)", // poisonandboneskilltab
		18: "+# to Summoning Skills (Necromancer only)",       // necromancersummoningskilltab
		24: "+# to Paladin Combat Skills (Paladin only)",      // palicombatskilltab
		25: "+# to Offensive Aura Skills (Paladin only)",      // offensiveaurasskilltab
		26: "+# to Defensive Aura Skills (Paladin only)",      // defensiveaurasskilltab
		32: "+# to Barbarian Combat Skills (Barbarian only)",  // barbcombatskilltab
		33: "+# to Mastery Skills (Barbarian only)",           // masteryesskilltab
		34: "+# to War Cry Skills (Barbarian only)",           // warcriesskilltab
		40: "+# to Druid Summoning Skills (Druid only)",       // druidsummoningskilltab
		41: "+# to Shapeshifting Skills (Druid only)",         // shapeshiftingskilltab
		42: "+# to Elemental Skills (Druid only)",             // elementalskilltab
		48: "+# to Traps Skills (Assassin only)",              // trapsskilltab
		49: "+# to Shadow Discipline Skills (Assassin only)",  // shadowdisciplinesskilltab
		50: "+# to Martial Arts Skills (Assassin only)",       // martialartsskilltab
	},
	189: {0: "Monster Fire Immunity is Sundered"},      // itempiercefireimmunity
	190: {0: "Monster Lightning Immunity is Sundered"}, // itempiercelightimmunity
	191: {0: "Monster Poison Immunity is Sundered"},    // itempiercepoisonimmunity
	192: {0: "Monster Physical Immunity is Sundered"},  // itempiercedamageimmunity
	193: {0: "Monster Magic Immunity is Sundered"},     // itempiercemagicimmunity
	194: {0: "Socketed (#)"},
	195: {
		1:    "#% Chance to cast level # [Skill] on attack", //itemskillonattack
		2:    "itemskillonattacklevel ????",
		3395: "#% Chance to cast level 3 Chain Lightning on attack",
	},
	196: {
		1: "#% Chance to cast level # [Skill] when you Kill an Enemy", //itemskillonkill
		2: "itemskillonkilllevel ????",
	},
	197: {
		1: "#% Chance to cast level # [Skill] when you Die", //itemskillondeath
		2: "itemskillondeathlevel ????",
	},
	198: {
		1:    "#% Chance to cast level # [Skill] on striking", //itemskillonhit
		2:    "itemskillonhitlevel ????",
		4225: "Amplify Damage on Hit", //amplifydamageonhit
	},
	199: {
		1: "#% Chance to cast level # [Skill] when you Level-Up",
		2: "itemskillonleveluplevel ????",
	},
	200: {0: "unused200 ????"},
	201: {
		1:    "#% Chance to cast level # [Skill] when struck", //itemskillongethit
		2:    "itemskillongethitlevel ????",
		5903: "#% Chance to cast level 15 Poison Nova when struck",
		7751: "#% Chance to cast level 7 Fist of Heavens when struck",
	},
	202: {0: "unused202 ????"},
	203: {0: "unused203 ????"},
	204: {
		1:     "itemchargedskill ????",
		2:     "itemchargedskilllevel ????",
		3461:  "Teleport (charged)",
		17795: "Venom level 3 (charged)",
	},
	205: {0: "unused204 ????"},
	206: {0: "unused205 ????"},
	207: {0: "unused206 ????"},
	208: {0: "unused207 ????"},
	209: {0: "unused208 ????"},
	210: {0: "unused209 ????"},
	211: {0: "unused210 ????"},
	212: {0: "unused211 ????"},
	213: {0: "unused212 ????"},
	214: {0: "+# Defense (Based on Character Level)"},                              // itemarmorperlevel
	215: {0: "+#% Enhanced Defense (Based on Character Level)"},                    // itemarmorpercentperlevel
	216: {0: "+# to Life (Based on Character Level)"},                              // itemmanaperlevel
	217: {0: "+# to Mana (Based on Character Level)"},                              // itemmanaperlevel
	218: {0: "+# to Maximum Damage (Based on Character Level)"},                    // itemmaxdamageperlevel
	219: {0: "+#% Enhanced Maximum Damage (Based on Character Level)"},             // itemmaxdamagepercentperlevel
	220: {0: "+# to Strength (Based on Character Level)"},                          // itemstrengthperlevel
	221: {0: "+# to Dexterity (Based on Character Level)"},                         // itemdexterityperlevel
	222: {0: "+# to Energy (Based on Character Level)"},                            // itemenergyperlevel
	223: {0: "+# to Vitality (Based on Character Level)"},                          // itemvitalityperlevel
	224: {0: "+# to Attack Rating (Based on Character Level)"},                     // itemtohitperlevel
	225: {0: "#% Bonus to Attack Rating (Based on Character Level)"},               // itemtohitpercentperlevel
	226: {0: "+# to Maximum Cold Damage (Based on Character Level)"},               // itemcolddamagemaxperlevel
	227: {0: "+# to Maximum Fire Damage (Based on Character Level)"},               // itemfiredamagemaxperlevel
	228: {0: "+# to Maximum Lightning Damage (Based on Character Level)"},          // itemltngdamagemaxperlevel
	229: {0: "+# to Maximum Poison Damage (Based on Character Level)"},             // itempoisdamagemaxperlevel
	230: {0: "Cold Resist +#% (Based on Character Level)"},                         // itemresistcoldperlevel
	231: {0: "Fire Resist +#% (Based on Character Level)"},                         // itemresistfireperlevel
	232: {0: "Lightning Resist +#% (Based on Character Level)"},                    // itemresistltngperlevel
	233: {0: "Poison Resist +#% (Based on Character Level)"},                       // itemresistpoisperlevel
	234: {0: "Absorbs Cold Damage (Based on Character Level)"},                     // itemabsorbcoldperlevel
	235: {0: "Absorbs Fire Damage (Based on Character Level)"},                     // itemabsorbfireperlevel
	236: {0: "Absorbs Lightning Damage (Based on Character Level)"},                // itemabsorbltngperlevel
	237: {0: "Absorbs Poison Damage (Based on Character Level)"},                   // itemabsorbpoisperlevel
	238: {0: "Attacker Takes Damage of # (Based on Character Level)"},              // itemthornsperlevel
	239: {0: "#% Extra Gold from Monsters (Based on Character Level)"},             // itemfindgoldperlevel
	240: {0: "#% Better Chance of Getting Magic Items (Based on Character Level)"}, // itemfindmagicperlevel
	241: {0: "Heal Stamina Plus #% (Based on Character Level)"},                    // itemregenstaminaperlevel
	242: {0: "+# Maximum Stamina (Based on Character Level)"},                      // itemstaminaperlevel
	243: {0: "+#% Damage to Demons (Based on Character Level)"},                    // itemdamagedemonperlevel
	244: {0: "+#% Damage to Undead (Based on Character Level)"},                    // itemdamageundeadperlevel
	245: {0: "+# to Attack Rating against Demons (Based on Character Level)"},      // itemtohitdemonperlevel
	246: {0: "+# to Attack Rating against Undead (Based on Character Level)"},      // itemtohitundeadperlevel
	247: {0: "#% Chance of Crushing Blow (Based on Character Level)"},              // itemcrushingblowperlevel
	248: {0: "#% Chance of Open Wounds (Based on Character Level)"},                // itemopenwoundsperlevel
	249: {0: "+# Kick Damage (Based on Character Level)"},                          // itemkickdamageperlevel
	250: {0: "#% Deadly Strike (Based on Character Level)"},                        // itemdeadlystrikeperlevel
	251: {0: "itemfindgemsperlevel ????"},
	252: {0: "Repairs 1 durability in # seconds"}, // itemreplenishdurability
	253: {0: "Replenishes quantity"},              // itemreplenishquantity
	254: {0: "Increased Stack Size"},              // itemextrastack
	255: {0: "itemfinditem ????"},
	256: {0: "itemslashdamage ????"},
	257: {0: "itemslashdamagepercent ????"},
	258: {0: "itemcrushdamage ????"},
	259: {0: "itemcrushdamagepercent ????"},
	260: {0: "itemthrustdamage ????"},
	261: {0: "itemthrustdamagepercent ????"},
	262: {0: "itemabsorbslash ????"},
	263: {0: "itemabsorbcrush ????"},
	264: {0: "itemabsorbthrust ????"},
	265: {0: "itemabsorbslashpercent ????"},
	266: {0: "itemabsorbcrushpercent ????"},
	267: {0: "itemabsorbthrustpercent ????"},
	268: {0: "+# Defense (Increases near [Day/Dusk/Night/Dawn])"},                              // itemarmorbytime
	269: {0: "+#% Enhanced Defense (Increases near [Day/Dusk/Night/Dawn])"},                    // itemarmorpercentbytime
	270: {0: "+# to Life (Increases near [Day/Dusk/Night/Dawn])"},                              // itemhpbytime
	271: {0: "+# to Mana (Increases near [Day/Dusk/Night/Dawn])"},                              // itemmanabytime
	272: {0: "+# to Maximum Damage (Increases near [Day/Dusk/Night/Dawn])"},                    // itemmaxdamagebytime
	273: {0: "+#% Enhanced Maximum Damage (Increases near [Day/Dusk/Night/Dawn])"},             // itemmaxdamagepercentbytime
	274: {0: "+# to Strength (Increases near [Day/Dusk/Night/Dawn])"},                          // itemstrengthbytime
	275: {0: "+# to Dexterity (Increases near [Day/Dusk/Night/Dawn])"},                         // itemdexteritybytime
	276: {0: "+# to Energy (Increases near [Day/Dusk/Night/Dawn])"},                            // itemenergybytime
	277: {0: "+# to Vitality (Increases near [Day/Dusk/Night/Dawn])"},                          // itemvitalitybytime
	278: {0: "+# to Attack Rating (Increases near [Day/Dusk/Night/Dawn])"},                     // itemtohitbytime
	279: {0: "+#% Bonus to Attack Rating (Increases near [Day/Dusk/Night/Dawn])"},              // itemtohitpercentbytime
	280: {0: "+# to Maximum Cold Damage (Increases near [Day/Dusk/Night/Dawn])"},               // itemcolddamagemaxbytime
	281: {0: "+# to Maximum Fire Damage (Increases near [Day/Dusk/Night/Dawn])"},               // itemfiredamagemaxbytime
	282: {0: "+# to Maximum Lightning Damage (Increases near [Day/Dusk/Night/Dawn])"},          // itemltngdamagemaxbytime
	283: {0: "+# to Maximum Poison Damage (Increases near [Day/Dusk/Night/Dawn])"},             // itempoisdamagemaxbytime
	284: {0: "Cold Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},                         // itemresistcoldbytime
	285: {0: "Fire Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},                         // itemresistfirebytime
	286: {0: "Lightning Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},                    // itemresistltngbytime
	287: {0: "Poison Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},                       // itemresistpoisbytime
	288: {0: "Absorbs Cold Damage (Increases near [Day/Dusk/Night/Dawn])"},                     // itemabsorbcoldbytime
	289: {0: "Absorbs Fire Damage (Increases near [Day/Dusk/Night/Dawn])"},                     // itemabsorbfirebytime
	290: {0: "Absorbs Lightning Damage (Increases near [Day/Dusk/Night/Dawn])"},                // itemabsorbltngbytime
	291: {0: "Absorbs Poison Damage (Increases near [Day/Dusk/Night/Dawn])"},                   // itemabsorbpoisbytime
	292: {0: "#% Extra Gold from Monsters (Increases near [Day/Dusk/Night/Dawn])"},             // itemfindgoldbytime
	293: {0: "#% Better Chance of Getting Magic Items (Increases near [Day/Dusk/Night/Dawn])"}, // itemfindmagicbytime
	294: {0: "Heal Stamina Plus #% (Increases near [Day/Dusk/Night/Dawn])"},                    // itemregenstaminabytime
	295: {0: "+# Maximum Stamina (Increases near [Day/Dusk/Night/Dawn])"},                      // itemstaminabytime
	296: {0: "+#% Damage to Demons (Increases near [Day/Dusk/Night/Dawn])"},                    // itemdamagedemonbytime
	297: {0: "+#% Damage to Undead (Increases near [Day/Dusk/Night/Dawn])"},                    // itemdamageundeadbytime
	298: {0: "+# to Attack Rating against Demons (Increases near [Day/Dusk/Night/Dawn])"},      // itemtohitdemonbytime
	299: {0: "+# to Attack Rating against Undead (Increases near [Day/Dusk/Night/Dawn])"},      // itemtohitundeadbytime
	300: {0: "#% Chance of Crushing Blow (Increases near [Day/Dusk/Night/Dawn])"},              // itemcrushingblowbytime
	301: {0: "#% Chance of Open Wounds (Increases near [Day/Dusk/Night/Dawn])"},                // itemopenwoundsbytime
	302: {0: "+# Kick Damage (Increases near [Day/Dusk/Night/Dawn])"},                          // itemkickdamagebytime
	303: {0: "#% Deadly Strike (Increases near [Day/Dusk/Night/Dawn])"},                        // itemdeadlystrikebytime
	304: {0: "itemfindgemsbytime ????"},
	305: {0: "-#% to Enemy Cold Resistance"},      // itempiercecold
	306: {0: "-#% to Enemy Fire Resistance"},      // itempiercefire
	307: {0: "-#% to Enemy Lightning Resistance"}, // itempierceltng
	308: {0: "-#% to Enemy Poison Resistance"},    // itempiercepois
	309: {0: "itemdamagevsmonster ????"},
	310: {0: "itemdamagepercentvsmonster ????"},
	311: {0: "itemtohitvsmonster ????"},
	312: {0: "itemtohitpercentvsmonster ????"},
	313: {0: "itemacvsmonster ????"},
	314: {0: "itemacpercentvsmonster ????"},
	315: {0: "firelength ????"},
	316: {0: "burningmin ????"},
	317: {0: "burningmax ????"},
	318: {0: "progressivedamage ????"},
	319: {0: "progressivesteal ????"},
	320: {0: "progressiveother ????"},
	321: {0: "progressivefire ????"},
	322: {0: "progressivecold ????"},
	323: {0: "progressivelightning ????"},
	324: {0: "itemextracharges ????"},
	325: {0: "progressivetohit ????"},
	326: {0: "poisoncount ????"},
	327: {0: "damageframerate ????"},
	328: {0: "pierceidx ????"},
	329: {0: "+#% to Fire Skill Damage"},          // passivefiremastery
	330: {0: "+#% to Lightning Skill Damage"},     // passiveltngmastery
	331: {0: "+#% to Cold Skill Damage"},          // passivecoldmastery
	332: {0: "+#% to Poison Skill Damage"},        // passivepoismastery
	333: {0: "-#% to Enemy Fire Resistance"},      // passivefirepierce
	334: {0: "-#% to Enemy Lightning Resistance"}, // passiveltngpierce
	335: {0: "-#% to Enemy Cold Resistance"},      // passivecoldpierce
	336: {0: "-#% to Enemy Poison Resistance"},    // passivepoispierce
	337: {0: "passivecriticalstrike ????"},
	338: {0: "passivedodge ????"},
	339: {0: "passiveavoid ????"},
	340: {0: "passiveevade ????"},
	341: {0: "passivewarmth ????"},
	342: {0: "passivemasterymeleeth ????"},
	343: {0: "passivemasterymeleedmg ????"},
	344: {0: "passivemasterymeleecrit ????"},
	345: {0: "passivemasterythrowth ????"},
	346: {0: "passivemasterythrowdmg ????"},
	347: {0: "passivemasterythrowcrit ????"},
	348: {0: "passiveweaponblock ????"},
	349: {0: "passivesummonresist ????"},
	350: {0: "modifierlistskill ????"},
	351: {0: "modifierlistlevel ????"},
	352: {0: "lastsenthppct ????"},
	353: {0: "sourceunittype ????"},
	354: {0: "sourceunitid ????"},
	355: {0: "shortparam1 ????"},
	356: {0: "questitemdifficulty ????"},
	357: {0: "passivemagmastery ????"},
	358: {0: "passivemagpierce ????"},
	359: {0: "skillcooldown ????"},
	360: {0: "skillmissiledamagescale ????"},
	555: {0: "All Resistances +#"},
}
