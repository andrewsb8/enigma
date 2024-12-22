package play

import (
	"enigma/src/awgame"
	"enigma/src/awio"
	"fmt"
)

func Play(map_file string, terrain_file string) {
	map_state := awio.GetMapState(map_file)
	terrain := awio.GetTerrain(terrain_file)
	game := awgame.ParseGameInfo(map_state, terrain)
	fmt.Print(game.Awmap.Tiles[0][1].Terrain_type, "\n")
}
