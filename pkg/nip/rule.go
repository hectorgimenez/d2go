package nip

const (
	OperandEqual             Operand = "=="
	OperandGreaterThan       Operand = ">"
	OperandGreaterOrEqualTo  Operand = ">="
	OperandLessThan          Operand = "<"
	OperandLessThanOrEqualTo Operand = "<="
	OperandNotEqualTo        Operand = "!="
	OperandAnd               Operand = "&&"
	OperandOr                Operand = "||"
	OperandNone              Operand = ""
)

type Rule struct {
	Properties  []Group
	Stats       []Group
	MaxQuantity []Group
}

type Keyword string
type Operand string

type Comparable struct {
	Keyword     string
	Comparison  Operand
	ValueInt    int
	ValueString string
	Operand
}

type Group struct {
	Comparable []Comparable
	Operand
}
