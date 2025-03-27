package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) TerrorZones() []area.ID {
	if gd == nil || gd.moduleBaseAddressPtr == 0 {
		return []area.ID{} // Return empty slice, not nil
	}

	tz := gd.moduleBaseAddressPtr + tzOnline

	// Use a temporary slice to collect current zones
	var currentZones []area.ID

	for i := 0; i < 8; i++ {
		tzArea := gd.ReadUInt(tz+uintptr(i*Uint32), Uint32)
		if tzArea != 0 {
			currentZones = append(currentZones, area.ID(tzArea))
		}
	}

	return currentZones
}

