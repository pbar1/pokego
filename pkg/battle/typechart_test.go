package battle

import "testing"

func TestTypeEffectiveness(t *testing.T) {
	type args struct {
		gen      Generation
		attacker DamageType
		defender DamageType
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Gen1_FireBeatsGrass", args{Gen1, TypeFire, TypeGrass}, 2.0},
		{"Gen1_WaterBeatsFire", args{Gen1, TypeWater, TypeFire}, 2.0},
		{"Gen1_GrassBeatsWater", args{Gen1, TypeGrass, TypeWater}, 2.0},
		{"Gen1_IceNeutralFire", args{Gen1, TypeIce, TypeFire}, 1.0},
		{"Gen1_GhostNotAffectPsychic", args{Gen1, TypeGhost, TypePsychic}, 0.0},
		{"Gen1_BugBeatsPoison", args{Gen1, TypeBug, TypePoison}, 2.0},
		{"Gen1_PoisonBeatsBug", args{Gen1, TypePoison, TypeBug}, 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeEffectiveness(tt.args.gen, tt.args.attacker, tt.args.defender); got != tt.want {
				t.Errorf("TypeEffectiveness() = %v, want %v", got, tt.want)
			}
		})
	}
}
