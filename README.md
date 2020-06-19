# pokego

Toy implementation of the Pokemon battle engine in Go.
Inspired by the disassembly of [Pokemon Ruby/Sapphire](https://github.com/pret/pokeruby).

## Example

Pikachu hurting itself in confusion?
```
> ./pokego
Pikachu used Thunderbolt on itself! Did 65 damage, current HP: 124
Pikachu used Thunderbolt on itself! Did 60 damage, current HP: 64
Pikachu used Thunderbolt on itself! Did 58 damage, current HP: 6
Pikachu used Thunderbolt on itself! Did 64 damage, current HP: 0
Pikachu fainted!
```

Charizard battling a surfing Venusaur?
```
> ./pokego
Charizard	hp: 268, atk: 175, def 163, speed: 207, special: 177
Venusaur	hp: 272, atk: 171, def 173, speed: 167, special: 207

TURN 1
	Charizard used Flamethrower on Venusaur! Did 186 damage, current HP: 86
	Venusaur used Surf on Charizard! Did 173 damage, current HP: 95
TURN 2
	Charizard used Flamethrower on Venusaur! Did 187 damage, current HP: 0
	Venusaur fainted!
```
