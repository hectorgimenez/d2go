package data

import "github.com/hectorgimenez/d2go/pkg/data/object"

type Object struct {
	ID UnitID
	object.Name
	IsHovered    bool
	Selectable   bool
	InteractType object.InteractType
	Position     Position
	Owner        string
}

type Objects []Object

func (o Objects) FindOne(name object.Name) (Object, bool) {
	for _, obj := range o {
		if obj.Name == name {
			return obj, true
		}
	}

	return Object{}, false
}

func (o Object) IsWaypoint() bool {
	switch o.Name {
	case object.WaypointPortal,
		object.Act1WildernessWaypoint,
		object.Act2Waypoint,
		object.Act3TownWaypoint,
		object.PandamoniumFortressWaypoint,
		object.Act2CellerWaypoint,
		object.Act2SewerWaypoint,
		object.Act3TravincalWaypoint,
		object.ValleyWaypoint,
		object.WorldstoneWaypoint,
		object.ExpansionWildernessWaypoint,
		object.IceCaveWaypoint,
		object.TempleWaypoint,
		object.InnerHellWaypoint,
		object.WaypointH,
		object.ExpansionWaypoint:
		return true
	}

	return false
}

func (o Object) IsPortal() bool {
	return o.Name == object.TownPortal
}

func (o Object) IsRedPortal() bool {
	return o.Name == object.PermanentTownPortal
}

func (o Object) IsChest() bool {
	switch o.Name {
	case 1, 3, 5, 6, 50, 51, 53, 54, 55, 56, 57, 58, 79, 87, 88, 89, 104, 105, 106, 107, 125, 126, 127, 128, 139, 140, 141,
		144, 146, 147, 148, 154, 155, 158, 159, 169, 171, 174, 175, 176, 177, 178, 181, 182, 183, 185, 186, 187, 188, 198,
		203, 204, 205, 223, 224, 225, 240, 241, 242, 243, 244, 247, 248, 266, 268, 270, 271, 272, 274, 284, 314, 315, 316,
		317, 326, 329, 330, 331, 332, 333, 334, 335, 336, 354, 355, 356, 360, 371, 372, 380, 381, 383, 384, 387, 388, 389,
		390, 391, 397, 405, 406, 407, 413, 416, 420, 424, 425, 430, 431, 432, 433, 454, 455, 463, 466, 485, 486, 487, 501,
		502, 504, 505, 518, 524, 525, 526, 529, 530, 531, 532, 533, 534, 535, 540, 541, 544, 545, 556, 580, 581:
		return true
	}

	return false
}
func (o Object) IsDoor() bool {
	switch o.Name {
	case object.DoorCathedralLeft,
		object.DoorCathedralRight,
		object.DoorCathedralDouble,
		object.DoorCourtyardLeft,
		object.DoorCourtyardRight,
		object.DoorGateLeft,
		object.DoorGateRight,
		object.DoorMonasteryDoubleRight,
		object.DoorWoodenLeft,
		object.DoorWoodenLeft2,
		object.DoorWoodenRight,
		object.IronGrateDoorLeft,
		object.IronGrateDoorRight,
		object.SlimeDoor1,
		object.SlimeDoor2,
		object.TombDoorLeft,
		object.TombDoorLeft2,
		object.TombDoorRight,
		object.TombDoorRight2,
		object.WoodenDoorLeft,
		object.WoodenDoorRight,
		object.WoodenGrateDoorLeft,
		object.WoodenGrateDoorRight,
		object.AndarielDoor,
		object.PenBreakableDoor,
		object.ArreatSummitDoorToWorldstone,
		object.SecretDoor1,
		object.ExpansionTownGate:
		return true
	}

	return false
}

func (o Object) IsSuperChest() bool {
	switch o.Name {
	case 104, 105, 106, 107, 181, 183, 580, 397, 387, 389, 390, 391, 455:
		return true
	}

	return false
}
