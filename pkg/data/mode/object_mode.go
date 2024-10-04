package mode

type ObjectMode uint32

const (
    ObjectModeIdle ObjectMode = iota
    ObjectModeOperating
    ObjectModeOpened
    ObjectModeSpecial1
    ObjectModeSpecial2
    ObjectModeSpecial3
    ObjectModeSpecial4
    ObjectModeSpecial5
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
