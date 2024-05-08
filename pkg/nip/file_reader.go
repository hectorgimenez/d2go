package nip

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
)

func ReadDir(path string) (Rules, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	rules := make([]Rule, 0)
	for _, file := range files {
		if !strings.HasSuffix(strings.ToLower(file.Name()), ".nip") || file.IsDir() {
			continue
		}

		newRules, err := ParseNIPFile(path + file.Name())
		if err != nil {
			return nil, err
		}

		rules = append(rules, newRules...)
	}

	return rules, nil
}

func ParseNIPFile(filePath string) (Rules, error) {
	fileToRead, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileToRead.Close()

	fileScanner := bufio.NewScanner(fileToRead)
	fileScanner.Split(bufio.ScanLines)

	rules := make([]Rule, 0)
	lineNumber := 0

	dummyItem := data.Item{
		ID:      516,
		Name:    "healingpotion",
		Quality: item.QualityNormal,
	}

	for fileScanner.Scan() {
		lineNumber++
		rule, err := NewRule(fileScanner.Text(), filePath, lineNumber)
		if errors.Is(err, ErrEmptyRule) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("error reading %s file at line %d: %w", filePath, lineNumber, err)
		}

		// We evaluate all the rules at startup to ensure no format errors, if there is a format error we will throw it now instead of during runtime
		_, err = rule.Evaluate(dummyItem)
		if err != nil {
			return nil, fmt.Errorf("error testing rule on [%s:%d]: %w", filePath, lineNumber, err)
		}
		rules = append(rules, rule)
	}

	return rules, nil
}

func sanitizeLine(rawLine string) string {
	l := strings.Split(rawLine, "//")
	line := strings.TrimSpace(l[0])
	line = strings.Join(strings.Fields(line), " ")
	line = strings.ReplaceAll(line, "'", "")

	// Fix possible wrong formatted lines
	line = strings.ReplaceAll(line, "=>", ">=")
	line = strings.ReplaceAll(line, "=<", "<=")
	line = strings.TrimSpace(strings.Trim(line, "&&"))

	return strings.ToLower(line)
}
