package npc

type ID int

const (
	Skeleton ID = iota
	Returned
	BoneWarrior
	BurningDead
	Horror
	Zombie
	HungryDead
	Ghoul
	DrownedCarcass
	PlagueBearer
	Afflicted
	Tainted
	Misshapen
	Disfigured
	Damned
	FoulCrow
	BloodHawk
	BlackRaptor
	CloudStalker
	Fallen
	Carver
	Devilkin
	DarkOne
	WarpedFallen
	Brute
	Yeti
	Crusher
	WailingBeast
	GargantuanBeast
	SandRaer
	Marauder
	Invader
	Infel
	Assailant
	Gorgon
	Gorgon2
	Gorgon3
	Gorgon4
	Ghost
	Wraith
	Specter
	Apparition
	DarkShape
	DarkHunter
	VileHunter
	DarkStalker
	BlackRogue
	FleshHunter
	DuneBeast
	RockDweller
	JungleHunter
	DoomApe
	TempleGuard
	MoonClan
	NightClan
	BloodClan
	HellClan
	DeathClan
	FallenShaman
	CarverShaman
	DevilkinShaman
	DarkShaman
	WarpedShaman
	QuillRat
	SpikeFiend
	ThornBeast
	RazorSpine
	JungleUrchin
	SandMaggot
	RockWorm
	Devourer
	GiantLamprey
	WorldKiller
	TombViper
	ClawViper
	Salamander
	PitViper
	SerpentMagus
	SandLeaper
	CaveLeaper
	TombCreeper
	TreeLurker
	RazorPitDemon
	Huntress
	SaberCat
	NightTiger
	HellCat
	Itchies
	BlackLocusts
	PlagueBugs
	HellSwarm
	DungSoldier
	SandWarrior
	Scarab
	SteelWeevil
	AlbinoRoach
	DriedCorpse
	Decayed
	Embalmed
	PreservedDead
	Cadaver
	HollowOne
	Guardian
	Unraveler
	HoradrimAncient
	BaalSubjectMummy
	ChaosHorde
	ChaosHorde2
	ChaosHorde3
	ChaosHorde4
	CarrionBird
	UndeadScavenger
	HellBuzzard
	WingedNightmare
	Sucker
	Feeder
	BloodHook
	BloodWing
	Gloam
	SwampGhost
	BurningSoul
	BlackSoul
	Arach
	SandFisher
	PoisonSpinner
	FlameSper
	SperMagus
	ThornedHulk
	BrambleHulk
	Thrasher
	Spikefist
	GhoulLord
	NightLord
	DarkLord
	BloodLord
	Banished
	DesertWing
	Fiend
	Gloombat
	BloodDiver
	DarkFamiliar
	RatMan
	Fetish
	Flayer
	SoulKiller
	StygianDoll
	DeckardCain
	Gheed
	Akara
	Chicken
	Kashya
	Rat
	Rogue
	HellMeteor
	Charsi
	Warriv
	Andariel
	Bird
	Bird2
	Bat
	DarkRanger
	VileArcher
	DarkArcher
	BlackArcher
	FleshArcher
	DarkSpearwoman
	VileLancer
	DarkLancer
	BlackLancer
	FleshLancer
	SkeletonArcher
	ReturnedArcher
	BoneArcher
	BurningDeadArcher
	HorrorArcher
	Warriv2
	Atma
	Drognan
	Fara
	Cow
	SandMaggotYoung
	RockWormYoung
	DevourerYoung
	GiantLampreyYoung
	WorldKillerYoung
	Camel
	Blunderbore
	Gorbelly
	Mauler
	Urdar
	SandMaggotEgg
	RockWormEgg
	DevourerEgg
	GiantLampreyEgg
	WorldKillerEgg
	Act2Male
	Act2Female
	Act2Child
	Greiz
	Elzix
	Geglash
	Jerhyn
	Lysander
	Act2Guard
	Act2Vendor
	Act2Vendor2
	FoulCrowNest
	BloodHawkNest
	BlackVultureNest
	CloudStalkerNest
	Meshif
	Duriel
	UndeadRatMan
	UndeadFetish
	UndeadFlayer
	UndeadSoulKiller
	UndeadStygianDoll
	DarkGuard
	DarkGuard2
	DarkGuard3
	DarkGuard4
	DarkGuard5
	BloodMage
	BloodMage2
	BloodMage3
	BloodMage4
	BloodMage5
	Maggot
	MummyGenerator
	Radament
	FireBeast
	IceGlobe
	LightningBeast
	PoisonOrb
	FlyingScimitar
	Zakarumite
	Faithful
	Zealot
	Sexton
	Cantor
	Heirophant
	Heirophant2
	Mephisto
	Diablo
	DeckardCain2
	DeckardCain3
	DeckardCain4
	SwampDweller
	BogCreature
	SlimePrince
	Summoner
	Tyrael
	Asheara
	Hratli
	Alkor
	Ormus
	Izual
	Halbu
	WaterWatcherLimb
	RiverStalkerLimb
	StygianWatcherLimb
	WaterWatcherHead
	RiverStalkerHead
	StygianWatcherHead
	Meshif2
	DeckardCain5
	Navi
	BloodRaven
	Bug
	Scorpion
	RogueScout
	Rogue2
	Rogue3
	GargoyleTrap
	ReturnedMage
	BoneMage
	BurningDeadMage
	HorrorMage
	RatManShaman
	FetishShaman
	FlayerShaman
	SoulKillerShaman
	StygianDollShaman
	Larva
	SandMaggotQueen
	RockWormQueen
	DevourerQueen
	GiantLampreyQueen
	WorldKillerQueen
	ClayGolem
	BloodGolem
	IronGolem
	FireGolem
	Familiar
	Act3Male
	NightMarauder
	Act3Female
	Natalya
	FleshSpawner
	StygianHag
	Grotesque
	FleshBeast
	StygianDog
	GrotesqueWyrm
	Groper
	Strangler
	StormCaster
	Corpulent
	CorpseSpitter
	MawFiend
	DoomKnight
	AbyssKnight
	OblivionKnight
	QuillBear
	SpikeGiant
	ThornBrute
	RazorBeast
	GiantUrchin
	Snake
	Parrot
	Fish
	EvilHole
	EvilHole2
	EvilHole3
	EvilHole4
	EvilHole5
	FireboltTrap
	HorzMissileTrap
	VertMissileTrap
	PoisonCloudTrap
	LightningTrap
	Kaelan
	InvisoSpawner
	DiabloClone
	SuckerNest
	FeederNest
	BloodHookNest
	BloodWingNest
	Guard
	MiniSper
	BonePrison
	BonePrison2
	BonePrison3
	BonePrison4
	BoneWall
	CouncilMember
	CouncilMember2
	CouncilMember3
	Turret
	Turret2
	Turret3
	Hydra
	Hydra2
	Hydra3
	MeleeTrap
	SevenTombs
	Decoy
	Valkyrie
	Act2Guard3
	IronWolf
	Balrog
	PitLord
	VenomLord
	NecroSkeleton
	NecroMage
	Griswold
	CompellingOrbNpc
	Tyrael2
	DarkWanderer
	NovaTrap
	SpiritMummy
	LightningSpire
	FireTower
	Slinger
	SpearCat
	NightSlinger
	HellSlinger
	Act2Guard4
	Act2Guard5
	ReturnedMage2
	BoneMage2
	BaalColdMage
	HorrorMage2
	ReturnedMage3
	BoneMage3
	BurningDeadMage2
	HorrorMage3
	ReturnedMage4
	BoneMage4
	BurningDeadMage3
	HorrorMage4
	HellBovine
	Window
	Window2
	SpearCat2
	NightSlinger2
	RatMan2
	Fetish2
	Flayer2
	SoulKiller2
	StygianDoll2
	MephistoSpirit
	TheSmith
	TrappedSoul
	TrappedSoul2
	Jamella
	Izual2
	RatMan3
	Malachai
	Hephasto
	WakeOfDestruction
	ChargedBoltSentry
	LightningSentry
	BladeCreeper
	InvisiblePet
	InfernoSentry
	DeathSentry
	ShadowWarrior
	ShadowMaster
	DruHawk
	DruSpiritWolf
	DruFenris
	SpiritOfBarbs
	HeartOfWolverine
	OakSage
	DruPlaguePoppy
	DruCycleOfLife
	VineCreature
	DruBear
	Eagle
	Wolf
	Bear
	BarricadeDoor
	BarricadeDoor2
	PrisonDoor
	BarricadeTower
	RotWalker
	ReanimatedHorde
	ProwlingDead
	UnholyCorpse
	DefiledWarrior
	SiegeBeast
	CrushBiest
	BloodBringer
	GoreBearer
	DeamonSteed
	SnowYeti
	SnowYeti2
	SnowYeti3
	SnowYeti4
	WolfRer
	WolfRer2
	WolfRer3
	MinionExp
	SlayerExp
	IceBoar
	FireBoar
	HellSpawn
	IceSpawn
	GreaterHellSpawn
	GreaterIceSpawn
	FanaticMinion
	BerserkSlayer
	ConsumedIceBoar
	ConsumedFireBoar
	FrenziedHellSpawn
	FrenziedIceSpawn
	InsaneHellSpawn
	InsaneIceSpawn
	SuccubusExp
	VileTemptress
	StygianHarlot
	HellTemptress
	BloodTemptress
	Dominus
	VileWitch
	StygianFury
	BloodWitch
	HellWitch
	OverSeer
	Lasher
	OverLord
	BloodBoss
	HellWhip
	MinionSpawner
	MinionSlayerSpawner
	MinionBoarSpawner
	MinionBoarSpawner2
	MinionSpawnSpawner
	MinionBoarSpawner3
	MinionBoarSpawner4
	MinionSpawnSpawner2
	Imp
	Imp2
	Imp3
	Imp4
	Imp5
	CatapultS
	CatapultE
	CatapultSiege
	CatapultW
	FrozenHorror
	FrozenHorror2
	FrozenHorror3
	FrozenHorror4
	FrozenHorror5
	BloodLord2
	BloodLord3
	BloodLord4
	BloodLord5
	BloodLord6
	Larzuk
	Drehya
	Malah
	NihlathakTown
	QualKehk
	CatapultSpotterS
	CatapultSpotterE
	CatapultSpotterSiegeName
	CatapultSpotterW
	DeckardCain6
	Tyrael3
	Act5Combatant
	Act5Combatant2
	BarricadeWallRight
	BarricadeWallLeft
	Nihlathak
	Drehya2
	EvilHut
	DeathMauler
	DeathMauler2
	DeathMauler3
	DeathMauler4
	DeathMauler5
	POW
	Act5Townguard
	Act5Townguard2
	AncientStatue
	AncientStatueNpc2
	AncientStatueNpc3
	AncientBarbarian
	AncientBarbarian2
	AncientBarbarian3
	BaalThrone
	BaalCrab
	BaalTaunt
	PutrDefiler
	PutrDefiler2
	PutrDefiler3
	PutrDefiler4
	PutrDefiler5
	PainWorm
	PainWorm2
	PainWorm3
	PainWorm4
	PainWorm5
	Bunny
	CouncilMemberBall
	VenomLord2
	BaalCrabToStairs
	Act5Hireling1Hand
	Act5Hireling2Hand
	BaalTentacle
	BaalTentacle2
	BaalTentacle3
	BaalTentacle4
	BaalTentacle5
	InjuredBarbarian
	InjuredBarbarian2
	InjuredBarbarian3
	BaalCrabClone
	BaalsMinion
	BaalsMinion2
	BaalsMinion3
	WorldstoneEffect
	BurningDeadArcher2
	BoneArcher2
	BurningDeadArcher3
	ReturnedArcher2
	HorrorArcher2
	Afflicted2
	Tainted2
	Misshapen2
	Disfigured2
	Damned2
	MoonClan2
	NightClan2
	HellClan2
	BloodClan2
	DeathClan2
	FoulCrow2
	BloodHawk2
	BlackRaptor2
	CloudStalker2
	ClawViper2
	PitViper2
	Salamander2
	TombViper2
	SerpentMagus2
	Marauder2
	Infel2
	SandRaer2
	Invader2
	Assailant2
	DeathMauler6
	QuillRat2
	SpikeFiend2
	RazorSpine2
	CarrionBird2
	ThornedHulk2
	Slinger2
	Slinger3
	Slinger4
	VileArcher2
	DarkArcher2
	VileLancer2
	DarkLancer2
	BlackLancer2
	Blunderbore2
	Mauler2
	ReturnedMage5
	BurningDeadMage4
	ReturnedMage6
	HorrorMage5
	BoneMage5
	HorrorMage6
	HorrorMage7
	Huntress2
	SaberCat2
	CaveLeaper2
	TombCreeper2
	Ghost2
	Wraith2
	Specter2
	SuccubusExp2
	HellTemptress2
	Dominus2
	HellWitch2
	VileWitch2
	Gloam2
	BlackSoul2
	BurningSoul2
	Carver2
	Devilkin2
	DarkOne2
	CarverShaman2
	DevilkinShaman2
	DarkShaman2
	BoneWarrior2
	Returned2
	Gloombat2
	Fiend2
	BloodLord7
	BloodLord8
	Scarab2
	SteelWeevil2
	Flayer3
	StygianDoll3
	SoulKiller3
	Flayer4
	StygianDoll4
	SoulKiller4
	FlayerShaman2
	StygianDollShaman2
	SoulKillerShaman2
	TempleGuard2
	TempleGuard3
	Guardian2
	Unraveler2
	HoradrimAncient2
	HoradrimAncient3
	Zealot2
	Zealot3
	Heirophant3
	Heirophant4
	Grotesque2
	FleshSpawner2
	GrotesqueWyrm2
	FleshBeast2
	WorldKiller2
	WorldKillerYoung2
	WorldKillerEgg2
	SlayerExp2
	HellSpawn2
	GreaterHellSpawn2
	Arach2
	Balrog2
	PitLord2
	Imp6
	Imp7
	UndeadStygianDoll2
	UndeadSoulKiller2
	Strangler2
	StormCaster2
	MawFiend2
	BloodLord9
	GhoulLord2
	DarkLord2
	UnholyCorpse2
	DoomKnight2
	DoomKnight3
	OblivionKnight2
	OblivionKnight3
	Cadaver2
	UberMephisto
	UberDiablo
	UberIzual
	Lilith
	UberDuriel
	UberBaal
	EvilHut2
	DemonHole
	PitLord3
	OblivionKnight4
	Imp8
	HellSwarm2
	WorldKiller3
	Arach3
	SteelWeevil3
	HellTemptress3
	VileWitch3
	FleshHunter2
	DarkArcher3
	BlackLancer3
	HellWhip2
	Returned3
	HorrorArcher3
	BurningDeadMage5
	HorrorMage8
	BoneMage6
	HorrorMage9
	DarkLord3
	Specter3
	BurningSoul3
)
