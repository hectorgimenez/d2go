package data

type KeyBindings struct {
	CharacterScreen KeyBinding
	Inventory       KeyBinding
	HoradricCube    KeyBinding
	PartyScreen     KeyBinding
	MercenaryScreen KeyBinding
	MessageLog      KeyBinding
	QuestLog        KeyBinding
	HelpScreen      KeyBinding

	SkillTree           KeyBinding
	SkillSpeedBar       KeyBinding
	Skills              [16]KeyBinding
	SelectPreviousSkill KeyBinding
	SelectNextSkill     KeyBinding

	ShowBelt    KeyBinding
	UseBelt     [4]KeyBinding
	SwapWeapons KeyBinding

	Chat          KeyBinding
	Run           KeyBinding
	ToggleRunWalk KeyBinding
	StandStill    KeyBinding
	ForceMove     KeyBinding
	ShowItems     KeyBinding
	ShowPortraits KeyBinding

	Automap        KeyBinding
	CenterAutomap  KeyBinding
	FadeAutomap    KeyBinding
	PartyOnAutomap KeyBinding
	NamesOnAutomap KeyBinding
	ToggleMiniMap  KeyBinding

	SayHelp         KeyBinding
	SayFollowMe     KeyBinding
	SayThisIsForYou KeyBinding
	SayThanks       KeyBinding
	SaySorry        KeyBinding
	SayBye          KeyBinding
	SayNowYouDie    KeyBinding
	SayRetreat      KeyBinding

	ClearScreen   KeyBinding
	ClearMessages KeyBinding
	Zoom          KeyBinding
	LegacyToggle  KeyBinding
}

type KeyBinding struct {
	Key1 [2]byte
	Key2 [2]byte
}
