package battle

import "testing"

func TestTypeEffectiveness(t *testing.T) {
	type args struct {
		gen      Generation
		attacker PokemonType
		defender PokemonType
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Gen1_FireBeatsGrass", args{GEN1, TYPE_FIRE, TYPE_GRASS}, 2.0},
		{"Gen1_WaterBeatsFire", args{GEN1, TYPE_WATER, TYPE_FIRE}, 2.0},
		{"Gen1_GrassBeatsWater", args{GEN1, TYPE_GRASS, TYPE_WATER}, 2.0},
		{"Gen1_IceNeutralFire", args{GEN1, TYPE_ICE, TYPE_FIRE}, 1.0},
		{"Gen1_GhostNotAffectPsychic", args{GEN1, TYPE_GHOST, TYPE_PSYCHIC}, 0.0},
		{"Gen1_BugBeatsPoison", args{GEN1, TYPE_BUG, TYPE_POISON}, 2.0},
		{"Gen1_PoisonBeatsBug", args{GEN1, TYPE_POISON, TYPE_BUG}, 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeEffectiveness(tt.args.gen, tt.args.attacker, tt.args.defender); got != tt.want {
				t.Errorf("TypeEffectiveness() = %v, want %v", got, tt.want)
			}
		})
	}
}
