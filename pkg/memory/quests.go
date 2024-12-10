package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/quest"
)

func (gd *GameReader) getQuests(questBytes []byte) quest.Quests {
	return quest.Quests{
		// Act 1
		quest.Act1DenOfEvil:             gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 0, 1)),
		quest.Act1SistersBurialGrounds:  gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 1, 1)),
		quest.Act1ToolsOfTheTrade:       gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 2, 1)),
		quest.Act1TheSearchForCain:      gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 3, 1)),
		quest.Act1TheForgottenTower:     gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 4, 1)),
		quest.Act1SistersToTheSlaughter: gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 5, 1)),
		// Act 2
		quest.Act2RadamentsLair:    gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 8, 1)),
		quest.Act2TheHoradricStaff: gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 9, 1)),
		quest.Act2TaintedSun:       gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 10, 1)),
		quest.Act2ArcaneSanctuary:  gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 11, 1)),
		quest.Act2TheSummoner:      gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 12, 1)),
		quest.Act2TheSevenTombs:    gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 13, 1)),
		// Act 3
		quest.Act3LamEsensTome:          gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 16, 1)),
		quest.Act3KhalimsWill:           gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 17, 1)),
		quest.Act3BladeOfTheOldReligion: gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 18, 1)),
		quest.Act3TheGoldenBird:         gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 19, 1)),
		quest.Act3TheBlackenedTemple:    gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 20, 1)),
		quest.Act3TheGuardian:           gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 21, 1)),
		// Act 4
		quest.Act4TheFallenAngel: gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 24, 1)),
		quest.Act4TerrorsEnd:     gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 25, 1)),
		quest.Act4HellForge:      gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 26, 1)),
		// Act 5
		quest.Act5SiegeOnHarrogath:    gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 34, 1)),
		quest.Act5RescueOnMountArreat: gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 35, 1)),
		quest.Act5PrisonOfIce:         gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 36, 1)),
		quest.Act5BetrayalOfHarrogath: gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 37, 1)),
		quest.Act5RiteOfPassage:       gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 38, 1)),
		quest.Act5EveOfDestruction:    gd.getQuestStatus(ReadUIntFromBuffer(questBytes, 39, 1)),
	}
}
func (gd *GameReader) getQuestStatus(questStatusInt uint) []quest.Status {
	var activeStates []quest.Status
	activeStates = append(activeStates, quest.Status(questStatusInt))
	return activeStates
}

//outdated

/* func (gd *GameReader) getQuestStatus(questBytes []byte) []quest.Status {
 combinedValue := uint16(questBytes[0])<<8 | uint16(questBytes[1])

 binaryRepresentation := fmt.Sprintf("%016b", combinedValue)

 var activeStates []quest.Status
 for i, char := range binaryRepresentation {
if char == '1' {
if 15-i >= 0 && i < 16 {
 activeStates = append(activeStates, quest.Status(15-i))
 }
 }
 }

 return activeStates
}*/
