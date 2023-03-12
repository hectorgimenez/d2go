package itemfilter

import "github.com/hectorgimenez/d2go/pkg/nip"

type evaluationChain struct {
	links []link
}

type link struct {
	Result  bool
	Operand nip.Operand
}

func (ch *evaluationChain) Add(result bool, operand nip.Operand) {
	ch.links = append(ch.links, link{
		Result:  result,
		Operand: operand,
	})
}

func (ch *evaluationChain) Evaluate() bool {
	if len(ch.links) == 1 {
		return ch.links[0].Result
	}

	result := ch.links[0].Result
	operand := ch.links[0].Operand
	for i, eval := range ch.links {
		if i == 0 {
			continue
		}

		if i < len(ch.links) {
			switch operand {
			case nip.OperandAnd:
				result = result && eval.Result
			case nip.OperandOr:
				result = result || eval.Result
			}
		}
		operand = eval.Operand
	}

	return result
}
