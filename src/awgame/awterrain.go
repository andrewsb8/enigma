package awgame

import (
	"log"
	"strconv"
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
			game.Awmap.Tiles[i] = append(game.Awmap.Tiles[i], CreateTile())
			GetTerrainFromID(row[j], game.Awmap.Tiles[i][j])
		}
	}
}

func CreateTile() *Tile {
	tile := Tile{}
	tile.Can_capture = false
	return &tile
}

func GetTerrainFromID(terrain_id string, tile *Tile) {
	int_id, err := strconv.Atoi(terrain_id)
	if err != nil {
		log.Fatal("Bad terrain id. Expected integer. Check terrain file.")
	} else {
		tile.Terrain_id = int_id
	}

	if int_id == 1 {
		tile.Terrain_type = "plain"
	} else if int_id == 2 {
		tile.Terrain_type = "mountain"
	} else if int_id == 3 {
		tile.Terrain_type = "woods"
	} else if int_id >= 4 && int_id <= 14 {
		tile.Terrain_type = "river"
	} else if int_id >= 15 && int_id <= 27 {
		tile.Terrain_type = "road"
	} else if int_id == 28 {
		tile.Terrain_type = "sea"
	} else if int_id >= 29 && int_id <= 32 {
		tile.Terrain_type = "shoals"
	} else if int_id == 33 {
		tile.Terrain_type = "shoals"
	} else if (int_id >= 101 && int_id <= 110) || int_id == 113 || int_id == 114 {
		tile.Terrain_type = "pipe"
	} else if int_id == 111 || int_id == 112 {
		tile.Terrain_type = "silo"
	} else if int_id == 34 {
		tile.Terrain_type = "city"
		tile.Can_capture = true
	} else if int_id == 35 { // these represent neutral properties
		tile.Terrain_type = "base"
		tile.Can_capture = true
	} else if int_id == 36 {
		tile.Terrain_type = "airport"
		tile.Can_capture = true
	} else if int_id == 37 {
		tile.Terrain_type = "port"
		tile.Can_capture = true
	} else if int_id == 145 {
		tile.Terrain_type = "lab"
		tile.Can_capture = true
	} else if int_id == 133 {
		tile.Terrain_type = "comm tower"
		tile.Can_capture = true
	}
}
