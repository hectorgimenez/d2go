package quest

type Quest int16
type Status int16

const (
	StatusNotCompleted Status = 0
	StatusCompleted    Status = 1

	Act1DenOfEvil             Quest = 0
	Act1SistersBurialGrounds  Quest = 1
	Act1ToolsOfTheTrade       Quest = 2
	Act1TheSearchForCain      Quest = 3
	Act1TheForgottenTower     Quest = 4
	Act1SistersToTheSlaughter Quest = 5
	Act2RadamentsLair         Quest = 6
	Act2TheHoradricStaff      Quest = 7
	Act2TaintedSun            Quest = 8
	Act2ArcaneSanctuary       Quest = 9
	Act2TheSummoner           Quest = 10
	Act2TheSevenTombs         Quest = 11
	Act3LamEsensTome          Quest = 12
	Act3KhalimsWill           Quest = 13
	Act3BladeOfTheOldReligion Quest = 14
	Act3TheGoldenBird         Quest = 15
	Act3TheBlackenedTemple    Quest = 16
	Act3TheGuardian           Quest = 17
	Act4TheFallenAngel        Quest = 18
	Act4HellForge             Quest = 19
	Act4TerrorsEnd            Quest = 20
	Act5SiegeOnHarrogath      Quest = 21
	Act5RescueOnMountArreat   Quest = 22
	Act5PrisonOfIce           Quest = 23
	Act5BetrayalOfHarrogath   Quest = 24
	Act5RiteOfPassage         Quest = 25
	Act5EveOfDestruction      Quest = 26
)

type Quests map[Quest]Status
