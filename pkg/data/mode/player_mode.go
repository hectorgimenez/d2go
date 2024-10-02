package mode

type PlayerMode uint32

const (
	Death PlayerMode = iota
	StandingOutsideTown
	Walking
	Running
	GettingHit
	StandingInTown
	WalkingInTown
	Attacking1
	Attacking2
	Blocking
	CastingSkill
	ThrowingItem
	Kicking
	UsingSkill1
	UsingSkill2
	UsingSkill3
	UsingSkill4
	Dead
	SkillActionSequence
	KnockedBack
)
