package nip

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/hectorgimenez/d2go/pkg/data"
)

type Rule struct {
	RawLine    string
	Filename   string
	LineNumber int
	Enabled    bool
}

func New(rawRule string, filename string, lineNumber int) (Rule, error) {
	rawRule = ruleLineCleanup(rawRule)
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

func (r Rule) Evaluate(item data.Item) bool {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matchLine := strings.ReplaceAll(r.RawLine, "#", "&&")
	rawProperties := re.FindAllStringSubmatch(r.RawLine, -1)

	for _, prop := range rawProperties {
		switch prop[1] {
		case "type":
			matchLine = strings.ReplaceAll(matchLine, "[type]", fmt.Sprintf("%d", item.Type))
		case "quality":
			matchLine = strings.ReplaceAll(matchLine, "[quality]", fmt.Sprintf("%d", item.Quality))
		case "class":
			// TODO: ???
		case "name":
			matchLine = strings.ReplaceAll(matchLine, "[name]", fmt.Sprintf("%s", item.Name))
		case "flag":
			matchLine = strings.ReplaceAll(matchLine, prop[0], fmt.Sprintf("%d", 1))
		case "color":
			// TODO: ???
		default:
			value := 0
			for statID, stat := range item.Stats {
				if strings.EqualFold(statID.String(), prop[1]) {
					value = stat.Value
				}
			}
			matchLine = strings.ReplaceAll(matchLine, prop[0], fmt.Sprintf("%d", value))
		}
	}

	output, err := expr.Eval(matchLine, nil)
	if err != nil {
		return false
	}

	fmt.Print(output)
	return false
}
