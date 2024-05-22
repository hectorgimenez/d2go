package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) TerrorZones() (areas []area.ID) {
	tz := gd.moduleBaseAddressPtr + tzOnline

	for i := 0; i < 8; i++ {
		tzArea := gd.ReadUInt(tz+uintptr(i*Uint32), Uint32)
		if tzArea != 0 {
			areas = append(areas, area.ID(tzArea))
		}
	}

	return
}
