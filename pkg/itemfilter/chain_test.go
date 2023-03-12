package itemfilter

import (
	"testing"

	"github.com/hectorgimenez/d2go/pkg/nip"
)

func Test_evaluationChain_Evaluate(t *testing.T) {
	tests := []struct {
		name  string
		links []link
		want  bool
	}{
		{
			name: "Given one single link with true value should return true",
			links: []link{
				{true, nip.OperandNone},
			},
			want: true,
		},
		{
			name: "Given one single link with false value should return false",
			links: []link{
				{false, nip.OperandNone},
			},
			want: false,
		},
		{
			name: "Given different links using OR and one valid value should be true",
			links: []link{
				{false, nip.OperandOr},
				{true, nip.OperandNone},
			},
			want: true,
		},
		{
			name: "Given different links using OR and no valid value should be false",
			links: []link{
				{false, nip.OperandOr},
				{false, nip.OperandNone},
			},
			want: false,
		},
		{
			name: "Given different links using AND and one valid value should be false",
			links: []link{
				{true, nip.OperandAnd},
				{false, nip.OperandNone},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &evaluationChain{
				links: tt.links,
			}
			if got := ch.Evaluate(); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
