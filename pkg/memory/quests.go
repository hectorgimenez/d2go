package memory

import "github.com/hectorgimenez/d2go/pkg/data/quest"

func (gd *GameReader) getQuests(questBytes []byte) quest.Quests {
	return quest.Quests{
		quest.Act1DenOfEvil:             gd.calculateQuestStatus(questBytes[3]),
		quest.Act1SistersBurialGrounds:  gd.calculateQuestStatus(questBytes[5]),
		quest.Act1ToolsOfTheTrade:       gd.calculateQuestStatus(questBytes[7]),
		quest.Act1TheSearchForCain:      gd.calculateQuestStatus(questBytes[9]),
		quest.Act1TheForgottenTower:     gd.calculateQuestStatus(questBytes[11]),
		quest.Act1SistersToTheSlaughter: gd.calculateQuestStatus(questBytes[13]),
	}
}

func (gd *GameReader) calculateQuestStatus(questByte byte) quest.Status {
	if questByte == 0x10 {
		return quest.StatusCompleted
	}

	return quest.StatusNotCompleted
}
