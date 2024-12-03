package awgame

import (
	"strings"
)

func ParseTerrain(terrain []string, game *Game) {
	// terrain file ends with empty line, so -1
	game.Awmap.Map_height = len(terrain) - 1
	for i := 0; i < game.Awmap.Map_height; i++ {
		game.Awmap.Tiles = append(game.Awmap.Tiles, []*Tile{})
		row := strings.Split(terrain[i], ",")
		if i == 0 {
			game.Awmap.Map_width = len(row)
		}
		for j := 0; j < len(row); j++ {
			game.Awmap.Tiles[i] = append(game.Awmap.Tiles[i], &Tile{})
			GetTerrainFromID(row[j], game.Awmap.Tiles[i][j])
		}
	}
}

func GetTerrainFromID(terrain_id string, tile *Tile) {
	switch terrain_id { //abstract this out to different function in own file?
	case "34":
		tile.Terrain_id = 34
		tile.Terrain_type = "city"
		tile.Can_capture = true
	case "1": // plain
		tile.Terrain_id = 1
		tile.Terrain_type = "plain"
		tile.Can_capture = false
	default:
		break
	}
}
