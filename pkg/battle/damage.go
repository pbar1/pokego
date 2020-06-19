package battle

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CalculateDamage(move *Move, attacker, defender *PartyPokemon) uint {
	random := float64(rand.Intn(39)+217) / float64(255)
	stab := 1.0
	if move.Type == attacker.Base.Type1 || move.Type == attacker.Base.Type2 {
		stab = 1.5
	}
	typeEff := TypeEffectiveness(GEN1, move.Type, defender.Base.Type1) * TypeEffectiveness(GEN1, move.Type, defender.Base.Type2)
	return uint((((2.0*float64(attacker.Level)/5.0+2.0)*float64(move.Power)*float64(attacker.BattleSpecial)/float64(defender.BattleSpecial))/50.0 + 2.0) * random * stab * typeEff)
}
