package packet

type OpCodeInfo struct {
	Name   string
	Size   int
	Fields []FieldInfo
}

type FieldInfo struct {
	Size int
	Name string
}

var (
	OpCodesAck = make(map[uint8]OpCodeInfo)
	OpCodesReq = make(map[uint8]OpCodeInfo)
)

func init() {
	initAck()
	initReq()
}

func initAck() {

	OpCodesAck[0] = OpCodeInfo{
		Name:   "GameLoading",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[1] = OpCodeInfo{
		Name: "GameFlags",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "Difficulty"},
			{Size: 4, Name: "ArenaFlags"},
			{Size: 1, Name: "IsExpansion"},
			{Size: 1, Name: "IsLadder"},
		},
	}

	OpCodesAck[2] = OpCodeInfo{
		Name:   "LoadStarted",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[3] = OpCodeInfo{
		Name: "LoadAct",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 1, Name: "ActId"},
			{Size: 4, Name: "MapId"},
			{Size: 2, Name: "AreaId"},
			{Size: 4, Name: "AutoMap"},
		},
	}

	OpCodesAck[4] = OpCodeInfo{
		Name:   "LoadComplete",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[5] = OpCodeInfo{
		Name:   "UnloadComplete",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[6] = OpCodeInfo{
		Name: "GameHandshake",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[7] = OpCodeInfo{
		Name: "RoomAdd",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "LeveNo"},
		},
	}

	OpCodesAck[8] = OpCodeInfo{
		Name: "RoomRemove",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "LeveNo"},
		},
	}

	OpCodesAck[9] = OpCodeInfo{
		Name: "AssignLvlWarp",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 1, Name: "WarpType"},
			{Size: 4, Name: "WarpId"},
			{Size: 1, Name: "WarpClassId"},
			{Size: 2, Name: "WarpX"},
			{Size: 2, Name: "WarpY"},
		},
	}

	OpCodesAck[10] = OpCodeInfo{
		Name: "RemoveObject",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[11] = OpCodeInfo{
		Name: "GameHandshake2",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[13] = OpCodeInfo{
		Name: "PlayerStop",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "HitClass"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "UnitHitClass"},
			{Size: 1, Name: "UnitLife"},
		},
	}

	OpCodesAck[14] = OpCodeInfo{
		Name: "ObjectState",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "PortalFalgs"},
			{Size: 1, Name: "FlagIsTargetable"},
			{Size: 4, Name: "UnitState"},
		},
	}

	OpCodesAck[15] = OpCodeInfo{
		Name: "PlayerMove",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "MoveType"},
			{Size: 2, Name: "TargetX"},
			{Size: 2, Name: "TargetY"},
			{Size: 1, Name: "UnitHitClass"},
			{Size: 2, Name: "UnitX"},
			{Size: 2, Name: "UnitY"},
		},
	}
	OpCodesAck[16] = OpCodeInfo{
		Name: "PlayerStop",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}
	OpCodesAck[17] = OpCodeInfo{
		Name: "ReportKill",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "Overlay"},
		},
	}

	OpCodesAck[18] = OpCodeInfo{
		Name: "PlayerInGame",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "CharClass"},
			{Size: 1, Name: "CharName[16]"},
		},
	}

	OpCodesAck[19] = OpCodeInfo{
		Name: "PlayerLeaveGame",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[20] = OpCodeInfo{
		Name: "NpcHit",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "AnimationType"},
			{Size: 2, Name: "UnknownA"},
			{Size: 1, Name: "UnknownB"},
			{Size: 1, Name: "LifePercentage"},
			{Size: 2, Name: "TargetX"},
			{Size: 2, Name: "TargetY"},
		},
	}

	OpCodesAck[21] = OpCodeInfo{
		Name: "ReassignPlayer",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "Flag"},
		},
	}

	OpCodesAck[22] = OpCodeInfo{
		Name: "UnitsCoordsUpdate",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "Unused1"},
			{Size: 1, Name: "Unused2"},
			{Size: 1, Name: "Count"},
			{Size: -1, Name: "UnitInfo[Count]"},
		},
	}

	OpCodesAck[23] = OpCodeInfo{
		Name: "Unknown_23",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "UnitX"},
			{Size: 2, Name: "UnitY"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesAck[24] = OpCodeInfo{
		Name: "HPMPUpdate2",
		Size: 15,
		Fields: []FieldInfo{
			{Size: 14, Name: "BitStream[14]"},
		},
	}

	OpCodesAck[25] = OpCodeInfo{
		Name: "PickupGold",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "Value"},
		},
	}

	OpCodesAck[26] = OpCodeInfo{
		Name: "AddExpBYTE",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "Value"},
		},
	}

	OpCodesAck[27] = OpCodeInfo{
		Name: "AddExpWORD",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 2, Name: "Value"},
		},
	}

	OpCodesAck[28] = OpCodeInfo{
		Name: "SetExpDWORD",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "Value"},
		},
	}

	OpCodesAck[29] = OpCodeInfo{
		Name: "AttributeSetBYTE",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 1, Name: "Attrib"},
			{Size: 1, Name: "Value"},
		},
	}

	OpCodesAck[30] = OpCodeInfo{
		Name: "AttributeSetWORD",
		Size: 4,
		Fields: []FieldInfo{
			{Size: 1, Name: "Attrib"},
			{Size: 2, Name: "Value"},
		},
	}

	OpCodesAck[31] = OpCodeInfo{
		Name: "AttributeSetDWORD",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "Attrib"},
			{Size: 4, Name: "Value"},
		},
	}

	OpCodesAck[32] = OpCodeInfo{
		Name: "AttributeUpdate",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Attrib"},
			{Size: 4, Name: "Value"},
		},
	}

	OpCodesAck[33] = OpCodeInfo{
		Name: "ItemUpdate_OSkill",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 1, Name: "Delete"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "BaseLevel"},
			{Size: 1, Name: "BonusLevel"},
			{Size: 1, Name: "Padding"},
		},
	}

	OpCodesAck[34] = OpCodeInfo{
		Name: "ItemUpdate_Skill",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 1, Name: "Delete"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "Quantity"},
			{Size: 1, Name: "Padding"},
			{Size: 1, Name: "Body"},
		},
	}

	OpCodesAck[35] = OpCodeInfo{
		Name: "SetSkill",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Hand"},
			{Size: 2, Name: "SkillId"},
			{Size: 4, Name: "ItemUnitId"},
		},
	}

	OpCodesAck[36] = OpCodeInfo{
		Name: "SkillPointsUsed",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "Skill"},
			{Size: 1, Name: "Amount"},
		},
	}

	OpCodesAck[37] = OpCodeInfo{
		Name: "SkillPointAssignedToSkill",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 1, Name: "Skill"},
		},
	}

	OpCodesAck[38] = OpCodeInfo{
		Name: "Chat",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "ChatType"},
			{Size: 1, Name: "LangCode"},
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "ChatColor"},
			{Size: 1, Name: "ChatSubType"},
			{Size: -1, Name: "sNick"},
			{Size: -1, Name: "sMessage"},
		},
	}

	OpCodesAck[39] = OpCodeInfo{
		Name: "NPCInfo",
		Size: 40,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Count"},
			{Size: 1, Name: "Unknown"},
			{Size: 32, Name: "UnitMessages"},
		},
	}

	OpCodesAck[40] = OpCodeInfo{
		Name: "PlayerQuestInfo",
		Size: 103,
		Fields: []FieldInfo{
			{Size: 1, Name: "UpdateType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "ActionType"},
			{Size: 96, Name: "QuestBitStream[96]"},
		},
	}

	OpCodesAck[41] = OpCodeInfo{
		Name: "GameQuestLog",
		Size: 97,
		Fields: []FieldInfo{
			{Size: 96, Name: "QuestBitStream[96]"},
		},
	}

	OpCodesAck[42] = OpCodeInfo{
		Name: "NPCTransaction",
		Size: 15,
		Fields: []FieldInfo{
			{Size: 1, Name: "TradeType"},
			{Size: 1, Name: "Result"},
			{Size: 4, Name: "Unused"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "InventoryGold"},
		},
	}

	OpCodesAck[43] = OpCodeInfo{
		Name: "GameQuest",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "QuestId"},
			{Size: 4, Name: "Unknown1"},
			{Size: 4, Name: "Unknown2"},
		},
	}

	OpCodesAck[44] = OpCodeInfo{
		Name: "PlaySound",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "Type"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "SoundId"},
		},
	}
	OpCodesAck[45] = OpCodeInfo{
		Name:   "UnlockD2Character",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[46] = OpCodeInfo{
		Name:   "D2CharacterPhaseEnd",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[47] = OpCodeInfo{
		Name: "NPCSetMode",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Mode"},
			{Size: 2, Name: "StateInfo"},
		},
	}

	OpCodesAck[48] = OpCodeInfo{
		Name: "NPCSetModeEx",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Mode"},
			{Size: 2, Name: "StateInfo"},
			{Size: 4, Name: "UnitId2"},
		},
	}

	OpCodesAck[49] = OpCodeInfo{
		Name: "NPCSetUniqueState",
		Size: 14,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "StateId"},
			{Size: 4, Name: "StateData1"},
			{Size: 4, Name: "StateData2"},
		},
	}

	OpCodesAck[50] = OpCodeInfo{
		Name: "NPCHeal",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "LifePercentage"},
		},
	}

	OpCodesAck[51] = OpCodeInfo{
		Name: "NPCAssign",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "TxtFileNo"},
			{Size: 1, Name: "Unk"},
		},
	}

	OpCodesAck[52] = OpCodeInfo{
		Name: "PartyUpdate",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 2, Name: "PartyId"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PartyFlags"},
			{Size: 1, Name: "PlayerLevel"},
		},
	}

	OpCodesAck[53] = OpCodeInfo{
		Name: "AssignMerc",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 2, Name: "MercNpcId"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "OwnerId"},
			{Size: 2, Name: "MercNameId"},
			{Size: 2, Name: "Unknown1"},
			{Size: 1, Name: "Unknown2"},
		},
	}

	OpCodesAck[54] = OpCodeInfo{
		Name: "AssignWarp",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[55] = OpCodeInfo{
		Name: "PlayNpcMessage",
		Size: 35,
		Fields: []FieldInfo{
			{Size: 2, Name: "Unknown1"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "MessageId"},
			{Size: 1, Name: "Unknown2"},
			{Size: 16, Name: "Name[16]"},
			{Size: 4, Name: "Unknown3"},
			{Size: 2, Name: "Unknown4"},
			{Size: 2, Name: "Unknown5"},
			{Size: 1, Name: "Unknown6"},
		},
	}

	OpCodesAck[56] = OpCodeInfo{
		Name: "QuestSpecial",
		Size: 15,
		Fields: []FieldInfo{
			{Size: 1, Name: "Type"},
			{Size: 4, Name: "Id"},
			{Size: 4, Name: "X"},
			{Size: 4, Name: "Y"},
			{Size: 1, Name: "State"},
		},
	}

	OpCodesAck[57] = OpCodeInfo{
		Name: "AssignObject",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "ObjectType"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "State"},
			{Size: 1, Name: "InteractionType"},
			{Size: 2, Name: "Unknown1"},
			{Size: 1, Name: "Unknown2"},
		},
	}

	OpCodesAck[58] = OpCodeInfo{
		Name: "AssignSkill",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "UnitType"},
			{Size: 2, Name: "UnitId"},
		},
	}

	OpCodesAck[59] = OpCodeInfo{
		Name: "UpdateSkill",
		Size: 14,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "SkillId"},
			{Size: 4, Name: "Amount"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesAck[60] = OpCodeInfo{
		Name: "SetSkill",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "Unknown1"},
			{Size: 4, Name: "Amount"},
		},
	}

	OpCodesAck[61] = OpCodeInfo{
		Name: "GameChat",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "PlayerUUID"},
			{Size: 1, Name: "ChatKind"},
			{Size: 1, Name: "Unknown1"},
			{Size: 4, Name: "Unknown2"},
			{Size: 1, Name: "Unknown3"},
			{Size: -1, Name: "Message"},
		},
	}
	OpCodesAck[62] = OpCodeInfo{
		Name: "UpdateItemStats",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "FullPacketSize"},
			{Size: -1, Name: "StatBitStream[nFullPacketSize - 2]"},
		},
	}

	OpCodesAck[63] = OpCodeInfo{
		Name: "UseStackableItem",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "SpellIcon"},
			{Size: 4, Name: "ItemId"},
			{Size: 2, Name: "SkillId"},
		},
	}

	OpCodesAck[64] = OpCodeInfo{
		Name: "ItemFlagSetter",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "ItemFlag"},
			{Size: 4, Name: "Remove"},
		},
	}

	OpCodesAck[66] = OpCodeInfo{
		Name: "ClearCursor",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[67] = OpCodeInfo{
		Name: "UpdateItemUI",
		Size: 15,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "ItemId"},
			{Size: 2, Name: "UiFlags"},
			{Size: 4, Name: "ActionType"},
		},
	}

	OpCodesAck[68] = OpCodeInfo{
		Name: "DeleteItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[69] = OpCodeInfo{
		Name: "ItemNameUpdate",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: -1, Name: "ItemName"},
		},
	}

	OpCodesAck[70] = OpCodeInfo{
		Name: "ItemAction",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 1, Name: "ActionId"},
			{Size: 4, Name: "Category"},
			{Size: 4, Name: "Id"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[71] = OpCodeInfo{
		Name: "Relator1",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 1, Name: "Gap"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "Padding[4]"},
		},
	}

	OpCodesAck[72] = OpCodeInfo{
		Name: "Relator2",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 1, Name: "Gap"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "Padding[4]"},
		},
	}

	OpCodesAck[73] = OpCodeInfo{
		Name:   "PlayerClearCursor",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[74] = OpCodeInfo{
		Name: "Relator1",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "StateType"},
		},
	}

	OpCodesAck[75] = OpCodeInfo{
		Name: "Relator2",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "StateType"},
		},
	}

	OpCodesAck[76] = OpCodeInfo{
		Name: "UnitCastSkillTarget",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 1, Name: "AttackerType"},
			{Size: 4, Name: "AttackerId"},
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "skilllevel"},
			{Size: 1, Name: "TargetType"},
			{Size: 4, Name: "TargetId"},
			{Size: 2, Name: "Zero"},
		},
	}

	OpCodesAck[77] = OpCodeInfo{
		Name: "UnitCastSkillXY",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 1, Name: "AttackerType"},
			{Size: 4, Name: "AttackerId"},
			{Size: 2, Name: "SkillId"},
			{Size: 2, Name: "Filler"},
			{Size: 1, Name: "skilllevel"},
			{Size: 2, Name: "TargetX"},
			{Size: 2, Name: "TargetY"},
			{Size: 2, Name: "Zero"},
		},
	}

	OpCodesAck[78] = OpCodeInfo{
		Name: "MercForHire",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 2, Name: "NameStringId"},
			{Size: 4, Name: "Seed"},
		},
	}

	OpCodesAck[79] = OpCodeInfo{
		Name:   "ClearMercList",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[80] = OpCodeInfo{
		Name: "QuestSpecial",
		Size: 15,
		Fields: []FieldInfo{
			{Size: 2, Name: "MessageType"},
			{Size: 2, Name: "Arg1"},
			{Size: 2, Name: "Arg2"},
			{Size: 2, Name: "Arg3"},
			{Size: 2, Name: "Arg4"},
			{Size: 2, Name: "Arg5"},
			{Size: 2, Name: "Arg6"},
		},
	}

	OpCodesAck[81] = OpCodeInfo{
		Name: "AddObject",
		Size: 14,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "UnitClass"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "State"},
			{Size: 1, Name: "Interaction"},
		},
	}

	OpCodesAck[82] = OpCodeInfo{
		Name: "PlayerQuestLog",
		Size: 42,
		Fields: []FieldInfo{
			{Size: 41, Name: "QuestBitStream[41]"},
		},
	}

	OpCodesAck[83] = OpCodeInfo{
		Name: "Darkness",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 4, Name: "Act"},
			{Size: 4, Name: "Angle"},
			{Size: 1, Name: "OnOff"},
		},
	}

	OpCodesAck[84] = OpCodeInfo{
		Name: "AssignSkill",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesAck[85] = OpCodeInfo{
		Name: "AssignSkillHotkey",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "HotkeyId"},
			{Size: 1, Name: "Unknown1"},
			{Size: 2, Name: "Unknown2"},
		},
	}

	OpCodesAck[86] = OpCodeInfo{
		Name: "UseSkillOnTarget",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesAck[87] = OpCodeInfo{
		Name: "NPCEnchant",
		Size: 14,
		Fields: []FieldInfo{
			{Size: 4, Name: "MonsterId"},
			{Size: 1, Name: "MonsterType"},
			{Size: 2, Name: "MonsterNameIDX"},
			{Size: 3, Name: "Enchant[3]"},
			{Size: 1, Name: "Filler"},
			{Size: 2, Name: "MonsterIsChampion"},
		},
	}

	OpCodesAck[88] = OpCodeInfo{
		Name: "OpenUI",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "UIType"},
			{Size: 1, Name: "Bool"},
		},
	}

	OpCodesAck[89] = OpCodeInfo{
		Name: "AssignPlayer",
		Size: 26,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "UnitType"},
			{Size: 16, Name: "Name"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
		},
	}

	OpCodesAck[90] = OpCodeInfo{
		Name: "EventMessages",
		Size: 40,
		Fields: []FieldInfo{
			{Size: 1, Name: "MessageType"},
			{Size: 1, Name: "Color"},
			{Size: 4, Name: "Arg"},
			{Size: 1, Name: "ArgTypes"},
			{Size: 16, Name: "Name1[16]"},
			{Size: 16, Name: "Name2[16]"},
		},
	}

	OpCodesAck[91] = OpCodeInfo{
		Name: "PlayerJoin",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "PacketLength"},
			{Size: 4, Name: "PlayerId"},
			{Size: 1, Name: "CharType"},
			{Size: 16, Name: "PlayerName"},
			{Size: 2, Name: "PlayerLevel"},
			{Size: 2, Name: "PartyId"},
			{Size: 2, Name: "Unused"},
			{Size: 2, Name: "Unknown"},
			{Size: 2, Name: "Unknown"},
			{Size: -1, Name: "Unknown[String]"},
			{Size: -1, Name: "Unknown[String]"},
		},
	}

	OpCodesAck[92] = OpCodeInfo{
		Name: "PlayerLeave",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "PlayerId"},
		},
	}

	OpCodesAck[93] = OpCodeInfo{
		Name: "QuestState",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "QuestId"},
			{Size: 1, Name: "AlertFlags"},
			{Size: 1, Name: "FilterStatus"},
			{Size: 2, Name: "Extra"},
		},
	}

	OpCodesAck[94] = OpCodeInfo{
		Name: "QuestsAvailability",
		Size: 38,
		Fields: []FieldInfo{
			{Size: 37, Name: "QuestBitStream[37]"},
		},
	}

	OpCodesAck[95] = OpCodeInfo{
		Name: "PortalFalgs",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "PortalFlags"},
		},
	}

	OpCodesAck[96] = OpCodeInfo{
		Name: "TownPortalState",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "State"},
			{Size: 1, Name: "AreaId"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[97] = OpCodeInfo{
		Name: "CanGoToAct",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "Act"},
		},
	}

	OpCodesAck[98] = OpCodeInfo{
		Name: "MakeUnitTargetable",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 1, Name: "UnitId"},
			{Size: 1, Name: "Unused"},
		},
	}

	OpCodesAck[99] = OpCodeInfo{
		Name: "WaypointMenu",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 2, Name: "Unknown"},
			{Size: 8, Name: "WayPointBitStream[8]"},
			{Size: 4, Name: "Padding"},
		},
	}

	OpCodesAck[100] = OpCodeInfo{
		Name: "WayPointCurrent",
		Size: 4,
		Fields: []FieldInfo{
			{Size: 2, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[101] = OpCodeInfo{
		Name: "PlayerKillCount",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 4, Name: "PlayerId"},
			{Size: 2, Name: "Count"},
		},
	}

	OpCodesAck[102] = OpCodeInfo{
		Name: "SetNpcUniqueId",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "UniqueId"},
		},
	}

	OpCodesAck[103] = OpCodeInfo{
		Name: "UnitMove",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "MoveType"},
			{Size: 2, Name: "DestX"},
			{Size: 2, Name: "DestY"},
			{Size: 2, Name: "OrigX"},
			{Size: 2, Name: "OrigY"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Count"},
			{Size: 1, Name: "MoveType?"},
			{Size: -1, Name: "PosData"},
		},
	}

	OpCodesAck[104] = OpCodeInfo{
		Name: "UnitMoveToUnit",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "MoveType"},
			{Size: 2, Name: "TargetX"},
			{Size: 2, Name: "TargetY"},
			{Size: 1, Name: "TargetType"},
			{Size: 4, Name: "TargetId"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 2, Name: "NotUsed"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[105] = OpCodeInfo{
		Name: "UnitState",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "State"},
			{Size: 2, Name: "UnitX"},
			{Size: 2, Name: "UnitY"},
			{Size: 1, Name: "UnitLife"},
			{Size: 1, Name: "HitClass"},
		},
	}

	OpCodesAck[106] = OpCodeInfo{
		Name: "NpcInteract",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "InteractType"},
		},
	}

	OpCodesAck[107] = OpCodeInfo{
		Name: "UnitAction2",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Action"},
			{Size: 1, Name: "Padding[6]"},
			{Size: 2, Name: "UnitX"},
			{Size: 2, Name: "UnitY"},
		},
	}

	OpCodesAck[108] = OpCodeInfo{
		Name: "UnitAttack",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "AttackType"},
			{Size: 4, Name: "TargetId"},
			{Size: 1, Name: "TargetType"},
			{Size: 2, Name: "TargetX"},
			{Size: 2, Name: "TargetY"},
		},
	}

	OpCodesAck[109] = OpCodeInfo{
		Name: "UnitStop",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "HpPercent"},
		},
	}

	OpCodesAck[110] = OpCodeInfo{
		Name: "NpcSetState",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "State"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
		},
	}

	OpCodesAck[111] = OpCodeInfo{
		Name: "NpcAction",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Action"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesAck[112] = OpCodeInfo{
		Name: "TerrorZoneNotify",
		Size: 123,
		Fields: []FieldInfo{
			{Size: 2, Name: "ZoneCount"},
			{Size: -1, Name: "ZoneData"},
		},
	}

	OpCodesAck[113] = OpCodeInfo{
		Name: "NpcMove",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[114] = OpCodeInfo{
		Name: "NpcMoveToTarget",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "TargetType"},
			{Size: 4, Name: "TargetId"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[115] = OpCodeInfo{
		Name: "MissileData",
		Size: 32,
		Fields: []FieldInfo{
			{Size: 4, Name: "Unused"},
			{Size: 2, Name: "MissileClassId"},
			{Size: 4, Name: "MissileX"},
			{Size: 4, Name: "MissileY"},
			{Size: 4, Name: "TargetX"},
			{Size: 4, Name: "TargetY"},
			{Size: 2, Name: "CurrentFrame"},
			{Size: 1, Name: "OwnerType"},
			{Size: 4, Name: "OwnerId"},
			{Size: 1, Name: "SkillLevel"},
			{Size: 1, Name: "PierceIdxValue"},
		},
	}

	OpCodesAck[116] = OpCodeInfo{
		Name: "PlayerCorpseAssign",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "Assign"},
			{Size: 4, Name: "OwnerId"},
			{Size: 4, Name: "CorpseId"},
		},
	}

	OpCodesAck[117] = OpCodeInfo{
		Name: "PlayerPartyInfo",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PartyId"},
			{Size: 2, Name: "CharLevel"},
			{Size: 2, Name: "Relationship"},
			{Size: 2, Name: "InParty"},
		},
	}
	OpCodesAck[118] = OpCodeInfo{
		Name: "PlayerInProximity",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[119] = OpCodeInfo{
		Name: "TradeAction",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "RequestType"},
		},
	}

	OpCodesAck[120] = OpCodeInfo{
		Name: "TradeAccepted",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 16, Name: "CharName[16]"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[121] = OpCodeInfo{
		Name: "GoldInTrade",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "OwnerId"},
			{Size: 4, Name: "Value"},
		},
	}

	OpCodesAck[122] = OpCodeInfo{
		Name: "SummonLog",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 1, Name: "Action"},
			{Size: 1, Name: "SkillId"},
			{Size: 2, Name: "SummonMonsterTxtNo"},
			{Size: 4, Name: "PlayerUnitId"},
			{Size: 4, Name: "SummonUnitId"},
		},
	}

	OpCodesAck[123] = OpCodeInfo{
		Name: "AssignHotkey",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "Slot"},
			{Size: 1, Name: "SkillId"},
			{Size: 1, Name: "Hand"},
			{Size: 4, Name: "ItemUnitId"},
		},
	}

	OpCodesAck[124] = OpCodeInfo{
		Name: "UseScroll",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "ScrollType"},
			{Size: 4, Name: "ScrollId"},
		},
	}

	OpCodesAck[125] = OpCodeInfo{
		Name: "SetItemFlags",
		Size: 18,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "AndValue"},
			{Size: 4, Name: "FlagsAfterAnd"},
		},
	}

	OpCodesAck[126] = OpCodeInfo{
		Name: "Cmncof",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "Unused[4]"},
		},
	}

	OpCodesAck[127] = OpCodeInfo{
		Name: "AllyPartyInfo",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 2, Name: "LifePercent"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "LevelNo"},
		},
	}

	OpCodesAck[128] = OpCodeInfo{
		Name: "NpcWantsInteract",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[129] = OpCodeInfo{
		Name: "AssignMerc",
		Size: 20,
		Fields: []FieldInfo{
			{Size: 1, Name: "SkillId"},
			{Size: 2, Name: "SummonType"},
			{Size: 4, Name: "OwnerId"},
			{Size: 4, Name: "MercId"},
			{Size: 4, Name: "Seed2"},
			{Size: 4, Name: "InitSeed"},
		},
	}

	OpCodesAck[130] = OpCodeInfo{
		Name: "PortalOwnership",
		Size: 29,
		Fields: []FieldInfo{
			{Size: 4, Name: "OwnerId"},
			{Size: 16, Name: "OwnerName[16]"},
			{Size: 4, Name: "PortalId1"},
			{Size: 4, Name: "PortalId2"},
		},
	}
	OpCodesAck[131] = OpCodeInfo{
		Name: "OverheadMessage",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "Unknown1"},
			{Size: 2, Name: "MessageType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Unknown2"},
			{Size: -1, Name: "Message"},
		},
	}

	OpCodesAck[132] = OpCodeInfo{
		Name: "NpcTransaction",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 1, Name: "TradeType"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "GoldInInventory"},
		},
	}

	OpCodesAck[133] = OpCodeInfo{
		Name: "PlaySound",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 2, Name: "SoundType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[134] = OpCodeInfo{
		Name: "UpdateItemStats",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "Unknown1"},
			{Size: 2, Name: "PacketLength"},
			{Size: -1, Name: "ItemStats"},
		},
	}

	OpCodesAck[135] = OpCodeInfo{
		Name: "UseStackableItem",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesAck[136] = OpCodeInfo{
		Name: "UseChargedItem",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 2, Name: "SkillId"},
			{Size: 4, Name: "TargetId"},
			{Size: 1, Name: "Unknown"},
		},
	}
	OpCodesAck[137] = OpCodeInfo{
		Name: "UniqueEvents",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "EventType"},
		},
	}

	OpCodesAck[138] = OpCodeInfo{
		Name: "NPCWantsInteract",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[139] = OpCodeInfo{
		Name: "PlayerRelationship",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "InParty"},
		},
	}

	OpCodesAck[140] = OpCodeInfo{
		Name: "RelationshipUpdate",
		Size: 11,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitIdA"},
			{Size: 4, Name: "UnitIdB"},
			{Size: 2, Name: "Status"},
		},
	}

	OpCodesAck[141] = OpCodeInfo{
		Name: "AssignPlayerToParty",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 4, Name: "PlayerId"},
			{Size: 2, Name: "PartyId"},
		},
	}

	OpCodesAck[142] = OpCodeInfo{
		Name: "CorpseAssign",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "Assign"},
			{Size: 4, Name: "OwnerId"},
			{Size: 4, Name: "CorpseId"},
		},
	}

	OpCodesAck[143] = OpCodeInfo{
		Name: "Pong",
		Size: 33,
		Fields: []FieldInfo{
			{Size: 4, Name: "Pong"},
			{Size: 4, Name: "Pong"},
			{Size: 4, Name: "Pong"},
			{Size: 4, Name: "TickCount"},
			{Size: 4, Name: "Pong"},
			{Size: 4, Name: "PongWardenRequest"},
			{Size: 4, Name: "PongWarden"},
			{Size: 4, Name: "PongWarden"},
		},
	}

	OpCodesAck[144] = OpCodeInfo{
		Name: "PartyAutoMapInfo",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "PlayerId"},
			{Size: 4, Name: "PlayerX"},
			{Size: 4, Name: "PlayerY"},
		},
	}

	OpCodesAck[145] = OpCodeInfo{
		Name: "NPCGossip",
		Size: 26,
		Fields: []FieldInfo{
			{Size: 1, Name: "Act"},
			{Size: 24, Name: "NpcId[12]"},
		},
	}

	OpCodesAck[146] = OpCodeInfo{
		Name: "RemoveItemDisplay",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesAck[147] = OpCodeInfo{
		Name: "UnKnownUnitSkill0x93",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 4, Name: "PlayerId"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Type"},
			{Size: 1, Name: "SkillPage"},
		},
	}

	OpCodesAck[148] = OpCodeInfo{
		Name: "SkillList",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "Count"},
			{Size: 4, Name: "UnitId"},
			{Size: -1, Name: "List"},
		},
	}

	OpCodesAck[149] = OpCodeInfo{
		Name: "HPMPUpdate",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 2, Name: "HP"},
			{Size: 2, Name: "MP"},
			{Size: 2, Name: "Stamina"},
			{Size: 2, Name: "X"},
			{Size: 2, Name: "Y"},
			{Size: 1, Name: "DX"},
			{Size: 1, Name: "DY"},
			{Size: 4, Name: "Unknown"},
		},
	}

	OpCodesAck[150] = OpCodeInfo{
		Name: "WalkVerify",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 8, Name: "DataBitStream"},
		},
	}

	OpCodesAck[151] = OpCodeInfo{
		Name:   "WeaponSwitchVerify",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[152] = OpCodeInfo{
		Name: "Evilhut",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitGUID"},
			{Size: 2, Name: "Value"},
		},
	}

	OpCodesAck[153] = OpCodeInfo{
		Name: "UnitSkillCastTarget",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 1, Name: "AttackerType"},
			{Size: 4, Name: "AttackerId"},
			{Size: 2, Name: "SkillId"},
			{Size: 1, Name: "SkillLevel"},
			{Size: 1, Name: "TargetType"},
			{Size: 4, Name: "TargetId"},
			{Size: 2, Name: "Zero"},
		},
	}

	OpCodesAck[154] = OpCodeInfo{
		Name: "UnitSkillCastXY",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 1, Name: "AttackerType"},
			{Size: 4, Name: "AttackerId"},
			{Size: 2, Name: "SkillId"},
			{Size: 2, Name: "Padding"},
			{Size: 1, Name: "SkillLevel"},
			{Size: 2, Name: "TargetX"},
			{Size: 2, Name: "TargetY"},
			{Size: 2, Name: "Zero"},
		},
	}

	OpCodesAck[155] = OpCodeInfo{
		Name: "MercReviveCosts",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 2, Name: "MercNameId"},
			{Size: 4, Name: "ReviveCosts"},
		},
	}

	OpCodesAck[156] = OpCodeInfo{
		Name: "AddWorldItem",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "Action"},
			{Size: 1, Name: "PacketLength"},
			{Size: 1, Name: "Category"},
			{Size: 4, Name: "UnitId"},
			{Size: -1, Name: "DataBitStream"},
		},
	}

	OpCodesAck[157] = OpCodeInfo{
		Name: "AssignItem",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "Action"},
			{Size: 1, Name: "PacketLength"},
			{Size: 1, Name: "Category"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "OwnerUnitType"},
			{Size: 4, Name: "OwnerUnitId"},
			{Size: -1, Name: "DataBitStream"},
		},
	}

	OpCodesAck[158] = OpCodeInfo{
		Name: "MercAttributeSetBYTE",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatId"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "Value"},
		},
	}

	OpCodesAck[159] = OpCodeInfo{
		Name: "MercAttributeSetWORD",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatId"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "Value"},
		},
	}

	OpCodesAck[160] = OpCodeInfo{
		Name: "MercAttributeSetDWORD",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatId"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "Value"},
		},
	}

	OpCodesAck[161] = OpCodeInfo{
		Name: "MercAddExpBYTE",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatId"},
			{Size: 4, Name: "MercId"},
			{Size: 1, Name: "Value"},
		},
	}

	OpCodesAck[162] = OpCodeInfo{
		Name: "MercAddExpWORD",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatID"},
			{Size: 4, Name: "MercId"},
			{Size: 2, Name: "Value"},
		},
	}

	OpCodesAck[163] = OpCodeInfo{
		Name: "SkillAuraStat",
		Size: 24,
		Fields: []FieldInfo{
			{Size: 1, Name: "AuraStatId"},
			{Size: 2, Name: "SkillId"},
			{Size: 2, Name: "SkillLevel"},
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "TargetType"},
			{Size: 4, Name: "TargetId"},
			{Size: 4, Name: "TargetX"},
			{Size: 4, Name: "TargetY"},
		},
	}

	OpCodesAck[164] = OpCodeInfo{
		Name: "BaalWave",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 2, Name: "MonStatId"},
		},
	}

	OpCodesAck[165] = OpCodeInfo{
		Name: "StateSkillMove",
		Size: 8,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "SkillId"},
		},
	}

	OpCodesAck[166] = OpCodeInfo{
		Name: "RunesTxt",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "MustBeZero"},
			{Size: 2, Name: "FullPacketSize"},
			{Size: 2, Name: "TxtRunesSize"},
			{Size: -1, Name: "BitStream[nFullPacketSize - 6]"},
		},
	}

	OpCodesAck[167] = OpCodeInfo{
		Name: "DelayState",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "StateId"},
		},
	}

	OpCodesAck[168] = OpCodeInfo{
		Name: "SetState",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "PacketLength"},
			{Size: 1, Name: "StateId"},
			{Size: -1, Name: "DataBitStream"},
		},
	}

	OpCodesAck[169] = OpCodeInfo{
		Name: "EndState",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "StateId"},
		},
	}

	OpCodesAck[170] = OpCodeInfo{
		Name: "SetMultiStates",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "PacketLength"},
			{Size: -1, Name: "DataBitStream"},
		},
	}

	OpCodesAck[171] = OpCodeInfo{
		Name: "NPCHeal",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 1, Name: "UnitLife"},
		},
	}

	OpCodesAck[172] = OpCodeInfo{
		Name: "AddMonster",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "UnitClass"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 1, Name: "HpPercent"},
			{Size: 1, Name: "PacketLength"},
			{Size: -1, Name: "DataBitStream"},
		},
	}

	OpCodesAck[173] = OpCodeInfo{
		Name: "AddMonster2",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "UnitClass"},
			{Size: 2, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "Animation"},
			{Size: 1, Name: "HpPercent"},
		},
	}

	OpCodesAck[174] = OpCodeInfo{
		Name: "Warden",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "StreamSize"},
			{Size: -1, Name: "DataBitStream"},
		},
	}

	OpCodesAck[175] = OpCodeInfo{
		Name: "StartLogOn",
		Size: -1,
		Fields: []FieldInfo{
			{Size: -1, Name: "UseCompression"},
		},
	}

	OpCodesAck[176] = OpCodeInfo{
		Name:   "ConnectionTerminated",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[177] = OpCodeInfo{
		Name:   "StartGame",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesAck[178] = OpCodeInfo{
		Name: "GamesInfo",
		Size: 53,
		Fields: []FieldInfo{
			{Size: 16, Name: "Unknown"},
			{Size: 16, Name: "Unknown"},
			{Size: 16, Name: "Unknown"},
			{Size: 2, Name: "ClientsCount"},
			{Size: 2, Name: "GameToken"},
		},
	}

	OpCodesAck[179] = OpCodeInfo{
		Name: "DownloadSave",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "ChunkSize"},
			{Size: 1, Name: "First"},
			{Size: 4, Name: "FullSize"},
			{Size: -1, Name: "Stream[ChunkSize]"},
		},
	}

	OpCodesAck[180] = OpCodeInfo{
		Name: "ConnectionRefused",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "Reason"},
		},
	}

}

