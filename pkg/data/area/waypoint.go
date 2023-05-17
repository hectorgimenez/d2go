package area

// WPAddresses represents the addresses of the waypoints in the game UI and the linked areas between them
var WPAddresses = map[Area]WPAddress{
	// Act 1
	RogueEncampment: {1, 1, nil},
	ColdPlains:      {1, 2, []Area{RogueEncampment, BloodMoor}},
	StonyField:      {1, 3, []Area{ColdPlains}},
	DarkWood:        {1, 4, []Area{StonyField, UndergroundPassageLevel1}},
	BlackMarsh:      {1, 5, []Area{DarkWood}},
	OuterCloister:   {1, 6, []Area{BlackMarsh, TamoeHighland, MonasteryGate}},
	JailLevel1:      {1, 7, []Area{OuterCloister, Barracks}},
	InnerCloister:   {1, 8, []Area{JailLevel1, JailLevel2, JailLevel3}},
	CatacombsLevel2: {1, 9, []Area{InnerCloister, Cathedral, CatacombsLevel1}},
	// Act 2
	LutGholein:           {2, 1, nil},
	SewersLevel2Act2:     {2, 2, []Area{LutGholein, SewersLevel1Act2}},
	DryHills:             {2, 3, []Area{LutGholein, RockyWaste}},
	HallsOfTheDeadLevel2: {2, 4, []Area{DryHills, HallsOfTheDeadLevel1}},
	FarOasis:             {2, 5, []Area{DryHills}},
	LostCity:             {2, 6, []Area{FarOasis}},
	PalaceCellarLevel1:   {2, 7, []Area{LutGholein, HaremLevel1, HaremLevel2}},
	ArcaneSanctuary:      {2, 8, []Area{PalaceCellarLevel1, PalaceCellarLevel2, PalaceCellarLevel3}},
	CanyonOfTheMagi:      {2, 9, []Area{ArcaneSanctuary}},
	// Act 3
	KurastDocks:         {3, 1, nil},
	SpiderForest:        {3, 2, []Area{KurastDocks}},
	GreatMarsh:          {3, 3, []Area{SpiderForest}},
	FlayerJungle:        {3, 4, []Area{GreatMarsh}},
	LowerKurast:         {3, 5, []Area{FlayerJungle}},
	KurastBazaar:        {3, 6, []Area{LowerKurast}},
	UpperKurast:         {3, 7, []Area{KurastBazaar}},
	Travincal:           {3, 8, []Area{UpperKurast, KurastCauseway}},
	DuranceOfHateLevel2: {3, 9, []Area{Travincal, DuranceOfHateLevel1}},
	// Act 4
	ThePandemoniumFortress: {4, 1, nil},
	CityOfTheDamned:        {4, 2, []Area{ThePandemoniumFortress, OuterSteppes, PlainsOfDespair}},
	RiverOfFlame:           {4, 3, []Area{CityOfTheDamned}},
	// Act 5
	Harrogath:               {5, 1, nil},
	FrigidHighlands:         {5, 2, []Area{Harrogath, BloodyFoothills}},
	ArreatPlateau:           {5, 3, []Area{FrigidHighlands}},
	CrystallinePassage:      {5, 4, []Area{ArreatPlateau}},
	GlacialTrail:            {5, 5, []Area{CrystallinePassage}},
	HallsOfPain:             {5, 6, []Area{Harrogath, NihlathaksTemple, HallsOfAnguish}},
	FrozenTundra:            {5, 7, []Area{GlacialTrail}},
	TheAncientsWay:          {5, 8, []Area{FrozenTundra}},
	TheWorldStoneKeepLevel2: {5, 9, []Area{TheAncientsWay, ArreatSummit, TheWorldStoneKeepLevel1}},
}

type WPAddress struct {
	Tab        int
	Row        int
	LinkedFrom []Area
}
