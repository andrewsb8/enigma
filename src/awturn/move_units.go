package awturn

import (
	"enigma/src/awgame"
)

func MoveUnits(player *awgame.Player) {
	for i := 0; i < len(player.Unit_ids); i++ {
		MoveUnit(player.Units[*player.Unit_ids[i]])
	}
}

func MoveUnit(unit *awgame.Unit) {
	// for now, just move units down one space
	unit.Y_position -= 1
}
