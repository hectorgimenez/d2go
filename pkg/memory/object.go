package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/mode"
	"sort"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/object"
	"github.com/hectorgimenez/d2go/pkg/utils"
)

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
				txtFileNo := gd.Process.ReadUInt(objectUnitPtr+0x04, Uint32)
				unitID := gd.Process.ReadUInt(objectUnitPtr+0x08, Uint32)

				// Read the object mode
				objectMode := mode.ObjectMode(gd.Process.ReadUInt(objectUnitPtr+0x0c, Uint32))

				// Coordinates (X, Y)
				pathPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x38, Uint64))
				posX := gd.Process.ReadUInt(pathPtr+0x10, Uint16)
				posY := gd.Process.ReadUInt(pathPtr+0x14, Uint16)

				unitDataPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x10, Uint64))
				// checking if this is a shrine
				shrineTextPtr := uintptr(gd.Process.ReadUInt(objectUnitPtr+0x0A, Uint64))
				shrineType := uint(0)
				interactType := uint(0)
				shrineData := object.ShrineData{}
				if shrineTextPtr > 0 {
					shrineType = gd.Process.ReadUInt(unitDataPtr+0x08, Uint8)
					shrineData = object.ShrineData{
						ShrineName: object.ShrineTypeNames[object.ShrineType(shrineType)],
						ShrineType: object.ShrineType(shrineType),
					}
				} else {
					interactType = gd.Process.ReadUInt(unitDataPtr+0x08, Uint8)
				}
				owner := gd.Process.ReadStringFromMemory(unitDataPtr+0x34, 32)

				obj := data.Object{
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
					Owner: owner,
					Mode:  objectMode,
				}
				objects = append(objects, obj)
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
