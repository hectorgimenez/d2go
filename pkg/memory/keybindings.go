package memory

import (
	"encoding/binary"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/skill"
)

func (gd *GameReader) GetKeyBindings() data.KeyBindings {
	blob := gd.ReadBytesFromMemory(gd.moduleBaseAddressPtr+0x1e0e934, 0x500)
	blobSkills := gd.ReadBytesFromMemory(gd.moduleBaseAddressPtr+0x21f6a30, 0x500)

	skillsKB := [16]data.SkillBinding{}
	for i := 0; i < 7; i++ {
		skillsKB[i] = data.SkillBinding{
			SkillID: skill.ID(binary.LittleEndian.Uint32(blobSkills[i*0x1c : i*0x1c+4])),
			KeyBinding: data.KeyBinding{
				Key1: [2]byte{blob[0x118+(i*0x14)], blob[0x119+(i*0x14)]},
				Key2: [2]byte{blob[0x122+(i*0x14)], blob[0x123+(i*0x14)]},
			},
		}
	}
	for i := 0; i < 9; i++ {
		skillIdx := i + 7
		skillsKB[skillIdx] = data.SkillBinding{
			SkillID: skill.ID(binary.LittleEndian.Uint32(blobSkills[skillIdx*0x1c : skillIdx*0x1c+4])),
			KeyBinding: data.KeyBinding{
				Key1: [2]byte{blob[0x384+(i*0x14)], blob[0x385+(i*0x14)]},
				Key2: [2]byte{blob[0x38e+(i*0x14)], blob[0x38f+(i*0x14)]},
			},
		}
	}

	belt := [4]data.KeyBinding{}
	for i := 0; i < 4; i++ {
		belt[i] = data.KeyBinding{
			Key1: [2]byte{blob[0x1b8+(i*0x14)], blob[0x1b9+(i*0x14)]},
			Key2: [2]byte{blob[0x1c2+(i*0x14)], blob[0x1c3+(i*0x14)]},
		}
	}

	return data.KeyBindings{
		CharacterScreen: data.KeyBinding{
			Key1: [2]byte{blob[0x00], blob[0x01]},
			Key2: [2]byte{blob[0xa], blob[0xb]},
		},
		Inventory: data.KeyBinding{
			Key1: [2]byte{blob[0x14], blob[0x15]},
			Key2: [2]byte{blob[0x1e], blob[0x1f]},
		},
		HoradricCube: data.KeyBinding{
			Key1: [2]byte{blob[0x4b0], blob[0x4b1]},
			Key2: [2]byte{blob[0x4ba], blob[0x4bb]},
		},
		PartyScreen: data.KeyBinding{
			Key1: [2]byte{blob[0x28], blob[0x29]},
			Key2: [2]byte{blob[0x32], blob[0x33]},
		},
		MercenaryScreen: data.KeyBinding{
			Key1: [2]byte{blob[0x438], blob[0x439]},
			Key2: [2]byte{blob[0x442], blob[0x443]},
		},
		MessageLog: data.KeyBinding{
			Key1: [2]byte{blob[0x3c], blob[0x3d]},
			Key2: [2]byte{blob[0x46], blob[0x47]},
		},
		QuestLog: data.KeyBinding{
			Key1: [2]byte{blob[0x50], blob[0x51]},
			Key2: [2]byte{blob[0x5a], blob[0x5b]},
		},
		HelpScreen: data.KeyBinding{
			Key1: [2]byte{blob[0x78], blob[0x79]},
			Key2: [2]byte{blob[0x82], blob[0x83]},
		},
		SkillTree: data.KeyBinding{
			Key1: [2]byte{blob[0xf0], blob[0xf1]},
			Key2: [2]byte{blob[0xfa], blob[0xfb]},
		},
		SkillSpeedBar: data.KeyBinding{
			Key1: [2]byte{blob[0x104], blob[0x105]},
			Key2: [2]byte{blob[0x10e], blob[0x10f]},
		},
		Skills: skillsKB,
		SelectPreviousSkill: data.KeyBinding{
			Key1: [2]byte{blob[0x2f8], blob[0x2f9]},
			Key2: [2]byte{blob[0x302], blob[0x303]},
		},
		SelectNextSkill: data.KeyBinding{
			Key1: [2]byte{blob[0x30c], blob[0x30d]},
			Key2: [2]byte{blob[0x316], blob[0x317]},
		},
		ShowBelt: data.KeyBinding{
			Key1: [2]byte{blob[0x1a4], blob[0x1a5]},
			Key2: [2]byte{blob[0x1ae], blob[0x1af]},
		},
		UseBelt: belt,
		SwapWeapons: data.KeyBinding{
			Key1: [2]byte{blob[0x35c], blob[0x35d]},
			Key2: [2]byte{blob[0x366], blob[0x367]},
		},
		Chat: data.KeyBinding{
			Key1: [2]byte{blob[0x64], blob[0x65]},
			Key2: [2]byte{blob[0x6e], blob[0x6f]},
		},
		Run: data.KeyBinding{
			Key1: [2]byte{blob[0x294], blob[0x295]},
			Key2: [2]byte{blob[0x29e], blob[0x29f]},
		},
		ToggleRunWalk: data.KeyBinding{
			Key1: [2]byte{blob[0x2a8], blob[0x2a9]},
			Key2: [2]byte{blob[0x2b2], blob[0x2b3]},
		},
		StandStill: data.KeyBinding{
			Key1: [2]byte{blob[0x2bc], blob[0x2bd]},
			Key2: [2]byte{blob[0x2c6], blob[0x2c7]},
		},
		ForceMove: data.KeyBinding{
			Key1: [2]byte{blob[0x49c], blob[0x49d]},
			Key2: [2]byte{blob[0x4a6], blob[0x4a7]},
		},
		ShowItems: data.KeyBinding{
			Key1: [2]byte{blob[0x2d0], blob[0x2d1]},
			Key2: [2]byte{blob[0x2da], blob[0x2db]},
		},
		ShowPortraits: data.KeyBinding{
			Key1: [2]byte{blob[0x348], blob[0x349]},
			Key2: [2]byte{blob[0x352], blob[0x353]},
		},
		Automap: data.KeyBinding{
			Key1: [2]byte{blob[0x8c], blob[0x8d]},
			Key2: [2]byte{blob[0x96], blob[0x97]},
		},
		CenterAutomap: data.KeyBinding{
			Key1: [2]byte{blob[0xa0], blob[0xa1]},
			Key2: [2]byte{blob[0xaa], blob[0xab]},
		},
		FadeAutomap: data.KeyBinding{
			Key1: [2]byte{blob[0xb4], blob[0xb5]},
			Key2: [2]byte{blob[0xbe], blob[0xbf]},
		},
		PartyOnAutomap: data.KeyBinding{
			Key1: [2]byte{blob[0xc8], blob[0xc9]},
			Key2: [2]byte{blob[0xd2], blob[0xd3]},
		},
		NamesOnAutomap: data.KeyBinding{
			Key1: [2]byte{blob[0xdc], blob[0xdd]},
			Key2: [2]byte{blob[0xe6], blob[0xe7]},
		},
		ToggleMiniMap: data.KeyBinding{
			Key1: [2]byte{blob[0x370], blob[0x371]},
			Key2: [2]byte{blob[0x37a], blob[0x37b]},
		},
		SayHelp: data.KeyBinding{
			Key1: [2]byte{blob[0x208], blob[0x209]},
			Key2: [2]byte{blob[0x212], blob[0x213]},
		},
		SayFollowMe: data.KeyBinding{
			Key1: [2]byte{blob[0x21c], blob[0x21d]},
			Key2: [2]byte{blob[0x226], blob[0x227]},
		},
		SayThisIsForYou: data.KeyBinding{
			Key1: [2]byte{blob[0x230], blob[0x231]},
			Key2: [2]byte{blob[0x23a], blob[0x23b]},
		},
		SayThanks: data.KeyBinding{
			Key1: [2]byte{blob[0x244], blob[0x245]},
			Key2: [2]byte{blob[0x24e], blob[0x24f]},
		},
		SaySorry: data.KeyBinding{
			Key1: [2]byte{blob[0x258], blob[0x259]},
			Key2: [2]byte{blob[0x262], blob[0x263]},
		},
		SayBye: data.KeyBinding{
			Key1: [2]byte{blob[0x26c], blob[0x26d]},
			Key2: [2]byte{blob[0x276], blob[0x277]},
		},
		SayNowYouDie: data.KeyBinding{
			Key1: [2]byte{blob[0x280], blob[0x281]},
			Key2: [2]byte{blob[0x28a], blob[0x28b]},
		},
		SayRetreat: data.KeyBinding{
			Key1: [2]byte{blob[0x44c], blob[0x44d]},
			Key2: [2]byte{blob[0x456], blob[0x457]},
		},
		ClearScreen: data.KeyBinding{
			Key1: [2]byte{blob[0x2e4], blob[0x2e5]},
			Key2: [2]byte{blob[0x2ee], blob[0x2ef]},
		},
		ClearMessages: data.KeyBinding{
			Key1: [2]byte{blob[0x320], blob[0x321]},
			Key2: [2]byte{blob[0x32a], blob[0x32b]},
		},
		Zoom: data.KeyBinding{
			Key1: [2]byte{blob[0x474], blob[0x475]},
			Key2: [2]byte{blob[0x47e], blob[0x47f]},
		},
		LegacyToggle: data.KeyBinding{
			Key1: [2]byte{blob[0x488], blob[0x489]},
			Key2: [2]byte{blob[0x492], blob[0x493]},
		},
	}
}
