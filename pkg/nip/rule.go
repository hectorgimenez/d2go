package nip

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

const (
	RuleResultFullMatch RuleResult = 1
	RuleResultPartial   RuleResult = 2
	RuleResultNoMatch   RuleResult = 3
)

var (
	fixedPropsRegexp = regexp.MustCompile(`(\[(type|quality|class|name|flag|color|prefix|suffix)]\s*(<=|<|>|>=|!=|==)\s*([a-zA-Z0-9]+))`)
	statsRegexp      = regexp.MustCompile(`\[(.*?)]`)
	maxQtyRegexp     = regexp.MustCompile(`(\[maxquantity]\s*(<=|<|>|>=|!=|==)\s*([0-9]+))`)
)

type Rule struct {
	RawLine       string // Original line, don't use it for evaluation
	Filename      string
	LineNumber    int
	Enabled       bool
	maxQuantity   int
	stage1        *vm.Program
	stage2        *vm.Program
	requiredStats []string
}

type RuleResult int
type Rules []Rule

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

var fixedPropsList = map[string]int{"type": 0, "quality": 0, "class": 0, "name": 0, "flag": 0, "color": 0, "prefix": 0, "suffix": 0}

func NewRule(rawRule string, filename string, lineNumber int) (Rule, error) {
	rule := sanitizeLine(rawRule)

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

	r := Rule{
		RawLine:     rawRule,
		Filename:    filename,
		LineNumber:  lineNumber,
		Enabled:     true,
		maxQuantity: maxQuantity,
	}

	parts := strings.Split(rule, "#")

	if len(parts) > 0 {
		stage1 := strings.TrimSpace(parts[0])
		if stage1 != "" {
			line, err := replaceStringPropertiesInStage1(stage1)
			if err != nil {
				return Rule{}, err
			}

			line = strings.ReplaceAll(line, "[", "")
			line = strings.ReplaceAll(line, "]", "")
			program, err := expr.Compile(line, expr.Env(fixedPropsList))
			if err != nil {
				return Rule{}, fmt.Errorf("error compiling rule stage1: %w", err)
			}
			r.stage1 = program
		}
	}

	if len(parts) > 1 {
		stage2 := strings.TrimSpace(parts[1])
		if stage2 != "" {
			// Extract stats before removing brackets for compilation
			r.requiredStats = getRequiredStatsForRule(stage2)

			statsMap := make(map[string]int)
			for _, prop := range r.requiredStats {
				statsMap[prop] = 0
			}

			// Normalize whitespace around operators in parenthesized expressions
			stage2 = normalizeParenthesizedExpressions(stage2)

			// Remove brackets for compilation
			compileReady := strings.ReplaceAll(stage2, "[", "")
			compileReady = strings.ReplaceAll(compileReady, "]", "")

			program, err := expr.Compile(compileReady, expr.Env(statsMap))
			if err != nil {
				return Rule{}, fmt.Errorf("error compiling rule stage2: %w, expression: %s", err, compileReady)
			}
			r.stage2 = program
		}
	}

	return r, nil
}

func normalizeParenthesizedExpressions(expr string) string {
	// Normalize common operators
	expr = strings.ReplaceAll(expr, "||", " || ")
	expr = strings.ReplaceAll(expr, "&&", " && ")
	expr = strings.ReplaceAll(expr, "==", " == ")
	expr = strings.ReplaceAll(expr, "!=", " != ")
	expr = strings.ReplaceAll(expr, ">=", " >= ")
	expr = strings.ReplaceAll(expr, "<=", " <= ")

	// Fix extra spaces
	expr = regexp.MustCompile(`\s+`).ReplaceAllString(expr, " ")

	// Normalize parentheses spacing
	expr = strings.ReplaceAll(expr, "( ", "(")
	expr = strings.ReplaceAll(expr, " )", ")")

	return expr
}

