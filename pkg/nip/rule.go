package nip

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

var (
	fixedPropsRegexp = regexp.MustCompile(`(\[(type|quality|class|name|flag|color)]\s*(<=|<|>|>=|!=|==)\s*([a-zA-Z]+))`)
	statsRegexp      = regexp.MustCompile(`\[(.*?)]`)
	maxQtyRegexp     = regexp.MustCompile(`(\[maxquantity]\s*(<=|<|>|>=|!=|==)\s*([0-9]+))`)
)

type Rule struct {
	RawLine       string // Original line, don't use it for evaluation
	SanitizedLine string
	Filename      string
	LineNumber    int
	Enabled       bool
	maxQuantity   int
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
	rule := sanitizeLine(rawRule)

	maxQuantity := 0
	for _, prop := range maxQtyRegexp.FindAllStringSubmatch(rule, -1) {
		mxQty, err := strconv.Atoi(prop[3])
		if err != nil {
			return Rule{}, fmt.Errorf("error parsing maxquantity value %s: %w", prop[3], err)
		}
		maxQuantity = mxQty
		rule = strings.ReplaceAll(rule, prop[0], "")
	}

	// Sanitize again, just in case we messed up the rule while parsing maxquantity
	rule = sanitizeLine(rule)
	if rule == "" {
		return Rule{}, ErrEmptyRule
	}

	return Rule{
		RawLine:       rawRule,
		SanitizedLine: rule,
		Filename:      filename,
		LineNumber:    lineNumber,
		Enabled:       true,
		maxQuantity:   maxQuantity,
	}, nil
}

func (r Rule) Evaluate(it data.Item) (bool, error) {
	line := r.SanitizedLine
	for _, prop := range fixedPropsRegexp.FindAllStringSubmatch(line, -1) {
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
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", it.Desc().Tier()))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", classAliases[prop[4]]))
			line = strings.ReplaceAll(line, prop[0], partialReplace)
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
	return r.maxQuantity
}
