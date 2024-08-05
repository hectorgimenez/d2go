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
		want    RuleResult
		wantErr bool
	}{
		{
			name: "Basic rule with posion dmg, ethereal is not specified as a condition so it should be ignored",
			fields: fields{
				RawLine:    "[name] == smallcharm && [quality] == magic  # (([poisonlength]*25)*[poisonmaxdam])/256 >= 123",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					ID:       603,
					Name:     "SmAlLCharM",
					Quality:  item.QualityMagic,
					Ethereal: true,
					Stats: []stat.Data{
						{ID: stat.PoisonLength, Value: 20},
						{ID: stat.PoisonMaxDamage, Value: 100},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Complex rule with flags",
			fields: fields{
				RawLine:    "[type] == armor && [quality] <= superior && [flag] != ethereal # ([itemmaxdurabilitypercent] == 0 || [itemmaxdurabilitypercent] == 15) && ([sockets] == 0 || [sockets] == 3 || [sockets] == 4)",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					ID:       373,
					Name:     "mageplate",
					Quality:  item.QualitySuperior,
					Ethereal: false,
					Stats: []stat.Data{
						{ID: stat.MaxDurabilityPercent, Value: 15},
						{ID: stat.NumSockets, Value: 4},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Armor with +3 Sorc skills",
			fields: fields{
				RawLine: "[type] == armor # [sorceressskills] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:   373,
					Name: "mageplate",
					Stats: []stat.Data{
						{ID: stat.AddClassSkills, Value: 3, Layer: 1},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Armor with +3 Glacial Spike",
			fields: fields{
				RawLine: "[type] == armor  # [skillglacialspike] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:   373,
					Name: "mageplate",
					Stats: []stat.Data{
						{ID: stat.SingleSkill, Value: 3, Layer: 55},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Unid item matching base stats should return partial match",
			fields: fields{
				RawLine: "[type] == armor && [quality] == magic # [defense] == 200",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Identified: false,
					Name:       "mageplate",
					Quality:    item.QualityMagic,
				},
			},
			want: RuleResultPartial,
		},
		{
			name: "Basic rule without stats or maxquantity",
			fields: fields{
				RawLine: "[type] == assassinclaw && [class] == elite && [quality] == magic # #",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         187,
					Identified: false,
					Name:       "GreaterTalons",
					Quality:    item.QualityMagic,
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Basic rule for a white superior item with enhanceddefense",
			fields: fields{
				RawLine: "[type] == armor && [quality] == superior # [enhanceddefense] >= 15 #",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         373,
					Name:       "mageplate",
					Quality:    item.QualitySuperior,
					Stats: []stat.Data{
						{ID: stat.EnhancedDefense, Value: 15},
						{ID: stat.Defense, Value: 301},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Basic rule for a white superior item with enhanceddamage",
			fields: fields{
				RawLine: "[type] == sword # [enhanceddamage] >= 15 #",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         234,
					Name:       "colossusblade",
					Quality:    item.QualitySuperior,
					Stats: []stat.Data{
						{ID: stat.EnhancedDamage, Value: 15},
						{ID: stat.MinDamage, Value: 28},
						{ID: stat.MaxDamage, Value: 74},
						{ID: stat.TwoHandedMinDamage, Value: 66},
						{ID: stat.TwoHandedMaxDamage, Value: 132},
					},
				},
			},
			want: RuleResultFullMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewRule(tt.fields.RawLine, tt.fields.Filename, tt.fields.LineNumber)
			require.NoError(t, err)
			got, err := r.Evaluate(tt.args.item)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		rawRule    string
		filename   string
		lineNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    Rule
		wantErr bool
	}{
		{
			name: "Ensure [color] returns error, not supported yet",
			args: args{
				rawRule: "[type] == armor && [color] == 1000 && [quality] == magic",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRule(tt.args.rawRule, tt.args.filename, tt.args.lineNumber)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
			}
		})
	}
}

func BenchmarkEvaluate(b *testing.B) {
	it := data.Item{
		ID:      0,
		Name:    "Axe",
		Quality: item.QualitySuperior,
	}

	rule, err := NewRule(
		"[type] == amulet && [quality] == crafted # ([shapeshiftingskilltab] >= 2 || [elementalskilltab] >= 2 || [druidsummoningskilltab] >= 2) && [fcr] >= 10 && ([strength]+[maxhp]+[maxmana] >= 60 || [dexterity]+[maxhp]+[maxmana] >= 60 || [strength]+[dexterity]+[maxhp] >= 50 || [strength]+[dexterity]+[maxmana] >= 55)",
		"test",
		1,
	)
	require.NoError(b, err)

	for n := 0; n < b.N; n++ {
		rule.Evaluate(it)
	}
}
