package area

// WPAddresses represents the addresses of the waypoints in the game UI and the linked areas between them
var WPAddresses = map[ID]WPAddress{
	// Act 1
	RogueEncampment: {1, 1, nil},
	ColdPlains:      {1, 2, []ID{RogueEncampment, BloodMoor}},
	StonyField:      {1, 3, []ID{ColdPlains}},
	DarkWood:        {1, 4, []ID{StonyField, UndergroundPassageLevel1}},
	BlackMarsh:      {1, 5, []ID{DarkWood}},
	OuterCloister:   {1, 6, []ID{BlackMarsh, TamoeHighland, MonasteryGate}},
	JailLevel1:      {1, 7, []ID{OuterCloister, Barracks}},
	InnerCloister:   {1, 8, []ID{JailLevel1, JailLevel2, JailLevel3}},
	CatacombsLevel2: {1, 9, []ID{InnerCloister, Cathedral, CatacombsLevel1}},
	// Act 2
	LutGholein:           {2, 1, nil},
	SewersLevel2Act2:     {2, 2, []ID{LutGholein, SewersLevel1Act2}},
	DryHills:             {2, 3, []ID{LutGholein, RockyWaste}},
	HallsOfTheDeadLevel2: {2, 4, []ID{DryHills, HallsOfTheDeadLevel1}},
	FarOasis:             {2, 5, []ID{DryHills}},
	LostCity:             {2, 6, []ID{FarOasis}},
	PalaceCellarLevel1:   {2, 7, []ID{LutGholein, HaremLevel1, HaremLevel2}},
	ArcaneSanctuary:      {2, 8, []ID{PalaceCellarLevel1, PalaceCellarLevel2, PalaceCellarLevel3}},
	CanyonOfTheMagi:      {2, 9, []ID{ArcaneSanctuary}},
	// Act 3
	KurastDocks:         {3, 1, nil},
	SpiderForest:        {3, 2, []ID{KurastDocks}},
	GreatMarsh:          {3, 3, []ID{SpiderForest}},
	FlayerJungle:        {3, 4, []ID{GreatMarsh}},
	LowerKurast:         {3, 5, []ID{FlayerJungle}},
	KurastBazaar:        {3, 6, []ID{LowerKurast}},
	UpperKurast:         {3, 7, []ID{KurastBazaar}},
	Travincal:           {3, 8, []ID{UpperKurast, KurastCauseway}},
	DuranceOfHateLevel2: {3, 9, []ID{Travincal, DuranceOfHateLevel1}},
	// Act 4
	ThePandemoniumFortress: {4, 1, nil},
	CityOfTheDamned:        {4, 2, []ID{ThePandemoniumFortress, OuterSteppes, PlainsOfDespair}},
	RiverOfFlame:           {4, 3, []ID{CityOfTheDamned}},
	// Act 5
	Harrogath:               {5, 1, nil},
	FrigidHighlands:         {5, 2, []ID{Harrogath, BloodyFoothills}},
	ArreatPlateau:           {5, 3, []ID{FrigidHighlands}},
	CrystallinePassage:      {5, 4, []ID{ArreatPlateau}},
	GlacialTrail:            {5, 5, []ID{CrystallinePassage}},
	HallsOfPain:             {5, 6, []ID{Harrogath, NihlathaksTemple, HallsOfAnguish}},
	FrozenTundra:            {5, 7, []ID{GlacialTrail}},
	TheAncientsWay:          {5, 8, []ID{FrozenTundra}},
	TheWorldStoneKeepLevel2: {5, 9, []ID{TheAncientsWay, ArreatSummit, TheWorldStoneKeepLevel1}},
}

type WPAddress struct {
	Tab        int
	Row        int
	LinkedFrom []ID
}
