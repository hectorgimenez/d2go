package mode

type NpcMode uint32

const (
	NpcDeath NpcMode = iota
	NpcStandingStill
	NpcWalking
	NpcGettingHit
	NpcAttacking1
	NpcAttacking2
	NpcBlocking
	NpcCastingSpell
	NpcUsingSkill1
	NpcUsingSkill2
	NpcUsingSkill3
	NpcUsingSkill4
	NpcDead
	NpcKnockedBack
	NpcActionSequence
	NpcRunning
)
