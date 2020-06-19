package main

import (
	"fmt"

	"github.com/pbar1/pokego/pkg/battle"
)

// this is some ugly code -- I know -- it's just a POC
func main() {
	// pikaConfusion()
	charizardVsVenusaur()
}

func pikaConfusion() {
	pika := battle.PartyPokemon{
		Base:             battle.PKMN_PIKACHU,
		Level:            81,
		IVHP:             7,
		IVAttack:         8,
		IVDefense:        13,
		IVSpecialAttack:  9,
		IVSpecialDefense: 9,
		IVSpeed:          5,
		IVSpecial:        9,
		EVHP:             22850,
		EVAttack:         23140,
		EVDefense:        17280,
		EVSpecialAttack:  19625,
		EVSpecialDefense: 19625,
		EVSpeed:          24795,
		EVSpecial:        19625,
		Move1:            battle.MOVE_THUNDERBOLT,
		Move2:            battle.MOVE_NONE,
		Move3:            battle.MOVE_NONE,
		Move4:            battle.MOVE_NONE,
	}
	pika.InitializeStats()

	// pikaJSON, _ := json.Marshal(pika)
	// fmt.Println(string(pikaJSON))

	for {
		damage := battle.CalculateDamage(&pika.Move1, &pika, &pika)
		if damage > pika.BattleHP {
			pika.BattleHP = 0
		} else {
			pika.BattleHP -= damage
		}
		fmt.Printf("%s used %s on itself! Did %d damage, current HP: %d\n", pika.Base.Name, pika.Move1.Name, damage, pika.BattleHP)
		if pika.BattleHP == 0 {
			fmt.Printf("%s fainted!\n", pika.Base.Name)
			break
		}
	}
}

func charizardVsVenusaur() {
	charizard := battle.PartyPokemon{
		Base:             battle.PKMN_CHARIZARD,
		Level:            100,
		IVHP:             1,
		IVAttack:         1,
		IVDefense:        1,
		IVSpecialAttack:  1,
		IVSpecialDefense: 1,
		IVSpeed:          1,
		IVSpecial:        1,
		EVHP:             0,
		EVAttack:         0,
		EVDefense:        0,
		EVSpecialAttack:  0,
		EVSpecialDefense: 0,
		EVSpeed:          0,
		EVSpecial:        0,
		Move1:            battle.MOVE_FLAMETHROWER,
		Move2:            battle.MOVE_NONE,
		Move3:            battle.MOVE_NONE,
		Move4:            battle.MOVE_NONE,
	}
	charizard.InitializeStats()

	venusaur := battle.PartyPokemon{
		Base:             battle.PKMN_VENUSAUR,
		Level:            100,
		IVHP:             1,
		IVAttack:         1,
		IVDefense:        1,
		IVSpecialAttack:  1,
		IVSpecialDefense: 1,
		IVSpeed:          1,
		IVSpecial:        1,
		EVHP:             0,
		EVAttack:         0,
		EVDefense:        0,
		EVSpecialAttack:  0,
		EVSpecialDefense: 0,
		EVSpeed:          0,
		EVSpecial:        0,
		Move1:            battle.MOVE_SURF,
		Move2:            battle.MOVE_NONE,
		Move3:            battle.MOVE_NONE,
		Move4:            battle.MOVE_NONE,
	}
	venusaur.InitializeStats()

	fmt.Printf("%s\thp: %d, atk: %d, def %d, speed: %d, special: %d\n", charizard.Base.Name, charizard.BattleHP, charizard.BattleAttack, charizard.BattleDefense, charizard.BattleSpeed, charizard.BattleSpecial)
	fmt.Printf("%s\thp: %d, atk: %d, def %d, speed: %d, special: %d\n\n", venusaur.Base.Name, venusaur.BattleHP, venusaur.BattleAttack, venusaur.BattleDefense, venusaur.BattleSpeed, venusaur.BattleSpecial)

	for turn := 1; true; turn++ {

		fmt.Printf("TURN %d\n", turn)

		// find out turn order, of course speed tie is not implemented
		var first *battle.PartyPokemon
		var second *battle.PartyPokemon
		if charizard.BattleSpeed >= venusaur.BattleSpeed {
			first = &charizard
			second = &venusaur
		} else {
			first = &venusaur
			second = &charizard
		}

		// hard-coding some moves
		firstDamageOnSecond := battle.CalculateDamage(&first.Move1, first, second)
		if firstDamageOnSecond > second.BattleHP {
			second.BattleHP = 0
		} else {
			second.BattleHP -= firstDamageOnSecond
		}
		fmt.Printf("\t%s used %s on %s! Did %d damage, current HP: %d\n", first.Base.Name, first.Move1.Name, second.Base.Name, firstDamageOnSecond, second.BattleHP)
		if second.BattleHP == 0 {
			fmt.Printf("\t%s fainted!\n", second.Base.Name)
			break
		}

		secondDamageOnFirst := battle.CalculateDamage(&second.Move1, second, first)
		if secondDamageOnFirst > first.BattleHP {
			first.BattleHP = 0
		} else {
			first.BattleHP -= secondDamageOnFirst
		}
		fmt.Printf("\t%s used %s on %s! Did %d damage, current HP: %d\n", second.Base.Name, second.Move1.Name, first.Base.Name, secondDamageOnFirst, first.BattleHP)
		if first.BattleHP == 0 {
			fmt.Printf("\t%s fainted!\n", first.Base.Name)
			break
		}

	}
}
