package mode

import "strings"

type ObjectMode uint32

const (
	ObjectModeIdle ObjectMode = 1 << iota
	ObjectModeOperating
	ObjectModeOpened
	ObjectModeSpecial1
	ObjectModeSpecial2
	ObjectModeSpecial3
	ObjectModeSpecial4
	ObjectModeSpecial5
)

// Has  specific mode is set
func (m ObjectMode) Has(mode ObjectMode) bool {
	return m&mode != 0
}

// String representation of the mode
func (m ObjectMode) String() string {
	modes := []string{}
	if m.Has(ObjectModeIdle) {
		modes = append(modes, "Idle")
	}
	if m.Has(ObjectModeOperating) {
		modes = append(modes, "Operating")
	}
	if m.Has(ObjectModeOpened) {
		modes = append(modes, "Opened")
	}
	if m.Has(ObjectModeSpecial1) {
		modes = append(modes, "Special1")
	}
	if m.Has(ObjectModeSpecial2) {
		modes = append(modes, "Special2")
	}
	if m.Has(ObjectModeSpecial3) {
		modes = append(modes, "Special3")
	}
	if m.Has(ObjectModeSpecial4) {
		modes = append(modes, "Special4")
	}
	if m.Has(ObjectModeSpecial5) {
		modes = append(modes, "Special5")
	}
	return strings.Join(modes, "|")
}
