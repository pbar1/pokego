package battle

import "math"

type PartyPokemon struct {
	Base                 BasePokemon
	Level                uint8
	IVHP                 uint8
	IVAttack             uint8
	IVDefense            uint8
	IVSpecialAttack      uint8
	IVSpecialDefense     uint8
	IVSpeed              uint8
	IVSpecial            uint8
	EVHP                 uint
	EVAttack             uint
	EVDefense            uint
	EVSpecialAttack      uint
	EVSpecialDefense     uint
	EVSpeed              uint
	EVSpecial            uint
	Move1                Move
	Move2                Move
	Move3                Move
	Move4                Move
	BattleHP             uint
	BattleAttack         uint
	BattleDefense        uint
	BattleSpecialAttack  uint
	BattleSpecialDefense uint
	BattleSpeed          uint
	BattleSpecial        uint
}

func (p *PartyPokemon) InitializeStats() {
	p.BattleHP = calcStatGen1(true, p.Base.HP, p.IVHP, p.Level, p.EVHP)
	p.BattleAttack = calcStatGen1(false, p.Base.Attack, p.IVAttack, p.Level, p.EVAttack)
	p.BattleDefense = calcStatGen1(false, p.Base.Defense, p.IVDefense, p.Level, p.EVDefense)
	p.BattleSpecialAttack = calcStatGen1(false, p.Base.SpecialAttack, p.IVSpecialAttack, p.Level, p.EVSpecialAttack)
	p.BattleSpecialDefense = calcStatGen1(false, p.Base.SpecialDefense, p.IVSpecialDefense, p.Level, p.EVSpecialDefense)
	p.BattleSpeed = calcStatGen1(false, p.Base.Speed, p.IVSpeed, p.Level, p.EVSpeed)
	p.BattleSpecial = calcStatGen1(false, p.Base.Special, p.IVSpecial, p.Level, p.EVSpecial)
}

func calcStatGen1(isHP bool, base, iv, level uint8, ev uint) uint {
	if isHP {
		return uint(math.Floor((float64(base)+float64(iv))*2.0+math.Floor(math.Sqrt(float64(ev))/4.0))*float64(level)/100.0) + uint(level) + 10
	}
	return uint(math.Floor((float64(base)+float64(iv))*2.0+math.Floor(math.Sqrt(float64(ev))/4.0))*float64(level)/100.0) + 5
}
