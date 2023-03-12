package itemfilter

import (
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/nip"
)

var qualities = map[string]item.Quality{
	nip.QualityLowQuality: item.QualityLowQuality,
	nip.QualityNormal:     item.QualityNormal,
	nip.QualitySuperior:   item.QualitySuperior,
	nip.QualityMagic:      item.QualityMagic,
	nip.QualitySet:        item.QualitySet,
	nip.QualityRare:       item.QualityRare,
	nip.QualityUnique:     item.QualityUnique,
	nip.QualityCrafted:    item.QualityCrafted,
}
