package battle

import (
	"math"
)

// NewBattlePokemon constructs a BattlePokemon from the options provided
// https://www.smogon.com/ingame/guides/rby_gsc_stats
func NewBattlePokemon(opts NewBattlePokemonOpts) BattlePokemon {
	dvHP := calcDVHP(opts.DVAtk, opts.DVDef, opts.DVSpe, opts.DVSpc)
	statHP := calcStatHP(opts.Base.BaseHP, dvHP, opts.StatExpHP, opts.Level)
	bp := BattlePokemon{
		DexNo:        opts.Base.DexNo,
		Level:        opts.Level,
		Type1:        opts.Base.Type1,
		Type2:        opts.Base.Type2,
		Move1:        opts.Move1,
		Move2:        opts.Move2,
		Move3:        opts.Move3,
		Move4:        opts.Move4,
		Status:       StatusNone,
		StatHP:       statHP,
		StatAtk:      calcStatOther(opts.Base.BaseAtk, opts.DVAtk, opts.StatExpAtk, opts.Level),
		StatDef:      calcStatOther(opts.Base.BaseDef, opts.DVDef, opts.StatExpDef, opts.Level),
		StatSpe:      calcStatOther(opts.Base.BaseSpe, opts.DVSpe, opts.StatExpSpe, opts.Level),
		StatSpc:      calcStatOther(opts.Base.BaseSpc, opts.DVSpc, opts.StatExpSpc, opts.Level),
		CurrentHP:    statHP,
		SleepCounter: 0,
	}
	return bp
}

// calcDVHP calculates the DV value for the HP stat by taking the least significant bit
// (in binary) for each DV and then concatenating them in order of Attack, Defense,
// Speed, and finally Special.
func calcDVHP(dvAtk, dvDef, dvSpe, dvSpc int) int {
	return (dvAtk&1)<<3 + (dvDef&1)<<2 + (dvSpe&1)<<1 + dvSpc&1
}

// calcStatHP calculates the in-battle max HP stat.
func calcStatHP(base, dv, statexp, level int) int {
	return int(((float64(base)+float64(dv))*2.0+float64(int(math.Ceil(math.Sqrt(float64(statexp)))/4.0)))*float64(level)/100.0) + level + 10
}

// calcStatOther calculates the in-battle stat for Attack, Defense, Speed, or Special.
func calcStatOther(base, dv, statexp, level int) int {
	return int(((float64(base)+float64(dv))*2.0+float64(int(math.Ceil(math.Sqrt(float64(statexp)))/4.0)))*float64(level)/100.0) + 5
}

func (p *ActivePokemon) Reset() {
	p.LvAtk = 0
	p.LvDef = 0
	p.LvSpe = 0
	p.LvSpc = 0
	p.LvAcc = 0
	p.LvEvasion = 0
	p.ConfusionCounter = 0
	p.ToxSeedCounter = 0
	p.SubstituteCurrentHP = 0
	p.ReflectUp = false
	p.LightScreenUp = false
	// reset types, moves
}
