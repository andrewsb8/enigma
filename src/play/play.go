package play

import (
	"enigma/src/awio"
)

func Play(map_file string) {
	awio.ReadMap(map_file)
}
