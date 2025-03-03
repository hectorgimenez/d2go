package memory

import (
	"fmt"
	"math"
	"time"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

type GameReader struct {
	offset Offset
	Process

	monstersLastUpdate  time.Time
	inventoryLastUpdate time.Time
	objectsLastUpdate   time.Time

	cachedMonsters  data.Monsters
	cachedInventory data.Inventory
	cachedObjects   []data.Object
}

var WidgetStateFlags = map[string]uint64{
	"WeaponSwap": 0xF2D7CF8E9CC08212,
}

func NewGameReader(process Process) *GameReader {
	return &GameReader{
		offset:              calculateOffsets(process),
		Process:             process,
		monstersLastUpdate:  time.Time{},
		inventoryLastUpdate: time.Time{},
		objectsLastUpdate:   time.Time{},
	}
}

func (gd *GameReader) GetData() data.Data {
	if gd.offset.UnitTable == 0 {
		gd.offset = calculateOffsets(gd.Process)
	}

	// Always refresh core player data
	rawPlayerUnits := gd.GetRawPlayerUnits()
	mainPlayerUnit := rawPlayerUnits.GetMainPlayer()
	pu := gd.GetPlayerUnit(mainPlayerUnit)
	hover := gd.hoveredData()

	now := time.Now()

	// Conditionally update monsters
	monsters := gd.cachedMonsters
	if now.Sub(gd.monstersLastUpdate) > 200*time.Millisecond {
		monsters = gd.Monsters(pu.Position, hover)
		gd.cachedMonsters = monsters
		gd.monstersLastUpdate = now
	}

	// Conditionally update inventory 500ms
	// Except when hovering over an item
	inventory := gd.cachedInventory
	if now.Sub(gd.inventoryLastUpdate) > 500*time.Millisecond ||
		(hover.IsHovered && hover.UnitType == 4) { // 4 = Item type
		inventory = gd.Inventory(rawPlayerUnits, hover)
		gd.cachedInventory = inventory
		gd.inventoryLastUpdate = now
	}

	// Conditionally update objects
	objects := gd.cachedObjects
	if now.Sub(gd.objectsLastUpdate) > 200*time.Millisecond {
		objects = gd.Objects(pu.Position, hover)
		gd.cachedObjects = objects
		gd.objectsLastUpdate = now
	}

	// Always update other critical data
	corpseUnit := rawPlayerUnits.GetCorpse()
	roster := gd.getRoster(rawPlayerUnits)
	openMenus := gd.openMenus()

	// Quests
	// q1 := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+0x22E2978, Uint64))
	// q2 := uintptr(gd.Process.ReadUInt(q1, Uint64))
	// q2 := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+0x22F1E79, Uint64))
	gameQuestsBytes := gd.Process.ReadBytesFromMemory(gd.moduleBaseAddressPtr+0x22F1E79, 85)
	// gameQuestsBytes = gameQuestsBytes[3:]

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
		Monsters:       monsters,
		Corpses:        gd.Corpses(pu.Position, hover),
		PlayerUnit:     pu,
		Inventory:      inventory,
		Objects:        objects,
		Entrances:      gd.Entrances(pu.Position, hover),
		OpenMenus:      openMenus,
		Roster:         roster,
		HoverData:      hover,
		TerrorZones:    gd.TerrorZones(),
		Quests:         gd.getQuests(gameQuestsBytes),
		KeyBindings:    gd.GetKeyBindings(),
		LegacyGraphics: gd.LegacyGraphics(),
		IsIngame:       gd.IsIngame(),

		// These use the Panel Manager which is heavy to read. Use the functions below instead.
		//IsOnline:       		   gd.IsOnline(),
		//IsInCharCreationScreen:  gd.IsInCharacterCreationScreen(),
		//IsInLobby:               gd.IsInLobby(),
		//IsInCharSelectionScreen: gd.IsInCharacterSelectionScreen(),
		HasMerc:          gd.HasMerc(),
		ActiveWeaponSlot: gd.GetActiveWeaponSlot(),
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

// GetPanel returns a Panel object from the specified path (starting from the root panel)
func (gd *GameReader) GetPanel(panelPath ...string) data.Panel {
	if len(panelPath) == 0 {
		return data.Panel{}
	}

	// Get all panels
	allPanels := gd.ReadAllPanels()

	// Start with the first panel in the path
	firstPanelName := panelPath[0]
	currentPanel, exists := allPanels[firstPanelName]
	if !exists {
		// Panel not found at top level
		return data.Panel{}
	}

	// Traverse the path from left to right
	for i := 1; i < len(panelPath); i++ {
		childName := panelPath[i]
		nextPanel, exists := currentPanel.PanelChildren[childName]
		if !exists {
			return data.Panel{}
		}
		currentPanel = nextPanel
	}

	return currentPanel
}

func (gd *GameReader) InCharacterSelectionScreen() bool {
	panel := gd.GetPanel("CharacterSelectPanel")
	return panel.PanelName != "" && panel.PanelEnabled && panel.PanelVisible
}

func (gd *GameReader) GetSelectedCharacterName() string {
	return gd.Process.ReadStringFromMemory(gd.Process.moduleBaseAddressPtr+0x222D0A8, 0)
}

func (gd *GameReader) LegacyGraphics() bool {
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x2227998, Uint64) == 1
}

func (gd *GameReader) IsOnline() bool {
	panel := gd.GetPanel("MainMenuPanel", "SecondaryContextButton")
	return panel.PanelName != "" && panel.PanelEnabled && panel.PanelVisible
}

func (gd *GameReader) IsIngame() bool {
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x22E51D0, 1) == 1
}

