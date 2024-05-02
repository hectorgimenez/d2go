package nip

import (
	"testing"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

func TestRule_Evaluate(t *testing.T) {
	type fields struct {
		RawLine    string
		Filename   string
		LineNumber int
		Enabled    bool
	}
	type args struct {
		item data.Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "[name] == smallcharm && [quality] == magic # (([poisonlength]*25)*[poisonmaxdam])/256 >= 123",
			fields: fields{
				RawLine:    "[name] == smallcharm && [quality] == magic # (([poisonlength]*25)*[poisonmaxdam])/256 >= 123",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					Name:    "SmallCharm",
					Quality: item.QualityMagic,
					Stats: map[stat.ID]stat.Data{
						stat.PoisonLength:    {Value: 20},
						stat.PoisonMaxDamage: {Value: 100},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rule{
				RawLine:    tt.fields.RawLine,
				Filename:   tt.fields.Filename,
				LineNumber: tt.fields.LineNumber,
				Enabled:    tt.fields.Enabled,
			}
			if got := r.Evaluate(tt.args.item); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
