package memory

import (
	"sort"

	"github.com/hectorgimenez/d2go/pkg/data/entrance"
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
			// Read minimal data first to check object type
			objectType := gd.Process.ReadUInt(objectUnitPtr+0x00, Uint32)

			if objectType == 2 {
				rawTxtFileNo := gd.Process.ReadUInt(objectUnitPtr+0x04, Uint32) // Extract actual txtFileNo
				txtFileNo := rawTxtFileNo & 0xFFFF
				unitID := gd.Process.ReadUInt(objectUnitPtr+0x08, Uint32)
				objectMode := mode.ObjectMode(gd.Process.ReadUInt(objectUnitPtr+0x0c, Uint32))
				unitDataPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x10, Uint64))

				//This offset gives timer for each mode to keep progress in real time. exemple: Mode.Operating fresh timer then Mode.Opened new timer (for objects)
				// timerValue := uint32(gd.Process.ReadUInt(objectUnitPtr+0x5C, Uint32))

				// Path and position data
				pathPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x38, Uint64))
				// Coordinates (X, Y)
				posX := gd.Process.ReadUInt(pathPtr+0x10, Uint16)
				posY := gd.Process.ReadUInt(pathPtr+0x14, Uint16)

				var shrineData object.ShrineData
				var portalData object.PortalData
				interactType := gd.Process.ReadUInt(unitDataPtr+0x08, Uint8)
				owner := gd.Process.ReadStringFromMemory(unitDataPtr+0x34, 32)

				// Handle portals
				if isPortal(int(txtFileNo)) {
					destArea := area.ID(gd.Process.ReadUInt(unitDataPtr+0x08, Uint8))
					portalData.DestArea = destArea
					// Handle Shrines
				} else {
					shrineTextPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x0A, Uint64))
					if shrineTextPtr > 0 {
						shrineType := gd.Process.ReadUInt(unitDataPtr+0x08, Uint8)
						shrineData = object.ShrineData{
							ShrineName: object.ShrineTypeNames[object.ShrineType(shrineType)],
							ShrineType: object.ShrineType(shrineType),
						}
					}
				}
				// Handle objects
				objects = append(objects, data.Object{
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
				})
			}
			objectUnitPtr = uintptr(gd.Process.ReadUInt(objectUnitPtr+0x158, Uint64))
		}
	}

	if len(objects) > 0 {
		sort.SliceStable(objects, func(i, j int) bool {
			distanceI := utils.DistanceFromPoint(playerPosition, objects[i].Position)
			distanceJ := utils.DistanceFromPoint(playerPosition, objects[j].Position)
			return distanceI < distanceJ
		})
	}

	return objects
}
func (gd *GameReader) Entrances(playerPosition data.Position, hover data.HoverData) []data.Entrance {
	baseAddr := gd.Process.moduleBaseAddressPtr + gd.offset.UnitTable + (5 * 1024)
	unitTableBuffer := gd.Process.ReadBytesFromMemory(baseAddr, 128*8)

	var entrances []data.Entrance

	for i := 0; i < 128; i++ {
		entranceOffset := 8 * i
		entranceUnitPtr := uintptr(ReadUIntFromBuffer(unitTableBuffer, uint(entranceOffset), Uint64))

		for entranceUnitPtr > 0 {

			if entranceType := gd.Process.ReadUInt(entranceUnitPtr+0x00, Uint32); entranceType == 5 {
				txtFileNo := gd.Process.ReadUInt(entranceUnitPtr+0x04, Uint32)
				unitID := gd.Process.ReadUInt(entranceUnitPtr+0x08, Uint32)

				pathPtr := uintptr(gd.Process.ReadUInt(entranceUnitPtr+0x38, Uint64))
				posX := gd.Process.ReadUInt(pathPtr+0x10, Uint16)
				posY := gd.Process.ReadUInt(pathPtr+0x14, Uint16)

				entrances = append(entrances, data.Entrance{
					ID:        data.UnitID(unitID),
					Name:      entrance.Name(txtFileNo),
					IsHovered: data.UnitID(unitID) == hover.UnitID && hover.UnitType == 5 && hover.IsHovered,
					Position: data.Position{
						X: int(posX),
						Y: int(posY),
					},
				})
			}
			entranceUnitPtr = uintptr(gd.Process.ReadUInt(entranceUnitPtr+0x158, Uint64))
		}
	}

	return entrances
}
