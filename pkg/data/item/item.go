package item

type Description struct {
	ID              int
	Name            string
	Code            string
	NormalCode      string // Normal
	UberCode        string // Exceptional
	UltraCode       string // Elite
	InventoryWidth  int
	InventoryHeight int
	MinDefense      int
	MaxDefense      int
	Type            string
}

func (d Description) Tier() Tier {
	if d.Code == d.UltraCode {
		return TierElite
	}

	if d.Code == d.UberCode {
		return TierExceptional
	}

	return TierNormal
}

func (d Description) GetType() Type {
	return ItemTypes[d.Type]
}
