package battle

import (
	"math/rand"
	"time"
)

// TODO: might want to move this, leaving it here so it's not forgotten
func init() {
	rand.Seed(time.Now().Unix())
}

// -----------------------------------------------------------------------------
// Helper functions
// -----------------------------------------------------------------------------

// pctChance returns true `chance`% of the time
func pctChance(chance int) bool {
	return rand.Intn(100)+1 <= chance
}

// byteChance returns true `chance`/256 of the time. It's supposed to allow for
// a 1/256 chance for things to be false even at a "100%" chance equivalent.
func byteChance(chance int) bool {
	return rand.Intn(256)+1 <= chance
}

func hasSub(target *ActivePokemon) bool {
	return target.SubstituteCurrentHP > 0
}

func hasSubOrStatus(target *ActivePokemon) bool {
	return target.SubstituteCurrentHP > 0 || target.Pokemon.Status != StatusNone
}

func statMod(delta int, targetStatLevel *int) bool {
	if *targetStatLevel == 6 {
		return false
	}
	*targetStatLevel += delta
	if *targetStatLevel > 6 {
		*targetStatLevel = 6
	} else if *targetStatLevel < -6 {
		*targetStatLevel = -6
	}
	return true
}

// -----------------------------------------------------------------------------
// Move effect functions
//
// See: https://github.com/pret/pokered/blob/master/constants/move_effect_constants.asm
// See: https://github.com/pret/pokered/blob/master/engine/battle/effects.asm
//
// {stat}_(UP|DOWN)(1|2) means that the move raises the user's (or lowers the target's) corresponding stat modifier by 1 (or 2) stages
// {status condition}_SIDE_EFFECT means that the move has a side chance of causing that condition
// {status condition}_EFFECT means that the move causes the status condition every time it hits the target
// -----------------------------------------------------------------------------

// -------------------------------------
// No effect
// -------------------------------------

// NO_ADDITIONAL_EFFECT
func NoAdditionalEffect(move *Move, target *ActivePokemon, board *Board) bool {
	return true
}

// PAY_DAY_EFFECT
func PayDayEffect(move *Move, target *ActivePokemon, board *Board) bool {
	return true
}

// SPLASH_EFFECT
func SplashEffect(move *Move, target *ActivePokemon, board *Board) bool {
	return true
}

// SWITCH_AND_TELEPORT_EFFECT
func SwitchAndTeleportEffect(move *Move, target *ActivePokemon, board *Board) bool {
	return true
}

// -------------------------------------
// Stat modifiers
// -------------------------------------

// ATTACK_UP1_EFFECT
func AttackUp1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+1, &target.LvAtk)
}

// DEFENSE_UP1_EFFECT
func DefenseUp1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+1, &target.LvDef)
}

// SPEED_UP1_EFFECT
func SpeedUp1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+1, &target.LvSpe)
}

// SPECIAL_UP1_EFFECT
func SpecialUp1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+1, &target.LvSpc)
}

// ACCURACY_UP1_EFFECT
func AccuracyUp1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+1, &target.LvAcc)
}

// EVASION_UP1_EFFECT
func EvasionUp1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+1, &target.LvEvasion)
}

// ATTACK_DOWN1_EFFECT
func AttackDown1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-1, &target.LvAtk)
}

// DEFENSE_DOWN1_EFFECT
func DefenseDown1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-1, &target.LvDef)
}

// SPEED_DOWN1_EFFECT
func SpeedDown1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-1, &target.LvSpe)
}

// SPECIAL_DOWN1_EFFECT
func SpecialDown1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-1, &target.LvSpc)
}

// ACCURACY_DOWN1_EFFECT
func AccuracyDown1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-1, &target.LvAcc)
}

// EVASION_DOWN1_EFFECT
func EvasionDown1Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-1, &target.LvEvasion)
}

// ATTACK_UP2_EFFECT
func AttackUp2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+2, &target.LvAtk)
}

// DEFENSE_UP2_EFFECT
func DefenseUp2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+2, &target.LvDef)
}

