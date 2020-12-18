package battle

import (
	"math/rand"
	"time"
)

// TODO: might want to move this, leaving it here so it's not forgotten
func init() {
	rand.Seed(time.Now().Unix())
}

// pctChance returns true `chance`% of the time
func pctChance(chance int) bool {
	return rand.Intn(100)+1 <= chance
}

func hasSubOrStatus(target *BoardPosition) bool {
	return target.SubstituteRemainingHP > 0 || target.Pokemon.Status != STATUS_NONE
}

/*
	taken from: https://github.com/pret/pokered/blob/master/constants/move_effect_constants.asm

	{stat}_(UP|DOWN)(1|2) means that the move raises the user's (or lowers the target's) corresponding stat modifier by 1 (or 2) stages
	{status condition}_SIDE_EFFECT means that the move has a side chance of causing that condition
	{status condition}_EFFECT means that the move causes the status condition every time it hits the target
*/

// NO_ADDITIONAL_EFFECT
func NoAdditionalEffect(board *Board, target *BoardPosition) bool {
	return true
}

// POISON_SIDE_EFFECT1
func PoisonSideEffect1(board *Board, target *BoardPosition) bool {
	if hasSubOrStatus(target) || target.Pokemon.Type1 == TYPE_POISON || target.Pokemon.Type2 == TYPE_POISON {
		return false
	}
	if pctChance(20) {
		target.Pokemon.Status = STATUS_POISONED
	}
	return true
}

// DRAIN_HP_EFFECT
// BURN_SIDE_EFFECT1
// FREEZE_SIDE_EFFECT
// PARALYZE_SIDE_EFFECT1
// EXPLODE_EFFECT  Explosion, Self Destruct
// DREAM_EATER_EFFECT
// MIRROR_MOVE_EFFECT
// ATTACK_UP1_EFFECT
// DEFENSE_UP1_EFFECT
// SPEED_UP1_EFFECT
// SPECIAL_UP1_EFFECT
// ACCURACY_UP1_EFFECT
// EVASION_UP1_EFFECT
// PAY_DAY_EFFECT
// SWIFT_EFFECT
// ATTACK_DOWN1_EFFECT
// DEFENSE_DOWN1_EFFECT
// SPEED_DOWN1_EFFECT
// SPECIAL_DOWN1_EFFECT
// ACCURACY_DOWN1_EFFECT
// EVASION_DOWN1_EFFECT
// CONVERSION_EFFECT
// HAZE_EFFECT
// BIDE_EFFECT
// THRASH_PETAL_DANCE_EFFECT
// SWITCH_AND_TELEPORT_EFFECT
// TWO_TO_FIVE_ATTACKS_EFFECT
// FLINCH_SIDE_EFFECT1
// SLEEP_EFFECT
// POISON_SIDE_EFFECT2
// BURN_SIDE_EFFECT2
// PARALYZE_SIDE_EFFECT2
// FLINCH_SIDE_EFFECT2
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
// ATTACK_UP2_EFFECT
// DEFENSE_UP2_EFFECT
// SPEED_UP2_EFFECT
// SPECIAL_UP2_EFFECT
// ACCURACY_UP2_EFFECT
// EVASION_UP2_EFFECT
// HEAL_EFFECT  Recover, Softboiled, Rest
// TRANSFORM_EFFECT
// ATTACK_DOWN2_EFFECT
// DEFENSE_DOWN2_EFFECT
// SPEED_DOWN2_EFFECT
// SPECIAL_DOWN2_EFFECT
// ACCURACY_DOWN2_EFFECT
// EVASION_DOWN2_EFFECT
// LIGHT_SCREEN_EFFECT
// REFLECT_EFFECT
// POISON_EFFECT
// PARALYZE_EFFECT
// ATTACK_DOWN_SIDE_EFFECT
// DEFENSE_DOWN_SIDE_EFFECT
// SPEED_DOWN_SIDE_EFFECT
// SPECIAL_DOWN_SIDE_EFFECT
// CONFUSION_SIDE_EFFECT
// TWINEEDLE_EFFECT
// SUBSTITUTE_EFFECT
// HYPER_BEAM_EFFECT
// RAGE_EFFECT
// MIMIC_EFFECT
// METRONOME_EFFECT
// LEECH_SEED_EFFECT
// SPLASH_EFFECT
// DISABLE_EFFECT
