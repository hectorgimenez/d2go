package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data"
)

func NewPanel(panelPtr uintptr, panelParent string, depth int, gd *GameReader) *data.Panel { //itemUnitPtr = uintptr(gd.Process.ReadUInt(itemUnitPtr+0x158, Uint64))
	panel := &data.Panel{
		PanelPtr:      panelPtr,
		PanelName:     gd.Process.ReadStringFromMemory(uintptr(gd.Process.ReadUInt(panelPtr+0x08, Uint64)), 0),
		PanelEnabled:  gd.Process.ReadUInt(panelPtr+0x50, Uint8) != 0,
		PanelVisible:  gd.Process.ReadUInt(panelPtr+0x51, Uint8) != 0,
		PtrChild:      uintptr(gd.Process.ReadUInt(panelPtr+0x58, Uint64)),
		NumChildren:   int(gd.Process.ReadUInt(panelPtr+0x60, Uint8)),
		ExtraText:     gd.Process.ReadStringFromMemory(panelPtr+0xA0, 0),
		ExtraText2:    gd.Process.ReadStringFromMemory(uintptr(gd.Process.ReadUInt(panelPtr+0x290, Uint64)), 0),
		ExtraText3:    gd.Process.ReadStringFromMemory(uintptr(gd.Process.ReadUInt(panelPtr+0x88, Uint64)), 0),
		PanelParent:   panelParent,
		PanelChildren: make([]data.Panel, 0),
		Depth:         depth,
	}
	if panel.NumChildren > 0 && panel.NumChildren < 50 {
		readPanel(panel.PtrChild, panel.NumChildren, &panel.PanelChildren, panel.PanelName, depth+1, gd)
	}
	return panel
}

func GetText(p data.Panel) string {
	text1 := cleanString(p.ExtraText)
	text2 := cleanString(p.ExtraText2)
	text3 := cleanString(p.ExtraText3)

	if text3 != "" && isASCII(text3) {
		return text3
	}
	if text2 != "" && isASCII(text2) {
		return text2
	}
	if text1 != "" && isASCII(text1) {
		return text1
	}
	return ""
}

func cleanString(input string) string {
	return input // Replace newlines and carriage returns as needed
}

func isASCII(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

func (gd *GameReader) ReadAllPanels() []data.Panel {
	base := gd.Process.moduleBaseAddressPtr + gd.offset.PanelManagerContainerOffset
	panelStructPtr := uintptr(gd.Process.ReadUInt(base, Uint64))
	panelPtr := uintptr(gd.Process.ReadUInt(panelStructPtr+0x58, Uint64))
	numChildren := int(gd.Process.ReadUInt(panelStructPtr+0x60, Uint8))

	panels := make([]data.Panel, 0)
	depth := 0
	// recursively read all panels, starting with the Root panel
	readPanel(panelPtr, numChildren, &panels, "Root", depth, gd)
	return panels
}

func readPanel(panelPtr uintptr, numChildren int, panels *[]data.Panel, panelParent string, depth int, gd *GameReader) {
	for i := 0; i < numChildren; i++ {
		panelStructPtr := uintptr(gd.Process.ReadUInt(uintptr(uint64(panelPtr)+uint64(i*8)), Uint64))
		thisPanel := NewPanel(panelStructPtr, panelParent, depth, gd)
		*panels = append(*panels, *thisPanel)
	}
}
