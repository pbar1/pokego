package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pbar1/pokego/pkg/battle"
)

// gen1demoCmd represents the gen1demo command
var gen1demoCmd = &cobra.Command{
	Use:   "gen1demo",
	Short: "Generation 1 demo",
	Long:  `Generation 1 demo`,
	Run: func(cmd *cobra.Command, args []string) {
		board := battle.Board{
			Player1: DemoTrainerBlue,
			Player2: DemoTrainerOak,
		}

		p1idx := 0
		p2idx := 0
		emptypkm := battle.BattlePokemon{}
		for {
			if p1idx > 6 {
				fmt.Println("Player 2 wins")
				break
			}
			if p2idx > 6 {
				fmt.Println("Player 1 wins")
				break
			}

			if board.Player1.Party[p1idx] == emptypkm {
				p1idx++
			}
			if board.Player2.Party[p2idx] == emptypkm {
				p2idx++
			}

			if board.Player1.Party[p1idx].StatSpe > board.Player2.Party[p2idx].StatSpe {
				fmt.Printf("%s attacks first\n", board.Player1.Party[p1idx].DexNo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(gen1demoCmd)

	// gen1demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var (
	DemoTrainerBlue = battle.Player{
		Party: [6]battle.BattlePokemon{
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_PIDGEOT,
				Level:      600,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_WINGATTACK,
				Move2:      battle.MOVE_MIRRORMOVE,
				Move3:      battle.MOVE_SKYATTACK,
				Move4:      battle.MOVE_WHIRLWIND,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_ALAKAZAM,
				Level:      59,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_PSYBEAM,
				Move2:      battle.MOVE_PSYCHIC,
				Move3:      battle.MOVE_REFLECT,
				Move4:      battle.MOVE_RECOVER,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_RHYDON,
				Level:      61,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_LEER,
				Move2:      battle.MOVE_TAILWHIP,
				Move3:      battle.MOVE_FURYATTACK,
				Move4:      battle.MOVE_HORNDRILL,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_GYARADOS,
				Level:      63,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_DRAGONRAGE,
				Move2:      battle.MOVE_HYDROPUMP,
				Move3:      battle.MOVE_HYPERBEAM,
				Move4:      battle.MOVE_LEER,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_EXEGGUTOR,
				Level:      61,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_HYPNOSIS,
				Move2:      battle.MOVE_BARRAGE,
				Move3:      battle.MOVE_STOMP,
				Move4:      battle.Move{}, // TODO
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_CHARIZARD,
				Level:      65,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_FIREBLAST,
				Move2:      battle.MOVE_RAGE,
				Move3:      battle.MOVE_SLASH,
				Move4:      battle.MOVE_FIRESPIN,
			}),
		},
		Active: battle.ActivePokemon{},
	}

	DemoTrainerOak = battle.Player{
		Party: [6]battle.BattlePokemon{
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_TAUROS,
				Level:      66,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_TAKEDOWN,
				Move2:      battle.MOVE_LEER,
				Move3:      battle.MOVE_RAGE,
				Move4:      battle.MOVE_TAILWHIP,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_EXEGGUTOR,
				Level:      67,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_STOMP,
				Move2:      battle.MOVE_BARRAGE,
				Move3:      battle.MOVE_HYPNOSIS,
				Move4:      battle.Move{}, // TODO
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_ARCANINE,
				Level:      68,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_TAKEDOWN,
				Move2:      battle.MOVE_EMBER,
				Move3:      battle.MOVE_ROAR,
				Move4:      battle.MOVE_LEER,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_BLASTOISE,
				Level:      69,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_BITE,
				Move2:      battle.MOVE_WITHDRAW,
				Move3:      battle.MOVE_SKULLBASH,
				Move4:      battle.MOVE_HYDROPUMP,
			}),
			battle.NewBattlePokemon(battle.NewBattlePokemonOpts{
				Base:       battle.PKM_GYARADOS,
				Level:      70,
				DVAtk:      0,
				DVDef:      0,
				DVSpe:      0,
				DVSpc:      0,
				StatExpHP:  0,
				StatExpAtk: 0,
				StatExpDef: 0,
				StatExpSpe: 0,
				StatExpSpc: 0,
				Move1:      battle.MOVE_HYDROPUMP,
				Move2:      battle.MOVE_DRAGONRAGE,
				Move3:      battle.MOVE_LEER,
				Move4:      battle.MOVE_HYPERBEAM,
			}),
		},
		Active: battle.ActivePokemon{},
	}
)