func (r Rule) Evaluate(it data.Item) (RuleResult, error) {
	// Stage 1: Basic properties evaluation
	stage1Props := make(map[string]int)
	for prop := range fixedPropsList {
		switch prop {
		case "type":
			stage1Props["type"] = it.Type().ID
		case "quality":
			stage1Props["quality"] = int(it.Quality)
		case "class":
			stage1Props["class"] = int(it.Desc().Tier())
		case "name":
			stage1Props["name"] = it.ID
		case "flag":
			stage1Props["flag"] = map[bool]int{true: 1, false: 0}[it.Ethereal]
		case "prefix":
			if it.Affixes.Rare.Prefix != 0 {
				stage1Props["prefix"] = int(it.Affixes.Rare.Prefix)
			}
			for _, prefix := range it.Affixes.Magic.Prefixes {
				if prefix != 0 {
					stage1Props["prefix"] = int(prefix)
					break
				}
			}
		case "suffix":
			if it.Affixes.Rare.Suffix != 0 {
				stage1Props["suffix"] = int(it.Affixes.Rare.Suffix)
			}
			for _, suffix := range it.Affixes.Magic.Suffixes {
				if suffix != 0 {
					stage1Props["suffix"] = int(suffix)
					break
				}
			}
		case "color":
			// TODO: Not supported yet
		}
	}

	// Check if stage1 exists before evaluating
	if r.stage1 == nil {
		return RuleResultNoMatch, fmt.Errorf("stage1 program is nil")
	}

	// Let's evaluate first stage
	stage1Result, err := expr.Run(r.stage1, stage1Props)
	if err != nil {
		return RuleResultNoMatch, fmt.Errorf("error evaluating rule stage1: %w", err)
	}

	// If stage1 does not match, we can stop here, nothing else to match
	if !stage1Result.(bool) {
		return RuleResultNoMatch, nil
	}

	// If we have no stage2 (no stat requirements), allow full match even for unidentified items
	if r.stage2 == nil {
		return RuleResultFullMatch, nil
	}

	// From here on we have stat requirements - return partial match for unidentified items
	if !it.Identified {
		return RuleResultPartial, nil
	}

	stage2Props := make(map[string]int)
	stage2 := ""
	if len(strings.Split(r.RawLine, "#")) > 1 {
		stage2 = strings.ToLower(strings.Split(r.RawLine, "#")[1])
	}

	// Preprocess stage2 to see if certain stats are being compared to zero
	zeroCheckStats := make(map[string]bool)
	for _, statName := range r.requiredStats {
		if strings.Contains(stage2, "["+statName+"] == 0") ||
			strings.Contains(stage2, "["+statName+"]==0") {
			zeroCheckStats[statName] = true
		}
	}

	// Handle resist sums
	hasAnyResist := false

	// Detect if this is a rule with a resist sum expression
	// We need to check both for direct addition/subtraction and parenthesized expressions
	isResistSum := false
	if strings.Contains(stage2, "resist") {
		isResistSum = strings.Contains(stage2, "+") || strings.Contains(stage2, "-") ||
			(strings.Contains(stage2, "(") && strings.Contains(stage2, ")"))
	}

	if isResistSum {
		// Check if the item has any resist stats at all
		for _, statName := range r.requiredStats {
			if !strings.Contains(statName, "resist") {
				continue
			}
			statData, found := statAliases[statName]
			if !found {
				continue
			}
			layer := 0
			if len(statData) > 1 {
				layer = statData[1]
			}

			if itemStat, found := it.FindStat(stat.ID(statData[0]), layer); found && itemStat.Value != 0 {
				hasAnyResist = true
				break
			}
		}
	}

	// Evaluate each required stat
	for _, statName := range r.requiredStats {
		statData, found := statAliases[statName]
		if !found {
			return RuleResultNoMatch, fmt.Errorf("property %s is not valid or not supported", statName)
		}

		layer := 0
		if len(statData) > 1 {
			layer = statData[1]
		}

		// Use the FindStat method which handles both Stats and BaseStats
		statFound := false
		var statValue int

		if itemStat, found := it.FindStat(stat.ID(statData[0]), layer); found {
			statValue = itemStat.Value
			statFound = true
		}
		// Special handling for stats not found
		if !statFound {
			// Check if this is a resist stat
			isResistStat := strings.Contains(statName, "resist")

			// For resist stats in a sum expression when no resists are present, don't match
			if isResistStat && isResistSum && !hasAnyResist {
				return RuleResultNoMatch, nil
			}

			// For all other stats, default to 0 to allow proper evaluation of OR conditions
			stage2Props[statName] = 0
		} else {
			// Stat was found, use its value
			stage2Props[statName] = statValue
		}
	}

	res, err := expr.Run(r.stage2, stage2Props)
	if err != nil {
		return RuleResultNoMatch, fmt.Errorf("error evaluating rule stage2: %w", err)
	}

	// 100% rule match, we can return here
	if res.(bool) {
		return RuleResultFullMatch, nil
	}

	return RuleResultNoMatch, nil
}

func replaceStringPropertiesInStage1(stage1 string) (string, error) {
	baseProperties := fixedPropsRegexp.FindAllStringSubmatch(stage1, -1)
	for _, prop := range baseProperties {
		replaceWith := ""
		switch prop[2] {
		case "type":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", item.ItemTypes[typeAliases[prop[4]]].ID))
		case "quality":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", qualityAliases[prop[4]]))
		case "class":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", classAliases[prop[4]]))
		case "name":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", item.GetIDByName(prop[4])))
		case "flag":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", 1))
		case "prefix", "suffix":
			// Handle prefix/suffix IDs
			replaceWith = strings.ReplaceAll(prop[0], prop[4], prop[4])
		case "color":
			// TODO: Not supported yet
			return "", fmt.Errorf("property %s is not supported yet", prop[2])
		}

		if replaceWith != "" {
			stage1 = strings.ReplaceAll(stage1, prop[0], replaceWith)
		}
	}

	return stage1, nil
}

func getRequiredStatsForRule(line string) []string {
	statsList := make([]string, 0)
	statsFound := make(map[string]bool)

	for _, statName := range statsRegexp.FindAllStringSubmatch(line, -1) {
		if !statsFound[statName[1]] {
			statsList = append(statsList, statName[1])
			statsFound[statName[1]] = true
		}
	}
	return statsList
}

// MaxQuantity returns the maximum quantity of items that character can have, 0 means no limit
func (r Rule) MaxQuantity() int {
	return r.maxQuantity
}
