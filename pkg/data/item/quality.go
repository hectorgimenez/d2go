package item

type Quality int

const (
	QualityLowQuality Quality = 0x01
	QualityNormal     Quality = 0x02
	QualitySuperior   Quality = 0x03
	QualityMagic      Quality = 0x04
	QualitySet        Quality = 0x05
	QualityRare       Quality = 0x06
	QualityUnique     Quality = 0x07
	QualityCrafted    Quality = 0x08
)

func (q Quality) ToString() string {
	switch q {
	case QualityLowQuality:
		return "LowQuality"
	case QualityNormal:
		return "Normal"
	case QualitySuperior:
		return "Superior"
	case QualityMagic:
		return "Magic"
	case QualitySet:
		return "Set"
	case QualityRare:
		return "Rare"
	case QualityUnique:
		return "Unique"
	case QualityCrafted:
		return "Crafted"
	}

	return "UnknownItemQuality"
}
