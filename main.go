package main

import (
	"fmt"

	"github.com/pbar1/pokego/pkg/battle"
)

func main() {

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
