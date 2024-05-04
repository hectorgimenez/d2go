package nip

import (
	"fmt"
	"regexp"
	"slices"
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
	RawLine     string // Original line, don't use it for evaluation
	Stages      [3]string
	Filename    string
	LineNumber  int
	Enabled     bool
	maxQuantity int
}

type Rules []Rule

type RuleResult int

var (
	RuleResultFullMatch RuleResult = 1
	RuleResultPartial   RuleResult = 2
	RuleResultNoMatch   RuleResult = 3
)

func (r Rules) EvaluateAll(it data.Item) (Rule, RuleResult) {
	bestMatch := RuleResultNoMatch
	bestMatchingRule := Rule{}
	for _, rule := range r {
		if rule.Enabled {
			result, err := rule.Evaluate(it)
			if err != nil {
				continue
			}
			if result == RuleResultFullMatch {
				return rule, result
			}
			if result == RuleResultPartial {
				bestMatch = result
				bestMatchingRule = rule
			}
		}
	}

	return bestMatchingRule, bestMatch
}

func New(rawRule string, filename string, lineNumber int) (Rule, error) {
	rule := sanitizeLine(rawRule)

	// Check for not supported stats
	for _, prop := range statsRegexp.FindAllStringSubmatch(rule, -1) {
		if slices.Contains(notSupportedStats, prop[1]) {
			return Rule{}, fmt.Errorf("property %s is not supported, please remove it", prop[1])
		}
	}

	// Try to get the maxquantity value and purge it from the rule, we can not evaluate it
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

	stages := [3]string{}
	for i, stg := range strings.Split(rule, "#") {
		stages[i] = strings.TrimSpace(stg)
	}

	return Rule{
		RawLine:     rawRule,
		Stages:      stages,
		Filename:    filename,
		LineNumber:  lineNumber,
		Enabled:     true,
		maxQuantity: maxQuantity,
	}, nil
}

func (r Rule) Evaluate(it data.Item) (RuleResult, error) {
	stage1 := r.Stages[0]
	baseProperties := fixedPropsRegexp.FindAllStringSubmatch(stage1, -1)
	for _, prop := range baseProperties {
		switch prop[2] {
		case "type":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", typeAliases[it.TypeAsString()]))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", typeAliases[prop[4]]))
			stage1 = strings.ReplaceAll(stage1, prop[0], partialReplace)
		case "quality":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", it.Quality))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", qualityAliases[prop[4]]))
			stage1 = strings.ReplaceAll(stage1, prop[0], partialReplace)
		case "class":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", it.Desc().Tier()))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", classAliases[prop[4]]))
			stage1 = strings.ReplaceAll(stage1, prop[0], partialReplace)
		case "name":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", item.GetIDByName(string(it.Name))))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", item.GetIDByName(prop[4])))
			stage1 = strings.ReplaceAll(stage1, prop[0], partialReplace)
		case "flag":
			partialReplace := strings.ReplaceAll(prop[0], fmt.Sprintf("[%s]", prop[2]), fmt.Sprintf("%d", map[bool]int{true: 1, false: 0}[it.Ethereal]))
			partialReplace = strings.ReplaceAll(partialReplace, prop[4], fmt.Sprintf("%d", 1))
			stage1 = strings.ReplaceAll(stage1, prop[0], partialReplace)
		case "color":
			// TODO: Not supported yet
		}
	}

	// Let's evaluate first stage
	stage1Result, err := expr.Eval(stage1, nil)
	if err != nil {
		return RuleResultNoMatch, fmt.Errorf("error evaluating rule: %w", err)
	}

	// Let's go with other stats now
	// TODO: properties are missing (enhanceddefense, enhanceddamage, etc)
	stage2 := r.Stages[1]
	stage2Result := true
	if stage2 != "" {
		for _, statName := range statsRegexp.FindAllStringSubmatch(stage2, -1) {
			statData, found := statAliases[statName[1]]
			if !found {
				return RuleResultNoMatch, fmt.Errorf("property %s is not valid or not supported", statName[1])
			}

			layer := 0
			if len(statData) > 1 {
				layer = statData[1]
			}
			itemStat, _ := it.FindStat(stat.ID(statData[0]), layer)
			// By default, value will be 0 is stat is not found, it's okay for evaluation purposes.
			stage2 = strings.ReplaceAll(stage2, statName[0], fmt.Sprintf("%d", itemStat.Value))
		}

		res, err := expr.Eval(stage2, nil)
		if err != nil {
			return RuleResultNoMatch, fmt.Errorf("error evaluating rule: %w", err)
		}
		stage2Result = res.(bool)
	}

	// 100% rule match, we can return here
	if stage1Result.(bool) && stage2Result {
		return RuleResultFullMatch, nil
	}

	// If Stage 1 matches and the item is NOT identified, let's return a partial match, we can not be 100% sure if all the stats are matching
	if stage1Result.(bool) && !it.Identified {
		return RuleResultPartial, nil
	}

	return RuleResultNoMatch, nil
}

// MaxQuantity returns the maximum quantity of items that character can have, 0 means no limit
func (r Rule) MaxQuantity() int {
	return r.maxQuantity
}
