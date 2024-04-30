package itemfilter

import (
	"testing"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/nip"
	"github.com/stretchr/testify/require"
)

func TestEvaluate(t *testing.T) {
	type args struct {
		i       data.Item
		nipRule string // not the best test but too lazy to write rules manually
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				i: data.Item{
					Name:     item.Name("lightplatedboots"),
					Quality:  item.QualityUnique,
					Ethereal: false,
				},
				nipRule: "[name] == lightplatedboots && [quality] == unique && [flag] != ethereal # [enhanceddefense] == 60 // goblin toe",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule, err := nip.ParseLine(tt.args.nipRule, "dummy", 0)
			require.NoError(t, err)

			if _, got := Evaluate(tt.args.i, []nip.Rule{rule}); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
