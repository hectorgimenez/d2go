package memory

import (
	"math"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

type GameReader struct {
	offset Offset
	Process
}

func NewGameReader(process Process) *GameReader {
	return &GameReader{
		offset:  calculateOffsets(process),
		Process: process,
	}
}

func (gd *GameReader) GetData() data.Data {
	if gd.offset.UnitTable == 0 {
		gd.offset = calculateOffsets(gd.Process)
	}

	rawPlayerUnits := gd.GetRawPlayerUnits()
	roster := gd.getRoster(rawPlayerUnits)
	mainPlayerUnit := rawPlayerUnits.GetMainPlayer()

	pu := gd.GetPlayerUnit(mainPlayerUnit)
	hover := gd.hoveredData()

	// Quests
	// q1 := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+0x22E2978, Uint64))
	// q2 := uintptr(gd.Process.ReadUInt(q1, Uint64))
	// q2 := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+0x22F1E79, Uint64))
	gameQuestsBytes := gd.Process.ReadBytesFromMemory(gd.moduleBaseAddressPtr+0x22F1E79, 85)

	// gameQuestsBytes = gameQuestsBytes[3:]

	corpseUnit := rawPlayerUnits.GetCorpse()
	d := data.Data{
		Corpse: data.Corpse{
			Found:     corpseUnit.Address != 0,
			IsHovered: corpseUnit.IsHovered,
			Position:  corpseUnit.Position,
			States:    corpseUnit.States,
		},
		Game: data.OnlineGame{
			LastGameName:     gd.LastGameName(),
			LastGamePassword: gd.LastGamePass(),
			FPS:              gd.FPS(),
		},
		Monsters:                gd.Monsters(pu.Position, hover),
		Corpses:                 gd.Corpses(pu.Position, hover),
		PlayerUnit:              pu,
		Inventory:               gd.Inventory(rawPlayerUnits, hover),
		Objects:                 gd.Objects(pu.Position, hover),
		Entrances:               gd.Entrances(pu.Position, hover),
		OpenMenus:               gd.openMenus(),
		Widgets:                 gd.UpdateWidgets(),
		Roster:                  roster,
		HoverData:               hover,
		TerrorZones:             gd.TerrorZones(),
		Quests:                  gd.getQuests(gameQuestsBytes),
		KeyBindings:             gd.GetKeyBindings(),
		LegacyGraphics:          gd.LegacyGraphics(),
		IsOnline:                gd.IsOnline(),
		IsIngame:                gd.IsIngame(),
		IsInCharCreationScreen:  gd.IsInCharacterCreationScreen(),
		IsInLobby:               gd.IsInLobby(),
		IsInCharSelectionScreen: gd.IsInCharacterSelectionScreen(),
		HasMerc:                 gd.HasMerc(),
	}

	return d
}

func (gd *GameReader) InGame() bool {
	player := gd.GetRawPlayerUnits().GetMainPlayer()

	return player.UnitID > 0 && player.Position.X > 0 && player.Position.Y > 0 && player.Area > 0
}

func (gd *GameReader) openMenus() data.OpenMenus {
	uiBase := gd.Process.moduleBaseAddressPtr + gd.offset.UI - 0xA

	buffer := gd.Process.ReadBytesFromMemory(uiBase, 0x16D)

	isMapShown := gd.Process.ReadUInt(gd.Process.moduleBaseAddressPtr+gd.offset.UI, Uint8)

	return data.OpenMenus{
		Inventory:      buffer[0x01] != 0,
		LoadingScreen:  buffer[0x168] != 0,
		NPCInteract:    buffer[0x08] != 0,
		NPCShop:        buffer[0x0B] != 0,
		Stash:          buffer[0x18] != 0,
		Waypoint:       buffer[0x13] != 0,
		MapShown:       isMapShown != 0,
		SkillTree:      buffer[0x04] != 0,
		NewSkills:      buffer[0x07] != 0,
		NewStats:       buffer[0x06] != 0,
		Character:      buffer[0x02] != 0,
		QuitMenu:       buffer[0x09] != 0,
		Cube:           buffer[0x19] != 0,
		SkillSelect:    buffer[0x03] != 0,
		Anvil:          buffer[0x0D] != 0,
		MercInventory:  buffer[0x1E] != 0,
		BeltRows:       buffer[0x1A] != 0,
		QuestLog:       buffer[0xE] != 0,
		PortraitsShown: buffer[0x1D] != 0,
		ChatOpen:       buffer[0x05] != 0,
	}
}

func (gd *GameReader) hoveredData() data.HoverData {
	hoverAddressPtr := gd.Process.moduleBaseAddressPtr + gd.offset.Hover
	hoverBuffer := gd.Process.ReadBytesFromMemory(hoverAddressPtr, 12)
	isUnitHovered := ReadUIntFromBuffer(hoverBuffer, 0, Uint16)
	if isUnitHovered > 0 {
		hoveredType := ReadUIntFromBuffer(hoverBuffer, 0x04, Uint32)
		hoveredUnitID := ReadUIntFromBuffer(hoverBuffer, 0x08, Uint32)

		return data.HoverData{
			IsHovered: true,
			UnitID:    data.UnitID(hoveredUnitID),
			UnitType:  int(hoveredType),
		}
	}

	return data.HoverData{}
}

func (gd *GameReader) getStatsList(statListPtr uintptr) stat.Stats {
	statsListBuffer := gd.ReadBytesFromMemory(statListPtr, 0x10)
	statList := ReadUIntFromBuffer(statsListBuffer, 0, Uint64)
	statCount := ReadUIntFromBuffer(statsListBuffer, 0x08, Uint64)
	if statCount == 0 {
		return []stat.Data{}
	}

	var stats = make([]stat.Data, 0)

	statBuffer := gd.Process.ReadBytesFromMemory(uintptr(statList), statCount*10)
	for i := 0; i < int(statCount); i++ {
		offset := uint(i * 8)

		statLayer := ReadUIntFromBuffer(statBuffer, offset, Uint16)
		statEnum := ReadUIntFromBuffer(statBuffer, offset+0x2, Uint16)
		statValue := ReadIntFromBuffer(statBuffer, offset+0x4, Uint32)

		value := statValue
		switch stat.ID(statEnum) {
		case stat.Life,
			stat.MaxLife,
			stat.Mana,
			stat.MaxMana,
			stat.Stamina,
			stat.MaxStamina:
			value = statValue >> 8
		case stat.ColdLength,
			stat.PoisonLength:
			value = statValue / 25
		case stat.DeadlyStrikePerLevel:
			value = int(float64(statValue) / .8)
		case stat.HitCausesMonsterToFlee:
			value = int(float64(statValue) / 1.28)
		case stat.AttackRatingUndeadPerLevel:
			value = statValue / 2
		case stat.MagicFindPerLevel,
			stat.ExtraGoldPerLevel,
			stat.DamageDemonPerLevel,
			stat.DamageUndeadPerLevel,
			stat.DefensePerLevel,
			stat.MaxDamagePerLevel,
			stat.MaxDamagePercentPerLevel,
			stat.StrengthPerLevel,
			stat.DexterityPerLevel,
			stat.VitalityPerLevel,
			stat.ThornsPerLevel:
			value = int(math.Max(float64(statValue/8), 1))
		case stat.LifePerLevel,
			stat.ManaPerLevel:
			value = int(math.Max(float64(statValue/2048), 1))
		case stat.ReplenishDurability, stat.ReplenishQuantity:
			value = int(math.Max(float64(2/statValue), 1))
		case stat.RegenStaminaPerLevel:
			value = int(statValue) * 10

		case stat.LevelRequirePercent:
			value = int(statValue) * -1
		case stat.AttackRatingPerLevel:
			value = int(math.Max(float64(statValue), 15))
		}

		stats = append(stats, stat.Data{
			ID:    stat.ID(statEnum),
			Value: value,
			Layer: int(statLayer),
		})
	}

	return stats
}

// TODO: Take a look to better ways to get this data, now it's very flakky, is just a random memory position + not in game
func (gd *GameReader) InCharacterSelectionScreen() bool {
	cs_visible, err := gd.IsWidgetVisible("CharacterSelectPanel")
	if err != nil {
	}
	return cs_visible
}

func (gd *GameReader) GetSelectedCharacterName() string {
	return gd.Process.ReadStringFromMemory(gd.Process.moduleBaseAddressPtr+0x222D0A8, 0)
}

func (gd *GameReader) LegacyGraphics() bool {
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x2227998, Uint64) == 1
}

