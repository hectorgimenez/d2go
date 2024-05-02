package itemfilter

import (
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data/stat"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/nip"
)

func Evaluate(i data.Item, rules []nip.Rule) (nip.Rule, bool) {
	for _, r := range rules {
		if !evaluateGroups(i, r.Properties, checkProperty) {
			// Properties not matching, skipping
			continue
		}

		// We can not check stats, item is not identified, but properties matching
		if !i.Identified {
			return nip.Rule{}, true
		}

		if evaluateGroups(i, r.Stats, checkStat) {
			return r, true
		}
	}

	return nip.Rule{}, false
}

func evaluateGroups(i data.Item, groups []nip.Group, evalFunc func(i data.Item, prop nip.Comparable) bool) bool {
	groupChain := evaluationChain{}
	for _, group := range groups {
		propChain := evaluationChain{}
		for _, st := range group.Comparable {
			propChain.Add(evalFunc(i, st), st.Operand)
		}
		groupChain.Add(propChain.Evaluate(), group.Operand)
	}

	return groupChain.Evaluate()
}

func checkStat(i data.Item, cmp nip.Comparable) bool {
	st, found := stats[cmp.Keyword]
	if !found {
		// pass it, just in case...
		return true
	}

	itemStat, found := i.Stats[stat.ID(st[0])]
	if !found {
		return false
	}

	if !compare(itemStat.Value, cmp.ValueInt, cmp.Comparison) {
		return false
	}

	if len(st) == 1 {
		return true
	}

	return st[1] == itemStat.Layer
}

func checkProperty(i data.Item, prop nip.Comparable) bool {
	switch prop.Keyword {
	case nip.PropertyType:
		return strings.EqualFold(i.TypeAsString(), prop.ValueString)
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
