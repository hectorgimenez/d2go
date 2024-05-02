package nip

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadDir(path string) ([]Rule, error) {
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

func ParseNIPFile(filePath string) ([]Rule, error) {
	fileReader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)

	rules := make([]Rule, 0)
	lineNumber := 0
	for fileScanner.Scan() {
		lineNumber++
		rule, err := New(fileScanner.Text(), filePath, lineNumber)
		if errors.Is(err, ErrEmptyRule) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("error reading %s file at line %d: %w", filePath, lineNumber, err)
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func ruleLineCleanup(rawLine string) string {
	l := strings.Split(rawLine, "//")
	rawLine = strings.TrimSpace(l[0])
	rawLine = strings.ReplaceAll(rawLine, "'", "")

	return strings.ToLower(rawLine)
}
