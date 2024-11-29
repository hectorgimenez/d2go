package memory

import (
	"sort"

	"github.com/hectorgimenez/d2go/pkg/data/mode"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/object"
	"github.com/hectorgimenez/d2go/pkg/utils"
)

func isPortal(txtFileNo int) bool {
	desc, ok := object.Desc[txtFileNo]
	return ok && desc.Name == "Portal"
}

func (gd *GameReader) Objects(playerPosition data.Position, hover data.HoverData) []data.Object {
	baseAddr := gd.Process.moduleBaseAddressPtr + gd.offset.UnitTable + (2 * 1024)
	unitTableBuffer := gd.Process.ReadBytesFromMemory(baseAddr, 128*8)

	var objects []data.Object
	for i := 0; i < 128; i++ {
		objectOffset := 8 * i
		objectUnitPtr := uintptr(ReadUIntFromBuffer(unitTableBuffer, uint(objectOffset), Uint64))
		for objectUnitPtr > 0 {
			objectType := gd.Process.ReadUInt(objectUnitPtr+0x00, Uint32)
			if objectType == 2 {
				rawTxtFileNo := gd.Process.ReadUInt(objectUnitPtr+0x04, Uint32)
				txtFileNo := rawTxtFileNo & 0xFFFF // Extract actual txtFileNo

				unitID := gd.Process.ReadUInt(objectUnitPtr+0x08, Uint32)
				// Read the object mode
				objectMode := mode.ObjectMode(gd.Process.ReadUInt(objectUnitPtr+0x0c, Uint32))

				//This offset gives timer for each mode to keep progress in real time. exemple: Mode.Operating fresh timer then Mode.Opened new timer (for objects)
				// timerValue := uint32(gd.Process.ReadUInt(objectUnitPtr+0x5C, Uint32))

				// Path and position data
				pathPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x38, Uint64))
				// Coordinates (X, Y)
				posX := gd.Process.ReadUInt(pathPtr+0x10, Uint16)
				posY := gd.Process.ReadUInt(pathPtr+0x14, Uint16)

				unitDataPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x10, Uint64))

				var destArea area.ID
				var shrineData object.ShrineData
				var interactType uint
				var portalData object.PortalData

				owner := gd.Process.ReadStringFromMemory(unitDataPtr+0x34, 32)

				if isPortal(int(txtFileNo)) {
					destArea = area.ID(gd.Process.ReadUInt(unitDataPtr+0x08, Uint8))

					portalData = object.PortalData{
						DestArea: destArea,
					}
				} else {
					// Handle shrines and other objects

					shrineTextPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x0A, Uint64))
					// Checking if this is a shrine
					if shrineTextPtr > 0 {
						shrineType := gd.Process.ReadUInt(unitDataPtr+0x08, Uint8)
						shrineData = object.ShrineData{
							ShrineName: object.ShrineTypeNames[object.ShrineType(shrineType)],
							ShrineType: object.ShrineType(shrineType),
						}
					} else {
						interactType = gd.Process.ReadUInt(unitDataPtr+0x08, Uint8)
					}
				}

				object := data.Object{
					ID:           data.UnitID(unitID),
					Name:         object.Name(int(txtFileNo)),
					IsHovered:    data.UnitID(unitID) == hover.UnitID && hover.UnitType == 2 && hover.IsHovered,
					InteractType: object.InteractType(interactType),
					Shrine:       shrineData,
					Selectable:   objectMode == mode.ObjectModeIdle,
					Position: data.Position{
						X: int(posX),
						Y: int(posY),
					},
					Owner:      owner,
					Mode:       objectMode,
					PortalData: portalData,
				}
				objects = append(objects, object)
			}
			objectUnitPtr = uintptr(gd.Process.ReadUInt(objectUnitPtr+0x150, Uint64))
		}
	}

	sort.SliceStable(objects, func(i, j int) bool {
		distanceI := utils.DistanceFromPoint(playerPosition, objects[i].Position)
		distanceJ := utils.DistanceFromPoint(playerPosition, objects[j].Position)
		return distanceI < distanceJ
	})

	return objects
}
