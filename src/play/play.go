package play

import (
	"enigma/src/awgame"
	"enigma/src/awio"
	"enigma/src/awturn"
)

func Play(map_file string, terrain_file string, turns int) {
	map_state := awio.GetMapState(map_file)
	terrain := awio.GetTerrain(terrain_file)
	game := awgame.ParseGameInfo(map_state, terrain)

	for i := 0; i < turns; i++ {
		for j := 0; j < game.Num_players; j++ {
			awturn.MoveUnits(game.Players[j])
		}
	}
}
