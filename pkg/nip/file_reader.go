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
		rule, err := ParseLine(fileScanner.Text())
		if errors.Is(err, errEmptyLine) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("error reading %s file at line %d: %w", filePath, lineNumber, err)
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
