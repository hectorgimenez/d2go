package object

type InteractType uint

const (
	InteractTypeNone   InteractType = 0x00
	InteractTypeTrap   InteractType = 0x04
	InteractTypeShrine InteractType = 0x08
	InteractTypeLocked InteractType = 0x80
)
