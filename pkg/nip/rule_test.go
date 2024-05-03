package nip

import (
	"testing"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/stretchr/testify/require"
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
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Basic rule with posion dmg, ethereal is not specified as a condition so it should be ignored",
			fields: fields{
				RawLine:    "[name] == smallcharm && [quality] == magic # (([poisonlength]*25)*[poisonmaxdam])/256 >= 123",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					Name:     "SmAlLCharM",
					Quality:  item.QualityMagic,
					Ethereal: true,
					Stats: []stat.Data{
						{ID: stat.PoisonLength, Value: 20},
						{ID: stat.PoisonMaxDamage, Value: 100},
					},
				},
			},
			want: true,
		},
		{
			name: "Complex rule with flags and enhanced defense",
			fields: fields{
				RawLine:    "[type] == armor && [quality] <= superior && [flag] != ethereal # [enhanceddefense] >= 15 && ([itemmaxdurabilitypercent] == 0 || [itemmaxdurabilitypercent] == 15) && ([sockets] == 0 || [sockets] == 3 || [sockets] == 4)",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					Name:     "mageplate",
					Quality:  item.QualitySuperior,
					Ethereal: false,
					Stats: []stat.Data{
						{ID: stat.EnhancedDefense, Value: 20},
						{ID: stat.MaxDurabilityPercent, Value: 15},
						{ID: stat.NumSockets, Value: 4},
					},
				},
			},
			want: true,
		},
		{
			name: "Armor with +3 Sorc skills",
			fields: fields{
				RawLine: "[type] == armor && [sorceressskills] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Name: "mageplate",
					Stats: []stat.Data{
						{ID: stat.AddClassSkills, Value: 3, Layer: 1},
					},
				},
			},
			want: true,
		},
		{
			name: "Armor with +3 Glacial Spike",
			fields: fields{
				RawLine: "[type] == armor && [skillglacialspike] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Name: "mageplate",
					Stats: []stat.Data{
						{ID: stat.SingleSkill, Value: 3, Layer: 55},
					},
				},
			},
			want: true,
		},
		{
			name: "Ensure [color] returns error, not supported yet",
			fields: fields{
				RawLine: "[type] == armor && [color] == 1000 && [quality] == magic",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Name:    "mageplate",
					Quality: item.QualityMagic,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := New(tt.fields.RawLine, tt.fields.Filename, tt.fields.LineNumber)
			require.NoError(t, err)
			got, err := r.Evaluate(tt.args.item)
			if !tt.wantErr {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