// SPEED_UP2_EFFECT
func SpeedUp2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+2, &target.LvSpe)
}

// SPECIAL_UP2_EFFECT
func SpecialUp2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+2, &target.LvSpc)
}

// ACCURACY_UP2_EFFECT
func AccuracyUp2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+2, &target.LvAcc)
}

// EVASION_UP2_EFFECT
func EvasionUp2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	return statMod(+2, &target.LvEvasion)
}

// ATTACK_DOWN2_EFFECT
func AttackDown2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-2, &target.LvAtk)
}

// DEFENSE_DOWN2_EFFECT
func DefenseDown2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-2, &target.LvDef)
}

// SPEED_DOWN2_EFFECT
func SpeedDown2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-2, &target.LvSpe)
}

// SPECIAL_DOWN2_EFFECT
func SpecialDown2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-2, &target.LvSpc)
}

// ACCURACY_DOWN2_EFFECT
func AccuracyDown2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-2, &target.LvAcc)
}

// EVASION_DOWN2_EFFECT
func EvasionDown2Effect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	return statMod(-2, &target.LvEvasion)
}

// ATTACK_DOWN_SIDE_EFFECT
func AttackDownSideEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	if byteChance(85) {
		return statMod(-1, &target.LvAtk)
	}
	return true
}

// DEFENSE_DOWN_SIDE_EFFECT
func DefenseDownSideEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	if byteChance(85) {
		return statMod(-1, &target.LvDef)
	}
	return true
}

// SPEED_DOWN_SIDE_EFFECT
func SpeedDownSideEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	if byteChance(85) {
		return statMod(-1, &target.LvSpe)
	}
	return true
}

// SPECIAL_DOWN_SIDE_EFFECT
func SpecialDownSideEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSub(target) {
		return false
	}
	if byteChance(85) {
		return statMod(-1, &target.LvSpc)
	}
	return true
}

// LIGHT_SCREEN_EFFECT
// TODO: this effect should show up before damage is calculated
func LightScreenEffect(move *Move, target *ActivePokemon, board *Board) bool {
	return false
}

// REFLECT_EFFECT
// TODO: this effect should show up before damage is calculated
func ReflectEffect(move *Move, target *ActivePokemon, board *Board) bool {
	return false
}

// -------------------------------------
// Non-volatile status conditions
// -------------------------------------

// POISON_SIDE_EFFECT1
func PoisonSideEffect1(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || target.Pokemon.Type1 == TypePoison || target.Pokemon.Type2 == TypePoison {
		return false
	}
	if pctChance(20) {
		target.Pokemon.Status = StatusPoison
	}
	return true
}

// BURN_SIDE_EFFECT1
func BurnSideEffect1(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || move.Type == target.Pokemon.Type1 || move.Type == target.Pokemon.Type2 {
		return false
	}
	if pctChance(10) {
		target.Pokemon.Status = StatusBurn
		// TODO: halve attack
	}
	return true
}

// FREEZE_SIDE_EFFECT
func FreezeSideEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || move.Type == target.Pokemon.Type1 || move.Type == target.Pokemon.Type2 {
		return false
	}
	if pctChance(10) {
		target.Pokemon.Status = StatusFreeze
	}
	return true
}

// PARALYZE_SIDE_EFFECT1
func ParalyzeSideEffect1(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || move.Type == target.Pokemon.Type1 || move.Type == target.Pokemon.Type2 {
		return false
	}
	if pctChance(10) {
		target.Pokemon.Status = StatusParalyze
		// TODO: quarter speed
	}
	return true
}

// SLEEP_EFFECT
// TODO: cannot move on turn affected pokemon wakes up
func SleepEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if move.Name == "Rest" {
		target.Pokemon.Status = StatusSleep
		target.Pokemon.SleepCounter = 2
		return true
	}
	if target.Pokemon.Status != StatusNone {
		return false
	}
	target.Pokemon.Status = StatusSleep
	target.Pokemon.SleepCounter = rand.Intn(7) + 1
	return true
}

