package mode

type ObjectMode uint32

const (
	ObjectModeIdle      ObjectMode = 0
	ObjectModeOperating ObjectMode = 1
	ObjectModeOpened    ObjectMode = 2
	ObjectModeSpecial1  ObjectMode = 3
	ObjectModeSpecial2  ObjectMode = 4
	ObjectModeSpecial3  ObjectMode = 5
	ObjectModeSpecial4  ObjectMode = 6
	ObjectModeSpecial5  ObjectMode = 7
)

func (m ObjectMode) String() string {
	switch m {
	case ObjectModeIdle:
		return "Idle"
	case ObjectModeOperating:
		return "Operating"
	case ObjectModeOpened:
		return "Opened"
	case ObjectModeSpecial1:
		return "Special1"
	case ObjectModeSpecial2:
		return "Special2"
	case ObjectModeSpecial3:
		return "Special3"
	case ObjectModeSpecial4:
		return "Special4"
	case ObjectModeSpecial5:
		return "Special5"
	default:
		return "Unknown"
	}
}
