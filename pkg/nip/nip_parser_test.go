package nip

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const nipLine = "[type] == boots && [quality] == rare # [frw] >= 10 && [fireresist] >= 10 && ([lightresist]+[coldresist] >= 10 && [dexterity] >= 1 && [fireresist]+[poisonresist] >= 10) // this is a comment"

func Test_parseLine(t *testing.T) {
	rules, err := ParseLine(nipLine)
	require.NoError(t, err)

	expected := Rule{
		Properties: []Group{
			{
				Comparable: []Comparable{
					{
						Keyword:     "type",
						Comparison:  OperandEqual,
						ValueString: "boots",
					},
				},
				Operand: OperandAnd,
			},
			{
				Comparable: []Comparable{
					{
						Keyword:     "quality",
						Comparison:  OperandEqual,
						ValueString: "rare",
					},
				},
			},
		},
		Stats: []Group{
			{
				Comparable: []Comparable{
					{
						Keyword:    "frw",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   10,
					},
				},
				Operand: OperandAnd,
			},
			{
				Comparable: []Comparable{
					{
						Keyword:    "fireresist",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   10,
					},
				},
				Operand: OperandAnd,
			},
			{
				Comparable: []Comparable{
					{
						Keyword:    "lightresist",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   10,
						Operand:    OperandOr,
					},
					{
						Keyword:    "coldresist",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   10,
						Operand:    OperandAnd,
					},
					{
						Keyword:    "dexterity",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   1,
						Operand:    OperandAnd,
					},
					{
						Keyword:    "fireresist",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   10,
						Operand:    OperandOr,
					},
					{
						Keyword:    "poisonresist",
						Comparison: OperandGreaterOrEqualTo,
						ValueInt:   10,
					},
				},
			},
		},
	}

	assert.Equal(t, expected, rules)
}
