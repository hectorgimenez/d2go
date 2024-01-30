package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/difficulty"
	"github.com/hectorgimenez/d2go/pkg/data/quest"
)

func (gd *GameReader) getQuests(questBytes []byte) map[difficulty.Difficulty]quest.Quests {
	return map[difficulty.Difficulty]quest.Quests{
		difficulty.Normal: {
			// Act 1
			quest.Act1DenOfEvil:             gd.calculateQuestStatus(questBytes[3]),
			quest.Act1SistersBurialGrounds:  gd.calculateQuestStatus(questBytes[5]),
			quest.Act1ToolsOfTheTrade:       gd.calculateQuestStatus(questBytes[7]),
			quest.Act1TheSearchForCain:      gd.calculateQuestStatus(questBytes[9]),
			quest.Act1TheForgottenTower:     gd.calculateQuestStatus(questBytes[11]),
			quest.Act1SistersToTheSlaughter: gd.calculateQuestStatus(questBytes[13]),
			// Act 2
			quest.Act2RadamentsLair:    gd.calculateQuestStatus(questBytes[17]), // wrong
			quest.Act2TheHoradricStaff: gd.calculateQuestStatus(questBytes[19]),
			quest.Act2TaintedSun:       gd.calculateQuestStatus(questBytes[21]), // wrong
			quest.Act2ArcaneSanctuary:  gd.calculateQuestStatus(questBytes[23]),
			quest.Act2TheSummoner:      gd.calculateQuestStatus(questBytes[25]), // wrong
			quest.Act2TheSevenTombs:    gd.calculateQuestStatus(questBytes[27]),
		},
	}
}

func (gd *GameReader) calculateQuestStatus(questByte byte) quest.Status {
	if questByte == 0x10 {
		return quest.StatusCompleted
	}

	return quest.StatusNotCompleted
}
