package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) TerrorZones() (areas []area.Area) {
	tz := gd.moduleBaseAddressPtr + 0x299E2D8

	for i := 0; i < 7; i++ {
		tzArea := gd.ReadUInt(tz+uintptr(i*Uint32), Uint32)
		if tzArea != 0 {
			areas = append(areas, area.Area(tzArea))
		}
	}

	return
}
