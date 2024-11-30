package play

import (
	"enigma/src/awgame"
	"enigma/src/awio"
)

func Play(map_file string) {
	map_state := awio.GetMapState(map_file)
	game := awgame.Game{}
	game.Awmap.Map_state = map_state
	awgame.ParseMapState(game)
}
