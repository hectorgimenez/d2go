package memory

import (
	"fmt"
	"github.com/hectorgimenez/d2go/pkg/data/quest"
)

func (gd *GameReader) getQuests(questBytes []byte) quest.Quests {
	return quest.Quests{
		// Act 1
		quest.Act1DenOfEvil:             gd.getQuestStatus([2]byte(questBytes[0:2])),
		quest.Act1SistersBurialGrounds:  gd.getQuestStatus([2]byte(questBytes[2:4])),
		quest.Act1ToolsOfTheTrade:       gd.getQuestStatus([2]byte(questBytes[4:6])),
		quest.Act1TheSearchForCain:      gd.getQuestStatus([2]byte(questBytes[6:8])),
		quest.Act1TheForgottenTower:     gd.getQuestStatus([2]byte(questBytes[8:10])),
		quest.Act1SistersToTheSlaughter: gd.getQuestStatus([2]byte(questBytes[10:12])),
		// Act 2
		quest.Act2RadamentsLair:    gd.getQuestStatus([2]byte(questBytes[16:18])),
		quest.Act2TheHoradricStaff: gd.getQuestStatus([2]byte(questBytes[18:20])),
		quest.Act2TaintedSun:       gd.getQuestStatus([2]byte(questBytes[20:22])),
		quest.Act2ArcaneSanctuary:  gd.getQuestStatus([2]byte(questBytes[22:24])),
		quest.Act2TheSummoner:      gd.getQuestStatus([2]byte(questBytes[24:26])),
		quest.Act2TheSevenTombs:    gd.getQuestStatus([2]byte(questBytes[26:28])),
		// Act 3
		quest.Act3LamEsensTome:          gd.getQuestStatus([2]byte(questBytes[32:34])),
		quest.Act3KhalimsWill:           gd.getQuestStatus([2]byte(questBytes[34:36])),
		quest.Act3BladeOfTheOldReligion: gd.getQuestStatus([2]byte(questBytes[36:38])),
		quest.Act3TheGoldenBird:         gd.getQuestStatus([2]byte(questBytes[38:40])),
		quest.Act3TheBlackenedTemple:    gd.getQuestStatus([2]byte(questBytes[40:42])),
		quest.Act3TheGuardian:           gd.getQuestStatus([2]byte(questBytes[42:44])),
		// Act 4
		quest.Act4TheFallenAngel: gd.getQuestStatus([2]byte(questBytes[48:50])),
		quest.Act4TerrorsEnd:     gd.getQuestStatus([2]byte(questBytes[50:52])),
		quest.Act4HellForge:      gd.getQuestStatus([2]byte(questBytes[52:54])),
		// Act 5
		quest.Act5SiegeOnHarrogath:    gd.getQuestStatus([2]byte(questBytes[68:70])),
		quest.Act5RescueOnMountArreat: gd.getQuestStatus([2]byte(questBytes[70:72])),
		quest.Act5PrisonOfIce:         gd.getQuestStatus([2]byte(questBytes[72:74])),
		quest.Act5BetrayalOfHarrogath: gd.getQuestStatus([2]byte(questBytes[74:76])),
		quest.Act5RiteOfPassage:       gd.getQuestStatus([2]byte(questBytes[76:78])),
		quest.Act5EveOfDestruction:    gd.getQuestStatus([2]byte(questBytes[78:80])),
	}
}

func (gd *GameReader) getQuestStatus(questBytes [2]byte) []quest.Status {
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
}
