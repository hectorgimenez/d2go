package pickit

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	errEmptyLine      = errors.New("empty line")
	operandRegex      = regexp.MustCompile("(\\|{2}|\\&{2})")
	comparableRegex   = regexp.MustCompile("(\\={2}|\\<\\=|\\>\\=|\\>|\\<|\\!\\=)")
	propertyNameRegex = regexp.MustCompile("\\[(.*)\\]")
)

func parseLine(line string) (Rule, error) {
	line = lineCleanup(line)
	if line == "" {
		return Rule{}, errEmptyLine
	}

	lineSplit := strings.Split(line, "#")
	properties, err := buildGroups(lineSplit[0])
	if err != nil {
		return Rule{}, err
	}

	var stats []Group
	if len(lineSplit) > 1 {
		stats, err = buildGroups(lineSplit[1])
		if err != nil {
			return Rule{}, err
		}
	}

	var maxQuantity []Group
	if len(lineSplit) > 2 {
		maxQuantity, err = buildGroups(lineSplit[2])
		if err != nil {
			return Rule{}, err
		}
	}

	return Rule{
		Properties:  properties,
		Stats:       stats,
		MaxQuantity: maxQuantity,
	}, nil
}

func lineCleanup(line string) string {
	l := strings.Split(line, "//")
	line = strings.TrimSpace(l[0])

	return strings.ToLower(line)
}

func buildGroups(block string) ([]Group, error) {
	operands := operandRegex.FindAllString(block, -1)
	groupLines := operandRegex.Split(block, -1)

	operands, groupLines = reGroupParentheses(operands, groupLines)

	if len(operands) != len(groupLines)-1 {
		return nil, errors.New("syntax error")
	}

	groups := make([]Group, 0)
	for i, groupLine := range groupLines {
		cmp, err := buildGroupComparables(groupLine)
		if err != nil {
			return nil, err
		}

		group := Group{Comparable: cmp}
		if i < len(operands) {
			group.Operand = Operand(operands[i])
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func reGroupParentheses(operands, groups []string) ([]string, []string) {
	cleanOperands := make([]string, 0)
	cleanGroups := make([]string, 0)

	inGroup := false
	currentGroupContent := ""
	for i, group := range groups {
		group = strings.TrimSpace(group)
		if strings.Contains(group, "(") || (inGroup && !strings.Contains(group, ")")) {
			inGroup = true
			currentGroupContent += strings.ReplaceAll(group, "(", "") + operands[i]
			continue
		}
		if strings.Contains(group, ")") {
			inGroup = false
			currentGroupContent += strings.ReplaceAll(group, ")", "")
			cleanGroups = append(cleanGroups, currentGroupContent)
			currentGroupContent = ""
			if i < len(operands) {
				cleanOperands = append(cleanOperands, operands[i])
			}
			continue
		}
		if !inGroup {
			cleanGroups = append(cleanGroups, group)
			if i < len(operands) {
				cleanOperands = append(cleanOperands, operands[i])
			}
		}
	}

	return cleanOperands, cleanGroups
}

func buildGroupComparables(group string) ([]Comparable, error) {
	operands := operandRegex.FindAllString(group, -1)
	properties := operandRegex.Split(group, -1)
	if len(properties)-1 != len(operands) {
		return nil, errors.New("invalid NIP block")
	}

	comparables := make([]Comparable, 0)
	for i, prop := range properties {
		comparisonSymbol := comparableRegex.FindString(prop)
		values := comparableRegex.Split(prop, -1)
		if len(values) != 2 {
			return nil, errors.New("invalid NIP line")
		}

		aggregatedProperties := strings.Split(values[0], "+")
		for propNum, prop := range aggregatedProperties {
			propertyNameGr := propertyNameRegex.FindStringSubmatch(prop)
			propertyName := strings.TrimSpace(propertyNameGr[1])

			cmp := Comparable{
				Keyword:    propertyName,
				Comparison: Operand(strings.TrimSpace(comparisonSymbol)),
			}
			if len(aggregatedProperties) > 0 && propNum < len(aggregatedProperties)-1 {
				cmp.Operand = OperandOr
			} else if i < len(operands) {
				cmp.Operand = Operand(operands[i])
			}

			value := strings.TrimSpace(values[1])
			if val, err := strconv.Atoi(value); err == nil {
				cmp.ValueInt = val
			} else {
				cmp.ValueString = value
			}

			comparables = append(comparables, cmp)
		}
	}

	return comparables, nil
}
