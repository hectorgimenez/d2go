package area

type ID int

type Area struct {
	ID
	Name string
}

func (a ID) IsTown() bool {
	switch a {
	case RogueEncampment, LutGholein, KurastDocks, ThePandemoniumFortress, Harrogath:
		return true
	}

	return false
}

func (a ID) CanBeTerrorized() bool {
	_, canBeTerrorized := CanBeTerrorized[a]

	return canBeTerrorized
}

func (a ID) Act() int {
	if a < 40 {
		return 1
	}
	if a >= 40 && a < 75 {
		return 2
	}
	if a >= 75 && a < 103 {
		return 3
	}
	if a >= 103 && a < 109 {
		return 4
	}

	return 5
}

func (a ID) Area() Area {
	return Areas[a]
}

const (
	Abaddon                  ID = 125
	AncientTunnels           ID = 65
	ArcaneSanctuary          ID = 74
	ArreatPlateau            ID = 112
	ArreatSummit             ID = 120
	Barracks                 ID = 28
	BlackMarsh               ID = 6
	BloodMoor                ID = 2
	BloodyFoothills          ID = 110
	BurialGrounds            ID = 17
	CanyonOfTheMagi          ID = 46
	CatacombsLevel1          ID = 34
	CatacombsLevel2          ID = 35
	CatacombsLevel3          ID = 36
	CatacombsLevel4          ID = 37
	Cathedral                ID = 33
	CaveLevel1               ID = 9
	CaveLevel2               ID = 13
	ChaosSanctuary           ID = 108
	CityOfTheDamned          ID = 106
	ClawViperTempleLevel1    ID = 58
	ClawViperTempleLevel2    ID = 61
	ColdPlains               ID = 3
	Crypt                    ID = 18
	CrystallinePassage       ID = 113
	DarkWood                 ID = 5
	DenOfEvil                ID = 8
	DisusedFane              ID = 95
	DisusedReliquary         ID = 99
	DrifterCavern            ID = 116
	DryHills                 ID = 42
	DuranceOfHateLevel1      ID = 100
	DuranceOfHateLevel2      ID = 101
	DuranceOfHateLevel3      ID = 102
	DurielsLair              ID = 73
	FarOasis                 ID = 43
	FlayerDungeonLevel1      ID = 88
	FlayerDungeonLevel2      ID = 89
	FlayerDungeonLevel3      ID = 91
	FlayerJungle             ID = 78
	ForgottenReliquary       ID = 96
	ForgottenSands           ID = 134
	ForgottenTemple          ID = 97
	ForgottenTower           ID = 20
	FrigidHighlands          ID = 111
	FrozenRiver              ID = 114
	FrozenTundra             ID = 117
	FurnaceOfPain            ID = 135
	GlacialTrail             ID = 115
	GreatMarsh               ID = 77
	HallsOfAnguish           ID = 122
	HallsOfPain              ID = 123
	HallsOfTheDeadLevel1     ID = 56
	HallsOfTheDeadLevel2     ID = 57
	HallsOfTheDeadLevel3     ID = 60
	HallsOfVaught            ID = 124
	HaremLevel1              ID = 50
	HaremLevel2              ID = 51
	Harrogath                ID = 109
	HoleLevel1               ID = 11
	HoleLevel2               ID = 15
	IcyCellar                ID = 119
	InfernalPit              ID = 127
	InnerCloister            ID = 32
	JailLevel1               ID = 29
	JailLevel2               ID = 30
	JailLevel3               ID = 31
	KurastBazaar             ID = 80
	KurastCauseway           ID = 82
	KurastDocks              ID = 75
	LostCity                 ID = 44
	LowerKurast              ID = 79
	LutGholein               ID = 40
	MaggotLairLevel1         ID = 62
	MaggotLairLevel2         ID = 63
	MaggotLairLevel3         ID = 64
	MatronsDen               ID = 133
	Mausoleum                ID = 19
	MonasteryGate            ID = 26
	MooMooFarm               ID = 39
	NihlathaksTemple         ID = 121
	None                     ID = 0
	OuterCloister            ID = 27
	OuterSteppes             ID = 104
	PalaceCellarLevel1       ID = 52
	PalaceCellarLevel2       ID = 53
	PalaceCellarLevel3       ID = 54
	PitLevel1                ID = 12
	PitLevel2                ID = 16
	PitOfAcheron             ID = 126
	PlainsOfDespair          ID = 105
	RiverOfFlame             ID = 107
	RockyWaste               ID = 41
	RogueEncampment          ID = 1
	RuinedFane               ID = 98
	RuinedTemple             ID = 94
	SewersLevel1Act2         ID = 47
	SewersLevel1Act3         ID = 92
	SewersLevel2Act2         ID = 48
	SewersLevel2Act3         ID = 93
	SewersLevel3Act2         ID = 49
	SpiderCave               ID = 84
	SpiderCavern             ID = 85
	SpiderForest             ID = 76
	StonyField               ID = 4
	StonyTombLevel1          ID = 55
	StonyTombLevel2          ID = 59
	SwampyPitLevel1          ID = 86
	SwampyPitLevel2          ID = 87
	SwampyPitLevel3          ID = 90
	TalRashasTomb1           ID = 66
	TalRashasTomb2           ID = 67
	TalRashasTomb3           ID = 68
	TalRashasTomb4           ID = 69
	TalRashasTomb5           ID = 70
	TalRashasTomb6           ID = 71
	TalRashasTomb7           ID = 72
	TamoeHighland            ID = 7
	TheAncientsWay           ID = 118
	ThePandemoniumFortress   ID = 103
	TheWorldstoneChamber     ID = 132
	TheWorldStoneKeepLevel1  ID = 128
	TheWorldStoneKeepLevel2  ID = 129
	TheWorldStoneKeepLevel3  ID = 130
	ThroneOfDestruction      ID = 131
	TowerCellarLevel1        ID = 21
	TowerCellarLevel2        ID = 22
	TowerCellarLevel3        ID = 23
	TowerCellarLevel4        ID = 24
	TowerCellarLevel5        ID = 25
	Travincal                ID = 83
	Tristram                 ID = 38
	UberTristram             ID = 136
	UndergroundPassageLevel1 ID = 10
	UndergroundPassageLevel2 ID = 14
	UpperKurast              ID = 81
	ValleyOfSnakes           ID = 45
	MapsAncientTemple        ID = 137
	MapsDesecratedTemple     ID = 138
	MapsFrigidPlateau        ID = 139
	MapsInfernalTrial        ID = 140
	MapsRuinedCitadel        ID = 141
)
