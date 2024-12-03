package play

import (
	"enigma/src/awgame"
	"enigma/src/awio"
)

func Play(map_file string, terrain_file string) {
	map_state := awio.GetMapState(map_file)
	terrain := awio.GetTerrain(terrain_file)
	game := awgame.Game{}
	awgame.ParseMap(map_state, terrain, &game)
}