func (gd *GameReader) IsInLobby() bool {
	panel := gd.GetPanel("LobbyBackgroundPanel")
	return panel.PanelName != "" && panel.PanelEnabled && panel.PanelVisible
}

func (gd *GameReader) IsInCharacterSelectionScreen() bool {
	panel := gd.GetPanel("CharacterSelectPanel")
	return panel.PanelName != "" && panel.PanelEnabled && panel.PanelVisible
}

func (gd *GameReader) IsInCharacterCreationScreen() bool {
	panel := gd.GetPanel("CharacterCreatePanel")
	return panel.PanelName != "" && panel.PanelEnabled && panel.PanelVisible
}

func (gd *GameReader) GetCharacterList() []string {
	containerPanel := gd.GetPanel("CharacterSelectPanel", "Background", "CharacterList", "View", "Container")
	if containerPanel.PanelName == "" || containerPanel.NumChildren == 0 {
		return []string{}
	}

	// Get the character names that are in the container children [ListView 0,1,2 (0 indexed)] -> children -> Name -> Extra Text 3
	characterNames := make([]string, containerPanel.NumChildren)
	for i := 0; i < containerPanel.NumChildren; i++ {
		characterNames[i] = containerPanel.PanelChildren[fmt.Sprintf("ListItem%d", i)].PanelChildren["Name"].ExtraText3
	}

	return characterNames
}

// IsDismissableModalPresent checks if there's a error popup present
func (gd *GameReader) IsDismissableModalPresent() (bool, string) {
	panel := gd.GetPanel("DismissableModal")

	if panel.PanelName == "" {
		return false, ""
	}

	modalText := panel.PanelChildren["Frame"].PanelChildren["Prompt"].ExtraText3
	return (panel.PanelName != "" && panel.PanelEnabled && panel.PanelVisible), modalText
}

func (gd *GameReader) LastGameName() string {
	return gd.ReadStringFromMemory(gd.moduleBaseAddressPtr+0x2587FB8, 0)
}

func (gd *GameReader) LastGamePass() string {
	return gd.ReadStringFromMemory(gd.moduleBaseAddressPtr+0x2588018, 0)
}

func (gd *GameReader) FPS() int {
	return int(gd.ReadUInt(gd.moduleBaseAddressPtr+gd.offset.FPS, Uint32))
}

func (gd *GameReader) HasMerc() bool {
	return gd.ReadUInt(gd.moduleBaseAddressPtr+0x22e51d0+0x12, Uint8) != 0
}

// GetWidgetState reference : https://github.com/ResurrectedTrader/ResurrectedTrade/blob/f121ec02dd3fbe1c574f713e5a0c2db92ccca821/ResurrectedTrade.AgentBase/Capture.cs#L618
func (gd *GameReader) GetWidgetState(stateFlag uint64) (int, error) {
	// Get widget states pointer
	stateFlags := uint64(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+gd.offset.WidgetStatesOffset, Uint64))
	if stateFlags == 0 {
		return 0, nil
	}

	v2 := uint64(gd.Process.ReadUInt(uintptr(stateFlags)+8, Uint64))
	if v2 == 0 {
		return 0, nil
	}

	flag := stateFlag
	v4 := uint64(0xC4CEB9FE1A85EC53) * ((uint64(0xFF51AFD7ED558CCD) * (flag ^ (flag >> 33))) ^ ((uint64(0xFF51AFD7ED558CCD) * (flag ^ (flag >> 33))) >> 33))
	v5 := (uint64(gd.Process.ReadUInt(uintptr(stateFlags), Uint64)) - 1) & (v4 ^ (v4 >> 33))
	v6 := uint64(gd.Process.ReadUInt(uintptr(v2)+uintptr(8*v5), Uint64))

	i := uintptr(v2) + uintptr(8*v5)

	for ; v6 != 0; v6 = uint64(gd.Process.ReadUInt(uintptr(v6), Uint64)) {
		if flag == uint64(gd.Process.ReadUInt(uintptr(v6)+8, Uint64)) {
			break
		}
		i = uintptr(v6)
	}

	ir := uint64(gd.Process.ReadUInt(i, Uint64))
	if ir != 0 {
		ptr1 := uint64(gd.Process.ReadUInt(uintptr(ir)+16, Uint64))
		ptr2 := uint64(gd.Process.ReadUInt(uintptr(ptr1)+16, Uint64))
		return int(gd.Process.ReadUInt(uintptr(ptr2), Uint8)), nil
	}

	return 0, nil
}

func (gd *GameReader) GetActiveWeaponSlot() int {
	state, err := gd.GetWidgetState(WidgetStateFlags["WeaponSwap"])
	if err != nil {
		return 0 // Default to primary weapons on error
	}
	return state
}