func (gd *GameReader) IsOnline() bool {
	// This represents which tab (Online/Offline) we're on in the Character Selection Screen
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x21864D0, 1) == 1
}

func (gd *GameReader) IsIngame() bool {
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x22E51D0, 1) == 1
}

func (gd *GameReader) IsInLobby() bool {
	widgets := gd.UpdateWidgets()
	if lobbyWidget, found := widgets["LobbyBackgroundPanel"]; found {
		return lobbyWidget["WidgetActive"].(bool) && lobbyWidget["WidgetVisible"].(bool)
	}

	return false
}

func (gd *GameReader) IsInCharacterSelectionScreen() bool {
	widgets := gd.UpdateWidgets()
	if csWidget, found := widgets["CharacterSelectPanel"]; found {
		return csWidget["WidgetActive"].(bool) && csWidget["WidgetVisible"].(bool)
	}

	return false
}

func (gd *GameReader) IsInCharacterCreationScreen() bool {
	widgets := gd.UpdateWidgets()
	if ccWidget, found := widgets["CharacterCreatePanel"]; found {
		return ccWidget["WidgetActive"].(bool) && ccWidget["WidgetVisible"].(bool)
	}

	return false
}

func (gd *GameReader) LastGameName() string {
	return gd.ReadStringFromMemory(gd.moduleBaseAddressPtr+0x2587FB8, 0)
}

