package battle

import (
	"reflect"
	"testing"
)

func Test_calcDVHP(t *testing.T) {
	type args struct {
		dvAtk int
		dvDef int
		dvSpe int
		dvSpc int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Mewtwo A",
			args: args{
				dvAtk: 14,
				dvDef: 5,
				dvSpe: 8,
				dvSpc: 6,
			},
			want: 4,
		},
		{
			name: "Mewtwo B",
			args: args{
				dvAtk: 2,
				dvDef: 6,
				dvSpe: 10,
				dvSpc: 12,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcDVHP(tt.args.dvAtk, tt.args.dvDef, tt.args.dvSpe, tt.args.dvSpc); got != tt.want {
				t.Errorf("calcDVHP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcStatHP(t *testing.T) {
	type args struct {
		base    int
		dv      int
		statexp int
		level   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Level 81 Pikachu HP",
			args: args{
				base:    35,
				dv:      7,
				statexp: 22850,
				level:   81,
			},
			want: 189,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcStatHP(tt.args.base, tt.args.dv, tt.args.statexp, tt.args.level); got != tt.want {
				t.Errorf("calcStatHP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcStatOther(t *testing.T) {
	type args struct {
		base    int
		dv      int
		statexp int
		level   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Level 81 Pikachu Attack",
			args: args{
				base:    55,
				dv:      8,
				statexp: 23140,
				level:   81,
			},
			want: 137,
		},
		{
			name: "Level 81 Pikachu Defense",
			args: args{
				base:    30,
				dv:      13,
				statexp: 17280,
				level:   81,
			},
			want: 101,
		},
		{
			name: "Level 81 Pikachu Speed",
			args: args{
				base:    90,
				dv:      5,
				statexp: 24795,
				level:   81,
			},
			want: 190,
		},
		{
			name: "Level 81 Pikachu Special",
			args: args{
				base:    50,
				dv:      9,
				statexp: 19625,
				level:   81,
			},
			want: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcStatOther(tt.args.base, tt.args.dv, tt.args.statexp, tt.args.level); got != tt.want {
				t.Errorf("calcStatOther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBattlePokemon(t *testing.T) {
	type args struct {
		opts NewBattlePokemonOpts
	}
	tests := []struct {
		name    string
		args    args
		want    BattlePokemon
		wantErr bool
	}{
		{
			name: "Level 81 Pikachu",
			args: args{opts: NewBattlePokemonOpts{
				Base: BasePokemon{
					DexNo:   25,
					Name:    "Pikachu",
					Type1:   TypeElectric,
					Type2:   TypeNone,
					BaseHP:  35,
					BaseAtk: 55,
					BaseDef: 30,
					BaseSpe: 90,
					BaseSpc: 50,
				},
				Level:      81,
				DVAtk:      8,
				DVDef:      13,
				DVSpe:      5,
				DVSpc:      9,
				StatExpHP:  22850,
				StatExpAtk: 23140,
				StatExpDef: 17280,
				StatExpSpe: 24795,
				StatExpSpc: 19625,
				Move1:      MOVE_HYPERBEAM,
				Move2:      MOVE_FIREBLAST,
				Move3:      MOVE_BLIZZARD,
				Move4:      MOVE_THUNDER,
			}},
			want: BattlePokemon{
				DexNo:        25,
				CurrentHP:    189,
				Level:        81,
				Status:       StatusNone,
				Type1:        TypeElectric,
				Type2:        TypeNone,
				Move1:        MOVE_HYPERBEAM,
				Move2:        MOVE_FIREBLAST,
				Move3:        MOVE_BLIZZARD,
				Move4:        MOVE_THUNDER,
				StatHP:       189,
				StatAtk:      137,
				StatDef:      101,
				StatSpe:      190,
				StatSpc:      128,
				SleepCounter: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBattlePokemon(tt.args.opts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBattlePokemon() got = %v, want %v", got, tt.want)
			}
		})
	}
}
