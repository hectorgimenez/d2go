package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/area"
)

func (gd *GameReader) TerrorZones() []area.ID {
	tz := gd.moduleBaseAddressPtr + tzOnline

	// Initialize an empty slice to hold only current terror zones -- Flush it to not keep previous when it changes
	areas := make([]area.ID, 0, 8)

	for i := 0; i < 8; i++ {
		tzArea := gd.ReadUInt(tz+uintptr(i*Uint32), Uint32)
		if tzArea != 0 {
			areas = append(areas, area.ID(tzArea))
		}
	}

	return areas
}

