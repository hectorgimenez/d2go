package nip

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

var (
	fixedPropsRegexp = regexp.MustCompile(`(\[(type|quality|class|name|flag|color)]\s*(<=|<|>|>=|!=|==)\s*([a-zA-Z]+))`)
	statsRegexp      = regexp.MustCompile(`\[(.*?)]`)
)

type Rule struct {
	RawLine    string
	Filename   string
	LineNumber int
	Enabled    bool
}

type Rules []Rule

func (r Rules) EvaluateAll(it data.Item) (Rule, bool) {
	for _, rule := range r {
		if rule.Enabled {
			result, err := rule.Evaluate(it)
			if err != nil {
				continue
			}
			if result {
				return rule, true
			}
		}
	}

	return Rule{}, false
}

func New(rawRule string, filename string, lineNumber int) (Rule, error) {
	rawRule = sanitizeLine(rawRule)
	if rawRule == "" {
		return Rule{}, ErrEmptyRule
	}

	return Rule{
		RawLine:    rawRule,
		Filename:   filename,
		LineNumber: lineNumber,
		Enabled:    true,
	}, nil
}

func (r Rule) Evaluate(it data.Item) (bool, error) {
	line := r.RawLine
	for _, prop := range fixedPropsRegexp.FindAllStringSubmatch(r.RawLine, -1) {
		switch prop[2] {
		case "type":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", typeAliases[it.TypeAsString()]))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", typeAliases[prop[4]]))
			line = strings.ReplaceAll(line, prop[0], partialReplace)
		case "quality":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", it.Quality))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", qualityAliases[prop[4]]))
			line = strings.ReplaceAll(line, prop[0], partialReplace)
		case "class":
			// TODO: ???
			//line = strings.ReplaceAll(line, fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", it.Class))
			line = strings.ReplaceAll(line, prop[4], fmt.Sprintf("%d", classAliases[prop[4]]))
		case "name":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", item.GetIDByName(string(it.Name))))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", item.GetIDByName(prop[4])))
			line = strings.ReplaceAll(line, prop[0], partialReplace)
		case "flag":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", map[bool]int{true: 1, false: 0}[it.Ethereal]))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", 1))
			line = strings.ReplaceAll(line, prop[0], partialReplace)
		case "color":
			// TODO: Not supported yet
		}
	}

	for _, statName := range statsRegexp.FindAllStringSubmatch(line, -1) {
		statData, found := statAliases[statName[1]]
		if !found {
			return false, fmt.Errorf("property %s is not valid or not supported", statName[1])
		}

		layer := 0
		if len(statData) > 1 {
			layer = statData[1]
		}
		itemStat, _ := it.FindStat(stat.ID(statData[0]), layer)
		// By default, value will be 0 is stat is not found, it's okay for evaluation purposes.
		line = strings.ReplaceAll(line, statName[0], fmt.Sprintf("%d", itemStat.Value))
	}

	output, err := expr.Eval(line, nil)
	if err != nil {
		return false, fmt.Errorf("error evaluating rule: %w", err)
	}

	return output.(bool), nil
}

// MaxQuantity returns the maximum quantity of items that character can have, 0 means no limit
func (r Rule) MaxQuantity() int {
	return 0
}
