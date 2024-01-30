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
)

type Quests map[Quest]Status
