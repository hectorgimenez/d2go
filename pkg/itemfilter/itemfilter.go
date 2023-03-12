package itemfilter

import (
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/nip"
)

func Evaluate(i data.Item, rules []nip.Rule) bool {
	for _, r := range rules {
		if !checkProperties(i, r.Properties) {
			// Properties not matching, skipping
			continue
		}

		// We can not check stats, item is not identified, but properties matching
		if !i.Identified {
			return true
		}

	}

	return false
}

func checkProperties(i data.Item, properties []nip.Group) bool {
	groupChain := evaluationChain{}
	for _, propGroup := range properties {
		propChain := evaluationChain{}
		for _, prop := range propGroup.Comparable {
			propChain.Add(checkProperty(i, prop), prop.Operand)
		}
		groupChain.Add(propChain.Evaluate(), propGroup.Operand)
	}

	return groupChain.Evaluate()
}

func checkProperty(i data.Item, prop nip.Comparable) bool {
	switch prop.Keyword {
	case nip.PropertyType:
		return strings.EqualFold(i.Type(), prop.ValueString)
	case nip.PropertyName:
		return strings.EqualFold(string(i.Name), prop.ValueString)
	case nip.PropertyClass:
		// TODO: Implement
	case nip.PropertyQuality:
		quality, found := qualities[prop.ValueString]
		if !found {
			return false
		}

		return compare(int(i.Quality), int(quality), prop.Comparison)
	case nip.PropertyFlag:
		if prop.Comparison == nip.OperandEqual && !i.Ethereal {
			return false
		}
		if prop.Comparison == nip.OperandNotEqualTo && i.Ethereal {
			return false
		}
	case nip.PropertyLevel:
		// TODO: Implement
	case nip.PropertyPrefix:
		// TODO: Implement
	case nip.PropertySuffix:
		// TODO: Implement
	}

	return true
}

func compare(val1, val2 int, operand nip.Operand) bool {
	switch operand {
	case nip.OperandEqual:
		return val1 == val2
	case nip.OperandGreaterThan:
		return val1 > val2
	case nip.OperandGreaterOrEqualTo:
		return val1 >= val2
	case nip.OperandLessThan:
		return val1 < val2
	case nip.OperandLessThanOrEqualTo:
		return val1 <= val2
	case nip.OperandNotEqualTo:
		return val1 != val2
	}

	return false
}