func (gd *GameReader) LastGamePass() string {
	return gd.ReadStringFromMemory(gd.moduleBaseAddressPtr+0x2588018, 0)
}

func (gd *GameReader) FPS() int {
	return int(gd.ReadUInt(gd.moduleBaseAddressPtr+0x2140DF4, 4))
}

func (gd *GameReader) HasMerc() bool {
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x22e51d0+0x12, Uint8) != 0
}
func (gd *GameReader) UpdateWidgets() map[string]map[string]interface{} {
	widgets := map[string]map[string]interface{}{}

	if gd.offset.PanelManagerContainerOffset == 0 {
		gd.offset = calculateOffsets(gd.Process)
	}

	// Assuming PanelManagerContainer address is known and stored in gd
	panelManagerContainerPtrAddr := gd.Process.moduleBaseAddressPtr + gd.offset.PanelManagerContainerOffset

	// Read the PanelManagerContainer pointer value
	panelManagerContainerAddr, err := gd.Process.ReadPointer(panelManagerContainerPtrAddr, 8)
	if err != nil {
		return widgets
	}
	// Read the Panel Managers WidgetContainer
	widgetContainer, err := gd.Process.ReadWidgetContainer(panelManagerContainerAddr, true)
	if err != nil {
		return widgets
	}
	// Read the list of child widgets
	childWidgets, err := gd.Process.ReadWidgetList(widgetContainer["ChildWidgetsListPointer"].(uintptr), int(widgetContainer["ChildWidgetSize"].(uint)))
	if err != nil {
		return widgets
	}
	return childWidgets
}

// IsWidgetVisible checks if any child widget on the PanelManager has the same name and has the Active and Visible booleans on.
func (gd *GameReader) IsWidgetVisible(widgetName string) (bool, error) {
	widgets := gd.UpdateWidgets()
	widget, exists := widgets[widgetName]
	if !exists {
		return false, nil
	}

	return widget["WidgetActive"].(bool) && widget["WidgetVisible"].(bool), nil
}