func initReq() {
	OpCodesReq[1] = OpCodeInfo{
		Name: "WalkToLocation",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 2, Name: "DestX"},
			{Size: 2, Name: "DestY"},
			{Size: 2, Name: "OriginX"},
			{Size: 2, Name: "OriginY"},
		},
	}

	OpCodesReq[2] = OpCodeInfo{
		Name: "WalkToUnit",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "OriginX"},
			{Size: 2, Name: "OriginY"},
		},
	}

	OpCodesReq[3] = OpCodeInfo{
		Name: "RunToLocation",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 2, Name: "DestX"},
			{Size: 2, Name: "DestY"},
			{Size: 2, Name: "OriginX"},
			{Size: 2, Name: "OriginY"},
		},
	}

	OpCodesReq[4] = OpCodeInfo{
		Name: "RunToUnit",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "OriginX"},
			{Size: 2, Name: "OriginY"},
		},
	}

	OpCodesReq[5] = OpCodeInfo{
		Name: "LeftSkillAtLocation",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "LocX"},
			{Size: 2, Name: "LocY"},
		},
	}

	OpCodesReq[6] = OpCodeInfo{
		Name: "LeftSkillAtUnit",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[7] = OpCodeInfo{
		Name: "LeftSkillAtUnitStandStill",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[8] = OpCodeInfo{
		Name: "InteractItemWithUnit",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesReq[9] = OpCodeInfo{
		Name: "PickupBufferItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[10] = OpCodeInfo{
		Name: "PickupUnitItem",
		Size: 10,
		Fields: []FieldInfo{
			{Size: 1, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[11] = OpCodeInfo{
		Name: "InitiateEntityChat",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesReq[12] = OpCodeInfo{
		Name: "RightSkillAtLocation",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "LocX"},
			{Size: 2, Name: "LocY"},
		},
	}

	OpCodesReq[13] = OpCodeInfo{
		Name: "RightSkillAtUnit",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[14] = OpCodeInfo{
		Name: "RightSkillAtUnitStandStill",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[15] = OpCodeInfo{
		Name: "OverheadMessage",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "Unknown"},
			{Size: -1, Name: "Message"},
		},
	}

	OpCodesReq[16] = OpCodeInfo{
		Name: "Chat",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "ChatType"},
			{Size: -1, Name: "Message"},
		},
	}

	OpCodesReq[17] = OpCodeInfo{
		Name: "PickChanMsg",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "ChatType"},
			{Size: -1, Name: "Message"},
		},
	}

	OpCodesReq[18] = OpCodeInfo{
		Name: "ChanWhisper",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "ChatType"},
			{Size: -1, Name: "Message"},
		},
	}

	OpCodesReq[19] = OpCodeInfo{
		Name: "InteractWithUnit",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "TargetUnitId"},
			{Size: 4, Name: "ExecuteeUnitId"},
		},
	}

	OpCodesReq[20] = OpCodeInfo{
		Name:   "JoinWarpReply",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[21] = OpCodeInfo{
		Name:   "LeaveChanReply",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[22] = OpCodeInfo{
		Name: "PickupItem",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
			{Size: 4, Name: "PosX"},
			{Size: 4, Name: "PosY"},
			{Size: 4, Name: "Cursor"},
		},
	}

	OpCodesReq[23] = OpCodeInfo{
		Name: "DropItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
		},
	}

	OpCodesReq[24] = OpCodeInfo{
		Name: "PutItemToInventory",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
			{Size: 4, Name: "PosX"},
			{Size: 4, Name: "PosY"},
			{Size: 4, Name: "InventoryId"},
		},
	}

	OpCodesReq[25] = OpCodeInfo{
		Name: "PullItemFromInventory",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
			{Size: 4, Name: "PosX"},
			{Size: 4, Name: "PosY"},
			{Size: 4, Name: "InventoryId"},
		},
	}

	OpCodesReq[26] = OpCodeInfo{
		Name: "PutItemToBody",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
			{Size: 4, Name: "BodyLocation"},
		},
	}

	OpCodesReq[27] = OpCodeInfo{
		Name: "ChangeItemSlot",
		Size: 12,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "Position"},
			{Size: 4, Name: "InventoryId"},
		},
	}

	OpCodesReq[28] = OpCodeInfo{
		Name: "PullItemFromBody",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
			{Size: 4, Name: "BodyLocation"},
		},
	}

	OpCodesReq[29] = OpCodeInfo{
		Name: "SwitchBeltItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[30] = OpCodeInfo{
		Name: "UseBeltItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[31] = OpCodeInfo{
		Name: "UseInventoryItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[32] = OpCodeInfo{
		Name: "UseItem",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "NumAdditionalItems"},
			{Size: 4, Name: "Item1UnitID"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesReq[33] = OpCodeInfo{
		Name: "SwitchBodyItem",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 4, Name: "TargetItemUnitId"},
			{Size: 4, Name: "SwitchItemUnitId"},
			{Size: 4, Name: "Unknown[0x00]"},
			{Size: 4, Name: "Unknown[0xFF]"},
			{Size: 4, Name: "TargetItemLocation"},
		},
	}

	OpCodesReq[34] = OpCodeInfo{
		Name: "IdentifyItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[35] = OpCodeInfo{
		Name: "ItemToBelt",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitID"},
			{Size: 4, Name: "BeltPosX"},
		},
	}

	OpCodesReq[36] = OpCodeInfo{
		Name: "RemoveBeltItem",
		Size: 24,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosX"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosXCurrent"},
			{Size: 1, Name: "ItemPosXPrevious"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosXCurrent"},
			{Size: 1, Name: "ItemPosXPrevious"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosXCurrent"},
			{Size: 1, Name: "ItemPosXPrevious"},
		},
	}

	OpCodesReq[37] = OpCodeInfo{
		Name: "CharToNpc",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "NpcId"},
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "Unknown"},
		},
	}

	OpCodesReq[38] = OpCodeInfo{
		Name: "NpcToChar",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[38] = OpCodeInfo{
		Name: "UseItemOnUnit",
		Size: 36,
		Fields: []FieldInfo{
			{Size: 1, Name: "Fix01"},
			{Size: 4, Name: "UnitID"},
			{Size: 1, Name: "UseOnMerc"},
			{Size: 1, Name: "FixFF"},
			{Size: 1, Name: "Fix00"},
			{Size: 1, Name: "Fix00"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "MercOrSelfUse"},
			{Size: 1, Name: "FixFF"},
			{Size: 1, Name: "ItemPosX"},
			{Size: 1, Name: "ItemPosY"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosXCurrent"},
			{Size: 1, Name: "ItemPosXPrevious"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosXCurrent"},
			{Size: 1, Name: "ItemPosXPrevious"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 1, Name: "ItemPosXCurrent"},
			{Size: 1, Name: "ItemPosXPrevious"},
		},
	}

	OpCodesReq[39] = OpCodeInfo{
		Name: "MoveStashGold",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitID"},
			{Size: 4, Name: "StashGold"},
			{Size: 4, Name: "InvGold"},
			{Size: 4, Name: "TransferAmount"},
		},
	}

	OpCodesReq[40] = OpCodeInfo{
		Name: "InsertItemToSocket",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemGUID"},
			{Size: 4, Name: "TargetItemGUID"},
			{Size: 2, Name: "Unknown"},
			{Size: 2, Name: "Unknown"},
			{Size: 4, Name: "InventoryId"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
		},
	}

	OpCodesReq[41] = OpCodeInfo{
		Name: "TakeWPOrPortal",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 1, Name: "WarpType"},
			{Size: 4, Name: "WarpId"},
			{Size: 4, Name: "Unknown"},
		},
	}

	OpCodesReq[42] = OpCodeInfo{
		Name: "MoveItemToCube",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitID"},
			{Size: 4, Name: "CubeUnitID"},
			{Size: 4, Name: "Unknown"},
			{Size: 2, Name: "Unknown"},
			{Size: 2, Name: "Unknown"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
		},
	}

	OpCodesReq[43] = OpCodeInfo{
		Name: "MakeEntityMove",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "EntityType"},
			{Size: 4, Name: "EntityId"},
		},
	}

	OpCodesReq[44] = OpCodeInfo{
		Name: "EntityAction",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "EntityType"},
			{Size: 4, Name: "EntityId"},
			{Size: 4, Name: "Action"},
		},
	}

	OpCodesReq[45] = OpCodeInfo{
		Name: "HealMerc",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesReq[46] = OpCodeInfo{
		Name: "HealOther",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
		},
	}

	OpCodesReq[47] = OpCodeInfo{
		Name: "NPCInit",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[48] = OpCodeInfo{
		Name: "NPCCancel",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[49] = OpCodeInfo{
		Name: "AssignSkillHotkey",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 2, Name: "Skill"},
			{Size: 4, Name: "ItemId"},
			{Size: 1, Name: "HotkeyId"},
		},
	}

	OpCodesReq[50] = OpCodeInfo{
		Name: "NPCBuy",
		Size: 24,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemPrice"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 4, Name: "NPCUnitID"},
			{Size: 2, Name: "ItemPosX"},
			{Size: 2, Name: "ItemPosY"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
			{Size: 1, Name: "SourceTab"},
			{Size: 1, Name: "TargetLocation"},
			{Size: 1, Name: "TransactionMode"},
		},
	}

	OpCodesReq[51] = OpCodeInfo{
		Name: "NPCSell",
		Size: 24,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemPrice"},
			{Size: 4, Name: "ItemUnitID"},
			{Size: 4, Name: "NPCUnitID"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
			{Size: 2, Name: "ItemPosX"},
			{Size: 2, Name: "ItemPosY"},
			{Size: 1, Name: "TargetTab"},
			{Size: 1, Name: "TargetLocation"},
			{Size: 1, Name: "TransactionMode"},
		},
	}

	OpCodesReq[52] = OpCodeInfo{
		Name: "NPCIdentifyItems",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 4, Name: "NPCUnitID"},
			{Size: 4, Name: "Unknown"},
			{Size: 4, Name: "CubeUnitID"},
			{Size: 4, Name: "Unknown"},
			{Size: 2, Name: "CubePosX"},
			{Size: 2, Name: "CubePosY"},
		},
	}

	OpCodesReq[53] = OpCodeInfo{
		Name: "NPCRepair",
		Size: 16,
		Fields: []FieldInfo{
			{Size: 1, Name: "Mode"},
			{Size: 1, Name: "ItemPosX"},
			{Size: 1, Name: "ItemPosY"},
			{Size: 4, Name: "NPCUnitID"},
			{Size: 4, Name: "RepairCosts"},
			{Size: 4, Name: "ItemUnitId"},
		},
	}

	OpCodesReq[54] = OpCodeInfo{
		Name: "QuestCompleted",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 1, Name: "QuestId"},
			{Size: 4, Name: "StateFlags"},
		},
	}

	OpCodesReq[55] = OpCodeInfo{
		Name: "MakeEntityInteractable",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "EntityType"},
			{Size: 4, Name: "EntityId"},
		},
	}

	OpCodesReq[56] = OpCodeInfo{
		Name: "NPCAction",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ActionId"},
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[57] = OpCodeInfo{
		Name: "AkaraRespec",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "AkaraID"},
		},
	}

	OpCodesReq[58] = OpCodeInfo{
		Name: "UseStatPoint",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatID"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
		},
	}
	OpCodesReq[59] = OpCodeInfo{
		Name: "UseSkillPoint",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 1, Name: "SkillID"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
			{Size: 1, Name: "Unknown"},
		},
	}

	OpCodesReq[60] = OpCodeInfo{
		Name: "SelectSkill",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillID"},
			{Size: 1, Name: "NULL"},
			{Size: 1, Name: "Hand"},
			{Size: 4, Name: "ItemGUID"},
		},
	}

	OpCodesReq[61] = OpCodeInfo{
		Name:   "ClearCursor",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[62] = OpCodeInfo{
		Name: "ActivateItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[63] = OpCodeInfo{
		Name:   "Terminate",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[64] = OpCodeInfo{
		Name:   "UpdateQuests",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[65] = OpCodeInfo{
		Name: "UnitInteractEx",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitId"},
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "Unknown"},
		},
	}

	OpCodesReq[66] = OpCodeInfo{
		Name:   "ReloadGame",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[67] = OpCodeInfo{
		Name: "RequestUnitUpdate",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitID"},
		},
	}

	OpCodesReq[68] = OpCodeInfo{
		Name: "DropGold",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "GoldAmount"},
		},
	}

	OpCodesReq[69] = OpCodeInfo{
		Name: "BindHotkeySkill",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "HotkeyId"},
			{Size: 2, Name: "SkillId"},
		},
	}

	OpCodesReq[70] = OpCodeInfo{
		Name: "PullItemFromSharedStash",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "OwnerId"},
			{Size: 2, Name: "FromPosX"},
			{Size: 2, Name: "FromPosY"},
		},
	}

	OpCodesReq[71] = OpCodeInfo{
		Name: "DropGold",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "InventoryGoldCount"},
			{Size: 4, Name: "DropGoldCount"},
		},
	}

	OpCodesReq[72] = OpCodeInfo{
		Name: "RelationRequest",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "EntityId"},
			{Size: 4, Name: "RelationType"},
		},
	}

	OpCodesReq[73] = OpCodeInfo{
		Name: "PartyRequest",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 2, Name: "PartyId"},
			{Size: 4, Name: "EntityId"},
		},
	}

	OpCodesReq[74] = OpCodeInfo{
		Name: "UpdatePlayerPos",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
		},
	}

	OpCodesReq[75] = OpCodeInfo{
		Name: "ToWaypoint",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "LevelNo"},
		},
	}

	OpCodesReq[76] = OpCodeInfo{
		Name: "SwapTwoHandedItem",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId1"},
			{Size: 4, Name: "ItemId2"},
			{Size: 4, Name: "BodyLocation1"},
			{Size: 4, Name: "BodyLocation2"},
		},
	}

	OpCodesReq[77] = OpCodeInfo{
		Name: "RemoveStackedItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[78] = OpCodeInfo{
		Name: "ItemStackToCursor",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "Amount"},
		},
	}

	OpCodesReq[79] = OpCodeInfo{
		Name: "ItemStackToGround",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "Amount"},
		},
	}

	OpCodesReq[80] = OpCodeInfo{
		Name: "SwapWeaponSlots",
		Size: 30,
		Fields: []FieldInfo{
			{Size: 4, Name: "LeftHandUnitId"},
			{Size: 4, Name: "RightHandUnitId"},
			{Size: 4, Name: "AltLeftHandUnitId"},
			{Size: 4, Name: "AltRightHandUnitId"},
			{Size: 4, Name: "Unkn"},
			{Size: 4, Name: "Unkn"},
			{Size: 2, Name: "LeftSkillId"},
			{Size: 2, Name: "RightSkillId"},
			{Size: 1, Name: "SwapSlotId"},
		},
	}

	OpCodesReq[81] = OpCodeInfo{
		Name: "RemoveBeltItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[82] = OpCodeInfo{
		Name: "NPCReviveMerc",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "NPCUnitId"},
			{Size: 4, Name: "MercNameId"},
			{Size: 4, Name: "MercReviveCosts"},
		},
	}

	OpCodesReq[83] = OpCodeInfo{
		Name: "InsertSocketItem",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "SocketedItemId"},
		},
	}

	OpCodesReq[84] = OpCodeInfo{
		Name: "QuickItemMove",
		Size: 21,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitId"},
			{Size: 4, Name: "FromInvPage"},
			{Size: 2, Name: "FromPosX"},
			{Size: 2, Name: "FromPosY"},
			{Size: 4, Name: "ToInvPage"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
		},
	}

	OpCodesReq[85] = OpCodeInfo{
		Name: "PutItemToSharedInventory",
		Size: 29,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitId"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
			{Size: 4, Name: "FromInvPage"},
			{Size: 2, Name: "FromPosX"},
			{Size: 2, Name: "FromPosY"},
			{Size: 4, Name: "ToInvPage"},
			{Size: 2, Name: "ToPosX"},
			{Size: 2, Name: "ToPosY"},
			{Size: 4, Name: "Unknown"},
		},
	}

	OpCodesReq[86] = OpCodeInfo{
		Name: "ScrollToBook",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ScrollId"},
			{Size: 4, Name: "BookId"},
		},
	}

	OpCodesReq[87] = OpCodeInfo{
		Name:   "CloseQuests",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[88] = OpCodeInfo{
		Name:   "CloseWaypoint",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[89] = OpCodeInfo{
		Name: "OpenCube",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "CubeId"},
		},
	}

	OpCodesReq[90] = OpCodeInfo{
		Name:   "OpenStash",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[91] = OpCodeInfo{
		Name: "ChangeSkillRightHand",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
		},
	}

	OpCodesReq[92] = OpCodeInfo{
		Name: "QuickItemDrop",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemUnitId"},
			{Size: 4, Name: "Unknown"},
			{Size: 4, Name: "FromInvPage"},
			{Size: 2, Name: "FromPosX"},
			{Size: 2, Name: "FromPosY"},
		},
	}

	OpCodesReq[93] = OpCodeInfo{
		Name: "ChangeSkillLeftHand",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
		},
	}

	OpCodesReq[94] = OpCodeInfo{
		Name:   "CloseAllWindows",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[95] = OpCodeInfo{
		Name: "InventoryToCube",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "CubeId"},
		},
	}

	OpCodesReq[96] = OpCodeInfo{
		Name: "CubeToInventory",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "CubeId"},
		},
	}

	OpCodesReq[97] = OpCodeInfo{
		Name:   "UnselectAllSkillHotkeys",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[98] = OpCodeInfo{
		Name: "SkillHotkeyToCursor",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "HotkeyId"},
		},
	}

	OpCodesReq[99] = OpCodeInfo{
		Name: "CursorToSkillHotkey",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 1, Name: "HotkeyId"},
			{Size: 1, Name: "SkillId"},
		},
	}

	OpCodesReq[100] = OpCodeInfo{
		Name: "CancelBuff",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "BuffId"},
			{Size: 2, Name: "EntityType"},
		},
	}

	OpCodesReq[101] = OpCodeInfo{
		Name: "DelayedState",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "StateId"},
			{Size: 2, Name: "EntityType"},
		},
	}

	OpCodesReq[102] = OpCodeInfo{
		Name: "SetPlayerMode",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "Mode"},
		},
	}

	OpCodesReq[103] = OpCodeInfo{
		Name:   "Stash_Arrange",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[104] = OpCodeInfo{
		Name:   "Stash_Collapse",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[105] = OpCodeInfo{
		Name: "Transmogrify",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[106] = OpCodeInfo{
		Name:   "ClickMercenaryBtn",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[107] = OpCodeInfo{
		Name: "ClickStashBtn",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "StashPage"},
		},
	}

	OpCodesReq[108] = OpCodeInfo{
		Name: "ClickBeltBtn",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "BeltSlot"},
		},
	}

	OpCodesReq[109] = OpCodeInfo{
		Name: "Ping",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "TickCount"},
			{Size: 4, Name: "Delay"},
			{Size: 4, Name: "WardenOrZero"},
		},
	}

	OpCodesReq[110] = OpCodeInfo{
		Name: "ClickQuestBtn",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "QuestId"},
		},
	}

	OpCodesReq[111] = OpCodeInfo{
		Name: "ActionConfirmation",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "Tick"},
		},
	}
	
	OpCodesReq[112] = OpCodeInfo{
		Name: "StatButtonPressed",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "StatType"},
		},
	}

	OpCodesReq[113] = OpCodeInfo{
		Name: "SkillButtonPressed",
		Size: 3,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
		},
	}

	OpCodesReq[114] = OpCodeInfo{
		Name: "SwitchInventoryPage",
		Size: 2,
		Fields: []FieldInfo{
			{Size: 1, Name: "PageNumber"},
		},
	}

	OpCodesReq[115] = OpCodeInfo{
		Name: "InventoryPosSwitch",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[116] = OpCodeInfo{
		Name:   "CancelTrade",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[117] = OpCodeInfo{
		Name:   "AcceptTrade",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[118] = OpCodeInfo{
		Name: "PutGoldInTrade",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "GoldAmount"},
		},
	}

	OpCodesReq[119] = OpCodeInfo{
		Name: "RemoveGoldFromTrade",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "GoldAmount"},
		},
	}

	OpCodesReq[120] = OpCodeInfo{
		Name: "TradeItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[121] = OpCodeInfo{
		Name: "CancelTradeItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[122] = OpCodeInfo{
		Name: "UpdateItemInTrade",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 1, Name: "TradeType"},
			{Size: 4, Name: "GoldAmount"},
		},
	}

	OpCodesReq[123] = OpCodeInfo{
		Name:   "CharacterPhaseEnd",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[124] = OpCodeInfo{
		Name: "NPCCancel",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "NPCId"},
		},
	}

	OpCodesReq[125] = OpCodeInfo{
		Name: "PlayerRelationAction",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "ActionType"},
			{Size: 2, Name: "PlayerId"},
			{Size: 4, Name: "Param"},
		},
	}

	OpCodesReq[126] = OpCodeInfo{
		Name: "RelationPartyAction",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 1, Name: "ActionType"},
			{Size: 2, Name: "PlayerId"},
			{Size: 4, Name: "Param"},
		},
	}

	OpCodesReq[127] = OpCodeInfo{
		Name: "AssignGameItem",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "OwnerId"},
			{Size: 4, Name: "TradeId"},
			{Size: 4, Name: "NewItemId"},
		},
	}

	OpCodesReq[128] = OpCodeInfo{
		Name: "PutBufferItemInStore",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
			{Size: 4, Name: "Cost"},
			{Size: 4, Name: "GambleId"},
		},
	}

	OpCodesReq[129] = OpCodeInfo{
		Name: "RemoveBufferItemFromStore",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[130] = OpCodeInfo{
		Name: "DynamicInteraction",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 2, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "ActionId"},
			{Size: -1, Name: "Data"},
		},
	}

	OpCodesReq[131] = OpCodeInfo{
		Name: "ClickButton2",
		Size: 4,
		Fields: []FieldInfo{
			{Size: 2, Name: "ButtonType"},
			{Size: 2, Name: "ButtonId"},
		},
	}

	OpCodesReq[132] = OpCodeInfo{
		Name: "ChatMsg",
		Size: -1,
		Fields: []FieldInfo{
			{Size: 1, Name: "ChatType"},
			{Size: -1, Name: "Message"},
		},
	}

	OpCodesReq[133] = OpCodeInfo{
		Name: "SpecialMove",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 2, Name: "MoveType"},
			{Size: 4, Name: "TargetId"},
		},
	}

	OpCodesReq[134] = OpCodeInfo{
		Name: "SpecialMoveAction",
		Size: 6,
		Fields: []FieldInfo{
			{Size: 2, Name: "ActionType"},
			{Size: 4, Name: "TargetId"},
		},
	}

	OpCodesReq[135] = OpCodeInfo{
		Name: "ChangeGameSettings",
		Size: 7,
		Fields: []FieldInfo{
			{Size: 2, Name: "SettingId"},
			{Size: 4, Name: "Value"},
		},
	}

	OpCodesReq[136] = OpCodeInfo{
		Name: "MoveFollower",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "FollowerId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
		},
	}

	OpCodesReq[137] = OpCodeInfo{
		Name:   "SwapWeapons",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[138] = OpCodeInfo{
		Name: "DropBufferItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[139] = OpCodeInfo{
		Name: "SpecialSkill",
		Size: 17,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesReq[140] = OpCodeInfo{
		Name: "SpecialSkillEx",
		Size: 13,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 4, Name: "UnitType"},
			{Size: 4, Name: "UnitId"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesReq[141] = OpCodeInfo{
		Name: "SpecialSkillEx2",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 2, Name: "PosX"},
			{Size: 2, Name: "PosY"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesReq[142] = OpCodeInfo{
		Name: "SpecialSkillEx3",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 2, Name: "SkillId"},
			{Size: 2, Name: "Unknown"},
		},
	}

	OpCodesReq[143] = OpCodeInfo{
		Name: "Scroll",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ScrollId"},
		},
	}

	OpCodesReq[144] = OpCodeInfo{
		Name:   "Resurrect",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[145] = OpCodeInfo{
		Name: "StaffInSocket",
		Size: 9,
		Fields: []FieldInfo{
			{Size: 4, Name: "StaffId"},
			{Size: 4, Name: "SocketId"},
		},
	}

	OpCodesReq[146] = OpCodeInfo{
		Name: "ItemToCursor",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[147] = OpCodeInfo{
		Name:   "BuyHealthPotion",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[148] = OpCodeInfo{
		Name:   "BuyManaPotion",
		Size:   1,
		Fields: []FieldInfo{},
	}

	OpCodesReq[149] = OpCodeInfo{
		Name: "SellItem",
		Size: 5,
		Fields: []FieldInfo{
			{Size: 4, Name: "ItemId"},
		},
	}

	OpCodesReq[150] = OpCodeInfo{
		Name:   "CainIdentifyItems",
		Size:   1,
		Fields: []FieldInfo{},
	}

}

func GetAckName(opCode uint8) string {
	if info, ok := OpCodesAck[opCode]; ok {
		return info.Name
	}
	return "nA."
}

func GetReqName(opCode uint8) string {
	if info, ok := OpCodesReq[opCode]; ok {
		return info.Name
	}
	return "nA."
}

func GetAckOpCode(name string) uint8 {
	for opCode, info := range OpCodesAck {
		if info.Name == name {
			return opCode
		}
	}
	return 0
}

func GetReqOpCode(name string) uint8 {
	for opCode, info := range OpCodesReq {
		if info.Name == name {
			return opCode
		}
	}
	return 0
}
