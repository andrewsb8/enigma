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
			game.Awmap.Tiles[i] = append(game.Awmap.Tiles[i], CreateTile(row[j]))
		}
	}
}

func CreateTile(terrain_id string) *Tile {
	tile := Tile{}
	tile.Can_capture = false
	tile.Capture_points = -1
	GetTerrainFromID(terrain_id, &tile)
	if tile.Can_capture {
		SetCaptureProperties(&tile)
	}
	return &tile
}

/*
Parse neutral terrain from map file. All of the nations have unique
codes for their properties once captured. These IDs can be parsed from
the map state file.
*/
func GetTerrainFromID(terrain_id string, tile *Tile) {
	int_id, err := strconv.Atoi(terrain_id)
	if err != nil {
		log.Fatal("Invalid terrain id. Check terrain file.")
	} else {
		tile.Terrain_id = int_id
	}

	if int_id == 1 {
		tile.Terrain_type = "plain"
		tile.Defense_stars = 1
		tile.Movement_cost_clear = [8]int{1, 1, 1, 2, 100, 100, 1, 100}
	} else if int_id == 2 {
		tile.Terrain_type = "mountain"
		tile.Defense_stars = 4
		tile.Movement_cost_clear = [8]int{2, 1, 100, 100, 100, 100, 1, 100}
	} else if int_id == 3 {
		tile.Terrain_type = "woods"
		tile.Defense_stars = 2
		tile.Movement_cost_clear = [8]int{1, 1, 2, 3, 100, 100, 1, 100}
	} else if int_id >= 4 && int_id <= 14 {
		tile.Terrain_type = "river"
		tile.Defense_stars = 0
		tile.Movement_cost_clear = [8]int{2, 1, 100, 100, 100, 100, 1, 100}
	} else if int_id >= 15 && int_id <= 27 {
		tile.Terrain_type = "road"
		tile.Defense_stars = 0
		tile.Movement_cost_clear = [8]int{1, 1, 1, 1, 100, 100, 1, 100}
	} else if int_id == 28 {
		tile.Terrain_type = "sea"
		tile.Defense_stars = 1
		tile.Movement_cost_clear = [8]int{100, 100, 100, 100, 1, 1, 1, 100}
	} else if int_id >= 29 && int_id <= 32 {
		tile.Terrain_type = "shoals"
		tile.Defense_stars = 0
		tile.Movement_cost_clear = [8]int{1, 1, 1, 1, 100, 1, 1, 100}
	} else if int_id == 33 {
		tile.Terrain_type = "reefs"
		tile.Defense_stars = 1
		tile.Movement_cost_clear = [8]int{100, 100, 100, 100, 2, 2, 1, 100}
	} else if (int_id >= 101 && int_id <= 110) || int_id == 113 || int_id == 114 {
		// pipe seems are pipes originally, but plains when broken
		tile.Terrain_type = "pipe"
		tile.Defense_stars = 0
		tile.Movement_cost_clear = [8]int{100, 100, 100, 100, 100, 100, 100, 1}
	} else if int_id == 111 || int_id == 112 {
		tile.Terrain_type = "silo"
		tile.Defense_stars = 3
		tile.Movement_cost_clear = [8]int{1, 1, 1, 1, 100, 100, 1, 100}
	} else if int_id == 34 { // these represent neutral properties
		tile.Terrain_type = "city"
		tile.Can_capture = true
	} else if int_id == 35 {
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
	} else {
		// leave as a placeholder to be changed when
		// parsing buildings in map state file
		tile.Terrain_type = "captured property"
		tile.Can_capture = true
	}
}

func SetCaptureProperties(tile *Tile) {
	if tile.Terrain_type == "headquarters" {
		tile.Defense_stars = 4
	} else {
		tile.Defense_stars = 3
	}
	tile.Capture_points = 20
	tile.Movement_cost_clear = [8]int{1, 1, 1, 1, 100, 100, 1, 100}
}