// POISON_SIDE_EFFECT2
func PoisonSideEffect2(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || target.Pokemon.Type1 == TypePoison || target.Pokemon.Type2 == TypePoison {
		return false
	}
	if pctChance(40) {
		target.Pokemon.Status = StatusPoison
	}
	return true
}

// BURN_SIDE_EFFECT2
func BurnSideEffect2(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || move.Type == target.Pokemon.Type1 || move.Type == target.Pokemon.Type2 {
		return false
	}
	if pctChance(30) {
		target.Pokemon.Status = StatusBurn
		// TODO: halve attack
	}
	return true
}

// PARALYZE_SIDE_EFFECT2
func ParalyzeSideEffect2(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || move.Type == target.Pokemon.Type1 || move.Type == target.Pokemon.Type2 {
		return false
	}
	if pctChance(30) {
		target.Pokemon.Status = StatusParalyze
		// TODO: quarter speed
	}
	return true
}

// POISON_EFFECT
func PoisonEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if hasSubOrStatus(target) || target.Pokemon.Type1 == TypePoison || target.Pokemon.Type2 == TypePoison {
		return false
	}
	target.Pokemon.Status = StatusPoison
	return true
}

// PARALYZE_EFFECT
func ParalyzeEffect(move *Move, target *ActivePokemon, board *Board) bool {
	if target.Pokemon.Status != StatusNone || move.Type == target.Pokemon.Type1 || move.Type == target.Pokemon.Type2 {
		return false
	}
	target.Pokemon.Status = StatusParalyze
	return true
}

// -------------------------------------
// Volatile status conditions
// -------------------------------------

// -------------------------------------
// TODO: unimplemented / other
// -------------------------------------

// DRAIN_HP_EFFECT
func DrainHPEffect(move *Move, target *ActivePokemon, board *Board) bool {
	panic("not implemented")
}

// EXPLODE_EFFECT  Explosion, Self Destruct
// DREAM_EATER_EFFECT
// MIRROR_MOVE_EFFECT
// SWIFT_EFFECT
// CONVERSION_EFFECT
// HAZE_EFFECT
// BIDE_EFFECT
// THRASH_PETAL_DANCE_EFFECT
// TWO_TO_FIVE_ATTACKS_EFFECT

// FLINCH_SIDE_EFFECT1
func FlinchSideEffect1(move *Move, target *ActivePokemon, board *Board) bool {
	if pctChance(10) {
		target.Flinch = true
	}
	return true
}

// FLINCH_SIDE_EFFECT2
func FlinchSideEffect2(move *Move, target *ActivePokemon, board *Board) bool {
	if pctChance(30) {
		target.Flinch = true
	}
	return true
}

// OHKO_EFFECT  moves like Horn Drill
// CHARGE_EFFECT  moves like Solar Beam
// SUPER_FANG_EFFECT
// SPECIAL_DAMAGE_EFFECT  Seismic Toss, Night Shade, Sonic Boom, Dragon Rage, Psywave
// TRAPPING_EFFECT  moves like Wrap
// FLY_EFFECT
// ATTACK_TWICE_EFFECT
// JUMP_KICK_EFFECT  Jump Kick and Hi Jump Kick effect
// MIST_EFFECT
// FOCUS_ENERGY_EFFECT
// RECOIL_EFFECT  moves like Double Edge
// CONFUSION_EFFECT  Confuse Ray, Supersonic (not the move Confusion)
// HEAL_EFFECT  Recover, Softboiled, Rest
// TRANSFORM_EFFECT
// CONFUSION_SIDE_EFFECT
// TWINEEDLE_EFFECT
// SUBSTITUTE_EFFECT
// HYPER_BEAM_EFFECT
// RAGE_EFFECT
// MIMIC_EFFECT
// METRONOME_EFFECT
// LEECH_SEED_EFFECT
// DISABLE_EFFECT
